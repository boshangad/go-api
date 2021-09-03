package cache

import (
	"context"
	"fmt"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/tend/wechatServer/ent"
	"github.com/tend/wechatServer/ent/appoption"
	"io/ioutil"
	"log"
	"time"
)

// AppToken struct contains *AppToken.Client
type AppToken struct {
	appId uint64
	entClient *ent.Client
	redis *cache.Redis
	memcache *cache.Memcache
	memory *cache.Memory
	appOption *ent.AppOption
}

// NewAppTokenCache NewAppToken create new AppToken
func NewAppTokenCache(appId uint64, entClient *ent.Client, redis *cache.Redis, memcache *cache.Memcache) *AppToken {
	return &AppToken{
		appId: appId,
		entClient: entClient,
		redis: redis,
		memory: cache.NewMemory(),
		memcache: memcache,
	}
}

// Get return cached value
func (mem *AppToken) Get(key string) interface{} {
	var item interface{}
	// 内存缓存
	if mem.memory != nil {
		if item = mem.memory.Get(key); item != nil {
			return item
		}
	}
	// redis 缓存
	if mem.redis != nil {
		if item = mem.redis.Get(key); item != nil {
			return item
		}
	}
	// memcache缓存
	if mem.memcache != nil {
		if item = mem.memcache.Get(key); item != nil {
			return item
		}
	}
	// 数据库中保存
	key = "token_" + key
	if mem.entClient != nil {
		re ,err := mem.entClient.AppOption.Query().
			Where(appoption.And(appoption.AppIDEQ(mem.appId), appoption.NameEQ(key))).
			First(context.Background())
		if err != nil {
			log.Fatalln("查询数据库失败", err)
			return nil
		}
		if re.ExpireTime > uint64(time.Now().Unix()) {
			return re.Name
		}
	}
	return nil
}

// IsExist check value exists in AppToken.
func (mem *AppToken) IsExist(key string) bool {
	if err := mem.Get(key); err != nil {
		return false
	}
	return true
}

// Set cached value with key and expire time.
func (mem *AppToken) Set(key string, val interface{}, timeout time.Duration) (err error) {
	// 数据库中保存
	s := fmt.Sprintf("%s => [%s]", key, fmt.Sprintf("%s", val))
	fmt.Printf(s)
	_ = ioutil.WriteFile("C:\\Users\\huanghu\\Desktop\\token.txt", []byte(s), 0666)
	if mem.entClient != nil {
		keyT := "token_" + key
		dApp ,err := mem.entClient.AppOption.Query().
			Where(appoption.And(appoption.AppIDEQ(mem.appId), appoption.NameEQ(keyT))).
			First(context.Background())
		e, _ := time.ParseDuration(fmt.Sprintf("+%ds", timeout))
		fmt.Println(err)
		if err != nil && dApp != nil {
			_, err = dApp.Update().SetValue(fmt.Sprintf("%s", val)).
				SetExpireTime(uint64(time.Now().Unix()) + uint64(e.Seconds())).
				Save(context.Background())
		} else {
			_, err = mem.entClient.AppOption.Create().
				SetAppID(mem.appId).
				SetName(keyT).
				SetTitle("token值").
				SetDescription("").
				SetValue(fmt.Sprintf("%s", val)).
				SetExpireTime(uint64(time.Now().Unix()) + uint64(e.Seconds())).
				Save(context.Background())
		}
		fmt.Println(err)
	}
	// redis 缓存
	if mem.redis != nil {
		err = mem.redis.Set(key, val, timeout)
		if err != nil {
			return nil
		}
	}
	// memcache缓存
	if mem.memcache != nil {
		err = mem.memcache.Set(key, val, timeout)
		if err != nil {
			return nil
		}
	}
	// 内存缓存
	if mem.memory != nil {
		err = mem.memory.Set(key, val, timeout)
	}
	return err
}

// Delete delete value in AppToken.
func (mem *AppToken) Delete(key string) error {
	// 数据库中保存
	if mem.entClient != nil {
		keyT := "token_" + key
		_ ,err := mem.entClient.AppOption.Query().
			Where(appoption.And(appoption.AppIDEQ(mem.appId), appoption.NameEQ(keyT))).
			First(context.Background())
		if err == nil {

		}
	}
	var err error
	// redis 缓存
	if mem.redis != nil {
		err = mem.redis.Delete(key)
		if err != nil {
			return nil
		}
	}
	// memcache缓存
	if mem.memcache != nil {
		err = mem.memcache.Delete(key)
		if err != nil {
			return nil
		}
	}
	// 内存缓存
	if mem.memory != nil {
		err = mem.memory.Delete(key)
	}
	return err
}
