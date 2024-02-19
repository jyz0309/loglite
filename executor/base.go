package executor

import (
	"github.com/jyz0309/loglite/types"
)

type Executor interface {
	Exec() (ResultSet, error)
}

type ResultSet interface {
	Keys() []string
	Next() ([][]types.Item, bool, error)
}

type BaseResultSet struct {
	keys []string
	ch   chan []types.Item
}

func (rs *BaseResultSet) Keys() []string {
	return rs.keys
}

func (rs *BaseResultSet) Next() ([][]types.Item, bool, error) {
	result := make([][]types.Item, 0)
	var done bool
	for {
		item, more := <-rs.ch
		if !more {
			done = true
			break
		} else {
			result = append(result, item)
			if len(result) == 4 {
				done = false
				break
			}
		}
	}
	return result, done, nil
}

func BuildBaseResultSet(keys []string, ch chan []types.Item) *BaseResultSet {
	return &BaseResultSet{
		ch:   ch,
		keys: keys,
	}
}
