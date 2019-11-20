package mmgo

import "gopkg.in/mgo.v2"

type MockedPipe struct {
	*mgo.Pipe
}

type Pipe interface {
	All(result interface{}) error
}

func (mp MockedPipe) All(result interface{}) error {
	return mp.Pipe.All(result)
}
