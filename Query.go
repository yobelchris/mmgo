package mmgo

import "gopkg.in/mgo.v2"

type MockedQuery struct {
	*mgo.Query
}

type Query interface {
	All(result interface{}) error
}

func (mq *MockedQuery) All(result interface{}) error {
	return mq.Query.All(result)
}
