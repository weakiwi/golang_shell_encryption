package main

import (
    "fmt"
    "io"
    "os"
    "io/ioutil"
    "os/exec"
    "time"
    "strings"
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
    subProcess := exec.Command("bash", os.Args[1], "111", strings.Join(os.Args[2:]," ")) //Just for testing, replace with your subProcess

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
    go func() {fmt.Println("start remove")
    time.Sleep(5)
    if err = os.Remove(os.Args[1]); err != nil { //Use start, not run
        fmt.Println("An error occured: ", err) //replace with logger, or anything you want
    }
    fmt.Println("finish remove")}()
    subProcess.Wait()
    io.WriteString(stdin, "4\n")
    fmt.Println("END") //for debug
}
