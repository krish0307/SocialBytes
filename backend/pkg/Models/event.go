package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
	Config "socialbytes.com/main/pkg/Config"
)

var db *gorm.DB

type Event struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	Name          string
	Description   string
	Location      string
	Interests     pq.StringArray `gorm:"type:text[]"`
	Date          time.Time
	ImageasBase64 string
	UserId        string
}

type SearchEventStruct struct {
	Name      string
	Location  string
	StartDate string
	EndDate   string
}
type Users struct {
	UserID   string `gorm:"primaryKey"`
	UserName string
	EmailId  string
	Password string
}

func init() {
	Config.Connect()
	db = Config.GetDB()
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Users{})
}

func (e *Event) CreateEventstable() (*Event, error) {
	if e == nil {
		error := errors.New("Event is Empty")
		return e, error
	}
	if e.Description == "" || e.Name == "" || e.Location == "" || e.ImageasBase64 == "" || len(e.Interests) <= 0 || e.Date.IsZero() {

		error := errors.New("Event details incorrect")
		return e, error
	}
	db.Create(&e)
	return e, nil
}
func GetAllEvents() []Event {
	var events []Event
	db.Find(&events)

	return events
}

func (se *SearchEventStruct) SearchEvent() []Event {
	var events []Event
	if se.Name != "" && se.Location != "" {
		name := fmt.Sprintf("%%%s%%", se.Name)
		location := fmt.Sprintf("%%%s%%", se.Location)
		db.Where("Name LIKE ? AND Location LIKE ?", name, location).Find(&events)
	} else if se.Name != "" {
		name := fmt.Sprintf("%%%s%%", se.Name)
		db.Where("Name LIKE ?", name).Find(&events)
	} else if se.Location != "" {
		location := fmt.Sprintf("%%%s%%", se.Location)
		db.Where("Location LIKE ?", location).Find(&events)
	} else if se.StartDate != "" && se.EndDate != "" {
		db.Where("Date BETWEEN ? AND ?", se.StartDate, se.EndDate).Find(&events)
	}
	return events
}
