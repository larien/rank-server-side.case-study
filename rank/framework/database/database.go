package database

import (
	"log"

	"github.com/juju/mgosession"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	mgo "gopkg.in/mgo.v2"
)

// Repository sets up database communication for Repository layer.
func Repository() *repository.MongoConn {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer mPool.Close()

	return repository.NewMongoConnection(mPool, config.MONGODB_DATABASE)
	// review := controller.NewReviewController(mongo)
}
