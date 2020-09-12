package main

import(
    "fmt"
)

func fops(args []string){
//    fmt.Printf("len=%d cap=%d %v\n", len(args), cap(args), args)
    if len(args)>=1{
        switch cmd := args[0]; cmd{
        case "help":
            help(args)
        case "linecount":
            linecount(args)
        default:
            fmt.Println("undefined command ", cmd)
        }
    }
}

func help(args []string){
    fmt.Println("test help command")
}

func linecount(args []string){
    fmt.Println("test linecount command")
}

func main(){
    fops([]string{"help"})
    fops([]string{"linecount"})
}