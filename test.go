package main

import (
	"time"

	"github.com/axiomzen/beanstalks-api/data"
	"github.com/axiomzen/beanstalks-api/model"
	"github.com/sirupsen/logrus"
)

// Run test
func test(db *data.DAL) {
	// Give the server some time to start
	time.Sleep(time.Second * 5)
	logrus.Info("Running tests...")
	runTests(db)
	logrus.Info("tests completed")
}

func runTests(db *data.DAL) {
	// Add a user to the DB directly
	user := &model.User{
		Name:           "Bruno",
		Email:          "bruno.bachmann@dapperlabs.com",
		HashedPassword: "blablabla",
		Tags:           []string{"Back-end", "Engineering"},
	}

	if err := db.CreateUser(user); err != nil {
		logrus.Fatal(err)
	}
}
