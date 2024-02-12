package rest

import (
	"encoding/hex"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/souravsk/BookMyTicket/src/lib/persistence"
)

type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

func NewEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler: databasehandler,
	}
}

func (eh *eventServiceHandler) FindEventHandler(c *gin.Context) {
	criteria := c.Param("SearchCriteria")
	searchkey := c.Param("search")

	if criteria == "" || searchkey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search criteria or search key"})
		return
	}

	var event persistence.Event
	var err error
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
		log.Printf("Error finding event: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (eh *eventServiceHandler) AllEventHandler(c *gin.Context) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		log.Printf("Error finding all events: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding all events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (eh *eventServiceHandler) NewEventHandler(c *gin.Context) {
	var event persistence.Event
	if err := c.BindJSON(&event); err != nil {
		log.Printf("Error decoding event data: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding event data"})
		return
	}

	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		log.Printf("Error persisting event: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error persisting event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}
