package jsonconvert 

import (
	"os"
	"log"
	"encoding/json" 
	"go-project-manager/pkg/mainreader"
)

func Convert(list mainreader.Datalist) {
	file := mainreader.File 
	newF, err := json.Marshal(list)
	if err != nil{
		log.Fatal(err)
	}
	err1 := os.WriteFile(file, newF, 0663)
	if err1 != nil {
		log.Fatal(err1)
	}
}
