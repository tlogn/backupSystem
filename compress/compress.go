package compress

import (
	"backupSystem/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func LocalCompress(w http.ResponseWriter, r *utils.Request){
	selectCompress := r.CompressPara.IsCompress
	if selectCompress {
		fmt.Fprintf(w, "%v", Compress(r.CompressPara.CompressPath))
	} else {
		fmt.Fprintf(w, "%v", UndoCompress(r.CompressPara.CompressPath))
	}
}

type Table struct {
	Key 	[]byte		`json:"key"`
	Value 	[]string	`json:"value"`
}

type TreeNode struct {
	left, right *TreeNode
	value 		int
	name 		byte
}
type PairList []*TreeNode
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value < p[j].value }

var (
	huffMap map[byte]string
	verseHuffMap map[string]byte
	table	*Table
)

func search(tail string, node *TreeNode) {
	if node.left == nil && node.right == nil {
		huffMap[node.name] = tail
		return
	}
	if node.left != nil {
		search(tail + "0", node.left)
	}
	if node.right != nil {
		search(tail + "1", node.right)
	}
}

func buildTree(srcPath string) ([]byte, error) {
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return nil, err
	}
	mapp := make(map[byte]int)
	for _, by := range file {
		if a, ok := mapp[by]; ok {
			mapp[by] = a + 1
		} else {
			mapp[by] = 1
		}
	}
	var arr PairList
	for i := 0;i <= 255;i++ {
		if a, ok := mapp[byte(i & 0xFF)]; ok {
			tmp := TreeNode{
				left:nil,
				right:nil,
				value:a,
				name:byte(i & 0xFF),
			}
			arr = append(arr, &tmp)
		} else {
			tmp := TreeNode{
				left:nil,
				right:nil,
				value:0,
				name:byte(i & 0xFF),
			}
			arr = append(arr, &tmp)
		}
	}

	sort.Sort(arr)
	for {
		if arr.Len() == 1 {
			break
		}
		tp := TreeNode{
			left:  arr[0],
			right: arr[1],
			value: arr[0].value + arr[1].value,
			name:  0,
		}
		arr = arr[1 : ]
		arr[0] = &tp
		sort.Sort(arr)
	}
	huffMap = make(map[byte]string)
	verseHuffMap = make(map[string]byte)
	now := arr[0]
	search("", now)
	table = &Table{}
	for key, value := range huffMap {
		tmp1 := key
		tmp2 := value
		table.Key = append(table.Key, tmp1)
		table.Value = append(table.Value, tmp2)
		verseHuffMap[tmp2] = tmp1
	}
	return file, nil
}

func Compress(srcPath string) string {
	file, err := buildTree(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}

	size := 0
	output := make([]byte,0)
	var tp byte
	for _, by := range file {
		for i:=0;i < len(huffMap[by]); i++ {
			if size == 0 {
				tp <<= 8
			}
			tp <<= 1
			if huffMap[by][i] == '1' {
				tp |= 0x1
			}
			if size == 7 {
				output = append(output, tp)
			}
			size ++
			size %= 8
		}
	}
	if size != 0 {
		tp <<= 8 - size
		output = append(output, tp)
	}
	fileHead := make([]byte,0)

	byteTree, err := json.Marshal(table)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	lenTree := len(byteTree)
	pathHeadLen := make([]byte, 4)
	pathHeadLen[3] =  (byte) (lenTree>>24) & 0xFF
	pathHeadLen[2] =  (byte) (lenTree>>16) & 0xFF
	pathHeadLen[1] =  (byte) (lenTree>>8) & 0xFF
	pathHeadLen[0] =  (byte) (lenTree) & 0xFF

	fileHead = append(fileHead, pathHeadLen...)
	fileHead = append(fileHead, byteTree...)
	fileHead = append(fileHead, byte(size))
	fileHead = append(fileHead, output...)

	dstPath := srcPath + ".ylx"
	ioutil.WriteFile(dstPath, fileHead, 0777)

	return utils.SucceedResponse()
}

func UndoCompress(srcPath string) string {
	if srcPath[len(srcPath) - 4 : ] != ".ylx" {
		return utils.ErrorResponse(errors.New("not compressed file"))
	}
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}

	pathHeadLen := file[: 4]
	lenTree := int(pathHeadLen[0]) + (int(pathHeadLen[1]) << 8) + (int(pathHeadLen[2]) << 16) + (int(pathHeadLen[3]) << 24)
	byteTree := file[4 : 4 + lenTree]
	table = &Table{}
	err = json.Unmarshal(byteTree, table)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	huffMap = make(map[byte]string)
	verseHuffMap = make(map[string]byte)
	for i, _ := range table.Key {
		key := table.Key[i]
		value := table.Value[i]
		huffMap[key] = value
		verseHuffMap[value] = key
	}
	file = file[4 + lenTree :]

	size := int(file[0])
	trueInput := file[1:]
	output := make([]byte, 0)
	var tp string
	for idx, by :=range trueInput {
		if idx == len(trueInput) - 1 {
			for i:=7;i>=8-size;i-- {
				if by & (0x1 << i) != 0 {
					tp = tp + "1"
				} else {
					tp = tp + "0"
				}
				if a, ok := verseHuffMap[tp];ok {
					output = append(output, a)
					tp = ""
				}
			}
		} else {
			for i:=7;i>=0;i-- {
				if by & (0x1 << i) != 0 {
					tp = tp + "1"
				} else {
					tp = tp + "0"
				}
				if a,ok := verseHuffMap[tp];ok {
					output = append(output, a)
					tp = ""
				}
			}
		}
	}
	err = ioutil.WriteFile(srcPath[ : len(srcPath) - 4], output,0777)
	if err != nil {
		return utils.ErrorResponse(err)
	}
	return utils.SucceedResponse()
}

