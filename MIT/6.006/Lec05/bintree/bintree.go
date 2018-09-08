package bintree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	LChild *Node
	RChild *Node
	Parent *Node
	Key    interface{}
	Data   interface{}
	Height int
}

type BinTree struct {
	Root *Node
}

func newSingleNode(key, data interface{}) *Node {
	return &Node{Key: key, Data: data}
}

// NewBinTree creates a bintree with a single root node
func NewBinTree(root *Node) *BinTree {
	return &BinTree{Root: root}
}

// HasLChild checks if n has left child
func (n *Node) HasLChild() bool { return n.LChild != nil }

// HasRChild checks if n has right child
func (n *Node) HasRChild() bool { return n.RChild != nil }

// IsLChild checks if n is a left child
func (n *Node) IsLChild() bool {
	return !n.IsRoot() && n == n.Parent.LChild
}

// IsRChild checks if n is a right child
func (n *Node) IsRChild() bool {
	return !n.IsRoot() && n == n.Parent.RChild
}

// IsRoot checks if n is root
func (n *Node) IsRoot() bool { return n.Parent == nil }

// Sibling gets n's sibling if exists
func (n *Node) Sibling() *Node {
	if n.IsRoot() {
		return nil
	}
	if n.IsLChild() {
		return n.Parent.RChild
	}
	return n.Parent.LChild
}

// AddLChild generates a new Node with the provided key and data, then adds to Node n as left child
func (n *Node) AddLChild(key, data interface{}) *Node {
	if n.HasLChild() {
		panic("node already has left child")
	}
	newNode := newSingleNode(key, data)
	n.LChild = newNode
	newNode.Parent = n
	if !n.HasRChild() {
		n.UpdateHeightAbove()
	}
	return newNode
}

// AddRChild generates a new Node with the provided key and data, then adds to Node n as right child
func (n *Node) AddRChild(key, data interface{}) *Node {
	if n.HasRChild() {
		panic("node already has right child")
	}
	newNode := newSingleNode(key, data)
	n.RChild = newNode
	newNode.Parent = n
	if !n.HasLChild() {
		n.UpdateHeightAbove()
	}
	return newNode
}

// AttachLSubTree will add left sub tree to n
func (n *Node) AttachLSubTree(left *Node) {
	if n.HasLChild() {
		panic("node already has left sub tree")
	}
	n.LChild = left
	left.Parent = n
	n.UpdateHeightAbove()
}

// AttachRSubTree will add right sub tree to n
func (n *Node) AttachRSubTree(right *Node) {
	if n.HasRChild() {
		panic("node already has right sub tree")
	}
	n.RChild = right
	right.Parent = n
	n.UpdateHeightAbove()
}

// UpdateHeightAbove updates height info for all the related nodes
func (n *Node) UpdateHeightAbove() {
	for max := n.maxHeightOfChildren(); n != nil && n.Height != max+1; {
		n.Height = max + 1
		n = n.Parent
	}
}

func (n *Node) maxHeightOfChildren() int {
	lH, rH := -1, -1
	if n.HasLChild() {
		lH = n.LChild.Height
	}
	if n.HasRChild() {
		rH = n.RChild.Height
	}
	return max(lH, rH)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Detach removes the whole sub tree
func (n *Node) Detach() {
	if n.IsRoot() {
		return
	}
	if n.IsLChild() {
		n.Parent.LChild = nil
		n.Parent.UpdateHeightAbove()
		n.Parent = nil
		return
	}
	n.Parent.RChild = nil
	n.Parent.UpdateHeightAbove()
	n.Parent = nil
}

// Size returns the count of sub tree's nodes
func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	count := 1
	if n.HasLChild() {
		count += n.LChild.Size()
	}
	if n.HasRChild() {
		count += n.RChild.Size()
	}
	return count
}

// Level returns the level of the node in the tree, root is level 0
func (n *Node) Level() int {
	l := 0
	for !n.IsRoot() {
		l++
		n = n.Parent
	}
	return l
}

// Option is a useful function to dealing with the node during the traversal of the tree
type Option func(n *Node)

// TravPre provides the pre-order traversal
func (n *Node) TravPre(opts ...Option) {
	for _, opt := range opts {
		opt(n)
	}
	if n.HasLChild() {
		n.LChild.TravPre(opts...)
	}
	if n.HasRChild() {
		n.RChild.TravPre(opts...)
	}
}

// TravIn provides the in-order traversal
func (n *Node) TravIn(opts ...Option) {
	if n.HasLChild() {
		n.LChild.TravIn(opts...)
	}
	for _, opt := range opts {
		opt(n)
	}
	if n.HasRChild() {
		n.RChild.TravIn(opts...)
	}
}

// TravPost provides the post-order traversal
func (n *Node) TravPost(opts ...Option) {
	if n.HasLChild() {
		n.LChild.TravIn(opts...)
	}
	if n.HasRChild() {
		n.RChild.TravIn(opts...)
	}
	for _, opt := range opts {
		opt(n)
	}
}

// TravLevel provides the level-order traversal
func (n *Node) TravLevel(opts ...Option) {
	queue := make([]*Node, 0, n.Size())
	queue = append(queue, n)
	for len(queue) > 0 {
		visitNode := queue[0]
		queue = queue[1:]
		if visitNode.HasLChild() {
			queue = append(queue, visitNode.LChild)
		}
		if visitNode.HasRChild() {
			queue = append(queue, visitNode.RChild)
		}
		for _, opt := range opts {
			opt(visitNode)
		}
	}
}

// Print 以子树节点个数的位数为一个基本单元的长度，打印子树的拓扑结构到标准输出
func (n *Node) Print() {
	n.PrintWithUnitSize(len(strconv.Itoa(n.Size())))
}

// Fprint 以树节点个数的位数为一个基本单元的长度，打印子树的拓扑结构到io.Writer
func (n *Node) Fprint(w io.Writer) {
	n.FprintWithUnitSize(w, len(strconv.Itoa(n.Size())))
}

// PrintWithUnitSize 以指定的长度为一个基本单元，打印子树的拓扑结构到标准输出
func (n *Node) PrintWithUnitSize(size int) {
	n.FprintWithUnitSize(os.Stdout, size)
}

// FprintWithUnitSize 以指定的长度为一个基本单元，打印子树的拓扑结构到io.Writer，树宽为节点数
func (n *Node) FprintWithUnitSize(w io.Writer, size int) {
	buf := bufio.NewWriter(w)
	if n == nil {
		buf.WriteString("Empty tree!")
		buf.Flush()
		return
	}
	if size <= 0 {
		panic("unit size can not be less than 1")
	}
	unitSpace := strings.Repeat(" ", size)
	unitHen := strings.Repeat("─", size)
	type nodeOffset struct {
		node   *Node
		offset int
	}
	q := make([]nodeOffset, 0, n.Size())
	q = append(q, nodeOffset{n, 0})
	prevlevel := 0

	delta := 0 // 由于没有孩子节点需要过继给后续有孩子节点的偏移量

	for len(q) > 0 {
		no := q[0]
		q = q[1:]
		n := no.node
		offset := no.offset
		if l := n.Level(); prevlevel != l {
			delta = 0 // 原二叉树与对应的完全二叉树中缺失节点造成的空格缺失数
			prevlevel = l
			buf.WriteString("\n")
		}

		var nodeLeftStr, nodeRightStr string
		// 由父亲节点遗留下来的前缀偏移量
		buf.WriteString(strings.Repeat(unitSpace, offset))
		if n.IsRChild() && n.Sibling() == nil {
			offset++                   // 缺失左兄弟导致原本兄弟与兄弟之间的一个单位空格缺失，需传给该节点的后代节点
			buf.WriteString(unitSpace) // 在打印右儿子之前补上该空格
		}
		if n.HasLChild() {
			offset += delta // 将偏移量加上由之前同层的叶子节点造成的空儿子节点引起的空格数，完成过继后置零
			delta = 0
			//如果该节点有左孩子，优先将偏移量转移至左孩子（因为从左往右打印）
			q = append(q, nodeOffset{n.LChild, offset})
			nodeLeftStr = strings.Repeat(unitSpace, n.LChild.LChild.Size()) +
				strings.Repeat(" ", size-1) + "┌" +
				strings.Repeat(unitHen, n.LChild.RChild.Size())
		}
		if n.HasRChild() {
			if n.HasLChild() {
				offset = 0 //若有左孩子，则偏移量已经转移至左孩子，无需再转移给右孩子
			} else { // 不得已转移给右孩子
				offset += delta
				delta = 0
			}
			q = append(q, nodeOffset{n.RChild, offset})
			nodeRightStr = strings.Repeat(unitHen, n.RChild.LChild.Size()) +
				strings.Repeat("─", size-1) + "┐" +
				strings.Repeat(unitSpace, n.RChild.RChild.Size())
		}
		if !n.HasLChild() && !n.HasRChild() { // 叶节点需保存当前的偏移量再加上二个单位的空格，+= 防止连续叶节点导致偏移量丢失
			delta += offset
			delta++
			delta++
		}

		buf.WriteString(nodeLeftStr)
		buf.WriteString(fmt.Sprintf("%*v", size, n.Key))
		buf.WriteString(nodeRightStr)
		buf.WriteString(unitSpace)
		if n.IsLChild() && n.Sibling() == nil {
			// 缺失左兄弟导致原本兄弟与兄弟之间的一个单位空格缺失，需传给该节点的后代节点
			buf.WriteString(unitSpace) // 在打印右儿子之前补上该空格
		}
	}
	buf.WriteString("\n")
	buf.Flush()
}
