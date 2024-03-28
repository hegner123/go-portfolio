package main

import (
//    "fmt"
//    "strings"
//    "bufio"
)


func parseTinaMd(content string) []KeyValue {

    slicedBytes := parseMdByBytes(content)
    tinaKeyValues := parseTinaKeyValues(slicedBytes)
    return tinaKeyValues
}
