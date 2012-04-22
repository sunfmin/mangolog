package mangolog

import (
	"github.com/paulbellamy/mango"
	"log"
	"time"
)

// Started GET "/look/women" for 111.86.201.25 at Sun Apr 22 11:21:28 +0000 2012
// Completed 200 OK in 932ms (Views: 738.4ms | ActiveRecord: 183.9ms)

func Logger(env mango.Env, app mango.App) (status mango.Status, headers mango.Headers, body mango.Body) {
	startTime := time.Now()
	r := env.Request()
	log.Printf("Started %s \"%s\" for %s at %s\n", r.Method, r.RequestURI, r.RemoteAddr, startTime.String())
	status, headers, body = app(env)
	execution := time.Now().Sub(startTime) / time.Millisecond
	log.Printf("Completed %d in %dms\n\n", status, execution)
	return
}
