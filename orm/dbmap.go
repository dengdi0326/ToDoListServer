package orm

import (
	"sync"
)

var DbmapServer = &dBMap{
	usermap:    make(map[string]string),
}

type dBMap struct {
	mutex       sync.RWMutex
	usermap     map[string]string
}

func (db *dBMap) Get(user string) (string, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	bool := false

	for name,_ := range db.usermap{
		if name == user{
			bool = true
			break
		}

		return db.usermap[user],bool
	}

	return "",bool
}

func (db *dBMap) Creater(user,pass string)bool{
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	db.usermap[user] = pass

	return true
}
