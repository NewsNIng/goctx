package context

import (
	"net/http"
	"sync"
	"time"
)

var (
	// 读写锁
	mutex sync.RWMutex
	// 数据持久层
	data = make(map[*http.Request]map[interface{}]interface{})
	// 数据时间持久层
	datat = make(map[*http.Request]int64)
)

func Set(r *http.Request, key, val interface{}) {
	mutex.Lock()
	if data[r] == nil {
		data[r] = make(map[interface{}]interface{})
		datat[r] = time.Now().Unix()
	}
	data[r][key] = val
	mutex.Unlock()
}

func Get(r *http.Request, key interface{}) interface{} {
	mutex.Lock()
	var value interface{}
	if ctx := data[r]; ctx != nil {
		value = ctx[key]
	} else {
		value = nil
	}
	mutex.Unlock()
	return value
}

func GetOk(r *http.Request, key interface{}) (interface{}, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := data[r]; ok {
		value, ok := data[r][key]
		return value, ok
	}
	return nil, false
}

func GetAll(r *http.Request) map[interface{}]interface{} {
	mutex.Lock()
	defer mutex.Unlock()
	if ctx, ok := data[r]; ok {
		result := make(map[interface{}]interface{}, len(ctx))
		for k, v := range ctx {
			result[k] = v
		}
		return result
	}
	return nil
}

func Delete(r *http.Request, key interface{}) {
	mutex.Lock()
	if data[r] != nil {
		delete(data[r], key)
	}
	mutex.Unlock()
}
func Clear(r *http.Request) {
	mutex.Lock()
	clear(r)
	mutex.Unlock()
}
func clear(r *http.Request) {
	delete(data, r)
	delete(datat, r)
}

func Purge(maxAge int) int {
	mutex.Lock()
	count := 0
	if maxAge <= 0 {
		count = len(data)
		data = make(map[*http.Request]map[interface{}]interface{})
		datat = make(map[*http.Request]int64)
	} else {
		min := time.Now().Unix() - int64(maxAge)
		for r := range data {
			if datat[r] < min {
				clear(r)
				count++
			}
		}
	}
	mutex.Unlock()
	return count
}

func ClearHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer Clear(r)
		h.ServeHTTP(w, r)
	})
}
