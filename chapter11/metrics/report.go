package metrics

import (
	"net/http"

	gometrics "github.com/rcrowley/go-metrics"
)

// ReportHandler will emit the current metrics in json format
func ReportHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	t := gometrics.GetOrRegisterTimer("reporthandler.writemetrics", nil)
	t.Time(func() {
		gometrics.WriteJSONOnce(gometrics.DefaultRegistry, w)
	})
}
