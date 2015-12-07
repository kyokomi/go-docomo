// +build appengine

package docomo

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
	"github.com/kyokomi/go-docomo/internal"
)

func init() {
	internal.RegisterContextClientFunc(contextClientAppEngine)
}

func contextClientAppEngine(ctx context.Context) (*http.Client, error) {
	return urlfetch.Client(ctx), nil
}
