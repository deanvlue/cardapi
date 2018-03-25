package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
	"net"
	"crypto/tls"
	"github.com/deanvlue/cardapi/main"
)

type CardsDAO struct {
	Server string
	Database string
	Password string
}

var db *mgo.Database

const (
	COLLECTION = "colcardart"
)


//Connect : connects to a mongo database
func (m *CardsDAO) Connect(){
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.documents.azure.com:10255", m.Database)}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: m.Database, // It can be anything
		Username: m.Database, // Username
		Password: m.Password, // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	session, err := mgo.DialWithInfo(dialInfo) 
	db = session.DB(m.Database)
}

func (m *CardsDAO) FindById(id string) (Card, error){

}