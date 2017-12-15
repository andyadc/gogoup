package main

/*
scrypt方案，scrypt是由著名的FreeBSD黑客Colin Percival为他的备份服务Tarsnap开发的。
*/
import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
)

func main() {
	dk, err := scrypt.Key([]byte("andyadc"), []byte(`salt`), 16384, 8, 1, 32)
	if err != nil {
		log.Fatal("scrypt error", err)
	}
	fmt.Println(dk)
	fmt.Println(string(dk))
}
