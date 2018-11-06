package main

import (
	"log"

	"github.com/juju/mgosession"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/controller"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/routing"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	mgo "gopkg.in/mgo.v2"
)

const port = ":8080"

func main() {
	Rank()
}

// Rank starts the routine for Rank's app.
func Rank() {

	// repo := database.Repository()

	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer mPool.Close()

	repo := repository.NewMongoConnection(mPool, config.MONGODB_DATABASE)
	cont := controller.NewReviewController(repo)
	router := routing.Router(cont)

	router.Run(port)
}
