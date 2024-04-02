package main

//import "fmt"


func parseTinaMd(content string) []KeyValue  {

    slicedBytes, stringFile := parseMdByBytes(content)
    tinaKeyValues := parseTinaKeyValues(slicedBytes, stringFile)
    
    return tinaKeyValues
}
