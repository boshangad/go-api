package db

import (
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"strings"
)

func DefaultClient() *ent.Client {
	key := strings.TrimSpace(global.G_CONFIG.DB.Default)
	if key == "" {
		global.G_LOG.Panic("")
	}
	return Client(key)
}

func Client(key string) *ent.Client {
	db, ok := global.G_DB[key]
	if !ok {
		global.G_LOG.Fatal("ddddd")
	}
	return db
}
