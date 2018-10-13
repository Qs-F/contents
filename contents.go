package contents

import (
	"errors"
	"fmt"
	"sort"
)

type Contents []Content

func New(csi interface{}) (*Contents, error) {
	// var c Contents
	// switch cs := csi.(type) {
	// case Contents:
	// 	for _, v := range cs {
	// 		if cv, ok := v.(Content); ok {
	// 			c = append(c, cv)
	// 		} else {
	// 			return nil, errors.New("Value does not satify Content interface")
	// 		}
	// 	}
	// default:
	// 	return nil, errors.New(spew.Sdump(cs))
	// return nil, errors.New("Cannot cast")
	// }
	// return &c, nil

	switch v := csi.(type) {
	case []interface{}:
		return nil, errors.New(fmt.Sprintf("%v", v))
	}
	return nil, errors.New("Cannot cast")
}

func (c *Contents) Len() int {
	return len(*c)
}

func (c *Contents) Less(i, j int) bool {
	return (*c)[i].Date().Before((*c)[j].Date())
}

func (c *Contents) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *Contents) SortByDate() {
	sort.Sort(c)
}

func (c *Contents) Reverse() {
	for i, j := 0, len(*c)-1; i < j; i, j = i+1, j-1 {
		(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
	}
}
