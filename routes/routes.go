package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/web-api/config"
	"github.com/saiprasadkrishnamurthy/web-api/controllers"
)

const apiBase = "/api"

// InitialiseAllRoutes initialises all routes in the API.
func InitialiseAllRoutes(r *httprouter.Router) {
	apiContext := apiBase + "/" + config.APIVersion()
	appController := controllers.BaseController{DB: config.GetDB(), NatsConn: config.GetNats()}

	// List all your controllers here.
	echoController(apiContext, r, appController)
	pubSubController(apiContext, r, appController)
}

func echoController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.EchoController{BaseController: baseController}
	r.GET(apiContext, c.Action(c.Echo))
	r.GET(apiContext+"/echos", c.Action(c.GetEchos))
}

func pubSubController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.PubSubController{BaseController: baseController}
	r.GET(apiContext+"/pub", c.Action(c.Pub))
}
