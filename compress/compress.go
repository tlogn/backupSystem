package compress

import (
	//"errors"
	"fmt"
	"io/ioutil"

	//"path/filepath"
	"sort"
)

//func Compress(w http.ResponseWriter, r *utils.Request){
//	comOrUncom := r.CompressPara.IsCompress
//	if comOrUncom {
//		fmt.Fprintf(w, "%v", compress(r))
//	} else {
//		fmt.Fprintf(w, "%v", undoCompress(r))
//	}
//}
type TreeNode struct {
	left,right * TreeNode
	value int
	name byte
}
type PairList []TreeNode

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value > p[j].value }

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
	for key, value := range mapp {
		tmp := TreeNode{
			left:nil,
			right:nil,
			value:value,
			name:key,
		}
		arr = append(arr, tmp)
	}
	sort.Sort(arr)
	fmt.Println(arr.Len())
	for _, node := range arr {
		fmt.Println(node.name,node.value)
	}
	return nil
}


//func compress(r *utils.Request) error {
//
//	srcPath, err := filepath.Abs(r.CompressPara.CompressPath)
//
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	if utils.IsDir(srcPath) {
//		err := errors.New("is a fucking dir")
//		log.Printf("%v %v", srcPath, err)
//		return err
//	}
//
//	file, err := ioutil.ReadFile(srcPath)
//	if err != nil {
//		log.Printf("read file %v error, %v",srcPath, err)
//		return err
//	}
//
//	err = ioutil.WriteFile(srcPath + ".ylxcompress", compressedFile, 0777)
//	return nil
//}
//
//func undoCompress(r *utils.Request) error {
//
//	return nil
//}

