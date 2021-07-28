package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	//hash()
	//doMerkle() //ea5576587697679a35181eff7912cddb66e13f1c37c79ff21a0eb99f3d25266f
	//doMerkle2()
}

func hash() {
	hashList := []string{
		"066853083017381c52c3833f4356e3d7fb3235f890e6c837d45cc4673efe6d0c",
		"07c5bf0a51de6235c6931045ad135ffc0123be0339f202dafb0ccdaa68bfda5c",
		"0e6763bf14b80a6e2f3dd6056f4c4c73a3f80328cfe9692cac86a27777a1f42a",
		"12327344182b6cdc3d31bba5e204646d2b962cf516e6483d9e88a404734bcff8",
		"156bf3fa07145a4616009c277320b0cf3b8771c4d0bdac69f72b44ec2dfc454f",
		"15f9e872bc45ae7a54dbe518d7448241f514ba4cc347511db3ce21aac28792c6",
		"1bf9ad7ce49adf6cbc707a689b6e17653151e95c1cd8a53f9fce54d3d51a2a24",
		"1e6614960ed894551b17c9e763a44bf182b492b131d4127db2dceebc8c09b967",
		"203a19ad813e9ae88569b09295a578a3b6e723b3e4afbdc93de5290b43a1e005",
		"22ced13392ccb51bb8cc78fbea9dc7bacaa43e4aceb67012d719a68b29a34211",
		"2695b78fff5dd0a5c65ca47b7c0599fb13813e0ee2d87d875f513eff4f571cc8",
		"279188dfed379880ee821f06260d7cc640a9658eecec0efeb977c115f664931e",
		"2b899cdfeddfa8341b2c66d838275380e0bf1724ee5ff92f0404cd691e028566",
		"2b9e35ef30e35516314d21a3e15539851a5613bb803708df6c37b523a9b8759d",
		"2f7ff8365335d36a965554f8e402ffb175c1bd135d98e2ad9cee274f5fe7330e",
		"3583d93c06e6b73b3b1215dd7d39d6042bfc5344bbf4c2413daaf125a591e730",
		"386c959fea351f35ac2b475ec9cac3d77437c39484fea0d593334026cb412b3f",
		"3c05691bd6d76a08867e80fa0533b7be2a12f65fca80e50a4175a37f7d008c0a",
		"3c7473fe808421b72a4c67da8db407f89a3dc284e3c53d05ca8b8bdd6e953659",
		"418b372107320cfb5a2c2e4dd9000f364670e10ea02872a218b91b4053e2ab76",
		"429300b6d8ace0ff48ab18669b6b36ae4e085ebb3d7f4b2c558cb23e3053cdf3",
		"441b8393ad5081aa47e14d413ae8fc07b3814078b1b95759210cc2b8e6d73cc6",
		"44a2bffee7c46c4b5fc11a4b4c929280638729d97a5a09d2fcea3fe1c46f0926",
		"478078e32d811fe03275c6fe0c1a31087fa02d6550dfa83884a32d46dfa146b0",
		"485f31128be9378e6bbb0bebe418f33eb3890c13d91226ced358c0925b48565a",
		"4ea169740e3d5ba348d0b50564e36dbe0568d1b9cbadc27613a28149237b2492",
		"50ec12943456048dbac348227e2f82406fb40583f0402e633fe3dd1d0421409f",
		"549eba6b2d89bca8ec143f60cb563d024785ffd49ab7efd3d45551307e9fc803",
		"589d40b55b6ec899970a5e5dc625e47d05632af87e62d222db1b3e37f21cf872",
		"589d40b55b6ec899970a5e5dc625e47d05632af87e62d222db1b3e37f21cf872",
		"5e6c43ff7fe7efa4dcf826901de573dbe847d63083a90d0f54e83795cc8cd483",
		"60e754f51ade6182c600c77c1011861e9d10d258cfdd7a93941763c6fb2c0520",
		"629feecc6e6afdb88c2f1ab18ed5fa309ef27fd798829452ef124d32744b3330",
		"62c9ffffcec2e74fad77d2c3212e3f5d2076b32daba9cd773296daedc547338f",
		"6357f24548303a95eb0838509f350b90495c3bf80e01606db7fd468e9fbf548f",
		"6768bbea607986c956d39ec5854d565177d5da8342ae40ad0983cd1b53f59e1e",
		"6a1363943c2bf8fe89cfee80d234786a8de33d13e3ea959981bb3d2b3f622ccc",
		"6e9b79e7af1886cb5a63e3a3bff295bab22d2441ed097ecb589b2907e6d6d257",
		"6fb719c83b1cdff635944490f7af4a1e3ea27e863cf9ec839f8c1f03a42d7120",
		"7121a15f18b3b551697fbf0b67a889468b8a5ca788e352e18563b4fcd9e38292",
		"71c2382f614a465bc00de1023e5bccbb7b9cd0074579bbfe00b7e3e24a94f901",
		"71d4502ad50e54da30d39fdd869ff3aa932e85bb36c414178d15cafeca2d2dda",
		"761ccdfcbb8f3284756cc27b27d34909a81d1c495274ad2c291cc41b1304d1b4",
		"8060aa268d477f4181dcb06a47d0d33564e4f0b4eb22cd756df12f16abdd0d95",
		"85b614b442a1999de775dd462696c13ba1287cae78f0a988e5601450315852a6",
		"86ef9be0272427d29822ea39e5ca7fc1cf329927cc343f9d61ebfd2107bcd53b",
		"887ac8f7eb2fc2bc7af840a53adc730674b75a269796256509e3df4c5f9c71ea",
		"8a4b0c91871cfe271927a86e3ec7fad209ba03fa90e19edee8a2ff110bc2707b",
		"8a8e4857398c0f184f5bf3df078249e39ff3586bfdf65777d534cb36fe5ce9dc",
		"8ee3739215d0e0cb77dce62e397dab169ab276b12df33044690d2cf3aa39cc7d",
		"9129c8612d407265f1b559be8dd3ddada682c90a283903793c18f9079fc0426a",
		"957c8760e5cb94cb48641e3689781f54f85b352799cd88e5d8e3da462208c6c6",
		"98f0cbbf4f648a17ae06456fcb72a390c50d043badc5e57053223393c1b4e516",
		"992184d0788794fdc774837a175e1d8e6a1aa7914efdf10715fd70dd9372b664",
		"9b7319ba97333dff53bc595e0a3d3ab829f3807831e82e8b0ee1ad19cc381d11",
		"a2312251528e55096a6f8b7e8d2ac312aaa08b17e2911c6a58e0b4a7008a2b1b",
		"a25f65b586ce1773ed5f2d9c4f9dc28eafaa8d56b2ca4cdb021887abe30889eb",
		"a370f95c54bcc9241493ed8a679dc7db820b264a62857fce2315e2669b9838c9",
		"a6c51ebe5474ccdef12c2134da36def02b76974a7132c05d4cf914b2e742ae30",
		"a9089731ea91b294dbfe9bb5a3a1b88e9f665f0a7ceff649692eb8145d65c539",
		"ab5cfe23377451ef13ba41dd18ef1bcc6fd61eb19264334435dd71f85bc66926",
		"adc6593ac0679348497a7d7465d4d503a14ba4ae92a5b90c9c72badd302238f7",
		"b5812ed9576750d21a393c0eaa7510814788673d24c64ba62aaf9699ce72e655",
		"bb8f39c9c89ca424cc05b488a34eab01e8678dbef090b426abc7e3d9700d456e",
		"bc20334bbaf8554e2c1edf9ad6e53acdaee7d9b3e6df82c7ec3092b50cad97ac",
		"bd498fcb96c97e946ea501520c202a2c5ec285cf0f739e1fcc5587fefcaccfa0",
		"bf28f9e56e65c16767c7206fa531f5f99ba69163bffdddd0acab1ed8e6c36d5d",
		"bfde61099c5ac9d2ce5a14278f675048b61ded2f12b9568b381c0d8a258ff6ce",
		"c2abece209cec54cac10387a63abad2a89802311a2ed6459720fea603e2505b2",
		"c5445ddf133cb6c6e50a4988273627c1aef90569e2a665ad1283a1cc5834e2c3",
		"ca591deb13370acb9ca71ec9b9affa5cfe05803410ba0749c9cd1bf4bcbb32d9",
		"cbc7a0d750a05d6c24840b92648ce0347efa2ba6f66a082a8786b335f6ad06e0",
		"cbe6618d7a158fe4365a46e3865a8c564ab7ff45691c86498e8744f1a16a21c1",
		"d035ce85a7ba7fcb5d8e7f870ffe0a3706eb1230d4e760bd69e9e70f717d92fb",
		"d74a8799916dd3e3b4776622b55af31b341e7b36fae4a83a90b374b91ad6300d",
		"d7769ad91f9f34dd3a36014f0dbf949841a49f921e63b1742eac941dba5bb133",
		"d783c49f45acfe4ddfb42c1e905d35d79cab4646abe48e3aa6b8e2ce7d0cb9d5",
		"d9574a6556e0db7f2af8db0b5f23214525e5abcdec494cad64eb2c675558115a",
		"dc3322e5f1c213c045feb4784046a933bf933d675938b511205fca37912b4d0d",
		"e21752ccce01f154a29537f0ee8ca0c9290b3dcac641ed20ac77cc00baf68c20",
		"e6456e3c44ca37178ea51cf29e891bc6e424f1a364454dda080e325f2cc4b77c",
		"e8e340f2e269546451fff9156f4b1d6ba2312328b77cc35a3e6e1f651ae8eb83",
		"eb070edbbc96d91b9e7b422d5860a86ab135737d73ee8207b0bf92821976a562",
		"ec23f3ac47ab283e2e3f912a8f2a51080918d9b7d0a25e34e6a43a4d8598e53f",
		"ec5d968a33a2e2d5f58c9613bd5c4570e73ea4d7887a7ce7a3bbcd320ea54a5b",
		"f06021ca4b54005eb0a0c4f34a3d356afcf07abdce6c4e308b34692b47e144ab",
		"f07a747dd817026f9a7d21bc96e8b7b0580a20aacee4dca4231f9d819855c764",
		"f297354707be0f403fe5d2b9b519a0008b4dd0453f43c7ac81548adef902415f",
		"f298561396961c78a80f29bc50d252682eee7fc3c2c4a936026541095dad54d6",
		"f50c0fc2c6a0bd3bfc3bcb7e11a2069c757888270837a93a522cc68eca6b9479",
		"f5c7940f011ac6e49b8b85b1b6c23c571e0889eda35b5597ec0a9962801c058d",
		"f6807919ba6eef0425f050cae1093d3aa4f4b175c4150a25ba60c08b684dcb20",
		"f88a0723ce7899f9dd5ea9acffe1e8023bcc0066ddcac57385c312d85b960dcc",
		"fd336781a4fee6e71aff1d0c046cfc59fb60cad996614c0683e19397a6bed6fa",
	}
	// 递增排序
	sort.Strings(hashList)
	for _, v := range hashList {
		fmt.Println(v)
	}
	// 构建默克尔树
	resList := getMerkleHash(hashList)
	//res := ReversHexStringToBytes(resList[0])
	fmt.Println("res:", resList[0])
}

func team() {
	str := `timyang🤯
specer
unclejia🥳
指尖下的幽灵
朱龙辉
jeffjing
link😂
汤志鸿
Mancy🙈
kylexiang
Park`
	list := strings.Split(str, "\n")
	for _, v := range list {
		data := sha256.Sum256([]byte(strings.TrimSpace(v)))
		fmt.Println(v, hex.EncodeToString(data[:]))
	}
}

// 数据文件计算得 sha256 hash 值
// 所有文件hash排序后，大小端字节转换
// 两两计算两次 sha256 hash
// 最终结果再大小端字节转换得默克尔根
func doMerkle() {
	hashList := getHashList()
	//for i, v := range hashList {
	//	bytes, _ := hex.DecodeString(v)
	//	ReverseBytes(bytes)
	//	hashList[i] = hex.EncodeToString(bytes)
	//}
	// 递增排序
	hashList = append(hashList, "86ef9be0272427d29822ea39e5ca7fc1cf329927cc343f9d61ebfd2107bcd53b")
	sort.Strings(hashList)
	fmt.Println(len(hashList))
	// 构建默克尔树
	resList := getMerkleHash(hashList)
	//res := ReversHexStringToBytes(resList[0])
	fmt.Println("res:", resList[0])
}

func doMerkle2() {
	hashList := getHashList()
	root := GeneraMerkleRoot(hashList)
	fmt.Println("root:", root)
}

func getHashList() []string {
	// 读取文件
	var hashList []string
	path := "/Users/zhangyikang/Documents/das3/das_data/" //"/Users/zhangyikang/Documents/das3/上链数据/"
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
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
		//hashDouble := sha256.Sum256(hashFirst[:])

		//sum := sha256.Sum256([]byte(list[i] + list[j]))
		res := hex.EncodeToString(hashFirst[:]) //fmt.Sprintf("%x", hashFirst)
		parentList = append(parentList, res)
	}
	return getMerkleHash(parentList)
}
