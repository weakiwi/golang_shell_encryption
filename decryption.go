package main

import (
    "fmt"
    "io"
    "os"
    "github.com/Tiked/FileEncryption"
    "io/ioutil"
    "os/exec"
    "flag"
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
    decryption_file := flag.String("e","/entrypoint.sh.enc","file path to decryption")
    flag.Parse()
    FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
    err := FileEncryption.Decrypter(*decryption_file)
    if err != nil {
      panic(err.Error())
    }
    origData := read3(FileEncryption.FilenameDeobfuscator(*decryption_file))
    err = os.Remove(FileEncryption.FilenameDeobfuscator(*decryption_file))
    if err != nil {
      panic(err.Error())
    }
    subProcess := exec.Command("bash", "-c", string(origData)) //Just for testing, replace with your subProcess

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
    fmt.Println(string(origData))
}
