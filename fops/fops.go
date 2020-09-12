package fops

import(
    "fmt"
    "os"
    "io"
    "bytes"
)

func Run(args []string){
    if len(args)>=1{
        remain := args[1:]
        switch cmd := args[0];  cmd{
        case "help":
            CmdHelp(remain)
        case "linecount":
            fmt.Println(CmdLineCount(remain))
        default:
            fmt.Println("undefined command ", cmd)
        }
    }
}

func CmdHelp(args []string){
    fmt.Println("test help command")
    fmt.Println("args: " ,args)
}

func CmdLineCount(args []string)string{
    switch args[0]{
    case "-f", "--file":
        file, err := os.Open(args[1])
        if err != nil{
            return fmt.Sprint(err)
        }
        count, err := linecount(file)
        if err != nil{
            return fmt.Sprint(err)
        }
        return fmt.Sprint(count)
    default:
        return "undefined error"
    }
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
