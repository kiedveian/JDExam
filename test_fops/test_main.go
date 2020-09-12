package main

import(
    "fmt"
    "fops"
)

func testCommand(args []string){
    fmt.Println("run commands: ",args)
    fops.Run(args)
}

func main(){
    testCommand([]string{"help"})
    testCommand([]string{"linecount", "-f", "myfile.txt"})
    testCommand([]string{"linecount", "-f", "no_find_file"})
}