package header

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bits-and-blooms/bitset"
)

const (
	PageSize        = 2 << 11
	HugepageSize    = 2 << 20
	RootfsBlockSize = 2 << 11
)

var (
	EmptyHugePage = make([]byte, HugepageSize)
	EmptyBlock    = make([]byte, RootfsBlockSize)
)

type Slicer interface {
	Slice(off, length int64) ([]byte, error)
}

func CreateDiff(
	source io.ReaderAt,
	blockSize int64,
	dirty *bitset.BitSet,
	base Slicer,
	diff io.Writer,
) (*bitset.BitSet, *bitset.BitSet, error) {
	b := make([]byte, blockSize)

	emptyBuf := EmptyBlock
	if blockSize == HugepageSize {
		emptyBuf = EmptyHugePage
	}

	empty := bitset.New(0)

	for i, e := dirty.NextSet(0); e; i, e = dirty.NextSet(i + 1) {
		_, err := source.ReadAt(b, int64(i)*blockSize)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading from source: %w", err)
		}

		if base != nil {
			// At this moment the template should be cached locally, because it was used when starting or during running the sandbox—that's why it is dirty.
			cacheBlock, err := base.Slice(int64(i)*blockSize, blockSize)
			if err != nil {
				return nil, nil, fmt.Errorf("error reading from cache: %w", err)
			}

			// If the block is the same as in the base it is not dirty.
			if bytes.Equal(b, cacheBlock) {
				dirty.Clear(uint(i))

				continue
			}
		}

		// If the block is empty, we don't need to write it to the diff.
		// Because we checked it does not equal to the base, so we keep it separately.
		if bytes.Equal(b, emptyBuf) {
			dirty.Clear(uint(i))
			empty.Set(uint(i))

			continue
		}

		_, err = diff.Write(b)
		if err != nil {
			return nil, nil, fmt.Errorf("error writing to diff: %w", err)
		}
	}

	return dirty, empty, nil
}
