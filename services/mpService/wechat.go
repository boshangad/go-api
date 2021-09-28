package mpService

import (
	"github.com/boshangad/go-api/ent"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	mpConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"sync"
)

var DefaultWechat Wechat
func init() {
	DefaultWechat = *NewWechat()
}

type Wechat struct {
	mu sync.Mutex
	wc *wechat.Wechat
	MiniPrograms map[string]*miniprogram.MiniProgram
}

// MiniProgram 获取微信小程序
func (that *Wechat) MiniProgram(app *ent.App) *miniprogram.MiniProgram {
	that.mu.Lock()
	if that.MiniPrograms == nil {
		that.MiniPrograms = make(map[string]*miniprogram.MiniProgram)
	}
	if data, ok := that.MiniPrograms[app.AppID]; ok {
		that.mu.Unlock()
		return data
	}
	mp := that.wc.GetMiniProgram(&mpConfig.Config{
		AppID: app.AppID,
		AppSecret: app.AppSecret,
	})
	that.MiniPrograms[app.AppID] = mp
	that.mu.Unlock()
	return mp
}

// Close 移除全部的数据
func (that *Wechat) Close() {
	if that.MiniPrograms != nil {
		for key, _ := range that.MiniPrograms {
			delete(that.MiniPrograms, key)
		}
	}
}

// NewWechat 实例化微信
func NewWechat() *Wechat {
	return &Wechat{
		wc: wechat.NewWechat(),
		MiniPrograms: make(map[string]*miniprogram.MiniProgram),
	}
}
