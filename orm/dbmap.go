package orm

import (
	"github.com/labstack/echo"
	"net/http"
	"sync"
)

var text = make(map[string][]string)

var DbmapServer = &dBMap{
	usermap: make(map[string]string),
}

type dBMap struct {
	mutex   sync.RWMutex
	usermap map[string]string
}

func (db *dBMap) Get(user string) (string, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	bool := false
	for name, _ := range db.usermap {
		if user == name {
			bool = true
			return db.usermap[user], bool
		}
	}
	return "", bool
}

func (db *dBMap) Creater(user, pass string) bool {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	db.usermap[user] = pass

	return true
}

func Add(c echo.Context, a string, b string) error {
	text[a] = append(text[a], b)
	return c.Render(http.StatusOK, "show.gtpl", text[a])
}
