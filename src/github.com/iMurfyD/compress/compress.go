package compress

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

func To(fromFile, toFile string) {
	//compresses data from fromFile to toFile
	data, err := ioutil.ReadFile(fromFile)
	if err != nil {
		fmt.Println("File is screwed up or does not exist")
	} else {
		fmt.Println("Compressing thingy", data)
		m := toMemory(data)
		compress(&m)
	}
}

func compress(uncompressedData *Memory) (*CompressedData, error) {
	size, err := determineBestByteSize(uncompressedData)
	if err != nil {
		return nil, errors.New("Could not determine best byte size")
	}
	comp := NewCompressedData(size)
	for i := 0; i < len(*uncompressedData); i += size {
		slice := (*uncompressedData)[i:size]
		fmt.Println(slice)
		comp.Add(slice)
		fmt.Println(comp)
		//TODO
		//will add "slice" to "comp" properly
		//fill in later
		//keep Println to check to make sure that it works
	}
	return comp, nil
}

func determineBestByteSize(uncompressedData *Memory) (int, error) {
	//represents pointer to uncompressedData
	uncompressedDataPt := *uncompressedData
	//Just filler for now, will fill in with actual algorthyms later
	var length int
	//checks to see if ends in newline, will disregard newline if it is there in algorthym
	if uncompressedDataPt[len(uncompressedDataPt)-1] == byte(10) {
		length = len(uncompressedDataPt) - 1
	} else {
		length = len(uncompressedDataPt)
	}

	fmt.Println(uncompressedDataPt[:length])

	var tokenNumbersAtVariousByteSizes []byteSize
	for i := 0; i < length-2; i++ {
		//TODO
		//make numOfTokens run in a seperate goroutine
		tokenNumbersAtVariousByteSizes = append(tokenNumbersAtVariousByteSizes, byteSize{tokens: numOfTokens(uncompressedData, i), bs: i})
	}
	max := maxByteSize(tokenNumbersAtVariousByteSizes)
	return max.bs, nil
}

func isNewForTokenList(tokens *[]Memory, data *Memory) bool {
	for _, t := range *tokens {
		if bytes.Equal(t, toBytes(*data)) {
			return false
		}
	}
	return true
}

//figures out the number of tokens in an bunch of Memory with a given byte size
func numOfTokens(data *Memory, bs int) int {
	var tokens_list []Memory
	var number_of_tokens int
	for i := 0; i < len(*data); i += bs {
		m := toMemory(toBytes(*data)[i : i+bs])
		if isNewForTokenList(&tokens_list, &m) {
			tokens_list = append(tokens_list, m)
			number_of_tokens++
		}
	}
	return number_of_tokens
}
