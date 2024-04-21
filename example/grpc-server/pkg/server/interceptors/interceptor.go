package interceptors

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	RequestIDKey = "x-request-id"
)

func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (_ interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("no metadata")
		}

		if vals := md.Get(RequestIDKey); len(vals) > 0 {
			ctx = context.WithValue(ctx, RequestIDKey, vals[0])
		} else {
			hasher := md5.New()
			_, _ = fmt.Fprint(hasher, []byte(fmt.Sprint(rand.Intn(1000))))

			ctx = context.WithValue(
				ctx,
				RequestIDKey,
				hex.EncodeToString(hasher.Sum(nil)),
			)
		}

		resp, err := handler(ctx, req)

		return resp, err
	}
}
