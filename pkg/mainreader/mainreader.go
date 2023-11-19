package mainreader

import (
        "os"
        "log"
	"encoding/json"
)

type Task struct {
	Task string `json:"Task"`
	Weight int `json:"Weight"`
	Completed bool `json:"Completed"`
}

type Data struct {
        ProjectName string `json:"ProjectName"`
        Progress int `json:"Progress"`
        ProjectMembers []string `json:"ProjectMembers"`
	Tasklength int `json:"TaskLength"`
	TasksCompleted bool `json:"TasksCompleted"`
	Tasklist []Task `json:"TaskList"`
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
