package main

import (
    "fmt"
    "io"
    "os"
    "github.com/Tiked/FileEncryption"
    "io/ioutil"
    "os/exec"
)

func read3(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    fd,err := ioutil.ReadAll(fi)  
    return string(fd)  
}

func main() {
    if os.Getenv("ENTRYPOINT_PUBLIC_KEY") != "" {
        FileEncryption.InitializeBlock([]byte(os.Getenv("ENTRYPOINT_PUBLIC_KEY")))
    } else {
        FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
    }
    err, ShellPath := FileEncryption.Decrypter(os.Args[1])
    if err != nil {
      panic(err.Error())
    }
    var ss []string;
    ss = append(ss,ShellPath)
    for i:=1;i<len(os.Args);i++{
        ss = append(ss,os.Args[i])
    }
    subProcess := exec.Command("bash") 

    stdin, err := subProcess.StdinPipe()
    if err != nil {
        fmt.Println(err) 
    }
    defer stdin.Close() 

    subProcess.Stdout = os.Stdout
    subProcess.Path = ShellPath
    subProcess.Args = os.Args[1:]
    subProcess.Stderr = os.Stderr

    fmt.Println("START")
    if err = subProcess.Start(); err != nil {
        fmt.Println("An error occured: ", err)
    }

    io.WriteString(stdin, "4\n")
    subProcess.Wait()
    fmt.Println("END")
    if err = os.Remove(ShellPath); err != nil {
        fmt.Println("An error occured: ", err)
    }
}
