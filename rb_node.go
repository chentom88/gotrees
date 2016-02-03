package gotrees

type RBNode struct {
	Color      RBNodeColor
	Parent     *RBNode
	RightChild *RBNode
	LeftChild  *RBNode
	Node       TreeNode
	ChildType  RBChildType
}

func (n *RBNode) SwitchColor() {
	if n.Color == RBColor_Red {
		n.Color = RBColor_Black
	} else {
		n.Color = RBColor_Red
	}
}

func (n *RBNode) SetLeftChild(child *RBNode) {
	if child == nil {
		n.LeftChild = child
		return
	}

	child.Parent = n
	n.LeftChild = child
	child.ChildType = RBChildType_Left
}

func (n *RBNode) SetRightChild(child *RBNode) {
	if child == nil {
		n.RightChild = child
		return
	}

	child.Parent = n
	n.RightChild = child
	child.ChildType = RBChildType_Right
}

func (n *RBNode) FindAncestorNodes() (parent, grandParent, uncle *RBNode) {
	parentNode := n.Parent
	var grandParentNode *RBNode = nil
	var uncleNode *RBNode = nil

	if parentNode.ChildType != RBChildType_Root {
		grandParentNode = parentNode.Parent

		if parentNode.ChildType == RBChildType_Right {
			uncleNode = grandParentNode.LeftChild
		} else {
			uncleNode = grandParentNode.RightChild
		}
	}

	return parentNode, grandParentNode, uncleNode
}