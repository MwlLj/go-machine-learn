package tree

type IValue interface {
	LessEqual(value interface{}) bool
}
