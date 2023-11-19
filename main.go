package main

import ( 
	"fmt"
	"slices"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
	"go-project-manager/pkg/progresschanger"
	"go-project-manager/pkg/taskmanager"
	"github.com/schollz/progressbar/v3"
	"os"
	"bufio"
	"strconv"
)

var list = mainreader.Listpaste()

func stringEntry() string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func intEntry() int{
	var n int;
	for{
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		entryStr := scanner.Text()
		entry, err := strconv.Atoi(entryStr)
		if err == nil{
			n = entry	
			break
		}else{
			fmt.Println("Enter a valid entry.")
		}
	}	
	return n
}

func editProjectMembers(entry int){
	length := len(list.Datalist[entry].ProjectMembers)
	for{
		fmt.Println("Do you want to edit the list of project members? ")
		entryStr3 := stringEntry()
		if entryStr3 == "y" || entryStr3 == "yes"{
			for i := 0; i<length; i++{
				fmt.Println("Enter a new name, or write same (so that it doesn't edit)")
				entryStr4 := stringEntry()
				if entryStr4 == "same" || entryStr4 == "Same"{
					continue
				}else{
					list.Datalist[entry].ProjectMembers[i] = entryStr4 
					jsonconvert.Convert(list)
				}
			}
		}else if entryStr3 == "n" || entryStr3 == "no"{ 
			break;
		}else{
			fmt.Println("Enter a valid entry.")
		}

	}
}

func addProjectMember(entry int){
	for{
		fmt.Println("Do you want to add a new Project Member?")
		entryStr1 := stringEntry()
		if entryStr1 == "y" || entryStr1 == "yes"{
			fmt.Println("Enter a new Project Member.")
			entryStr2 := stringEntry()
			list.Datalist[entry].ProjectMembers = append(list.Datalist[entry].ProjectMembers, entryStr2)			
			jsonconvert.Convert(list)
		}else if  entryStr1 == "n" || entryStr1 == "no"{
			break;
		}else{
			fmt.Println("Please enter a valid entry.")
		}
	}		
}

func viewLists(){
	list = mainreader.Listpaste()
	length := len(list.Datalist)
	for i := 0; i<length; i++ {
		taskmanager.ProgressEditer(i)
		fmt.Println("\n\nCurrent Project: ",list.Datalist[i].ProjectName)
        	fmt.Println("Your Progress: ")
        	bar := progressbar.NewOptions(100, progressbar.OptionSetWidth(list.Datalist[i].Progress), progressbar.OptionSetPredictTime(false))
        	bar.Add(list.Datalist[i].Progress)
        	fmt.Println("\nProject Members: ",list.Datalist[i].ProjectMembers)
	}
}

func editLists(){
	length := len(list.Datalist)
	fmt.Println("Choose the respective index number, in order to edit that specific project.")
	for i:=0; i < length; i++ {
		fmt.Println(i, ",", list.Datalist[i].ProjectName)
	} 
	// check your main func and add the progresschanger in this one, also try and create that progressName changer too by using scan
	for{
		entry := intEntry() 
		if entry < length{
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
							entryStr1 := stringEntry()
							list.Datalist[i].ProjectName = entryStr1
							jsonconvert.Convert(list)
							editProjectMembers(i)
							addProjectMember(i)
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
func CreateProject(){
	fmt.Println("What name do you want to give your project?")
	entryStr := stringEntry()
	var members []string; 
	for {
		fmt.Println("Enter a name of a Project Member, (no) - to stop adding.")
		entryStr1 := stringEntry()
		if entryStr1 == "no" || entryStr1 == "No"{
			break
		}else{
			members = append(members,entryStr1)
		}
	}
	newProject := mainreader.Data{
		ProjectName: entryStr,
		ProjectMembers: members,
	}	 
	list.Datalist = append(list.Datalist, newProject)
	jsonconvert.Convert(list)
}

func removeInstance(entry int){
	list.Datalist = slices.Delete(list.Datalist, entry, entry+1)
	jsonconvert.Convert(list)
}

func DeleteProject(){
	fmt.Println("Which instance do you want to delete?")
	length := len(list.Datalist)
	for i:=0;i<length; i++{
		fmt.Printf("%d %s",i,list.Datalist[i].ProjectName)	
	}
	for{
		entry := intEntry()
		if entry >= 0 && entry <= length{
			removeInstance(entry)
			break
		}else{
			fmt.Printf("Please enter a valid input.")
		}
	}
}

func AskScanner(){
	for{
		fmt.Println("\nChoose your Option: (V)-view project lists, (E)-edit project lists, (Exit)-exit the program, (A) - add tasks to each project, (M) - markdown completed tasks, (C) - create new Project, (D) - Delete Project")
		answer := stringEntry()
		if answer == "V"{
			viewLists() 
		}else if answer == "E"{
			editLists()
		}else if answer == "A"{
			taskmanager.TasksCreator()
		}else if answer == "M" {
			taskmanager.TaskEditer()
		}else if answer == "D"{
			// breaks so that it makes sure that .json is updated
			DeleteProject()
			break
		}else if answer == "C" || answer == "c"{
			// breaks so that it makes sure that the .json is updated 
			CreateProject()
			break
		}else if answer == "Exit" || answer == "exit"{
			break
		}else{
			fmt.Println("Please enter a valid response.")
		}
	}
}
	
func main(){
	AskScanner()
}
