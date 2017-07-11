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

	return db.usermap[user]
}
