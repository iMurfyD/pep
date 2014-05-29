package compress

import (
	"bytes"
)

//Instead of writing []byte for matching data of variable length
type memory []byte

//represents compressed buffer as a series of int's (index of memory) which each correspond to a block of memory
//once this works I should probably use a more effiecient data type
type CompressedData struct {
	data      []int
	key       []memory
	byte_size int
}

func NewCompressedData(bs int) *CompressedData {
	var c CompressedData
	c.data = []int{}
	c.key = []memory{}
	c.byte_size = bs
	return &c
}

func (c *CompressedData) Add(m memory) {
	if c.isNew(m) {
		c.key = append(c.key, m)
		c.data = append(c.data, len(c.key))
	} else {
		c.data = append(c.data, find(c.key, m))
	}
}

func find(arr []memory, val memory) int {
	found := false
	indx := -1
	for !found {
		indx++
		if bytes.Equal(arr[indx], val) {
			found = true
		}
	}
	return indx
}

func (c *CompressedData) isNew(m memory) bool {
	var is bool = true
	for i := 0; i < len(c.key); i++ {
		if bytes.Equal(c.key[i], m) {
			is = false
		}
	}
	return is
}

func (c *CompressedData) String() string {
	var ret string
	for _, val := range c.data {
		ret += string(val)
		ret += " "
	}
	return ret
}
