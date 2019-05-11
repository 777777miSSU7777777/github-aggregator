package foo

import "context"

type FooService interface {
	Bar(ctx context.Context, i int, s string) (string, error)
}
