package main
import (
    "github.com/Tiked/FileEncryption"
    "os"
)

func main() {
  FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
  err := FileEncryption.Encrypter(os.Args[1])
  if err != nil {
    panic(err.Error())
  }
  
}
