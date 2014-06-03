package compress

import (
	"bytes"
)

//represents compressed buffer as a series of int's (index of Memory) which each correspond to a block of Memory
//once this works I should probably use a more effiecient data type
type CompressedData struct {
	data      []int
	key       []Memory
	byte_size int
}

func NewCompressedData(bs int) *CompressedData {
	var c CompressedData
	c.data = []int{}
	c.key = []Memory{}
	c.byte_size = bs
	return &c
}

func (c *CompressedData) Add(m Memory) {
	if c.isNew(m) {
		c.key = append(c.key, m)
		c.data = append(c.data, len(c.key))
	} else {
		c.data = append(c.data, find(c.key, m))
	}
}

func find(arr []Memory, val Memory) int {
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

func (c *CompressedData) isNew(m Memory) bool {
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
