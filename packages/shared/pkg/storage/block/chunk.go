package block

import (
	"context"
	"errors"
	"fmt"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/e2b-dev/infra/packages/shared/pkg/utils"
)

const (
	// Chunks must always be bigger or equal to the block size.
	chunkSize = 2 * 1024 * 1024 // 2 MB
)

type chunker struct {
	ctx context.Context

	base  io.ReaderAt
	cache *cache

	size int64

	fetchers *utils.WaitMap
}

func newChunker(
	ctx context.Context,
	size,
	blockSize int64,
	base io.ReaderAt,
	cachePath string,
) (*chunker, error) {
	cache, err := newCache(size, blockSize, cachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file cache: %w", err)
	}

	chunker := &chunker{
		ctx:      ctx,
		size:     size,
		base:     base,
		cache:    cache,
		fetchers: utils.NewWaitMap(),
	}

	return chunker, nil
}

func (c *chunker) ReadAt(b []byte, off int64) (int, error) {
	slice, err := c.Slice(off, int64(len(b)))
	if err != nil {
		return 0, fmt.Errorf("failed to slice cache at %d-%d: %w", off, off+int64(len(b)), err)
	}

	return copy(b, slice), nil
}

func (c *chunker) Slice(off, length int64) ([]byte, error) {
	b, err := c.cache.Slice(off, length)
	if err == nil {
		return b, nil
	}

	if !errors.As(err, &ErrBytesNotAvailable{}) {
		return nil, fmt.Errorf("failed read from cache at offset %d: %w", off, err)
	}

	chunkErr := c.fetchToCache(off, length)
	if chunkErr != nil {
		return nil, fmt.Errorf("failed to ensure data at %d-%d: %w", off, off+length, chunkErr)
	}

	b, cacheErr := c.cache.Slice(off, length)
	if cacheErr != nil {
		return nil, fmt.Errorf("failed to read from cache after ensuring data at %d-%d: %w", off, off+length, cacheErr)
	}

	return b, nil
}

// fetchToCache ensures that the data at the given offset and length is available in the cache.
func (c *chunker) fetchToCache(off, len int64) error {
	var eg errgroup.Group

	blocks := listBlocks(off, off+len, chunkSize)

	for _, block := range blocks {
		start := block.start
		end := block.end

		eg.Go(func() error {
			return c.fetchers.Wait(block.start, func() error {
				select {
				case <-c.ctx.Done():
					return fmt.Errorf("error fetching range %d-%d: %w", start, end, c.ctx.Err())
				default:
				}

				b := make([]byte, end-start)

				_, err := c.base.ReadAt(b, start)
				if err != nil && !errors.Is(err, io.EOF) {
					return fmt.Errorf("failed to read chunk from base %d: %w", start, err)
				}

				_, cacheErr := c.cache.WriteAt(b, start)
				if cacheErr != nil {
					return fmt.Errorf("failed to write chunk %d to cache: %w", start, cacheErr)
				}

				return nil
			})
		})
	}

	err := eg.Wait()
	if err != nil {
		return fmt.Errorf("failed to ensure data at %d-%d: %w", off, off+len, err)
	}

	return nil
}

func (c *chunker) Close() error {
	return c.cache.Close()
}
