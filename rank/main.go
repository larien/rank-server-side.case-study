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

const port = ":8899"

func main() {
	Rank()
}

// Rank starts the routine for Rank's app.
func Rank() {

	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	log.Printf("Session with MongoDB host started")

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()
	log.Printf("Pool with MongoDB session set")

	repo := repository.New(pool, config.MONGODB_DATABASE)
	log.Printf("Repository layer created")

	if err := repo.InsertCategories(); err != nil {
		log.Fatal(err.Error())
	}

	controllers := controller.New(repo)
	log.Printf("Controller layer created")

	router := routing.Router(controllers)
	log.Printf("Routing endpoints set")

	router.Run(port)
	log.Printf("Running router on port %s", port)
}
