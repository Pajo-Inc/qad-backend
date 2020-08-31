package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "db", // endpoint without http
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)

	r := gin.Default()
	m := melody.New()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	r.Run(":5000")
}
