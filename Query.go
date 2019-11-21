package mmgo

import "gopkg.in/mgo.v2"

type MockedQuery struct {
	*mgo.Query
}

type Query interface {
	All(result interface{}) error
	Count() (n int, err error)
	One(result interface{}) (err error)
}

func (mq MockedQuery) All(result interface{}) error {
	return mq.Query.All(result)
}

func (mq MockedQuery) Count() (n int, err error) {
	return mq.Query.Count()
}

func (mq MockedQuery) One(result interface{}) (err error) {
	return mq.Query.One(result)
}