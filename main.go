package main

import ( 
	"fmt"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
	"go-project-manager/pkg/progresschanger"
	"github.com/schollz/progressbar/v3"
	"os"
	"bufio"
	"strconv"
)

func viewLists(){
	var list = mainreader.Listpaste()
	length := len(list.Datalist)
	for i := 0; i<length; i++ {
		fmt.Println("\n\nCurrent Project: ",list.Datalist[i].ProjectName)
        	fmt.Println("Your Progress: ")
        	bar := progressbar.NewOptions(100, progressbar.OptionSetWidth(list.Datalist[i].Progress), progressbar.OptionSetPredictTime(false))
        	bar.Add(list.Datalist[i].Progress)
        	fmt.Println("\nProject Members: ",list.Datalist[i].ProjectMembers)
	}
}

func editLists(){
	var list = mainreader.Listpaste()
	length := len(list.Datalist)
	fmt.Println("Choose the respective index number, in order to edit that specific project.")
	for i:=0; i < length; i++ {
		fmt.Println(i, ",", list.Datalist[i].ProjectName)
	} 
	// check your main func and add the progresschanger in this one, also try and create that progressName changer too by using scan
	for{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		entryStr := scanner.Text()
		entry, err := strconv.Atoi(entryStr)
		if entry < length && err == nil{
			for i := 0; i<length; i++{
				if entry == i{
					var numb int;
        				numb = progresschanger.RequestChange(list.Datalist[i].Progress)
        				list.Datalist[i].Progress = numb
        				jsonconvert.Convert(list)
					for{
						var request string
						fmt.Println("Do you want to change the ProjectName?")
						fmt.Scanf("%s", &request)
						if request == "yes" || request == "y"{
							fmt.Println("Now give it a new name.")
							scanner1 := bufio.NewScanner(os.Stdin)
							scanner1.Scan()
							entryStr1 := scanner1.Text()
							list.Datalist[i].ProjectName = entryStr1
							jsonconvert.Convert(list)
							break
						}else if request == "no" || request == "n"{
							break
						}else{
							fmt.Println("Enter a valid input")	
						}
					}
					break
				}
			}
		}else{
		fmt.Println("Enter a valid entry.")
		}
		break
	}	
}

func askScanner(){
	for{
		fmt.Println("\nChoose your Option: (V)-view project lists, (E)-edit project lists, (Exit)-exit the program")
		var answer string; 
		fmt.Scanf("%s", &answer);
		if answer == "V"{
			viewLists() 
		}else if answer == "E"{
			editLists()
		}else if answer == "Exit" || answer == "exit"{
			os.Exit(0)
		}else{
			fmt.Println("Please enter a valid response.")
		}
	}
}
	
func main(){
	askScanner()
}
