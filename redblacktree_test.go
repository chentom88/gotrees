package gotrees_test

import (
	. "gotrees"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"math/rand"
)

var _ = Describe("RedBlackTree", func() {
	var (
		rbTree, nilTree *RBTree
		randGen *rand.Rand = rand.New(rand.NewSource(9))
	)

	BeforeEach(func() {
		rbTree = &RBTree{}
	})

	AfterEach(func() {
		rbTree = nil
	})

	Context("is not initialized", func() {
		It("does not panic when item is inserted", func() {
			Expect(func() { nilTree.Insert(&IntegerNode { value:randGen.Int() }) }).ToNot(Panic())
		})

		It("does not panic when list is requested", func() {
			Expect(func() { nilTree.List() }).ToNot(Panic())
		})

		It("returns nil when list is requested", func() {
			Expect(nilTree.List()).To(BeNil())
		})
	})

	Context("has no items inserted" , func() {
		It("does not panic when asked for list of items", func() {
			Expect(func() { rbTree.List() }).ToNot(Panic())
		})
	})

	Context("is given invalid items to insert", func() {
		It("does not panic when given a nil item", func() {
			Expect(func() { rbTree.Insert(nil) }).ToNot(Panic())
		})

		It("does not insert nil nodes", func() {
			rbTree.Insert(nil)

			result := rbTree.List()
			Expect(len(result)).To(Equal(0))
		})
	})

	Context("is given valid items to insert", func() {
		DescribeTable("sorts items as they are added", func(numItems int) {
			for i := 0; i < numItems; i++ {
				rbTree.Insert(&IntegerNode{ value: randGen.Int() })
			}

			result := rbTree.List()

			Expect(ValidateRBTree(rbTree)).To(BeNil())
			Expect(len(result)).To(Equal(numItems))

			prevNode := result[0]
			for _, resultItem := range result[1:] {
				Expect(prevNode.Key() < resultItem.Key()).To(BeTrue())
				prevNode = resultItem
			}
		},
			Entry("just root node", 1),
			Entry("x-small number of items", 3),
			Entry("small number of items", 5),
			Entry("medium number of items", 15),
			Entry("large number of items", 10000),
			Entry("x-large number of items", 50000),
			Entry("xx-large number of items", 200000),
		)

		It("does not insert duplicate items", func() {
			rbTree.Insert(&IntegerNode{ value: 5 })
			rbTree.Insert(&IntegerNode{ value: 4 })
			rbTree.Insert(&IntegerNode{ value: 5 })
			rbTree.Insert(&IntegerNode{ value: 4 })

			result := rbTree.List()

			Expect(ValidateRBTree(rbTree)).To(BeNil())
			Expect(len(result)).To(Equal(2))
		})
	})
})

type IntegerNode struct {
	value int
}

func (i *IntegerNode) Key() float64 {
	return float64(i.value)
}

func (i *IntegerNode) String() string {
	return fmt.Sprintf("%d", i)
}