package mmgo

import "gopkg.in/mgo.v2"

type MockedCollection struct {
	*mgo.Collection
}

type Collection interface {
	Find(query interface{}) Query
	Pipe(pipeline interface{}) Pipe
	Insert(docs ...interface{}) error
	Update(selector interface{}, update interface{}) error
	DropCollection() error
}

func (mc MockedCollection) Find(query interface{}) Query {
	return &MockedQuery{mc.Collection.Find(query)}
}

func (mc MockedCollection) Insert(docs ...interface{}) error {
	return mc.Collection.Insert(docs...)
}

func (mc *MockedCollection) Update(selector interface{}, update interface{}) error {
	return mc.Collection.Update(selector, update)
}

func (mc MockedCollection) DropCollection() error {
	return mc.Collection.DropCollection()
}

func (mc MockedCollection) Pipe(pipeline interface{}) Pipe {
	return &MockedPipe{mc.Collection.Pipe(pipeline)}
}
