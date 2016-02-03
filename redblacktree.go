package gotrees

type listResult struct {
	slice []TreeNode
}

type RBTree struct {
	root *RBNode
}

func (rb *RBTree) Insert(newNode TreeNode) {
	if rb == nil || newNode == nil {
		return
	}

	if rb.root == nil {
		rb.root = &RBNode{
			Color:     RBColor_Black,
			Node:      newNode,
			ChildType: RBChildType_Root,
		}

		return
	}

	if newRBNode, rebalance := rb.insertNewNode(rb.root, newNode); rebalance == true {
		rb.rebalanceTree(newRBNode)
	}
}

func (rb *RBTree) List() []TreeNode {
	if rb == nil {
		return nil
	}

	result := &listResult{
		slice: make([]TreeNode, 0),
	}

	rb.traverse(rb.root, result)

	return result.slice
}

func (rb *RBTree) traverse(node *RBNode, result *listResult) {
	if node == nil || result == nil {
		return
	}

	rb.traverse(node.LeftChild, result)
	result.slice = append(result.slice, node.Node)
	rb.traverse(node.RightChild, result)
}

func (rb *RBTree) insertNewNode(parent *RBNode, newNode TreeNode) (*RBNode, bool) {
	if parent.Node.Key() == newNode.Key() {
		return nil, false
	}

	newRBNode := &RBNode{
		Parent: parent,
		Node:   newNode,
	}

	if parent.Node.Key() > newNode.Key() {
		if parent.LeftChild == nil {
			newRBNode.ChildType = RBChildType_Left
			parent.LeftChild = newRBNode

			return newRBNode, parent.Color == RBColor_Red
		}

		return rb.insertNewNode(parent.LeftChild, newNode)
	}

	if parent.RightChild == nil {
		newRBNode.ChildType = RBChildType_Right
		parent.RightChild = newRBNode
		return newRBNode, parent.Color == RBColor_Red
	}

	return rb.insertNewNode(parent.RightChild, newNode)
}

func (rb *RBTree) rebalanceTree(node *RBNode) {
	if node.ChildType == RBChildType_Root {
		node.Color = RBColor_Black
		return
	}

	parent, grandParent, uncle := node.FindAncestorNodes()

	var uncleColor = RBColor_Black
	if uncle != nil {
		uncleColor = uncle.Color
	}

	if grandParent == nil && parent.ChildType == RBChildType_Root {
		parent.Color = RBColor_Black
		return
	}

	if uncleColor == RBColor_Red {
		rb.balanceCase1(parent, grandParent, uncle)
	} else {
		if node.ChildType != parent.ChildType {
			rb.balanceCase2(node, parent, grandParent)

			// Parent becomes the new leaf node so need to recalculate ancestors
			// before moving to case 3
			node = parent
			parent, grandParent, uncle = node.FindAncestorNodes()
		}

		rb.balanceCase3(node, parent, grandParent)

		// Get ancestors against since rotation may have occurred
		parent, grandParent, uncle = node.FindAncestorNodes()
	}

	if grandParent != nil && grandParent.Parent != nil {
		if grandParent.ChildType != RBChildType_Root &&
		grandParent.Color == RBColor_Red && grandParent.Parent.Color == RBColor_Red {
			rb.rebalanceTree(grandParent)
		}
	}
}

func (rb *RBTree) balanceCase2(node, parent, grandParent *RBNode) {
	if node.ChildType == RBChildType_Right {
		tempChild := node.LeftChild
		grandParent.SetLeftChild(node)
		node.SetLeftChild(parent)
		parent.SetRightChild(tempChild)
	} else {
		tempChild := node.RightChild
		grandParent.SetRightChild(node)
		node.SetRightChild(parent)
		parent.SetLeftChild(tempChild)
	}
}

func (rb *RBTree) balanceCase3(node, parent, grandParent *RBNode) {
	origGrandParentType := grandParent.ChildType
	greatGrand := grandParent.Parent
	var tempChild *RBNode

	if parent.ChildType == RBChildType_Left {
		tempChild = parent.RightChild
		parent.SetRightChild(grandParent)
		grandParent.SetLeftChild(tempChild)
	} else {
		tempChild = parent.LeftChild
		parent.SetLeftChild(grandParent)
		grandParent.SetRightChild(tempChild)
	}

	if greatGrand != nil {
		if origGrandParentType == RBChildType_Right {
			greatGrand.SetRightChild(parent)
		} else {
			greatGrand.SetLeftChild(parent)
		}
	} else {
		parent.Parent = nil
		parent.ChildType = RBChildType_Root
		rb.root = parent
	}

	parent.SwitchColor()
	grandParent.SwitchColor()
}

func (rb *RBTree) balanceCase1(parent, grandParent, uncle *RBNode) {
	parent.SwitchColor()

	if grandParent.ChildType != RBChildType_Root {
		grandParent.SwitchColor()
	}

	uncle.SwitchColor()
}
