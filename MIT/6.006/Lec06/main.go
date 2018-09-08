package main

import (
	"fmt"

	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec05/bintree"
	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec06/avl"
)

func main() {
	keys := []int{16, 10, 25, 5, 11, 19, 28, 2, 8, 15, 17, 22, 27, 37, 4, 33}
	at := avl.NewAVL()
	for _, v := range keys {
		at.Insert(v, nil)
	}
	at.Print()
	at.Insert(13, nil)
	at.Print()
	at.Insert(30, nil)
	at.Print()
	at2 := avl.NewAVL()
	at2.Insert(100, nil)
	at2.Print()
	at2.Insert(95, nil)
	at2.Print()
	at2.Insert(85, nil)
	at2.Print()
	at2.Insert(75, nil)
	at2.Print()
	at2.Insert(65, nil)
	at2.Print()
	at2.Insert(55, nil)
	at2.Print()
	at2.Insert(45, nil)
	at2.Print()
	at2.Insert(35, nil)
	at2.Print()
	at2.Root.TravLevel(withprintheight())
	fmt.Println()

	keys2 := []int{16, 10, 25, 5, 11, 19, 28, 2, 8, 15, 17, 22, 27, 37, 4, 33}
	at = avl.NewAVL()
	for _, v := range keys2 {
		at.Insert(v, nil)
	}
	at.Print()
	at.Delete(8)
	at.Print()
	at.Delete(16)
	at.Print()
	at.Delete(19)
	at.Print()
	at.Root.TravLevel(withprintheight())
	fmt.Println()
	at2 = avl.NewAVL()
	at2.Insert(5, nil)
	at2.Insert(1, nil)
	at2.Insert(10, nil)
	at2.Insert(11, nil)
	at2.Print()
	at2.Delete(1)
	at2.Print()
	at2.Root.TravLevel(withprintheight())
}

func withprintheight() func(n *bintree.Node) {
	return func(n *bintree.Node) {
		fmt.Printf("%d ", n.Height)
	}
}
