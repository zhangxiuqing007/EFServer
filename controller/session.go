package controller

import (
	"EFServer/forum"
	"net/http"
	"time"

	"github.com/satori/uuid"
)

const cookie string = "userKey"

//SessionDict currentSessions
var SessionDict = make(map[string]*Session, 20)

//Session Session
type Session struct {
	UUID            string
	CreatedTime     int64
	LastRequestTime int64
	User            *forum.User
}

func createNewSession() *Session {
	session := new(Session)
	uid, _ := uuid.NewV4()
	session.UUID = uid.String()
	session.CreatedTime = time.Now().UnixNano()
	session.LastRequestTime = session.CreatedTime
	session.User = nil
	SessionDict[session.UUID] = session
	return session
}

func getExsitSession(r *http.Request) (session *Session) {
	cook, err := r.Cookie(cookie)
	if err != nil {
		return nil
	}
	v, ok := SessionDict[cook.Value]
	if !ok {
		return nil
	}
	return v
}

func getExsitOrCreateNewSession(w http.ResponseWriter, r *http.Request) *Session {
	session := getExsitSession(r)
	if session == nil {
		session = createNewSession()
		cook := &http.Cookie{
			Name:     cookie,
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, cook)
	}
	return session
}
