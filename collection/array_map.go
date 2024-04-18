package collection

import (
	"encoding/json"
	"strings"
)

// ArrayMap 保持存续的map
type ArrayMap[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

// NewArrayMap 创建并返回一个新的线程安全的 ArrayMap
func NewArrayMap[K comparable, V any]() *ArrayMap[K, V] {
	return &ArrayMap[K, V]{
		values: make(map[K]V),
		keys:   make([]K, 0),
	}
}

// Set 设置键值对，并保持键的顺序
func (am *ArrayMap[K, V]) Set(key K, value V) {
	if _, exists := am.values[key]; !exists {
		am.keys = append(am.keys, key)
	}
	am.values[key] = value
}

// Get 获取指定键的值
func (am *ArrayMap[K, V]) Get(key K) (V, bool) {
	val, exists := am.values[key]
	return val, exists
}

// Delete 删除指定的键
func (am *ArrayMap[K, V]) Delete(key K) bool {
	if _, exists := am.values[key]; exists {
		delete(am.values, key)
		// 查找并移除 key
		for i, k := range am.keys {
			if k == key {
				am.keys = append(am.keys[:i], am.keys[i+1:]...)
				return true
			}
		}
	}
	return false
}

// Has 判断key是否存在
func (am *ArrayMap[K, V]) Has(key K) bool {
	if _, exists := am.values[key]; exists {
		return true
	}
	return false
}

// Size 返回大小
func (am *ArrayMap[K, V]) Size() int {
	return len(am.keys)
}

// Clear 清空整个 map
func (am *ArrayMap[K, V]) Clear() {
	am.values = make(map[K]V)
	am.keys = make([]K, 0)
}

// Keys 返回所有键的有序切片
func (am *ArrayMap[K, V]) Keys() []K {
	return am.keys
}

func (am *ArrayMap[K, V]) UnmarshalJSON(data []byte) error {
	m := make(map[K]V)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	am.values = m
	keys := make([]K, 0)
	for key := range am.values {
		keys = append(keys, key)
	}
	am.keys = keys
	return nil
}

func (am *ArrayMap[K, V]) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	sb.WriteString("{")
	for i, k := range am.keys {
		mk, err := json.Marshal(k)
		if err != nil {
			continue
		}
		mv, err := json.Marshal(am.values[k])
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

func (am *ArrayMap[K, V]) GetValues() map[K]V {
	return am.values
}
