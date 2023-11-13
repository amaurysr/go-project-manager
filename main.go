package main

import ( 
	"fmt"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
	"go-project-manager/pkg/progresschanger"
	"github.com/schollz/progressbar/v3"
)
	
func main(){
	var list = mainreader.Listpaste()
	fmt.Println("\n\nCurrent Project: ",list.Datalist[0].ProjectName)
	fmt.Println("Your Progress: ")
	bar := progressbar.NewOptions(100, progressbar.OptionSetWidth(list.Datalist[0].Progress), progressbar.OptionSetPredictTime(false))
	bar.Add(list.Datalist[0].Progress)
	fmt.Println("\nProject Members: ",list.Datalist[0].ProjectMembers)
	var numb int;
	numb = progresschanger.RequestChange(list.Datalist[0].Progress)
	list.Datalist[0].Progress = numb
	jsonconvert.Convert(list)
}
