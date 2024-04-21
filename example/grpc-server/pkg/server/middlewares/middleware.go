package middlewares

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/Planck1858/golang101/examples/grpc-server/pkg/server"
)

const (
	RequestIDKey  = "x-request-id"
	XRequestIDKey = "X-Request-Id"
)

type requestIDMiddleware struct {
	handler http.Handler
}

func NewRequestIDMiddleware() server.Handler {
	return func(h http.Handler) http.Handler {
		return &requestIDMiddleware{
			handler: h,
		}
	}
}

func (m *requestIDMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqID := r.Header.Get(XRequestIDKey)
	if reqID == "" {
		hasher := md5.New()
		_, _ = fmt.Fprint(hasher, []byte(fmt.Sprint(rand.Intn(1000))))
		reqID = hex.EncodeToString(hasher.Sum(nil))
		r.Header.Set(XRequestIDKey, reqID)
	}

	ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
	r = r.WithContext(ctx)

	m.handler.ServeHTTP(w, r)
}
