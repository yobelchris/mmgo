package mmgo

import "gopkg.in/mgo.v2"

type MockedCollection struct {
	*mgo.Collection
}

type Collection interface {
	Find(query interface{}) Query
	Insert(docs ...interface{}) error
	DropCollection() error
}

func (mc MockedCollection) Find(query interface{}) Query {
	return &MockedQuery{mc.Collection.Find(query)}
}

func (mc MockedCollection) Insert(docs ...interface{}) error {
	return mc.Collection.Insert(docs...)
}

func (mc MockedCollection) DropCollection() error {
	return mc.Collection.DropCollection()
}
