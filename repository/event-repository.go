package repository

import (
	"festApp/entity"
	"gorm.io/gorm"
)

//EventRepository is a ....
type EventRepository interface {
	InsertEvent(b entity.Event) entity.Event
	UpdateEvent(b entity.Event) entity.Event
	DeleteEvent(b entity.Event)
	AllEvent() []entity.Event
	FindEventByID(eventID uint64) entity.Event
}

type eventConnection struct {
	connection *gorm.DB
}

//NewEventRepository creates an instance EventRepository
func NewEventRepository(dbConn *gorm.DB) EventRepository {
	return &eventConnection{
		connection: dbConn,
	}
}

func (db *eventConnection) InsertEvent(b entity.Event) entity.Event {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *eventConnection) UpdateEvent(b entity.Event) entity.Event {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *eventConnection) DeleteEvent(b entity.Event) {
	db.connection.Delete(&b)
}

func (db *eventConnection) FindEventByID(eventID uint64) entity.Event {
	var event entity.Event
	db.connection.Preload("User").Find(&event, eventID)
	return event
}

func (db *eventConnection) AllEvent() []entity.Event {
	var events []entity.Event
	db.connection.Preload("User").Find(&events)
	return events
}