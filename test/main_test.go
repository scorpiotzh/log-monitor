package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestReverseBytes(t *testing.T) {
	//hash:="1af24e4e49acc234785a8a66a6c3e0cfb234489d6b29d5fd2a995321d9f35965"
	//bytes,_ := hex.DecodeString(hash)
	//fmt.Println(len(hash),len(bytes))
	//ReverseBytes(bytes)
	//fmt.Println(hex.EncodeToString(bytes))
	hashValue := append([]byte("0c6dfe3e67c45cd437c8e690f83532fbd7e356433f83c3521c38173008536806"), []byte("67b9098cbceedcb27d12d431b192b482f14ba463e7c9171b5594d80e9614661e")...)
	hashFirst := sha256.Sum256(hashValue)
	hashDouble := sha256.Sum256(hashFirst[:])
	fmt.Println(hex.EncodeToString(hashDouble[:]))
}
