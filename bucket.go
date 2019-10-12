package bucket

import (
	"context"
	"fmt"
)

// Bucket interface
type Bucket interface {
	Wait(context.Context) error
}

// New returns a Bucket interface
func New(rate float64, depth int) Bucket {
	return &bucket{
		rate:  rate,
		depth: depth,
	}
}

type bucket struct {
	rate  float64
	depth int
}

func (b *bucket) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("wait error: %w", ctx.Err())
	default:
	}
	return nil
}
