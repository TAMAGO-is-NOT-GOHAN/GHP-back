package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go"
	"github.com/TAMAGO-is-NOT-GOHAN/GHP-back/event"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func main() {
	router := gin.Default()

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "tng-ghp-ok"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	Test(router)

	router.Run()
}

func Test(r *gin.Engine) {
	r.GET("GHP/test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!?!?!?!?!?!?!?!?!?!?!")
	})
}

func GetEvent(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/v1/event", func(c *gin.Context) {
		iter := client.Collection(c.Query("event_id")).Documents(ctx)
		var eventData event.Event
		eventID, _ := strconv.Atoi(c.Query("event_id"))
		eventData.ID = uint32(eventID)

		for {
			doc, err := iter.Next()
			if err != nil {
				log.Fatal(err)
			}

			if err == iterator.Done {
				break
			}
			if err != nil {
				return
			}

			doc.DataTo(&eventData)
		}

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
