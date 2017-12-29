package revel

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type JsonRawMap map[string]interface{}

var convertUTC = false
var timeLocation *time.Location

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

func (this JsonRawMap) GetInt(key string) int {
	return int(this.GetInt64(key))
}

func (this JsonRawMap) GetArray(key string) []interface{} {
	if val, ok := this[key]; ok {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			return this[key].([]interface{})
		}
	}
	return nil
}

func formatDate(val string) string {
	var t time.Time
	if timeLocation == nil {
		t, _ = time.Parse("2006-01-02", val)
	} else {
		t, _ = time.ParseInLocation("2006-01-02", val, timeLocation)
		if convertUTC {
			t = t.UTC()
		}
	}
	return t.Format(time.RFC3339)
}

func formatDateTime(val string) string {
	var t time.Time
	var err error
	if timeLocation == nil {
		t, err = time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			t, _ = time.Parse("2006-01-02 15:04", val)
		}
	} else {
		t, err = time.ParseInLocation("2006-01-02 15:04:05", val, timeLocation)
		if err != nil {
			t, _ = time.ParseInLocation("2006-01-02 15:04", val, timeLocation)
		}
	}
	if convertUTC {
		t = t.UTC()
	}
	return t.Format(time.RFC3339)
}

func SetDateToUTC(val bool) {
	convertUTC = val
}

func SetTimeLocation(loc *time.Location) {
	timeLocation = loc
}
