package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PubSubController pub sub examples.
type PubSubController struct {
	BaseController
}

// Pub publishes something.
func (c *PubSubController) Pub(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	doc := map[string]string{"Name": "Sai"}
	json, _ := json.Marshal(doc)
	err := c.NatsConn.Publish("foo", json)
	if err != nil {
		log.Println(err)
	}
	log.Println("Published message on subject foo")

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
	return nil // no error
}
