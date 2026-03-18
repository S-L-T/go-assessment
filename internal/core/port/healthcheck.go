package port

import "context"

type Healthcheck interface {
	IsAlive(ctx context.Context) bool
}
