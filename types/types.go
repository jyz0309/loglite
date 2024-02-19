package types

const (
	StrItemType = iota
	IntItemType
)

type Item interface {
	Type() int
	Val() string
}

type StrItem struct {
	val string
}

func (item StrItem) Type() int {
	return StrItemType
}

func (item StrItem) Val() string {
	return item.val
}

func BuildStrItem(val string) StrItem {
	return StrItem{
		val: val,
	}
}
