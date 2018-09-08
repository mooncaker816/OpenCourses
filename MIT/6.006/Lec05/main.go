package main

import (
	"fmt"

	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec05/bst"
)

func main() {
	keys := []int{16, 10, 25, 5, 11, 19, 28, 2, 8, 15, 17, 22, 27, 37, 4, 13, 33}
	bst := bst.NewBST()
	for _, v := range keys {
		bst.Insert(v, nil)
	}
	bst.Print()
	n, ok := bst.Search(37)
	fmt.Println(ok, n.Key)
	n, ok = bst.Search(36)
	fmt.Println(ok, n.Key)
	hot, _ := bst.Delete(16)
	bst.Print()
	if succ := bst.Successor(hot); succ != nil {
		fmt.Println("succ: ", succ.Key)
	}
	hot, _ = bst.Delete(28)
	bst.Print()
	if succ := bst.Successor(hot); succ != nil {
		fmt.Println("succ: ", succ.Key)
	}
	if pred := bst.Predecessor(hot); pred != nil {
		fmt.Println("pred: ", pred.Key)
	}
}
