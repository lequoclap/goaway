package main

import (
	"log" // System

	"./routers" // Local

	"github.com/urfave/negroni" // Third Party
)

func main() {
	router := routers.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Listening:")
	n.Run(":3001")
}
