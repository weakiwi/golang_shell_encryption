package main
import (
    "github.com/Tiked/FileEncryption"
    "os"
)

func main() {
  if os.Getenv("ENTRYPOINT_PUBLIC_KEY") != "" {
      FileEncryption.InitializeBlock([]byte(os.Getenv("ENTRYPOINT_PUBLIC_KEY")))
  } else {
      FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
  }
  err := FileEncryption.Encrypter(os.Args[1])
  if err != nil {
    panic(err.Error())
  }
  
}
