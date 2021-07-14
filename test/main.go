package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	doMerkle()
	doMerkle2()
}

// 数据文件计算得 sha256 hash 值
// 所有文件hash排序后，大小端字节转换
// 两两计算两次 sha256 hash
// 最终结果再大小端字节转换得默克尔根
func doMerkle() {
	hashList := getHashList()
	for i, v := range hashList {
		bytes, _ := hex.DecodeString(v)
		ReverseBytes(bytes)
		hashList[i] = hex.EncodeToString(bytes)
	}
	// 构建默克尔树
	resList := getMerkleHash(hashList)
	res := ReversHexStringToBytes(resList[0])
	fmt.Println("res:", hex.EncodeToString(res))
}

func doMerkle2() {
	hashList := getHashList()
	root := GeneraMerkleRoot(hashList)
	fmt.Println("root:", root)
}

func getHashList() []string {
	// 读取文件
	path := "/Users/zhangyikang/Downloads/部分写入数据/"
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var hashList []string
	for _, v := range fileList {
		if v.Name() == ".DS_Store" {
			continue
		}
		if !v.IsDir() {
			if hash, err := fileHash(path, v.Name()); err != nil {
				panic(err)
			} else {
				hashList = append(hashList, hash)
			}
		}
	}
	// 递增排序
	sort.Strings(hashList)
	return hashList
}

// 计算文件hash
func fileHash(path, name string) (string, error) {
	file, err := os.Open(path + name)
	if err != nil {
		return "", fmt.Errorf("open err:%s [%s]", err.Error(), name)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("copy err:%s [%s]", err.Error(), name)
	}
	sum := hash.Sum(nil)
	res := fmt.Sprintf("%x", sum)
	fmt.Printf("%s [%s]\n", res, name)
	return res, nil
}

// 递归计算默克尔根
func getMerkleHash(list []string) []string {
	if len(list) <= 1 {
		return list
	}
	var parentList []string
	for i, j := 0, 1; i < len(list); i, j = i+2, j+2 {
		if j == len(list) {
			j = i
		}
		//fmt.Println(i, j, len(list))
		left, _ := hex.DecodeString(list[i])
		right, _ := hex.DecodeString(list[j])
		hashValue := append(left, right...)
		hashFirst := sha256.Sum256(hashValue)
		hashDouble := sha256.Sum256(hashFirst[:])

		//sum := sha256.Sum256([]byte(list[i] + list[j]))
		res := fmt.Sprintf("%x", hashDouble)
		parentList = append(parentList, res)
	}
	return getMerkleHash(parentList)
}
