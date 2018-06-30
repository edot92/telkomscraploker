package db

import (
	"fmt"
	"time"

	// mgo "gopkg.in/mgo.v2"
	mgo "github.com/globalsign/mgo"
)

var SessionMgo *mgo.Session
var DBNAME string
var ColFb = "colfb"

// mongo ds161529.mlab.com:61529/goiotprod -u <dbuser> -p <dbpassword>
func InitDB() (*mgo.Session, error) {
initMongok:
	MongoDBHosts := "127.0.0.1:27017"
	AuthDatabase := "dblowongan"
	AuthUserName := ""
	AuthPassword := ""
	DBNAME = AuthDatabase
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  20 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Println(err)
		goto initMongok
	}
	fmt.Println("mongodb connected to " + MongoDBHosts + " DB " + AuthDatabase)
	// see https://github.com/go-mgo/mgo/issues/213

	// Optional. Switch the session to a monotonic behavior.
	mongoSession.SetMode(mgo.Monotonic, true)
	// mongoSession.SetMode(mgo.Strong, true)
	// isDebug, _ := beego.AppConfig.Bool("EnableDebugMongo")
	// if isDebug {
	// 	mgo.SetDebug(true)
	// 	var aLogger *log.Logger
	// 	aLogger = log.New(os.Stderr, "", log.LstdFlags)
	// 	mgo.SetLogger(aLogger)
	// }

	SessionMgo = mongoSession
	return mongoSession, nil
	// fmt.Println("Phone:", result.Phone)
}

func DBCopyMGO() *mgo.Session {
	return SessionMgo.Clone()
}
