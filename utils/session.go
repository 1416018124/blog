package utils

import (
	"github.com/beego/beego/v2/server/web/session"
)

func InitSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	var globalSessions *session.Manager
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}
