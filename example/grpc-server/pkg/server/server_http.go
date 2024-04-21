package server

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/flowchartsman/swaggerui"
	oapi "github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) runGrpcHTTPServer(imps []Implementation) (func(ctx context.Context) error, error) {
	gwServer, err := s.initGrpcHTTPServer(imps)
	if err != nil {
		return nil, fmt.Errorf("create gateway server: %w", err)
	}

	go func() {
		log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%d", s.httpPort)

		err := gwServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Error on start listen gRPC gateway: %s", err)
			s.cancel()
		}
	}()

	return gwServer.Shutdown, nil
}

func (s *Server) initGrpcHTTPServer(imps []Implementation) (*http.Server, error) {
	// server call options
	callOpts := make([]grpc.CallOption, 0)
	callOpts = append(callOpts, s.gatewayToGrpcCallOptions...)

	// create a client connection to the gRPC server we just started. This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		s.ctx,
		fmt.Sprintf("0.0.0.0:%d", s.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(callOpts...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	// gateway mux options
	serverOpts := []runtime.ServeMuxOption{}
	serverOpts = s.addIncomingHeaderMatcher(serverOpts)
	serverOpts = s.addForwardResponseOption(serverOpts)

	gwmux := runtime.NewServeMux(serverOpts...)

	// register all implementations
	for _, i := range imps {
		err = i.RegisterHTTP(s.ctx, gwmux, conn)
		if err != nil {
			return nil, fmt.Errorf("register implementation: %w", err)
		}
	}

	// http server mux
	publicMux := http.NewServeMux()
	publicMux.Handle("/", gwmux)

	// swagger
	err = s.addSwaggerToMux(publicMux, imps)
	if err != nil {
		return nil, fmt.Errorf("add swagger: %w", err)
	}

	// pprof
	if s.enablePprof {
		profiler(publicMux)
	}

	// http middlewares
	h := s.addGatewayMiddlewares(publicMux)

	// gRPC GW HTTP server
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.httpPort),
		Handler: h,
	}

	if s.readTimeout > 0 {
		gwServer.ReadTimeout = time.Duration(s.readTimeout) * time.Second
	}
	if s.writeTimeout > 0 {
		gwServer.WriteTimeout = time.Duration(s.writeTimeout) * time.Second
	}

	return gwServer, nil
}

func (s *Server) addSwaggerToMux(mux *http.ServeMux, imps []Implementation) error {
	for _, v := range imps {
		patt := "/swagger/"
		if len(imps) > 1 {
			patt = fmt.Sprintf("%s%s", patt, v.GetName())
		}

		op, err := oapi.Embedded([]byte(v.GetSpec()), []byte(v.GetSpec()))
		if err != nil {
			return fmt.Errorf("parse spec: %w", err)
		}

		if op.Spec().Info == nil {
			op.Spec().Info = &spec.Info{}
		}

		op.Spec().Info.Version = fmt.Sprintf("%s, %s", s.versionTag, s.versionCreatedAt.Format(versionTimeFormat))

		spc, err := op.Spec().MarshalJSON()
		if err != nil {
			return fmt.Errorf("marshal spec: %w", err)
		}

		mux.Handle(patt, http.StripPrefix("/swagger", swaggerui.Handler(spc)))
	}

	return nil
}

func (s *Server) addIncomingHeaderMatcher(opts []runtime.ServeMuxOption) []runtime.ServeMuxOption {
	for _, v := range s.gatewayIncomingHeaderMatchers {
		opts = append(opts, runtime.WithIncomingHeaderMatcher(runtime.HeaderMatcherFunc(v)))
	}

	return opts
}

func (s *Server) addForwardResponseOption(opts []runtime.ServeMuxOption) []runtime.ServeMuxOption {
	for _, v := range s.gatewayForwardResponseOptions {
		opts = append(opts, runtime.WithForwardResponseOption(v))
	}

	return opts
}

func (s *Server) addGatewayMiddlewares(mux http.Handler) http.Handler {
	for _, m := range s.gatewayMiddlewares {
		mux = m(mux)
	}

	return mux
}

func profiler(mux *http.ServeMux) {
	mux.HandleFunc("/pprof/", pprof.Index)
	mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/pprof/profile", pprof.Profile)
	mux.HandleFunc("/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/pprof/trace", pprof.Trace)
	mux.HandleFunc("/vars", expVars)

	mux.Handle("/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("/pprof/mutex", pprof.Handler("mutex"))
	mux.Handle("/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/pprof/block", pprof.Handler("block"))
	mux.Handle("/pprof/allocs", pprof.Handler("allocs"))
}

// Replicated from expvar.go as not public.
func expVars(w http.ResponseWriter, r *http.Request) {
	first := true
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}
