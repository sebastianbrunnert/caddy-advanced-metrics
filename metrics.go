package advanced_metrics

import (
	"net/http"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(AdvancedMetrics{})
	httpcaddyfile.RegisterHandlerDirective("advanced_metrics", parseAdvancedMetrics)
}

type AdvancedMetrics struct {
	PrometheusPort int
}

func (AdvancedMetrics) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.advanced_metrics",
		New: func() caddy.Module { return new(AdvancedMetrics) },
	}
}

func (AdvancedMetrics) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.NextArg() {
			return d.ArgErr()
		}
	}
	return nil
}

func (am AdvancedMetrics) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	lrw := NewLoggingResponseWriter(w)
	next.ServeHTTP(w, r)
	HandleRequest(am.PrometheusPort, r.Method, r.URL.Path, lrw.statusCode, r.Host)
	return nil
}

func getPort(d *caddyfile.Dispenser) int {
	var port int = 6611
	for d.Next() {
		for d.NextBlock(0) {
			switch d.Val() {
			case "port":
				if d.NextArg() {
					i, err := strconv.Atoi(d.Val())
					if err == nil {
						port = i
					}
				}
			}
		}
	}
	return port
}

func parseAdvancedMetrics(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	port := getPort(h.Dispenser)

	StartServer(port)

	return AdvancedMetrics{PrometheusPort: port}, nil
}

var (
	_ caddyfile.Unmarshaler       = (*AdvancedMetrics)(nil)
	_ caddyhttp.MiddlewareHandler = (*AdvancedMetrics)(nil)
)
