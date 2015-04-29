package docomo

import "golang.org/x/net/context"

type docomoKey string

func NewContext(ctx context.Context, apiKey string) context.Context {
	c := NewClient(apiKey)
	return context.WithValue(ctx, docomoKey("docomo"), c)
}

func FromContext(ctx context.Context) *Client {
	c, _ := ctx.Value(docomoKey("docomo")).(*Client)
	return c
}
