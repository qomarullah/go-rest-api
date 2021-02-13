package helpers

import (
	"net/http"
)

// RequestLogger used for httprouter request handler
type RequestLogger struct {
	Handle http.Handler
}

// ServeHTTP will server http & ouput request to log
func (rl RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	rl.Handle.ServeHTTP(w, r)
	//logger.Info("request", zap.String("Method", r.Method), zap.String("URLPath", r.URL.Path), zap.Duration("Time", time.Since(start)))
	//log.Printf("%s%s %s in %v", rl.Logger.Prefix(), r.Method, r.URL.Path, time.Since(start))
	//logger.Infow("Request", "Method", r.Method, "URLPath", r.URL.Path, "ProcTime", time.Since(start))

}
