package gotrees
import "fmt"

type rbValidator struct {
	blackNodeCount     int
	violatingNodeCount int
	redRedViolation    bool
	blackNodeViolation bool
	bstViolation       bool
}

func ValidateRBTree(tree *RBTree) error {
	if tree == nil {
		return nil
	}

	return ValidateRBTreeByRoot(tree.root)
}

func ValidateRBTreeByRoot(root *RBNode) error {
	if root == nil {
		return nil
	}

	validator := &rbValidator{
		blackNodeCount: -1,
	}

	validator.validateFromNode(root, 0)

	if validator.blackNodeViolation {
		return fmt.Errorf("Tree invalid because of differing black node counts of %d and %d",
					validator.blackNodeCount, validator.violatingNodeCount)
	}

	if validator.redRedViolation {
		return fmt.Errorf("Tree invalid because of red node with red child")
	}

	if validator.bstViolation {
		return fmt.Errorf("Tree invalid because nodes of incorrect order were found")
	}

	return nil
}

func (v* rbValidator) String() string {
	return fmt.Sprintf("%d:%t:%t:%t", v.blackNodeCount, v.redRedViolation, v.blackNodeViolation, v.bstViolation)
}

func (v *rbValidator) validateFromNode(node *RBNode, count int) {
	if v.blackNodeViolation || v.redRedViolation  || v.bstViolation {
		return
	}

	if node == nil {
		v.checkBlackNodeCount(count + 1)

		return
	}

	if v.checkForBSTViolation(node) {
		return
	}


	if node.Color == RBColor_Red {
		if v.checkForRedRedViolation(node) {
			return
		}
	} else {
		count = count + 1
	}

	v.validateFromNode(node.LeftChild, count)
	v.validateFromNode(node.RightChild, count)
}

func (v* rbValidator) checkBlackNodeCount(count int) {
	if v.blackNodeCount < 0 {
		v.blackNodeCount = count
	} else if v.blackNodeCount != count {
		v.violatingNodeCount = count
		v.blackNodeViolation = true
	}
}

func (v* rbValidator) checkForBSTViolation(node *RBNode) bool {
	if node.LeftChild != nil {
		if node.Node.Key() < node.LeftChild.Node.Key() {
			v.bstViolation = true
			return true
		}
	}

	if node.RightChild != nil {
		if node.Node.Key() > node.RightChild.Node.Key() {
			v.bstViolation = true
			return true
		}
	}

	return false
}

func (v* rbValidator) checkForRedRedViolation(node *RBNode) bool {
	if node.LeftChild != nil && node.LeftChild.Color == RBColor_Red {
		v.redRedViolation = true
		return true
	}

	if node.RightChild != nil && node.RightChild.Color == RBColor_Red {
		v.redRedViolation = true
		return true
	}

	return false
}
