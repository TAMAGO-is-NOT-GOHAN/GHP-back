package main

import (
	"net/http"
	"strconv"

	"github.com/TAMAGO-is-NOT-GOHAN/GHP-back/event"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	Test(router)

	router.Run()
}

func Test(r *gin.Engine) {
	r.GET("GHP/test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!?!?!?!?!?!?!?!?!?!?!")
	})
}

func GetEvent(r *gin.Engine) {
	r.GET("/v1/event", func(c *gin.Context) {
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEvent(r *gin.Engine) {

}

func PutEvent(r *gin.Engine) {

}

func PostEventJoin(r *gin.Engine) {

}

func GetEventNgDate(r *gin.Engine) {
	r.GET("/v1/event/ngdate", func(c *gin.Context) {
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEventNgDate(r *gin.Engine) {

}

func GetEventDeparture(r *gin.Engine) {
	r.GET("/v1/event/departure", func(c *gin.Context) {
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		c.JSON(http.StatusOK, eventData)
	})
}

func GetEventRoute(r *gin.Engine) {
	r.GET("/v1/event/route", func(c *gin.Context) {
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEventArrival(r *gin.Engine) {

}

func GetEventArrivalRank(r *gin.Engine) {
	r.GET("/v1/event/arrival/rank", func(c *gin.Context) {
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		c.JSON(http.StatusOK, eventData)
	})
}

func GetUserLocation(r *gin.Engine) {
	r.GET("/v1/user/location", func(c *gin.Context) {
		user, _ := strconv.Atoi(c.Query("user"))

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
}

func PostUserLocation(r *gin.Engine) {

}
