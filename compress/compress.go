package compress

import (
	"backupSystem/utils"
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"io/ioutil"
	//"path/filepath"
	"sort"
)

func LocalCompress(w http.ResponseWriter, r *utils.Request){
	comOrUncom := r.CompressPara.IsCompress
	if comOrUncom {
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
	left,right * TreeNode
	value int
	name byte
}
type PairList []*TreeNode

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value < p[j].value }
var hafmap map[byte]string
var versehafmap map[string]byte

func search(tail string, node *TreeNode) {
	if node.left == nil && node.right == nil {
		hafmap[node.name] = tail
		return
	}
	if node.left != nil {
		search(tail + "0", node.left)
	}
	if node.right != nil {
		search(tail + "1", node.right)
	}
}
func buildTree(srcPath string) error {
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return err
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
		if a, ok := mapp[byte(i&0xFF)]; ok {
			tmp := TreeNode{
				left:nil,
				right:nil,
				value:a,
				name:byte(i&0xFF),
			}
			arr = append(arr, &tmp)
		} else {
			tmp := TreeNode{
				left:nil,
				right:nil,
				value:0,
				name:byte(i&0xFF),
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
	hafmap = make(map[byte]string)
	now := arr[0]
	search("",now)
	table := Table{}
	for key,value := range hafmap {
		tmp1 := key
		tmp2 := value
		table.Key = append(table.Key, tmp1)
		table.Value = append(table.Value, tmp2)
	}
	bytes, err := json.Marshal(table)

	ioutil.WriteFile("tree",bytes,0777)
	return nil
}

func readTree() {
	inputTree, _ := ioutil.ReadFile("tree")
	table := Table{}
	hafmap = make(map[byte]string)
	_ = json.Unmarshal(inputTree, &table)
	for idx, key := range table.Key {
		hafmap[key] = table.Value[idx]
	}
}
func readInverseTree() {
	inputTree, _ := ioutil.ReadFile("tree")
	table := Table{}
	versehafmap = make(map[string]byte)
	_ = json.Unmarshal(inputTree, &table)
	for idx, key := range table.Key {
		versehafmap[table.Value[idx]] = key
	}
}
func Compress(srcPath string) string {
	readTree()
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	size := 0
	output := make([]byte,0)
	var tp byte
	for _, by := range file {
		for i:=0;i < len(hafmap[by]); i++ {
			if size == 0 {
				tp<<=8
			}
			tp<<=1
			if hafmap[by][i] == '1' {
				tp |= 0x1
			}
			if size == 7 {
				output = append(output, tp)
			}
			size+=1
			size%=8
		}
	}
	if size != 0 {
		tp <<= 8-size
		output = append(output, tp)
	}
	fileHead := make([]byte,0)
	fileHead = append(fileHead, (byte) (size))
	dstPath := srcPath + ".ylx"
	fileHead = append(fileHead, output...)
	ioutil.WriteFile(dstPath, fileHead, 0777)
	fmt.Println(size)
	return utils.SucceedResponse()
}

func UndoCompress(srcPath string) string {
	readInverseTree()
	file, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Println(err)
		return utils.ErrorResponse(err)
	}
	size := int(file[0])
	trueInput := file[1:]
	output := make([]byte,0)
	var tp string
	for idx,by :=range trueInput {
		fmt.Println(by)
		if idx == len(trueInput) - 1 {
			for i:=7;i>=8-size;i-- {
				fmt.Println(by&(0x1<<i),i)
				if by&(0x1<<i) != 0 {
					tp = tp + "1"
				} else {
					tp = tp + "0"
				}
				if a,ok := versehafmap[tp];ok {
					fmt.Println(tp,versehafmap[tp])
					output = append(output, a)
					tp = ""
				}
			}
		} else {
			for i:=7;i>=0;i-- {
				fmt.Println(by&(0x1<<i),i)
				if by&(0x1<<i) != 0 {
					tp = tp + "1"
				} else {
					tp = tp + "0"
				}
				if a,ok := versehafmap[tp];ok {
					fmt.Println(tp,versehafmap[tp])
					output = append(output, a)
					tp = ""
				}
			}
		}
	}
	ioutil.WriteFile(srcPath[:len(srcPath)-4],output,0777)
	return utils.SucceedResponse()
}

