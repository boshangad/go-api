package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/boshangad/v1/app/helpers"
)

type FileCache struct {
	// 缓存路径
	Director string `json:"director,omitempty" yaml:"director"`
	// 缓存键前缀
	KeyPrefix string `json:"keyPrefix,omitempty" yaml:"keyPrefix"`
	// 缓存文件前缀
	CacheFileSuffix string `json:"cacheFileSuffix,omitempty" yaml:"cacheFileSuffix"`
	// 目录层级
	DirectoryLevel int `json:"directoryLevel,omitempty" yaml:"directoryLevel"`
	// 缓存回收时间限制
	GcProbability int64 `json:"gcProbability,omitempty" yaml:"gcProbability"`
	// 文件权限模式
	FileMode int64 `json:"fileMode,omitempty" yaml:"fileMode"`
	// 目录权限
	DirMode uint64 `json:"dirMode,omitempty" yaml:"dirMode"`
}

func (that FileCache) getDirector() string {
	if that.Director == "" {
		return filepath.Join(filepath.Dir(os.Args[0]), "runtime/cache")
	}
	return that.Director
}

// 缓存文件名称
func (that FileCache) getCacheFileKey(key string) string {
	return buildingKey(key)
}

// 获取缓存文件路径
func (that FileCache) getCacheFilePath(key string) string {
	var (
		baseKey        = that.getCacheFileKey(key)
		fullKey        = that.KeyPrefix + baseKey
		directoryLevel = that.DirectoryLevel
		levelStr       = ""
		cacheFilename  = ""
	)
	if directoryLevel > 0 {
		levelStr = baseKey[0:directoryLevel]
		cacheFilename = path.Join(that.getDirector(), strings.Join(strings.Split(levelStr, ""), "/"), fullKey)
	} else {
		cacheFilename = path.Join(that.getDirector(), fullKey)
	}
	return cacheFilename
}

// 获取缓存数据
func (that FileCache) getData(key string) *CacheData {
	var (
		filename  = that.getCacheFilePath(key)
		cacheData = CacheData{}
		nowTime   = time.Now().Unix()
	)
	// 缓存文件存在，需要判断文件失效
	if helpers.IsFile(filename) {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Println(err)
			return nil
		}
		err = json.Unmarshal(data, &cacheData)
		if err != nil {
			log.Println(err)
			return nil
		}
		if cacheData.ExpireTime > 0 && cacheData.ExpireTime < nowTime {
			_ = that.Delete(key)
			return nil
		}
		return &cacheData
	}
	return nil
}

// 获取缓存数据
func (that FileCache) Get(key string) interface{} {
	var (
		cacheData = that.getData(key)
	)
	if cacheData != nil {
		return cacheData.Data
	}
	return nil
}

// 获取缓存数据
func (that FileCache) Set(key string, value interface{}, ttl int64) error {
	var (
		filename  = that.getCacheFilePath(key)
		cacheData = &CacheData{
			Filename:   filename,
			Data:       value,
			ExpireTime: ttl,
		}
	)
	if ttl > 0 {
		cacheData.ExpireTime = time.Now().Add(time.Duration(ttl)).Unix()
	}
	defer cacheData.Close()
	data, err := json.Marshal(that)
	if err != nil {
		return fmt.Errorf("json.marshal failed, err: %s", err)
	}
	_, err = cacheData.Write(data)
	return err
}

// 删除文件
func (that FileCache) Delete(key string) error {
	var (
		filename = that.getCacheFilePath(key)
	)
	if helpers.IsFile(filename) {
		return os.Remove(filename)
	}
	return nil
}

// 缓存是否存在
func (that FileCache) Exists(key string) bool {
	var (
		filename  = that.getCacheFilePath(key)
		cacheData = CacheData{}
		nowTime   = time.Now().Unix()
	)
	// 缓存文件存在，需要判断文件失效
	if helpers.IsFile(filename) {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Println(err)
			return false
		}
		err = json.Unmarshal(data, &cacheData)
		if err != nil {
			log.Println(err)
			return false
		}
		if cacheData.ExpireTime > 0 && cacheData.ExpireTime < nowTime {
			_ = that.Delete(key)
			return false
		}
		return true
	}
	return false
}

// 清空全部缓存文件
func (that FileCache) Flush() error {
	var (
		director = that.getDirector()
	)
	dirs, err := ioutil.ReadDir(director)
	if err == nil {
		for _, d := range dirs {
			_ = os.RemoveAll(path.Join([]string{director, d.Name()}...))
		}
	}
	return err
}
