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

	CreateEvent := &models.Event{}
	Utils.ParseBody(r, CreateEvent)
	event, err := CreateEvent.CreateEventstable()
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

// @Summary Get Events
// @Description Endpoint used to get all events from db.
// @Tags Events
// @Success 200 {object} models.Event
// @Failure 404 {object} object
// @Router /api/getEvents [get]
func GetEvents(w http.ResponseWriter, r *http.Request) {

	Utils.AddCorsHeaders(w, r)
	events := models.GetAllEvents()
	json.NewEncoder(w).Encode(events)
}

// @Summary Get Event
// @Description Endpoint used to get an events from db based on id.
// @Tags Events
// @Success 200 {object} models.Event
// @Failure 404 {object} object
// @Router /api/getEvent [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {

	Utils.AddCorsHeaders(w, r)
	var event models.Event
	if r.URL.Query().Has("id") {
		id := r.URL.Query()["id"][0]
		event = models.GetEventByID(id)
	}
	json.NewEncoder(w).Encode(event)
}
