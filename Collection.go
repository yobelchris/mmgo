package mmgo

import "gopkg.in/mgo.v2"

type MockedCollection struct {
	*mgo.Collection
}

type Collection interface {
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Pipe(pipeline interface{}) Pipe
	Insert(docs ...interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Remove(selector interface{}) error
	DropCollection() error
}

func (mc MockedCollection) Find(query interface{}) Query {
	return &MockedQuery{mc.Collection.Find(query)}
}

func (mc MockedCollection) Insert(docs ...interface{}) error {
	return mc.Collection.Insert(docs...)
}

func (mc MockedCollection) Update(selector interface{}, update interface{}) error {
	return mc.Collection.Update(selector, update)
}

func (mc MockedCollection) DropCollection() error {
	return mc.Collection.DropCollection()
}

func (mc MockedCollection) Pipe(pipeline interface{}) Pipe {
	return &MockedPipe{mc.Collection.Pipe(pipeline)}
}

func (mc MockedCollection) Remove(selector interface{}) error {
	return mc.Collection.Remove(selector)
}

func (mc MockedCollection) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return mc.Collection.UpdateAll(selector, update)
}

func (mc MockedCollection) FindId(id interface{}) Query {
	return &MockedQuery{mc.Collection.FindId(id)}
}

func (mc MockedCollection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return mc.Collection.Upsert(selector, update)
}