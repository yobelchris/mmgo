package mmgo

import "gopkg.in/mgo.v2"

type MockedDatabase struct {
	*mgo.Database
}

type Database interface {
	C(name string) Collection
}

func (md MockedDatabase) C(name string) Collection {
	return &MockedCollection{md.Database.C(name)}
}
