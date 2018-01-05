package FileEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
    //"bytes"
    m_rand "math/rand"
	"strings"
    "time"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
var plaintext []byte

var block cipher.Block
var key []byte

// Ext is the encrypted appended extension
var Ext = ".enc"

func main() {
	return
}

// InitializeBlock Sets up the encription with a key
func InitializeBlock(myKey []byte) {
	key = myKey
	block, _ = aes.NewCipher(key)

}
func initIV() (stream cipher.Stream, iv []byte) {
	iv = make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	stream = cipher.NewCTR(block, iv[:])
	return stream, iv
}
func initWithIV(myIv []byte) cipher.Stream {
	return cipher.NewCTR(block, myIv[:])
}

var src = m_rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}
// Decrypter decryps a file given its filepath
func Decrypter(path string) (err error, result string) {
	if block == nil {
		return errors.New("Need to Initialize Block first. Call: InitializeBlock(myKey []byte)"),""
	}

	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
		return err,""
	}
    thisDir := "/tmp/"+ "." + randStringBytesMaskImprSrc(32)
    if err = os.MkdirAll(thisDir, 0500); err != nil {
		fmt.Println("error:", err)
		return err,""
    }
	deobfPath := thisDir + "/" + "." + randStringBytesMaskImprSrc(64)
	outFile, err := os.OpenFile(deobfPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err, ""
	}

	iv := make([]byte, aes.BlockSize)
	io.ReadFull(inFile, iv[:])
	stream := initWithIV(iv)
	inFile.Seek(aes.BlockSize, 0) // Read after the IV

	reader := &cipher.StreamReader{S: stream, R: inFile}
   // buf := bytes.NewBuffer(nil)
   // if _, err = io.Copy(buf, reader); err != nil {
   //     return err, ""
   // }
	if _, err = io.Copy(outFile, reader); err != nil {
		fmt.Println(err)
	}
	inFile.Close()
	outFile.Close()

//	os.Remove(path)
	return nil, deobfPath
}

// Encrypter encrypts a file given its filepatth
func Encrypter(path string) (err error) {
	if block == nil {
		return errors.New("Need to Initialize Block first. Call: InitializeBlock(myKey []byte)")
	}

	inFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	obfuscatePath := FilenameObfuscator(path)
	outFile, err := os.OpenFile(obfuscatePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	fmt.Println(outFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}

	stream, iv := initIV()
	outFile.Write(iv)
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err = io.Copy(writer, inFile); err != nil {
		fmt.Println(err.Error())
	}
	inFile.Close()
	outFile.Close()
//	os.Remove(path)
	return nil
}

func FilenameObfuscator(path string) string {
	filenameArr := strings.Split(path, string(os.PathSeparator))
	filename := filenameArr[len(filenameArr)-1]
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))

	return path2 + string(os.PathSeparator) + filename + Ext

}
func FilenameDeobfuscator(path string) string {
	//get the path for the output
	opPath := strings.Trim(path, Ext)
	// Divide filepath
	filenameArr := strings.Split(opPath, string(os.PathSeparator))
	//Get  filename
	filename := filenameArr[len(filenameArr)-1]
	// get parent dir
	path2 := strings.Join(filenameArr[:len(filenameArr)-1], string(os.PathSeparator))
	return path2 + string(os.PathSeparator) + filename
}
