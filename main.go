package main

import (
	"log"
	"net/http"

	"github.com/TAMAGO-is-NOT-GOHAN/GHP-back/event"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var coors []event.Coordinate
var ngdata []event.NG
var users []event.User
var eventData event.Event

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://ghp.itok01.com/*"}
	router.Use(cors.New(config))

	Test(router)
	GetEvent(router)
	PostEvent(router)
	PostEventJoin(router)
	GetEventNgDate(router)
	PostEventJoin(router)
	GetUserLocation(router)
	PostUserLocation(router)

	router.Run()
}

func Test(r *gin.Engine) {
	r.GET("GHP/test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!?!?!?!?!?!?!?!?!?!?!")
	})
}

func GetEvent(r *gin.Engine) {
	r.GET("/v1/event", func(c *gin.Context) {
		eventData.Coors = coors

		c.JSON(http.StatusOK, eventData)
	})
}

func PostEvent(r *gin.Engine) {
	r.POST("/v1/event", func(c *gin.Context) {
		c.BindJSON(&eventData)

		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}

		uuidUI := u.String()
		eventData.ID = uuidUI

		c.JSON(http.StatusOK, `{"event_id":"`+uuidUI+`"}`)
	})
}

func PostEventJoin(r *gin.Engine) {
	type EventID struct {
		EventID string `json:"event_id"`
	}

	r.POST("/v1/event/join", func(c *gin.Context) {
		var tmp EventID
		user := c.Query("user")
		c.BindJSON(&tmp)

		users = append(users, event.User{tmp.EventID, user})

		c.JSON(http.StatusOK, `{"status": "ok"}`)
	})
}

func GetEventNgDate(r *gin.Engine) {
	r.GET("/v1/event/ngdate", func(c *gin.Context) {
		c.JSON(http.StatusOK, ngdata)
	})
}

func PostEventNgDate(r *gin.Engine) {
	var ng []event.NG

	r.POST("/v1/event/ngdate", func(c *gin.Context) {
		c.BindJSON(&ng)

		for _, e := range ng {
			ngdata = append(ngdata, e)
		}

		c.JSON(http.StatusOK, `{"status": "ok"}`)
	})
}

func GetUserLocation(r *gin.Engine) {
	r.GET("/v1/user/location", func(c *gin.Context) {
		c.JSON(http.StatusOK, coors)
	})
}

func PostUserLocation(r *gin.Engine) {
	var coor event.Coordinate

	r.POST("/v1/user/location", func(c *gin.Context) {
		c.BindJSON(&coor)
		flg := false

		if len(coors) == 0 {
			for i, e := range coors {
				if e.User == coor.User {
					coors[i].Latitude = coor.Latitude
					coors[i].Longitude = coor.Longitude
					flg = true
				}
			}
		}

		if !flg {
			coors = append(coors, coor)
		}

		c.JSON(http.StatusOK, `{"status":"ok"}`)
	})
}
