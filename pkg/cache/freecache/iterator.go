package freecachepackage

import (
	cachepackage "back/pkg/cache"
	"github.com/coocood/freecache"
)

type iterator struct {
	iter *freecache.Iterator
}

func (i *iterator) Next() *cachepackage.Entry {
	entry := i.iter.Next()
	if entry == nil {
		return nil
	}

	return &cachepackage.Entry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}
