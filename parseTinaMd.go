package main

//import "fmt"


func parseTinaMd(content string) map[string]interface{}  {

     stringFile := parseMdByBytes(content)
    tinaKeyValues := parseTinaKeyValues( stringFile)
    
    return tinaKeyValues
}
