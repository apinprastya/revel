package revel

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type JsonRawMap map[string]interface{}

func (this JsonRawMap) FixInt64(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if v1, k := this[v]; k && v1 != nil {
				if reflect.TypeOf(this[v]).Kind() != reflect.Float64 {
					this[v] = 0
					continue
				}
				this[v] = int64(this[v].(float64))
			}
		}
	}
}

func (this JsonRawMap) FixInt(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if v1, k := this[v]; k && v1 != nil {
				if reflect.TypeOf(this[v]).Kind() != reflect.Float64 {
					this[v] = 0
					continue
				}
				this[v] = int(this[v].(float64))
			}
		}
	}
}

func (this JsonRawMap) FormatDate(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if v1, k := this[v]; k && v1 != nil {
				if reflect.TypeOf(this[v]).Kind() != reflect.String {
					panic(errors.New("Keys does not string"))
				}
				this[v] = formatDate(this[v].(string))
			}
		}
	}
}

func (this JsonRawMap) FormatDateTime(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if v1, k := this[v]; k && v1 != nil {
				if reflect.TypeOf(this[v]).Kind() != reflect.String {
					panic(errors.New("Keys does not string"))
				}
				this[v] = formatDateTime(this[v].(string))
			}
		}
	}
}

func (this JsonRawMap) FormatBool(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if v1, k := this[v]; k && v1 != nil {
				if reflect.TypeOf(this[v]).Kind() == reflect.Float64 {
					f := this[v].(float64)
					if f == 0 {
						this[v] = false
					} else {
						this[v] = true
					}
				}
			}
		}
	}
}

func (this JsonRawMap) RemoveField(keys ...string) {
	if len(keys) > 0 {
		for _, v := range keys {
			if _, ok := this[v]; ok {
				delete(this, v)
			}
		}
	}
}

func (this JsonRawMap) Available(key string) bool {
	_, ok := this[key]
	return ok
}

func (this JsonRawMap) GetString(key string) string {
	if val, ok := this[key]; ok {
		if reflect.TypeOf(val).Kind() == reflect.String {
			return val.(string)
		}
	}
	return ""
}

func (this JsonRawMap) GetInt64(key string) int64 {
	if val, ok := this[key]; ok {
		if reflect.TypeOf(val).Kind() == reflect.String {
			v, _ := strconv.ParseInt(val.(string), 10, 0)
			return v
		} else if reflect.TypeOf(val).Kind() == reflect.Float64 {
			return int64(val.(float64))
		}
	}
	return 0
}

func formatDate(val string) string {
	t, _ := time.Parse("2006-01-02", val)
	return t.Format(time.RFC3339)
}

func formatDateTime(val string) string {
	t, _ := time.Parse("2006-01-02 15:04:05", val)
	return t.Format(time.RFC3339)
}
