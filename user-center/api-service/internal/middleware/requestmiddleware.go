package middleware

import (
	"context"
	"coupon-platform/common/util/system"
	"net/http"
)

type RequestMiddleware struct {
}

type Request struct {
	*http.Request
	ClientIP  string
	RouteAddr string
}

func NewRequestMiddleware() *RequestMiddleware {
	return &RequestMiddleware{}
}

func (m *RequestMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		ip := system.ClientIP(r)
		routeAddr := system.GetRouteAddr(r)
		ctx := r.Context()

		ctx = context.WithValue(ctx, "x-request", &Request{
			Request:   r,
			ClientIP:  ip,
			RouteAddr: routeAddr,
		})
		// Passthrough to next handler if need
		next(w, r.WithContext(ctx))
	}
}

func GetRequest(ctx context.Context) *Request {
	return ctx.Value("x-request").(*Request)
}
