package taskmanager 

import( "bufio"
	"fmt"
	"os"
	"strconv"
	"go-project-manager/pkg/mainreader"
	"go-project-manager/pkg/jsonconvert"
)

var MAX_WEIGHT int = 100 

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

func CheckCompletion(entry int){
	length := list.Datalist[entry].Tasklength
	var numbofCompleted int;
	for i:=0; i<length; i++{
		if list.Datalist[entry].Tasklist[i].Completed == true{
			numbofCompleted++
		}
	}
	if length == numbofCompleted{
		list.Datalist[entry].TasksCompleted = true
		jsonconvert.Convert(list)
	}	
}

func ProgressEditer(entry int){
	list.Datalist[entry].Tasklength = len(list.Datalist[entry].Tasklist)
        length := list.Datalist[entry].Tasklength
	var numbofCompleted int; 
        var progress int;
        for i := 0; i<length; i++{
                if list.Datalist[entry].Tasklist[i].Completed == true{
                        progress += list.Datalist[entry].Tasklist[i].Weight
			numbofCompleted++
                }
        }
        list.Datalist[entry].Progress = 0
	if length == numbofCompleted && length > 0{
		list.Datalist[entry].Progress = 100
	}else{
       		list.Datalist[entry].Progress = progress
	}
        jsonconvert.Convert(list)
}


func TaskMarkdown(){
	length := len(list.Datalist)
	// prevents running of panic error
	if length > 0{
        	fmt.Println("Choose the respective index number, in order to mark completion of task.")
        	for i:=0; i < length; i++ {
			fmt.Printf("[%d] %s\n",i,list.Datalist[i].ProjectName)
        	}
		for{
			entry := intEntry()
			CheckCompletion(entry)
			if entry < length {
				var data = list.Datalist[entry]
				// This is the same as list.Datalist[entry].Tasklist
				length0 := len(data.Tasklist)
				if length0 > 0 && list.Datalist[entry].TasksCompleted != true{
					fmt.Println("Choose the Task to mark as completed: ")
					for i := 0; i<length0; i++{
						if data.Tasklist[i].Completed == false{
							fmt.Printf("[%d.%d]  %s [ ]\n", entry, i, data.Tasklist[i].Task)
						}else if data.Tasklist[i].Completed == true{
							fmt.Printf("[%d.%d]  %s [X]\n", entry, i, data.Tasklist[i].Task)
						}
					}
					for{
						entry0 := intEntry()
						if entry0 <= length0 && data.Tasklist[entry0].Completed == false{
							data.Tasklist[entry0].Completed = true
							jsonconvert.Convert(list)
							ProgressEditer(entry)
							break
						}else{
							fmt.Printf("Please enter a valid response.\n")
						}
					}
				}
				break
			}
		}	
	//prevents panic from running
	}else{
		fmt.Printf("Error: There are no instances in your Datalist.\n")
		os.Exit(0)		
	}
}

func WeightedTasks(entry int){
	fmt.Println("How many tasks do you want to add to your Project?")
	var tasks int = intEntry()
	list.Datalist[entry].Tasklength += tasks
	tasklength := list.Datalist[entry].Tasklength
	var weight int = MAX_WEIGHT / tasklength  	
	for i:=0; i<tasks; i++{
		fmt.Println("Enter the task name: ")
		entryStr := stringEntry()
		// Create a new Task instance
    		newTask := mainreader.Task{
        		Task:      entryStr,
        		Weight:    weight,
        		Completed: false,
    		}	

    		// Append the new task to the Tasklist of the specified project
    		list.Datalist[entry].Tasklist = append(list.Datalist[entry].Tasklist, newTask)
	}
	jsonconvert.Convert(list)
	ProgressEditer(entry)
}

func TasksCreator(){
	fmt.Println("Enter the index in which you want to add the Tasklist.")
        length := len(list.Datalist)
	// Prevents StackOverflow from happening if .json file has no Datalist instances
	if length > 0{
        	for i := 0; i<length; i++{
        		fmt.Println(i, "," , list.Datalist[i].ProjectName)
        	}
        	entry := intEntry()
        	for{
           		if entry < length && entry >= 0{
                		WeightedTasks(entry)
                        	break
                	}else{
                		fmt.Println("Please enter a valid entry.")
                	}
        	}
	}else{
		fmt.Printf("Error: There is no item in this Datalist.\n")
		os.Exit(0)
	}
}
