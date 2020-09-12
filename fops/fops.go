package main

import(
    "fmt"
    "os"
    "io"
    "bytes"
)

func fops(args []string){
//    fmt.Printf("len=%d cap=%d %v\n", len(args), cap(args), args)
    if len(args)>=1{
        switch cmd := args[0]; cmd{
        case "help":
            cmdhelp(args)
        case "linecount":
            fmt.Println(cmdlinecount(args))
        default:
            fmt.Println("undefined command ", cmd)
        }
    }
}

func cmdhelp(args []string){
    fmt.Println("test help command")
}

func cmdlinecount(args []string)string{
    // TODO error handle
    var file *os.File
    switch args[1]{
    case "-f":
        f, err := os.Open(args[2])
        _ = err
        file = f
    case "--file":
        f, err := os.Open(args[2])
        file = f
        _ = err
    }
    count, error := linecount(file)
    _ = error
    return fmt.Sprint(count)
//    return "undefined error"
}

func linecount(flie io.Reader)(int, error){
    buf := make([]byte, 32*1024)
    result := 0
    lineSep := []byte{'\n'}
    
    for {
        count, err := flie.Read(buf)
        result += bytes.Count(buf[:count], lineSep)
        switch {
        case err == io.EOF:
            return result, nil
        case err != nil:
            return result, err
        }
    }
}

func main(){
    fops([]string{"help"})
    fops([]string{"linecount", "-f", "myfile.txt"})
}