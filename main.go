package main

import (
	"log"
	"net/http"

	"github.com/saiprasadkrishnamurthy/web-api/listeners"
	"github.com/saiprasadkrishnamurthy/web-api/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/web-api/config"
)

var r *httprouter.Router

func init() {
	r = httprouter.New()
	config.InitConfigs()
	routes.InitialiseAllRoutes(r)
	listeners.InitializeAllListeners(config.GetNats())
}

func main() {
	port := config.GetPort()
	log.Println(" Starting the Server on port: ", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalln(err)
	}
}
