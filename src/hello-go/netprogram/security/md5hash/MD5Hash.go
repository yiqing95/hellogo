package main

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	fmt.Printf("hash value : %x \n ", hashValue)
	hashSize := hash.Size()
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])

		fmt.Printf("%x ", val)
	}
	fmt.Println()

	// hmac
	hash2 := hmac.New(md5.New, []byte("myKey"))
	hash2.Write([]byte("this is the massage"))
	hash2Value := hash2.Sum(nil)
 
	// 十六进制，小写字母，每字节两个字符 
	fmt.Printf("keyed hash value : %x \n ", hash2Value)
	fmt.Println("hmac is ", hash2Value)
}
