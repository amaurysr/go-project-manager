package main

import ( 
	"fmt"
	"slices"
	"os"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
	"go-project-manager/pkg/progresschanger"
	"go-project-manager/pkg/taskmanager"
	"go-project-manager/pkg/input"
	"github.com/schollz/progressbar/v3"
)

var list = mainreader.Listpaste()

func editProjectMembers(entry int){
	length := len(list.Datalist[entry].ProjectMembers)
	for{
		fmt.Println("Do you want to edit the list of project members or do you want to delete a member? ")
		entryStr3 := input.StringEntry()
		if entryStr3 == "y" || entryStr3 == "yes"{
			for i := 0; i<length; i++{
				fmt.Println("Enter a new name, or write same (so that it doesn't edit)")
				entryStr4 := input.StringEntry()
				if entryStr4 == "same" || entryStr4 == "Same"{
					continue
				}else{
					list.Datalist[entry].ProjectMembers[i] = entryStr4 
					jsonconvert.Convert(list)
				}
			}
		}else if entryStr3 == "n" || entryStr3 == "no"{ 
			break;
		}else if entryStr3 == "d" || entryStr3 == "D" || entryStr3 == "delete" || entryStr3 == "Delete"{
			for{
				for i:= 0; i<length; i++{
					fmt.Printf("%d %v", i, list.Datalist[entry].ProjectMembers[i])
				}
				entryint := input.IntEntry()
				if entryint >= 0 || entryint > length{
					list.Datalist[entry].ProjectMembers = slices.Delete(list.Datalist[entry].ProjectMembers,entryint,entryint+1)
					jsonconvert.Convert(list)		
					break
				}
			}
		}else {
			fmt.Println("Enter a valid entry.")
		}
	}
}

func addProjectMember(entry int){
	for{
		fmt.Println("Do you want to add a new Project Member?")
		entryStr1 := input.StringEntry()
		if entryStr1 == "y" || entryStr1 == "yes"{
			fmt.Println("Enter a new Project Member.")
			entryStr2 := input.StringEntry()
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
	if length > 0 {
		fmt.Println("Choose the respective index number, in order to edit that specific project.")
		for i:=0; i < length; i++ {
			fmt.Println(i, ",", list.Datalist[i].ProjectName)
		} 
		// check your main func and add the progresschanger in this one, also try and create that progressName changer too by using scan
		for{
			entry := input.IntEntry() 
			if entry < length{
				for i := 0; i<length; i++{
					if entry == i{
        					numb := progresschanger.RequestChange(list.Datalist[i].Progress)
        					list.Datalist[i].Progress = numb
        					jsonconvert.Convert(list)
						for{
							fmt.Println("Do you want to change the ProjectName?")
							request := input.StringEntry()
							if request == "yes" || request == "y"{
								fmt.Println("Now give it a new name.")
								entryStr1 := input.StringEntry()
								list.Datalist[i].ProjectName = entryStr1
								jsonconvert.Convert(list)
								addProjectMember(i)
								editProjectMembers(i)
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
	}else{
		fmt.Printf("Error: There are no Datalist instances.\n")
		os.Exit(0)
	}
}
func CreateProject(){
	fmt.Println("What name do you want to give your project?")
	entryStr := input.StringEntry()
	var members []string; 
	for {
		fmt.Println("Enter a name of a Project Member, (no) - to stop adding.")
		entryStr1 := input.StringEntry()
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
	length := len(list.Datalist)
	if length > 0{
	fmt.Println("Which instance do you want to delete?")
	for i:=0;i<length; i++{
		fmt.Printf("%d %s\n",i,list.Datalist[i].ProjectName)	
	}
	for{
		entry := input.IntEntry()
		if entry >= 0 && entry <= length{
			removeInstance(entry)
			break
		}else{
			fmt.Printf("Please enter a valid input.\n")
		}
	}
	}else{
		fmt.Printf("Error: Datalist doesn't have any instances.\n")
	}
}

func AskScanner(){
	for{
		fmt.Println("\nChoose your Option: [V] - view project lists, [E] - edit project lists, [Exit] - exit the program, [A] - add tasks to each project, [M] - markdown completed tasks, [C] - create new Project, [D] - Delete Project")
		fmt.Printf("Enter your input: ")
		answer := input.StringEntry()
		if answer == "V"{
			viewLists() 
		}else if answer == "E"{
			editLists()
		}else if answer == "A"{
			taskmanager.TasksCreator()
		}else if answer == "M" {
			taskmanager.TaskMarkdown()
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
