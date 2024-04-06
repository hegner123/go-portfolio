package main

import (
    "fmt"
    "reflect"
)

func printProperties(obj interface{}) reflect.Value {
    val := reflect.ValueOf(obj)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    typ := val.Type()
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fmt.Printf("%s: %v\n", typ.Field(i).Name, field.Interface())
    }
    return val
}
