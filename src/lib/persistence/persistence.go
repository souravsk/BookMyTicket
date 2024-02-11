package persistence

// DatabaseHandler is an interface representing methods for interacting with a database.
type DatabaseHandler interface {
	// AddEvent adds a new event to the database and returns its ID.
	AddEvent(Event) ([]byte, error)

	// FindEvent retrieves an event from the database based on its ID.
	FindEvent([]byte) (Event, error)

	// FindEventByName retrieves an event from the database based on its name.
	FindEventByName(string) (Event, error)

	// FindAllAvailableEvents retrieves all available events from the database.
	FindAllAvailableEvents() ([]Event, error)
}
