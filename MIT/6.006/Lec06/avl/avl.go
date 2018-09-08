package avl

import (
	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec05/bintree"
	"github.com/mooncaker816/OpenCourses/MIT/6.006/Lec05/bst"
)

// AVL tree
type AVL struct {
	*bst.BST
}

func avlOK(n *bintree.Node) bool {
	lh, rh := -1, -1
	if n.HasLChild() {
		lh = n.LChild.Height
	}
	if n.HasRChild() {
		rh = n.RChild.Height
	}
	diff := lh - rh
	if diff > 1 || diff < -1 {
		return false
	}
	return true
}

// NewAVL returns a new empty avl tree with basic comparator
func NewAVL() *AVL {
	return &AVL{bst.NewBST()}
}

// NewAVLWithComparator returns a new empty avl tree with provided comparator
func NewAVLWithComparator(comp bst.Comparator) *AVL {
	return &AVL{bst.NewBSTWithComparator(comp)}
}

func (avl *AVL) Insert(key, data interface{}) (*bintree.Node, error) {
	n, err := avl.BST.Insert(key, data)
	if err != nil {
		return nil, err
	}
	hot := n.Parent
	avl.reBalance(hot, true)
	return n, nil
}

func (avl *AVL) reBalance(hot *bintree.Node, insert bool) {
	for g := hot; g != nil; g = g.Parent {
		if !avlOK(g) {
			x := g.Parent
			p := g.TallerChild()
			v := p.TallerChild()
			tmp := new(bintree.Node)
			if g.IsLChild() {
				x.LChild = rotateAt(v)
				tmp = x.LChild
			} else if g.IsRChild() {
				x.RChild = rotateAt(v)
				tmp = x.RChild
			} else {
				avl.Root = rotateAt(v)
				tmp = avl.Root
			}
			if insert {
				break
			} else {
				g = tmp
			}
		} else {
			g.UpdateHeight()
		}
	}
}

// connect34 connect nodes as below
//	 	   b
//		a	  c
//	  T0 T1 T2 T3
func connect34(a, b, c, t1, t2, t3, t4 *bintree.Node) *bintree.Node {
	a.LChild = t1
	if t1 != nil {
		t1.Parent = a
	}
	a.RChild = t2
	if t2 != nil {
		t2.Parent = a
	}
	a.UpdateHeight()
	c.LChild = t3
	if t3 != nil {
		t3.Parent = c
	}
	c.RChild = t4
	if t4 != nil {
		t4.Parent = c
	}
	c.UpdateHeight()
	b.LChild = a
	b.RChild = c
	a.Parent = b
	c.Parent = b
	b.UpdateHeight()
	return b
}

func rotateAt(v *bintree.Node) *bintree.Node {
	p := v.Parent
	g := p.Parent
	if v.IsLChild() {
		if p.IsLChild() {
			p.Parent = g.Parent
			return connect34(v, p, g, v.LChild, v.RChild, p.RChild, g.RChild)
		}
		if p.IsRChild() {
			v.Parent = g.Parent
			return connect34(g, v, p, g.LChild, v.LChild, v.RChild, p.RChild)
		}
	}
	if v.IsRChild() {
		if p.IsLChild() {
			v.Parent = g.Parent
			return connect34(p, v, g, p.LChild, v.LChild, v.RChild, g.RChild)
		}
		if p.IsRChild() {
			p.Parent = g.Parent
			return connect34(g, p, v, g.LChild, p.LChild, v.LChild, v.RChild)
		}
	}
	return nil
}

func (avl *AVL) Delete(key interface{}) (*bintree.Node, error) {
	hot, err := avl.BST.Delete(key)
	if err != nil {
		return nil, err
	}
	avl.reBalance(hot, false)
	return hot, nil
}
