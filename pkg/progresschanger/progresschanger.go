package progresschanger 

import(
	"fmt"
	"os"
)

func progressChanger(x int) (int){
   for {
        var entry int;
        fmt.Println("\n\nEnter a number greater than the current percentage: \n")
        fmt.Scanf("%d", &entry)
        if entry > x {
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

