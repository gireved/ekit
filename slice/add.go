package slice

import "my-ekit/internal/slice"

// Add 在index处添加元素  index的范围[0,len(src)]  如果index == len(src) 则表示向末尾添加元素
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}
