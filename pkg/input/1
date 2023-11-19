package input 

import(
	"bufio"
	"strconv"
	"fmt"
	"os"
)

func StringEntry() string{
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        return scanner.Text()
}

func IntEntry() int{
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
