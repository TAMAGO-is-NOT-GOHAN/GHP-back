package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/TAMAGO-is-NOT-GOHAN/GHP-back/event"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://ghp.itok01.com/*"},
	}))

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
	GetEvent(router, app)
	PostEvent(router, app)
	PostEventJoin(router, app)
	GetEventNgDate(router, app)
	PostEventJoin(router, app)
	GetEventDeparture(router)
	GetEventRoute(router)
	PostEventArrival(router)
	GetEventArrivalRank(router)
	GetUserLocation(router, app)
	PostUserLocation(router, app)

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
	defer client.Close()

	r.GET("/v1/event", func(c *gin.Context) {
		iter := client.Collection(c.Query("event_id")).Documents(ctx)
		var eventData event.Event
		eventData.ID = c.Query("event_id")

		for {
			doc, err := iter.Next()

			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			doc.DataTo(&eventData)
		}

		iter = client.Collection("coor").Documents(ctx)

		var coors []event.Coordinate

		for {
			doc, err := iter.Next()

			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			var tmp event.Coordinate

			doc.DataTo(&tmp)

			coors = append(coors, tmp)
		}

		eventData.Coors = coors

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEvent(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var eventData event.Event

	r.POST("/v1/event", func(c *gin.Context) {
		c.BindJSON(&eventData)

		eventID := string(eventData.ID)

		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}

		uuidUI := u.String()

		_, _, err = client.Collection(eventID).Add(ctx, map[string]interface{}{
			"ID":          uuidUI,
			"Date":        eventData.Date,
			"Name":        eventData.Name,
			"MaxPeople":   eventData.MaxPeople,
			"Description": eventData.Description,
		})
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, `{"event_id":"`+uuidUI+`"}`)
	})
}

func PostEventJoin(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	type EventID struct {
		EventID string `json:"event_id"`
	}

	r.POST("/v1/event/join", func(c *gin.Context) {
		var tmp EventID
		user := c.Query("user")
		c.BindJSON(&tmp)

		_, _, err := client.Collection(user).Add(ctx, map[string]interface{}{
			"ID": tmp.EventID,
		})
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, `{"status": "ok"}`)
	})
}

func GetEventNgDate(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var ng []event.NG

	r.GET("/v1/event/ngdate", func(c *gin.Context) {
		iter := client.Collection("ng").Documents(ctx)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			var tmp event.NG

			doc.DataTo(&tmp)
			ng = append(ng, tmp)
		}

		c.JSON(http.StatusOK, ng)
	})
}

func PostEventNgDate(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var ng []event.NG

	r.POST("/v1/event/ngdate", func(c *gin.Context) {
		c.BindJSON(&ng)

		for _, e := range ng {
			_, _, err := client.Collection("ng").Add(ctx, map[string]interface{}{
				"date": e,
			})
			if err != nil {
				log.Fatal(err)
			}
		}

		c.JSON(http.StatusOK, `{"status": "ok"}`)
	})
}

func GetEventDeparture(r *gin.Engine) {
	r.GET("/v1/event/departure", func(c *gin.Context) {
		var eventData event.Event
		eventData.ID = c.Query("event_id")

		c.JSON(http.StatusOK, eventData)
	})
}

func GetEventRoute(r *gin.Engine) {
	r.GET("/v1/event/route", func(c *gin.Context) {
		var eventData event.Event
		eventData.ID = c.Query("event_id")

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEventArrival(r *gin.Engine) {
	type EventID struct {
		EventID string `json:"event_id"`
	}
	r.POST("/v1/event/departure", func(c *gin.Context) {
		var tmp EventID
		c.BindJSON(&tmp)
		c.JSON(http.StatusOK, tmp)
	})
}

func GetEventArrivalRank(r *gin.Engine) {
	r.GET("/v1/event/arrival/rank", func(c *gin.Context) {
		var eventData event.Event
		eventData.ID = c.Query("event_id")

		c.JSON(http.StatusOK, eventData)
	})
}

func GetUserLocation(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var coors []event.Coordinate

	r.GET("/v1/user/location", func(c *gin.Context) {
		iter := client.Collection("coor").Documents(ctx)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return
			}

			var tmp event.Coordinate

			doc.DataTo(&tmp)

			coors = append(coors, tmp)
		}

		c.JSON(http.StatusOK, coors)
	})
}

func PostUserLocation(r *gin.Engine, app *firebase.App) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var coor event.Coordinate

	r.POST("/v1/user/location", func(c *gin.Context) {
		c.BindJSON(&coor)

		_, err := client.Collection("coor").Doc(coor.User).Set(ctx, map[string]interface{}{
			"user":      coor.User,
			"latitude":  coor.Latitude,
			"longitude": coor.Longitude,
		})

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, `{"status":"ok"}`)
	})
}

func PutEvent(r *gin.Engine) {

}
