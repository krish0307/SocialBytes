package Controllers

import (
	"encoding/json"
	"net/http"

	models "socialbytes.com/main/pkg/Models"
	"socialbytes.com/main/pkg/Utils"
)

var NewEvent models.Event

// @Summary Create Event
// @Description Endpoint used to create an entry of the event in db.
// @Tags Events
// @Success 200 {object} models.Event
// @Failure 404 {object} object
// @Router /api/createEvent [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {

	Utils.AddCorsHeaders(w, r)
	username, err := GetUserName(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.Write(response)
	}
	CreateEvent := &models.Event{}
	Utils.ParseBody(r, CreateEvent)
	event, err := CreateEvent.CreateEventstable(username.FirstName + " " + username.LastName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(CreateEvent)
		w.Write(response)
	} else {

		response, _ := json.Marshal(event)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}
