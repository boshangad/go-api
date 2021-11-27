package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// 缓存数据结构
type CacheData struct {
	// 文件路径
	Filename string `json:"-" yaml:"-"`
	// 失效时间
	ExpireTime int64 `json:"expireTime" yaml:"expireTime"`
	// 数据内容
	Data interface{} `json:"data" yaml:"data"`

	file *os.File
	mu   sync.Mutex
}

// 写入文件
func (that *CacheData) Write(p []byte) (n int, err error) {
	that.mu.Lock()
	defer that.mu.Unlock()
	if that.file == nil {
		if err = that.openExistingOrNew(); err != nil {
			return 0, err
		}
	}
	n, err = that.file.Write(p)
	return n, err
}

// Close implements io.Closer, and closes the current logfile.
func (that *CacheData) Close() error {
	that.mu.Lock()
	defer that.mu.Unlock()
	return that.close()
}

// close closes the file if it is open.
func (that *CacheData) close() error {
	if that.file == nil {
		return nil
	}
	err := that.file.Close()
	that.file = nil
	return err
}

// openExistingOrNew opens the logfile if it exists and if the current write
// would not put it over MaxSize.  If there is no such file or the write would
// put it over the MaxSize, a new file is created.
func (that *CacheData) openExistingOrNew() error {
	_, err := os.Stat(that.Filename)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(that.Filename), 0744)
		if err != nil {
			return fmt.Errorf("can't make directories for new file: %s", err)
		}
		f, err := os.OpenFile(that.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("can't open new file: %s", err)
		}
		that.file = f
	} else {
		file, err := os.OpenFile(that.Filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		that.file = file
	}
	return nil
}

// 实例化文件缓存
func NewCache(config map[string]interface{}) *FileCache {
	return &FileCache{}
}
