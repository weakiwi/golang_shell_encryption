package main
import (
    "github.com/Tiked/FileEncryption"
    "flag"
)

func main() {
  FileEncryption.InitializeBlock([]byte("a very very very very secret key"))
  encryption_file := flag.String("e", "/entrypoint.sh","file to encryption")
  flag.Parse()
  err := FileEncryption.Encrypter(*encryption_file)
  if err != nil {
    panic(err.Error())
  }
  
}
