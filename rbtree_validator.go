package gotrees
import (
	"fmt"
)

type rbValidator struct {
	blackNodeCount     int
}

func ValidateRBTree(tree *RBTree) error {
	if tree == nil {
		return nil
	}

	return ValidateRBSubtree(tree.root)
}

func ValidateRBSubtree(root *RBNode) error {
	if root == nil {
		return nil
	}

	validator := &rbValidator{
		blackNodeCount: -1,
	}

	return validator.validateFromNode(root, 0)
}

func (v *rbValidator) validateFromNode(node *RBNode, count int) error {
	if node == nil {
		if err := v.checkBlackNodeCount(count + 1); err != nil {
			return err
		}

		return nil
	}

	if err := v.checkForBSTViolation(node); err != nil {
		return err
	}

	if node.Color == RBColor_Red {
		if err := v.checkForRedRedViolation(node); err != nil {
			return err
		}
	} else {
		count = count + 1
	}

	if err := v.validateFromNode(node.LeftChild, count); err != nil {
		return err
	}

	if err := v.validateFromNode(node.RightChild, count); err != nil {
		return err
	}

	return nil
}

func (v* rbValidator) checkBlackNodeCount(count int) error {
	if v.blackNodeCount < 0 {
		v.blackNodeCount = count
	} else if v.blackNodeCount != count {
		return fmt.Errorf("Tree invalid because of differing black node counts of %d and %d",
			v.blackNodeCount, count)
	}

	return nil
}

func (v* rbValidator) checkForBSTViolation(node *RBNode) error {
	violation := false

	if node.LeftChild != nil {
		if node.Node.Key() < node.LeftChild.Node.Key() {
			violation = true
		}
	}

	if !violation && node.RightChild != nil {
		if node.Node.Key() > node.RightChild.Node.Key() {
			violation = true
		}
	}

	if violation {
		return fmt.Errorf("Tree invalid because nodes of incorrect order were found")
	}

	return nil
}

func (v* rbValidator) checkForRedRedViolation(node *RBNode) error {
	violation := false

	if node.LeftChild != nil && node.LeftChild.Color == RBColor_Red {
		violation = true
	}

	if !violation && node.RightChild != nil && node.RightChild.Color == RBColor_Red {
		violation = true
	}

	if violation {
		return fmt.Errorf("Tree invalid because of red node with red child")
	}

	return nil
}
