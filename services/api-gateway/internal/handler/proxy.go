package handler

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyHandler struct {
	services map[string]*url.URL
}

func NewProxyHandler(services map[string]*url.URL) *ProxyHandler {
	return &ProxyHandler{services: services}
}

func (h *ProxyHandler) Proxy(serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target, ok := h.services[serviceName]
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"error":"Service not available"}`))
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host

			// Forward auth headers
			if auth := r.Header.Get("Authorization"); auth != "" {
				req.Header.Set("Authorization", auth)
			}
		}

		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"error":"Service unavailable"}`))
		}

		proxy.ServeHTTP(w, r)
	}
}

func (h *ProxyHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := make(map[string]string)
	for name, u := range h.services {
		resp, err := http.Get(u.String() + "/health")
		if err != nil || resp.StatusCode != http.StatusOK {
			status[name] = "unhealthy"
		} else {
			status[name] = "healthy"
		}
		if resp != nil {
			resp.Body.Close()
		}
	}

	io.WriteString(w, `{"status":"healthy","services":`+mapToJSON(status)+`}`)
}

func mapToJSON(m map[string]string) string {
	pairs := make([]string, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, `"`+k+`":"`+v+`"`)
	}
	return `{` + strings.Join(pairs, ",") + `}`
}
