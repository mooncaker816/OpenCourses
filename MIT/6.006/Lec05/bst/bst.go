package bst

import (
	"errors"

	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec05/bintree"
)

type BST struct {
	Root *bintree.Node
	Comp Comparator
}

// NewBST 新建BST
func NewBST() *BST {
	return NewBSTWithComparator(BasicCompare)
}

// NewBSTWithComparator 新建自定义比较器的BST
func NewBSTWithComparator(comp Comparator) *BST {
	return &BST{Comp: comp}
}

// Search returns the (node,true) if the there is a node with the same key as provided.
// If there's no node with the same key, Search returns (node,false) where the node is parent node for inserting a new node with the key
func (bst *BST) Search(key interface{}) (*bintree.Node, bool) {
	n, result := bst.searchIn(bst.Root, key)
	if result == 0 {
		return n, true
	}
	return n, false
}

func (bst *BST) searchIn(n *bintree.Node, key interface{}) (*bintree.Node, int) {
	switch bst.Comp(key, n.Key) {
	case 0:
		return n, 0
	case -1:
		if n.HasLChild() {
			return bst.searchIn(n.LChild, key)
		}
		return n, -1
	case 1:
		if n.HasRChild() {
			return bst.searchIn(n.RChild, key)
		}
		return n, 1
	}
	return nil, 0
}

// Insert returns the new node if it is successfully inserted, otherwise it returns nil,err
func (bst *BST) Insert(key, data interface{}) (*bintree.Node, error) {
	if bst.Root == nil {
		bst.Root = &bintree.Node{Key: key, Data: data}
		return bst.Root, nil
	}
	n, result := bst.searchIn(bst.Root, key)
	switch result {
	case 0:
		return nil, errors.New("insert node with duplicate key")
	case -1:
		return n.AddLChild(key, data), nil
	case 1:
		return n.AddRChild(key, data), nil
	}
	return nil, errors.New("unknow compare result")
}

func subTreeMin(n *bintree.Node) *bintree.Node {
	for n.HasLChild() {
		n = n.LChild
	}
	return n
}

func subTreeMax(n *bintree.Node) *bintree.Node {
	for n.HasRChild() {
		n = n.RChild
	}
	return n
}

// Successor returns the next larger node of n if it exists.
func (bst *BST) Successor(n *bintree.Node) *bintree.Node {
	if n.HasRChild() {
		return subTreeMin(n.RChild)
	}
	for n.IsRChild() {
		n = n.Parent
	}
	return n.Parent
}

// Predecessor returns the prev smaller node of n if it exists.
func (bst *BST) Predecessor(n *bintree.Node) *bintree.Node {
	if n.HasLChild() {
		return subTreeMax(n.LChild)
	}
	for n.IsLChild() {
		n = n.Parent
	}
	return n.Parent
}

// Delete removes the node with the same key as provided and returns the "hot node" where its subtree is changed.
func (bst *BST) Delete(key interface{}) (hot *bintree.Node, err error) {
	n, result := bst.searchIn(bst.Root, key)
	if result != 0 {
		return nil, nil
	}
	// n has both left and right subtree
	if n.HasLChild() && n.HasRChild() {
		succ := bst.Successor(n)
		swapKeyData(n, succ)
		hot = succ.Parent
		if succ == n.RChild {
			hot = n
			if succ.HasRChild() {
				succ.RChild.Parent = n
			}
			n.RChild = succ.RChild
		} else if succ.HasRChild() {
			// hot = succ.Parent
			succ.RChild.Parent = hot
			hot.LChild = succ.RChild
		} else {
			// hot = succ.Parent
			hot.LChild = nil
		}
		succ.Parent, succ.LChild, succ.RChild = nil, nil, nil
		hot.UpdateHeightAbove()
		return hot, nil
	}
	// n only has left subtree
	if n.HasLChild() && !n.HasRChild() {
		if n.IsLChild() {
			n.Parent.LChild = n.LChild
		} else if n.IsRChild() {
			n.Parent.RChild = n.LChild
		} else {
			bst.Root = n.LChild
		}
		n.LChild.Parent = n.Parent
		hot = n.Parent
		n.Parent, n.LChild, n.RChild = nil, nil, nil
		hot.UpdateHeightAbove()
		return hot, nil
	}
	// n only has right subtree
	if !n.HasLChild() && n.HasRChild() {
		if n.IsLChild() {
			n.Parent.LChild = n.RChild
		} else if n.IsRChild() {
			n.Parent.RChild = n.RChild
		} else {
			bst.Root = n.RChild
		}
		n.RChild.Parent = n.Parent
		hot = n.Parent
		n.Parent, n.LChild, n.RChild = nil, nil, nil
		hot.UpdateHeightAbove()
		return hot, nil
	}
	// n is leaf(or single root)
	if n.IsLChild() {
		n.Parent.LChild = nil
	} else if n.IsRChild() {
		n.Parent.RChild = nil
	} else {
		bst.Root = nil
	}
	hot = n.Parent
	n.Parent, n.LChild, n.RChild = nil, nil, nil
	hot.UpdateHeightAbove()
	return hot, nil
}

func swapKeyData(n1, n2 *bintree.Node) {
	n1.Key, n1.Data, n2.Key, n2.Data = n2.Key, n2.Data, n1.Key, n1.Data
}

func (bst *BST) Print() {
	bst.Root.PrintWithUnitSize(2)
}

func (bst *BST) PrintWithUnitSize(size int) {
	bst.Root.PrintWithUnitSize(size)
}
