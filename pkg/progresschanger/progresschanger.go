package progresschanger 

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func progressChanger(x int) (int){
   scanner := bufio.NewScanner(os.Stdin)
		
   for {
        var entry int;
        fmt.Println("\n\nEnter a number greater than the current percentage: \n")
	scanner.Scan()
	entryStr := scanner.Text()
	entry, err := strconv.Atoi(entryStr)
	
        if entry > x && err == nil{
                x = entry;
        	break;  
        } 
        fmt.Println("\n\nEnter a valid entry.")
   }
   return x;	
}


func RequestChange(x int) (int){
        for {
        fmt.Println("\n\nDo you want to change your progress?\n")
        var response string;
        fmt.Scanf("%s", &response)
        if response == "yes"{
                x = progressChanger(x)
                break;
        }else if response == "no"{
                fmt.Println("Exitting")
                os.Exit(x)
        }else{
		fmt.Println("Enter a valid entry.")
	}
        }
        return x;
}

