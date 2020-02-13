package gotrees

type TreeNode interface {
	Key() float64
}

type RBNodeColor int64
type RBChildType int64

const (
	RBColor_Red   RBNodeColor = 0
	RBColor_Black RBNodeColor = 1

	RBChildType_Root  RBChildType = 0
	RBChildType_Left  RBChildType = 1
	RBChildType_Right RBChildType = 2

	Blah = 0
)
