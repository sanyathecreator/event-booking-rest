package models

import "time"

// Event represents a bookable event created by a user.
// Fields tagged `binding:"required"` are enforced when parsing JSON request bodies.
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       // ID of the user who created the event
}

var events = []Event{}

// Save appends the event to the in-memory slice.
// TODO: replace with a real database insert
func (e Event) Save() {
	events = append(events, e)
}

// GetAllEvents returns every event currently in the store.
func GetAllEvents() []Event {
	return events
}
