package service

import (
	"festApp/dto"
	"festApp/entity"
	"festApp/repository"
	"github.com/mashingan/smapping"
	"log"
)

//EventService is a ....
type EventService interface {
	Insert(b dto.EventCreateDTO) entity.Event
	Update(b dto.EventUpdateDTO) entity.Event
	Delete(b entity.Event)
	All() []entity.Event
	FindByID(eventID uint64) entity.Event
	//IsAllowedToEdit(userID string, eventID uint64) bool
}

type eventService struct {
	eventRepository repository.EventRepository
}

//NewEventService .....
func NewEventService(eventRepo repository.EventRepository) EventService {
	return &eventService{
		eventRepository: eventRepo,
	}
}

func (service *eventService) Insert(b dto.EventCreateDTO) entity.Event {
	event := entity.Event{}
	err := smapping.FillStruct(&event, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.eventRepository.InsertEvent(event)
	return res
}

func (service *eventService) Update(b dto.EventUpdateDTO) entity.Event {
	event := entity.Event{}

	err := smapping.FillStruct(&event, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.eventRepository.UpdateEvent(event)
	return res
}

func (service *eventService) Delete(b entity.Event) {
	service.eventRepository.DeleteEvent(b)
}

func (service *eventService) All() []entity.Event {
	return service.eventRepository.AllEvent()
}

func (service *eventService) FindByID(eventID uint64) entity.Event {
	return service.eventRepository.FindEventByID(eventID)
}

