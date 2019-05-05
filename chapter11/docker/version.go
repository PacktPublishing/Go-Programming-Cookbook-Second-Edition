package docker

import (
	"encoding/json"
	"net/http"
	"time"
)

// VersionInfo holds artifacts passed in
// at build time
type VersionInfo struct {
	Version   string
	BuildDate time.Time
	Uptime    time.Duration
}

// VersionHandler writes the latest version info
func VersionHandler(v *VersionInfo) http.HandlerFunc {
	t := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		v.Uptime = time.Since(t)
		vers, err := json.Marshal(v)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(vers)
	}
}
