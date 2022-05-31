package datastore

import (
	"github.com/sirupsen/logrus"
	"github.com/upper/db/v4/adapter/mysql"
)

var (
	Datastore *DatastoreT
)

func init() {
	Datastore = new(DatastoreT)

}
func New(settings mysql.ConnectionURL) {
	sess, err := mysql.Open(settings)
	if err != nil {
		logrus.Error("Error in New.", err)
	}
	//sess.SetLogging(false)
	sess.SetMaxOpenConns(1000)

	Datastore = &DatastoreT{Session: sess}
}
