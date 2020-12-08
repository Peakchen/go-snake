package common

import (
    "crypto/md5"
    "encoding/hex"

    "fmt"
    "reflect"
    "unsafe"
)

func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

func SizeVal(data interface{}) int {
    return sizeof(reflect.ValueOf(data))
}

func sizeof(v reflect.Value) (sum int) {
    switch v.Kind() {
    case reflect.Map:
        keys := v.MapKeys()
        for i := 0; i < len(keys); i++ {
            mapkey := keys[i]
            s := sizeof(mapkey)
            if s < 0 {
                return -1
            }
            sum += s
            s = sizeof(v.MapIndex(mapkey))
            if s < 0 {
                return -1
            }
            sum += s
        }

    case reflect.Slice, reflect.Array:
        for i, n := 0, v.Len(); i < n; i++ {
            s := sizeof(v.Index(i))
            if s < 0 {
                return -1
            }
            sum += s
        }

    case reflect.String:
        for i, n := 0, v.Len(); i < n; i++ {
            s := sizeof(v.Index(i))
            if s < 0 {
                return -1
            }
            sum += s
        }
  
    case reflect.Ptr, reflect.Interface:
        p := (*[]byte)(unsafe.Pointer(v.Pointer()))
        if p == nil {
            return
        }
        sum += sizeof(v.Elem())

    case reflect.Struct:
        for i, n := 0, v.NumField(); i < n; i++ {
            s := sizeof(v.Field(i))
            if s < 0 {
                return
            }
            sum += s
        }

    case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
        reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
        reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
        reflect.Int, reflect.Chan:

        sum += int(v.Type().Size())

    default:
        fmt.Println("input val Kind not found:", v.Kind())
    }

    return
}