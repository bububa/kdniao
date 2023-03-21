package util

import (
	"net/url"
	"strings"
	"sync"
)

var stringsBuilder = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	b := stringsBuilder.Get().(*strings.Builder)
	b.Reset()
	return b
}

func PutStringsBuilder(b *strings.Builder) {
	stringsBuilder.Put(b)
}

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func GetUrlValues() url.Values {
	vals := urlValuesPool.Get().(url.Values)
	for k := range vals {
		vals.Del(k)
	}
	return vals
}

func PutUrlValues(vals url.Values) {
	urlValuesPool.Put(vals)
}

func StringsJoin(strs ...string) string {
	var n int
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	if n <= 0 {
		return ""
	}
	builder := GetStringsBuilder()
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}
	ret := builder.String()
	PutStringsBuilder(builder)
	return ret
}
