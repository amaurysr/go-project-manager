package jsonconvert 

import (
	"os"
	"log"
	"encoding/json" 
	"go-project-manager/pkg/mainreader"
)

func Convert(list mainreader.Datalist) {
	newF, err := json.MarshalIndent(list, "", "   ")
	if err!=nil{
		log.Fatal(err)
	}
	err1 := os.WriteFile("simple.json", newF, 0664)
	if err1 != nil {
		log.Fatal(err1)
	}
}
