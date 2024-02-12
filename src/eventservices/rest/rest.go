package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/souravsk/BookMyTicket/src/lib/persistence"
)

func ServeAPI(endpoint, tlsendpoint string, databasehandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := NewEventHandler(databasehandler)
	router := gin.Default()

	// Define routes
	eventsGroup := router.Group("/events")
	{
		eventsGroup.GET("/:SearchCriteria/:search", handler.FindEventHandler)
		eventsGroup.GET("", handler.AllEventHandler)
		eventsGroup.POST("", handler.NewEventHandler)
	}

	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	go func() {
		httptlsErrChan <- router.RunTLS(tlsendpoint, "cert.pem", "key.pem")
	}()

	go func() {
		httpErrChan <- router.Run(endpoint)
	}()

	return httpErrChan, httptlsErrChan
}
