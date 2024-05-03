package handlers

import (
	"fmt"
	"net/http"

	v1 "github.com/Victor-Acrani/go-api-starterkit/app/service/demo-api/handlers/v1"
	"github.com/Victor-Acrani/go-api-starterkit/business/web/v1/mid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Log *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *mux.Router {
	// set gorilla mux
	r := mux.NewRouter()

	// *** API ***
	// create sub router for app v1
	appV1 := r.PathPrefix("/v1").Subrouter()

	// set middlewares
	// NOTE: the order of the middlewares afect the way a request is handled.
	// NOTE: middlewares run backwards
	appV1.Use(
		mid.RequestValues,
		mid.Logger(cfg.Log),
		mid.ErrorHandler,
		mid.Recover(cfg.Log),
	)
	// set routes
	appV1.HandleFunc("/health", v1.HealthCheck).Methods(http.MethodGet)

	printRoutes(r, cfg.Log)
	return r
}

// printRoutes prints all routes assigned to the router.
func printRoutes(r *mux.Router, log *zap.SugaredLogger) {
	log.Info("API routes")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		log.Info(fmt.Sprintf("%v %s", methods, path))
		return nil
	})
}
