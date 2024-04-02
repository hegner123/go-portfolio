package main

import (
//    "fmt"
//    "strings"
//    "bufio"
)


func parseTinaMd(content string) []KeyValue  {

    slicedBytes, stringFile := parseMdByBytes(content)
    tinaKeyValues := parseTinaKeyValues(slicedBytes, stringFile)
    return tinaKeyValues
}
