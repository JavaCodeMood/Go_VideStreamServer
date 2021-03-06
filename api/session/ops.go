package session

import(
	"time"
	"sync"
	"github.com/myproject/video_server/api/defs"
	"github.com/myproject/video_server/api/dbops"
	"github.com/myproject/video_server/api/utils"
)

var sessionMap *sync.Map

func NowInMilli() int64{
	return time.Now().UnixNano()/100000
}
func init(){
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB(){
	r , err := dbops.RetrieveAllSessions() 
	if err != nil{
		return 
	}

	r.Range(func(k,v interface{}) bool{
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true 
		})
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func GenerateNewSessionId(un string)(string){
	id , _ := utils.NewUUID()
	ct := NowInMilli()
	ttl := ct + 30*60*1000 
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id,ss)
	dbops.InertSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string)(string,bool){
	ss , ok := sessionMap.Load(sid)
	if ok{
		ct := NowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username,false
	}

	return "", true
}