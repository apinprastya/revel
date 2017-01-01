package revel

import (
"reflect"
"errors"
"time"
)

type JsonRawMap map[string]interface{}

func (this JsonRawMap) FixInt64(keys ...string) {
    if len(keys) > 0 {
        for _, v := range(keys) {
            if v1, k := this[v]; k && v1 != nil {
                if reflect.TypeOf(this[v]).Kind() != reflect.Float64 {
                    panic(errors.New("Keys does not Float, please check"))
                }
                this[v] = int64(this[v].(float64))
            }
        }
    }
}

func (this JsonRawMap) FixInt(keys ...string) {
    if len(keys) > 0 {
        for _, v := range(keys) {
            if v1, k := this[v]; k && v1 != nil {
                if reflect.TypeOf(this[v]).Kind() != reflect.Float64 {
                    panic(errors.New("Keys does not Float, please check"))
                }
                this[v] = int(this[v].(float64))
            }
        }
    }
}

func (this JsonRawMap) FormatDate(keys ...string) {
    if len(keys) > 0 {
        for _, v := range(keys) {
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
        for _, v := range(keys) {
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
        for _, v := range(keys) {
            if v1, k := this[v]; k  && v1 != nil {
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
        for _, v := range(keys) {
            if _, ok := this[v]; ok {
                delete(this, v)
            }
        }
    }
}

func formatDate(val string) string {
    t, _ := time.Parse("2006-01-02", val)
    return t.Format(time.RFC3339)
}

func formatDateTime(val string) string {
    t, _ := time.Parse("2006-01-02 15:04:05", val)
    return t.Format(time.RFC3339)
}
