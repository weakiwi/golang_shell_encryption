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
    // fmt.Println(string(fd))  
    return string(fd)  
}

func main() {
    FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
    err, ShellStrings := FileEncryption.Decrypter(os.Args[1])
    if err != nil {
      panic(err.Error())
    }
    subProcess := exec.Command("bash", "-c", ShellStrings) //Just for testing, replace with your subProcess

    stdin, err := subProcess.StdinPipe()
    if err != nil {
        fmt.Println(err) //replace with logger, or anything you want
    }
    defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

    subProcess.Stdout = os.Stdout
    subProcess.Stderr = os.Stderr

    fmt.Println("START") //for debug
    if err = subProcess.Start(); err != nil { //Use start, not run
        fmt.Println("An error occured: ", err) //replace with logger, or anything you want
    }

    io.WriteString(stdin, "4\n")
    subProcess.Wait()
    fmt.Println("END") //for debug
}
