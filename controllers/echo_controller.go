package controllers

import (
	"encoding/json"
	"net/http"

	"io/ioutil"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/web-api/models"
)

// EchoController echos.
type EchoController struct {
	BaseController
}

// Echo echo.
func (c *EchoController) Echo(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	url := "http://slowwly.robertomurray.co.uk/delay/1000/url/https://reqres.in/api/users?page=2"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bodyBytes)
	return nil // no error
}

// GetEchos from database.
// Echo echos.
// @Summary Get echo messages.
// @Description Get echos from the database.
// @Produce  json
// @Success 200 {object} models.Echo
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /echos [get]
func (c *EchoController) GetEchos(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	rows, _ := c.DB.Query("SELECT * FROM ECHO")
	var id, message string
	var echos []models.Echo = []models.Echo{}

	for rows.Next() {
		err := rows.Scan(&id, &message)
		if err != nil { /* error handling */
			log.Println(" Error while mapping row: ", err)
		}
		echos = append(echos, models.Echo{ID: id, Message: message})
	}

	json, _ := json.Marshal(echos)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(json)
	return nil // no error
}
