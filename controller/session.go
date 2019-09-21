package controller

import (
	"EFServer/forum"
	"EFServer/tool"
	"net/http"
	"time"
)

const cookieKey string = "sid"

//SessionDict 会话map
var SessionDict = make(map[string]*Session, 20)

//Session 会话
type Session struct {
	UUID            string
	CreatedTime     int64
	LastRequestTime int64
	PostSortType    int //帖子排序方式
	User            *forum.UserInDB
}

func createNewSession() *Session {
	session := new(Session)
	session.UUID = tool.NewUUID()
	session.CreatedTime = time.Now().UnixNano()
	session.LastRequestTime = session.CreatedTime
	session.User = nil
	SessionDict[session.UUID] = session
	return session
}

func getExsitSession(r *http.Request) *Session {
	cook, err := r.Cookie(cookieKey)
	if err != nil {
		return nil
	}
	v, ok := SessionDict[cook.Value]
	if !ok {
		return nil
	}
	return v
}

func getExsitOrCreateNewSession(w http.ResponseWriter, r *http.Request, recordTime bool) *Session {
	session := getExsitSession(r)
	if session == nil {
		session = createNewSession()
		http.SetCookie(w, &http.Cookie{
			Name:     cookieKey,
			Value:    session.UUID,
			HttpOnly: true,
			Path:     "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     cookieKey,
			Value:    session.UUID,
			HttpOnly: true,
			Path:     "/Theme",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     cookieKey,
			Value:    session.UUID,
			HttpOnly: true,
			Path:     "/Post",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     cookieKey,
			Value:    session.UUID,
			HttpOnly: true,
			Path:     "/User",
		})
	}
	if recordTime {
		session.LastRequestTime = time.Now().UnixNano()
	}
	return session
}
