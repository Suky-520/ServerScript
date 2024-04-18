package collection

import (
	"encoding/json"
	"strings"
	"sync"
)

// ConcurrentArrayMap 线程安全、保持存续的map
// 使用 sync.RWMutex 保证线程安全
type ConcurrentArrayMap[K comparable, V any] struct {
	sync.RWMutex
	keys   []K
	values map[K]V
}

// NewConcurrentArrayMap 创建并返回一个新的线程安全的 ConcurrentArrayMap
func NewConcurrentArrayMap[K comparable, V any]() *ConcurrentArrayMap[K, V] {
	return &ConcurrentArrayMap[K, V]{
		values: make(map[K]V),
		keys:   make([]K, 0),
	}
}

// Set 设置键值对，并保持键的顺序
func (cam *ConcurrentArrayMap[K, V]) Set(key K, value V) {
	cam.Lock()         // 加锁以确保线程安全
	defer cam.Unlock() // 解锁
	if _, exists := cam.values[key]; !exists {
		cam.keys = append(cam.keys, key)
	}
	cam.values[key] = value
}

// Get 获取指定键的值
func (cam *ConcurrentArrayMap[K, V]) Get(key K) (V, bool) {
	cam.RLock()         // 读锁
	defer cam.RUnlock() // 解读锁
	val, exists := cam.values[key]
	return val, exists
}

// Delete 删除指定的键
func (cam *ConcurrentArrayMap[K, V]) Delete(key K) bool {
	cam.Lock()
	defer cam.Unlock()

	if _, exists := cam.values[key]; exists {
		delete(cam.values, key)
		// 查找并移除 key
		for i, k := range cam.keys {
			if k == key {
				cam.keys = append(cam.keys[:i], cam.keys[i+1:]...)
				return true
			}
		}
	}
	return false
}

// Has 判断key是否存在
func (cam *ConcurrentArrayMap[K, V]) Has(key K) bool {
	if _, exists := cam.values[key]; exists {
		return true
	}
	return false
}

// Size 返回大小
func (cam *ConcurrentArrayMap[K, V]) Size() int {
	return len(cam.keys)
}

// Clear 清空整个 map
func (cam *ConcurrentArrayMap[K, V]) Clear() {
	cam.Lock()
	defer cam.Unlock()

	cam.values = make(map[K]V)
	cam.keys = make([]K, 0)
}

// Keys 返回所有键的有序切片
func (cam *ConcurrentArrayMap[K, V]) Keys() []K {
	cam.RLock()         // 读锁
	defer cam.RUnlock() // 解读锁
	return cam.keys
}

func (cam *ConcurrentArrayMap[K, V]) UnmarshalJSON(data []byte) error {
	m := make(map[K]V)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	cam.values = m
	keys := make([]K, 0)
	for key := range cam.values {
		keys = append(keys, key)
	}
	cam.keys = keys
	return nil
}

func (cam *ConcurrentArrayMap[K, V]) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	sb.WriteString("{")
	for i, k := range cam.keys {
		mk, err := json.Marshal(k)
		if err != nil {
			continue
		}
		mv, err := json.Marshal(cam.values[k])
		if err != nil {
			continue
		}
		if i == 0 {
			sb.WriteString(string(mk) + ":" + string(mv))
		} else {
			sb.WriteString("," + string(mk) + ":" + string(mv))
		}
	}
	sb.WriteString("}")
	return []byte(sb.String()), nil
}

func (cam *ConcurrentArrayMap[K, V]) GetValues() map[K]V {
	return cam.values
}
