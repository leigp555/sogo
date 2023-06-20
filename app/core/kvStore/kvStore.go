package kvStore

import (
	"log"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
	"strings"
	"sync"
)

// 定义一个全局键值对存储容器
var sMap sync.Map

func CreateKvStore() *kvStore {
	return &kvStore{}
}

// 定义一个容器结构体
type kvStore struct {
	lock *sync.Mutex
}

// Set  增加键值对
func (c *kvStore) Set(key string, value interface{}) (ret bool) {
	if _, exists := c.KeyIsExists(key); !exists {
		sMap.Store(key, value)
		ret = true
	} else {
		ret = false
		if variable.ZapLog == nil {
			log.Fatal(my_errors.ErrorStoreKeyAlreadyExist + ":" + key)
			return
		}
		// 程序启动初始化完成
		variable.ZapLog.Warn(my_errors.ErrorStoreKeyAlreadyExist + ":" + key)

	}
	return
}

// Delete  删除
func (c *kvStore) Delete(key string) {
	sMap.Delete(key)
}

// Get 查询值
func (c *kvStore) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

// KeyIsExists 判断键是否已存在
func (c *kvStore) KeyIsExists(key string) (interface{}, bool) {
	return sMap.Load(key)
}

// BulkDelete 批量删除相关键值对
func (c *kvStore) BulkDelete(keyPre string) {
	sMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPre) {
				sMap.Delete(keyName)
			}
		}
		return true
	})
}
