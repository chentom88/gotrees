package gotrees_test

import (
	. "gotrees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RBTreeValidator", func() {
	var (
		root 				*RBNode
		validationResult	error
	)

	BeforeEach(func() {
		root = buildValidTree()
	})

	AfterEach(func() {
		root = nil
		validationResult = nil
	})

	Context("Trees are valid", func() {
		It("returns that a nil tree is valid", func() {
			Expect(func() { validationResult = ValidateRBTree(nil) }).ToNot(Panic())
			Expect(validationResult).To(BeNil())
		})

		It("returns that a nil root is valid", func() {
			Expect(func() { validationResult = ValidateRBTreeByRoot(nil) }).ToNot(Panic())
			Expect(validationResult).To(BeNil())
		})

		It("correctly identifies a valid tree", func() {
			Expect(func() { validationResult = ValidateRBTreeByRoot(root) } ).ToNot(Panic())
			Expect(validationResult).To(BeNil())
		})
	})

	Context("Trees are invalid", func() {
		It("correctly identifies a black node count violation", func() {
			root.LeftChild.RightChild = nil

			Expect(func() { validationResult = ValidateRBTreeByRoot(root) } ).ToNot(Panic())
			Expect(validationResult).ToNot(BeNil())
		})

		It("correctly identifies a black node count violation", func() {
			root.LeftChild.RightChild.Color = RBColor_Red

			Expect(func() { validationResult = ValidateRBTreeByRoot(root) } ).ToNot(Panic())
			Expect(validationResult).ToNot(BeNil())
		})

		It("correctly identifies a BST violation", func() {
			root.LeftChild.LeftChild.Node = &IntegerNode{ value: 11 }

			Expect(func() { validationResult = ValidateRBTreeByRoot(root) } ).ToNot(Panic())
			Expect(validationResult).ToNot(BeNil())
		})
	})
})

func buildValidTree() *RBNode {
	root := &RBNode{
		ChildType: RBChildType_Root,
		Color: RBColor_Black,
		Node: &IntegerNode{ value: 5 },
	}

	node3 := &RBNode{
		Color: RBColor_Red,
		Node: &IntegerNode{ value: 3 },
	}

	node2 := &RBNode{
		Color: RBColor_Black,
		Node: &IntegerNode{ value: 2 },
	}

	node4 := &RBNode{
		Color: RBColor_Black,
		Node: &IntegerNode{ value: 4 },
	}

	node1 := &RBNode{
		Color: RBColor_Red,
		Node: &IntegerNode{ value: 1 },
	}

	root.SetLeftChild(node3)
	node3.SetLeftChild(node2)
	node3.SetRightChild(node4)
	node2.SetLeftChild(node1)

	node7 := &RBNode{
		Color: RBColor_Red,
		Node: &IntegerNode{ value: 7 },
	}

	node6 := &RBNode{
		Color: RBColor_Black,
		Node: &IntegerNode{ value: 6 },
	}

	node9 := &RBNode{
		Color: RBColor_Black,
		Node: &IntegerNode{ value: 9 },
	}

	node8 := &RBNode{
		Color: RBColor_Red,
		Node: &IntegerNode{ value: 8 },
	}

	node10 := &RBNode{
		Color: RBColor_Red,
		Node: &IntegerNode{ value: 10 },
	}

	root.SetRightChild(node7)
	node7.SetLeftChild(node6)
	node7.SetRightChild(node9)
	node9.SetLeftChild(node8)
	node9.SetRightChild(node10)

	return root
}

