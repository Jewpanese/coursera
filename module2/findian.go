package main

import (
        "fmt"
        "strings"
        "os"
        "bufio"
)

func main() {


                fmt.Println("Please enter the string: ")


                scanner := bufio.NewScanner(os.Stdin)
                scanner.Scan()
                userInput := scanner.Text()

                str:=strings.ToLower(userInput)

                indexI := strings.IndexByte(str,'i')
                indexA := strings.IndexByte(str,'a')
                indexN := strings.IndexByte(str,'n')

                if indexI == 0 && indexA > 0 &&  indexN == len(userInput)-1 {
                fmt.Println("Found!")
                } else {
                        fmt.Println("Not Found!")
                }

}
