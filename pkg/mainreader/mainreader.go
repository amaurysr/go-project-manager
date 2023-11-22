package mainreader

import (
	"fmt"
        "os"
	"go-project-manager/pkg/input"
        "log"
	"encoding/json"
	"slices"
)
// This is our .json file for reading and writing purposes
var File string
// This is for tracking our one time event
var Count int = 0

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

type Memlist struct{
	Filenames []string `json:"Filenames"`
}

func CreatorLoop(files Memlist) (int, Memlist){
	var entry int
	fmt.Printf("Give a name to the new .json file: \n")
        name := input.StringEntry()
        filename := fmt.Sprintf("%v.json", name)
       	files.Filenames = append(files.Filenames, filename)
	os.Create(filename)
        newF, err1 := json.Marshal(files)
        if err1 != nil{
        	log.Fatal(err1)
        }
       	err2 := os.WriteFile("filememory.json", newF, 0664)
        if err2 != nil {
        	log.Fatal(err2)
        }
	entry = len(files.Filenames) - 1
	return entry, files
}

func CurrentLoop(files Memlist) (int, Memlist){
	var entry int;
	if len(files.Filenames) > 0{
                fmt.Printf("Enter a File of your choice to edit: \n")
                for i:=0; i<len(files.Filenames); i++{
                        fmt.Printf("[%d] %v\n",i,files.Filenames[i])
                }
                for{
			fmt.Printf("Enter your choice: ")
                        entry = input.IntEntry()
                        if entry >= 0 && entry < len(files.Filenames){
                                break
                        }
                }
        }else{
		fmt.Printf("There are no files that exist, so now create a .json file.\n")
                entry, files = CreatorLoop(files)
        }
	return entry, files
}

func DeleteJSON(files Memlist) Memlist{
	length := len(files.Filenames)
	fmt.Printf("Please choose any of these files to Delete\n")
	for i:=0;i<length;i++{
		fmt.Printf("%d %v\n", i, files.Filenames[i])
	}
	for{
		entry := input.IntEntry()
		if entry >= 0 && entry < length{
			os.Remove(files.Filenames[entry])
			files.Filenames = slices.Delete(files.Filenames, entry, entry + 1)
        		newF, err1 := json.Marshal(files)
        		if err1 != nil{
                		log.Fatal(err1)
        		}
        		err2 := os.WriteFile("filememory.json", newF, 0664)
        		if err2 != nil {
                		log.Fatal(err2)
        		}
			break
		}else{
			fmt.Printf("Enter a valid input.")
		}
	}
	return files
}

func MemoryLoop() string{
	content, err := os.ReadFile("filememory.json")
	if err != nil{
		log.Fatal(err)
	}
	var files Memlist
	var entry int
	err0 := json.Unmarshal(content, &files)
	if err0 != nil{
		log.Fatal(err0)
	}
	for{
		fmt.Printf("Do you want to [C] - Create a new .JSON file, [U] - Use an existing .JSON file, [D] - Delete an existing .JSON file\n")
		fmt.Printf("Enter your choice: ")
		entryStr := input.StringEntry()
		if entryStr == "C" || entryStr == "c"{
			// insert func here
			entry, files = CreatorLoop(files) 
			break
		}else if entryStr == "U" || entryStr == "u"{
			entry, files = CurrentLoop(files)
			break
		}else if entryStr == "D" || entryStr == "d"{ 
			files = DeleteJSON(files)	
		}else{
			fmt.Printf("Enter a valid response.\n")
		}
	}
	return files.Filenames[entry]
}

func Listpaste() Datalist{
	// When this function is called for the first time, this if statement executes and global var FILE  maintains the file you've chosen
	if Count < 1{
		File = MemoryLoop()
		Count++ 
	}
	content, err := os.ReadFile(File)
        if err != nil{
		log.Fatal(err)
        }
	var list Datalist
        err0 := json.Unmarshal(content, &list)
        if err0 != nil{
		if len(list.Datalist) > 0{
			fmt.Printf("Error: While Unmarshalling\n")
			os.Exit(0)
		}
        }
	return list
}
