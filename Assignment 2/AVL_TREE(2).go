package avltree

import (
	"fmt"
)


type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(key int, value int) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree) Remove(key int) {
	t.root.remove(key)
}

func (t *AVLTree) Update(old_Key int, new_Key int, new_Value int) {
	t.root.remove(old_Key)
	t.root = t.root.add(new_Key, new_Value)
}

func (t *AVLTree) Search(key int) (node *AVLNode) {
	return t.root.search(key)
}

func (t *AVLTree) DIO() {                      //DIO - display in order
	t.root.DNIO()                          //DNIO - display nodes in order
}


type AVLNode struct {
	key   int
	Value int

	
	count_nodes int
	left   *AVLNode
	right  *AVLNode
}


func (n *AVLNode) add(key int, value int) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.Value = value
	}
	return n.rebalanceTree()
}

func (n *AVLNode) remove(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}


func (n *AVLNode) search(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}


func (n *AVLNode) DNIO() {
	if n.left != nil {
		n.left.DNIO()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.DNIO()
	}
}

func (n *AVLNode) getcount_nodes() int {
	if n == nil {
		return 0
	}
	return n.count_nodes
}

func (n *AVLNode) re_count_nodes() {
	n.count_nodes = 1 + max(n.left.getcount_nodes(), n.right.count_nodes())
}


func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.re_count_nodes()

	// check balance factor 

	balanceFactor := n.left.getcount_nodes() - n.right.getcount_nodes()
	if balanceFactor == -2 {
		
		if n.right.left.getcount_nodes() > n.right.right.getcount_nodes() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		
		if n.left.right.getcount_nodes() > n.left.left.getcount_nodes() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}


func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.re_count_nodes()
	return newRoot
}


func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.re_count_nodes()
	return newRoot
}


func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
    tree := new(avltree.AVLTree)

    keys := []int{5,7,4,2,3}
    for _, key := range keys {
        tree.Add(key, key*key)
    }   

    tree.Remove(3)
    tree.Update(4, 8, 8*8)
    tree.DIO()

    val := tree.Search(2).Value
}
