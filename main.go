package main

import ( 
	"fmt"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
	"go-project-manager/pkg/progresschanger"
)
	
func main(){
	var list = mainreader.Listpaste()
	fmt.Println("Current Project: ",list.Datalist[0].ProjectName)
	fmt.Println("Your Progress: ",list.Datalist[0].Progress)
	fmt.Println("Project Members: ",list.Datalist[0].ProjectMembers)
	var numb int;
	numb = progresschanger.RequestChange(list.Datalist[0].Progress)
	list.Datalist[0].Progress = numb
	jsonconvert.Convert(list)
}
