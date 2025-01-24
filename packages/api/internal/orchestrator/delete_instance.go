package orchestrator

import (
	"context"
)

func (o *Orchestrator) DeleteInstance(ctx context.Context, sandboxID string) bool {
	_, childSpan := o.tracer.Start(ctx, "delete-instance")
	defer childSpan.End()

	sbx, err := o.instanceCache.Get(sandboxID)
	if err != nil {
		return false
	}

	o.dns.Remove(sandboxID, sbx.Value().InternalID)
	return o.instanceCache.Kill(sandboxID)
}
