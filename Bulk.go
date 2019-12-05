package mmgo

import "gopkg.in/mgo.v2"

type MockedBulk struct {
	*mgo.Bulk
}

type Bulk interface {
	Run() (*mgo.BulkResult, error)
	Upsert(pairs ...interface{})
}

func (mb MockedBulk) Upsert(pairs ...interface{}) {
	mb.Bulk.Upsert(pairs...)
}

func (mb MockedBulk) Run() (*mgo.BulkResult, error) {
	return mb.Bulk.Run()
}