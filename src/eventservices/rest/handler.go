package rest

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/souravsk/BookMyTicket/src/lib/persistence"
)

// eventServiceHandler: Defines a struct representing an event service handler.
type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
	//dbhandler: A field of type persistence.DatabaseHandler used to interact with the database.
}

// NewEventHandler: A function that creates and returns a new instance of eventServiceHandler.
// It takes a databasehandler parameter and initializes the dbhandler field.
func NewEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

// FindEventHandler: A function to handle HTTP requests for finding an event based on certain criteria.
// Uses gin.Context for request and response handling.
func (eh *eventServiceHandler) FindEventHandler(c *gin.Context) {
	criteria := c.Param("SearchCriteria")
	searchkey := c.Param("search")

	if criteria == "" || searchkey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No search criteria or search key found"})
		return
	}

	var event persistence.Event
	var err error

	// Initialize event to an empty Event struct to avoid nil pointer dereference
	event = persistence.Event{}

	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// AllEventHandler: A function to handle HTTP requests for retrieving all available events.
func (eh *eventServiceHandler) AllEventHandler(c *gin.Context) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error occurred while trying to find all available events %s", err)})
		return
	}

	c.JSON(http.StatusOK, events)
}

// NewEventHandler: A function to handle HTTP requests for creating a new event.
func (eh *eventServiceHandler) NewEventHandler(c *gin.Context) {
	var event persistence.Event
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error occurred while decoding event data %s", err)})
		return
	}

	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error occurred while persisting event %d %s", id, err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}
