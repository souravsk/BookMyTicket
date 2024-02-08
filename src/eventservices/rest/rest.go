package rest

import "github.com/gin-gonic/gin"

func ServeAPI(endpoint string, dbhandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databasehandler) // Create a new event handler
	r := gin.Default()                          // Create a new default gin router

	eventsGroup := r.Group("/events") // Create a new group of routes
	{
		eventsGroup.GET("/:SearchCriteria/:search", handler.FindEventHandler)
		eventsGroup.GET("", handler.AllEventHandler)
		eventsGroup.POST("", handler.AddEventHandler)
	}
	return r.Run(endpoint) // Listen and serve on the specified endpoint
}
