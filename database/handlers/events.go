package handlers

import (
	"time"
	"gorm.io/gorm"
	"database/models"
)

func createEvent(db *gorm.DB, title, description string, startAt, endAt time.Time, location string, online bool) uint { // create new event
  event := models.Event{Title: title, Description: description, StartsAt: startAt, EndsAt: endAt, Location: location, Online: online}
  db.Select("ID","Title", "Description", "StartsAt", "EndsAt", "Location", "Online").Create(&event)
  return event.ID
}

func deleteEvent(db *gorm.DB, eventid uint){ // remove task = works
  db.Delete(&models.Event{}, eventid)
}

func addSubtaskToEvent() {

}