package mainreader

import (
        "os"
        "log"
	"encoding/json"
)
type Data struct {
        ProjectName string `json:"projectName"`
        Progress int `json:"progress"`
        ProjectMembers []string `json:"projectMembers"`
}

type Datalist struct{
        Datalist []Data `json:"Data"`
}

func Listpaste() Datalist{
	content, err := os.ReadFile("simple.json")
        if err != nil{
                log.Fatal(err)
        }
	var list Datalist
        err0 := json.Unmarshal(content, &list)
        if err0 != nil{
                log.Fatal(err0)
        }
	return list
}
