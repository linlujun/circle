package session

import (
	"circleTest/api/dbops"
	"circleTest/api/defs"
	"circleTest/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

//清除过期session
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

//从DB加载到缓存
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

//生成session
func GenerateNewSessionId(login_name string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 24*60*60*1000 // Severside session valid time: 1 day

	ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
	sessionMap.Store(id, ss)
	//写入session表
	dbops.InsertSession(id, ttl, login_name)

	return id
}

//获取session 并判断是否过期
//如果过期则清除
//考虑增加定期自动清除功能
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
