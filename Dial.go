package mmgo

import "gopkg.in/mgo.v2"

type MockedSession struct {
	*mgo.Session
}

type Session interface {
	DB(name string) Database
	Close()
	Copy() Session
}

func Dial(url string) (*Session, error) {
	ses, err := mgo.Dial(url)

	if err != nil {
		return nil, err
	}

	session := MockedSession{ses}

	var sessionIn Session

	sessionIn = &session

	return &sessionIn, nil
}

func IsDup(err error) bool {
	return mgo.IsDup(err)
}

func (ms *MockedSession) DB(name string) Database {
	return &MockedDatabase{ms.Session.DB(name)}
}

func (ms *MockedSession) Close() {
	ms.Session.Close()
}

func (ms *MockedSession) Copy() Session {
	return &MockedSession{ms.Session.Copy()}
}
