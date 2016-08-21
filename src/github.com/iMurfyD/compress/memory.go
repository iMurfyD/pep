package compress

// Instead of writing []byte for matching data of variable length
type Memory []byte

func toMemory(b []byte) Memory {
	var m Memory
	m = append(b[:len(b)])
	return m
}

func toBytes(m Memory) []byte {
	var bytes []byte
	for _, mbit := range m {
		bytes = append(bytes, mbit)
	}
	return bytes
}
