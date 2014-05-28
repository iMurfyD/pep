package compress

//Instead of writing []byte for matching data of variable length
type memory []byte

//represents compressed buffer as a series of int's which each correspond to a block of memory
//once this works I should probably use a more effiecient data type
type CompressedData struct {
	data      []int
	key       map[int]memory
	byte_size int
}

func NewCompressedData(bs int) *CompressedData {
	var c CompressedData
	c.data = []int{}
	c.key = make(map[int]memory)
	c.byte_size = bs
	return &c
}

func (c *CompressedData) Add() {

}
