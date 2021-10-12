package repositories

import (
	"errors"
	"github.com/XVRomanova/wb-tech-2/develop/dev11/internal/app/models"

	"time"
)

type Repository struct {
	store map[string][]models.Event
}

func New() *Repository {
	store := make(map[string][]models.Event, 0)

	repo := Repository{
		store: store,
	}

	return &repo
}

func (repo *Repository) AddEvent(event models.Event) error {

	if _, found := repo.store[event.Date]; !found {
		repo.store[event.Date] = append(repo.store[event.Date], event)
	} else {
		for _, value := range repo.store[event.Date] {
			if value == event {
				return errors.New("this event has already been added")
			}
		}
		repo.store[event.Date] = append(repo.store[event.Date], event)
	}
	return nil
}

func (repo *Repository) UpdateEvent(oldEventName string, event models.Event) error {

	for _, value := range repo.store[event.Date] {
		if value.UserID == event.UserID && value.NameEvent == oldEventName {
			value = event
			return nil
		}
	}

	return errors.New("this event already exists")
}

func (repo *Repository) DeleteEvent(event models.Event) {
	events := repo.store[event.Date]
	for i, value := range events {
		if value == event {
			events[i] = events[len(events)-1]
			events[len(events)-1] = models.Event{}
			events = events[:len(events)-1]
			return
		}
	}
}

func (repo *Repository) GetEventsForDay(date string) ([]models.Event, error) {
	if _, found := repo.store[date]; found {
		return repo.store[date], nil
	}
	return nil, errors.New("the day wasn't found")
}

func (repo *Repository) GetEventsForWeek(date string) ([]models.Event, error) {
	myDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	day := myDate.Day()

	var events []models.Event
	for key, value := range repo.store {
		currentDate, err := time.Parse("2006-01-02", key)
		if err != nil {
			return nil, err
		}
		currentDay := currentDate.Day()

		if currentDay >= day && currentDay < day+7 {
			events = append(events, value...)
		}
	}

	return events, nil
}

func (repo *Repository) GetEventsForMonth() []models.Event {
	var events []models.Event
	for _, value := range repo.store {
		events = append(events, value...)
	}

	return events
}
