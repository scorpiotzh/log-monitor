package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//当前区块下交易交易的hash
var txHashList = []string{
	"1af24e4e49acc234785a8a66a6c3e0cfb234489d6b29d5fd2a995321d9f35965",
	"bb6336875dd11992b530eec192dee059d96ddc6cd772a68480e3c8195f8d34d0",
	"d80525f2a62c8048657b00f3fe4bc08ecba5327a333e0717c13f6ec5fa48087e",
	"492bc7b02ec8de99a19a607b771fca968a4f8d0ee09d2659e18e8f1c43b9e0fd",
	"39d92827898375dd5ba27084be5bbd99d33284486664f68517bdfed48a9e4610",
	"c5db88a606df8015fd7869dd2e6ebca83e085e443385a0a9b7cb16380cc93327",
	"ec718c070f9459b3958897024d3dafa4680ab7d370db427f1223f7086af31ec7",
	"539f76e53da25775d56c5dcfc4bded7ed6ae4b48c34e2d558ea0aede9c1d5c53",
	"213e027865fa1befbe67db8f307ef00e17f8864131056ebad5026206247a09f8",
}

func main1() {
	//doMerkle() //0949d7f3c7021b60d6eacbdfaf9197a147526bc43c03ebc4971c9e31a3bef38f
	root := GeneraMerkleRoot(txHashList)
	fmt.Println("root:", root)
}

type MerkleNode struct {
	LeftNode  *MerkleNode
	RightNode *MerkleNode
	Data      []byte //保存当前节点哈希
}

type MerkleTree struct {
	Node *MerkleNode
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {

	mNode := new(MerkleNode)
	mNode.LeftNode = left
	mNode.RightNode = right
	//叶子节点
	if left == nil && right == nil {
		mNode.Data = data
	} else {
		//	对左右两侧分支节点的hash进行双哈希
		hashValue := append(left.Data, right.Data...)
		hashFirst := sha256.Sum256(hashValue)
		hashDouble := sha256.Sum256(hashFirst[:])
		mNode.Data = hashDouble[:]
	}
	return mNode
}

func NewMerkleTree(dataList [][]byte) *MerkleTree {
	//包含整个树上的节点的容器
	var nodes []MerkleNode
	//生成所有的叶子节点
	for _, data := range dataList {
		node := NewMerkleNode(nil, nil, data)
		nodes = append(nodes, *node)
	}

	j := 0
	//生成分支节点
	for nSize := len(dataList); nSize > 1; nSize = (nSize + 1) / 2 {
		//进行两两分组
		//i是左侧分支节点的索引。因为两个一组哈希 所以 i+=2
		for i := 0; i < nSize; i += 2 {
			//ii是跟i配套，凑成一组右侧分支节点的索引
			ii := min(i+1, nSize-1)

			node := NewMerkleNode(&nodes[j+i], &nodes[j+ii], nil)
			nodes = append(nodes, *node)
		}
		j += nSize
	}
	return &MerkleTree{&(nodes[len(nodes)-1])}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func GeneraMerkleRoot(txlist []string) string {
	txSlice := [][]byte{}
	for _, value := range txlist {
		txSlice = append(txSlice, ReversHexStringToBytes(value))
	}
	//将二维数组作为参数,通过NewMerkleTree()函数进行两两哈希处理
	hashedBytes := NewMerkleTree(txSlice).Node.Data
	//大小端颠倒后转为字符串
	return ReverseBytesToString(hashedBytes)
}

/**
字节数组大小端颠倒
*/
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

/**
将16进制字符串进行大小端颠倒
*/
func ReversHexStringToBytes(hexString string) []byte {
	bytes, _ := hex.DecodeString(hexString)
	ReverseBytes(bytes)
	return bytes
}

/**
字节数组大端和小端进行颠倒，转成字符串
*/
func ReverseBytesToString(bytes []byte) string {
	ReverseBytes(bytes)
	return hex.EncodeToString(bytes)
}
