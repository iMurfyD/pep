package compress

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
		//TODO
		//Don't know what type of file permission to pass in
		ioutil.WriteFile(toFile, m, os.ModeDevice)
	}
}

func compress(uncompressedData *Memory) (*CompressedData, error) {
	size, err := determineBestByteSize(uncompressedData)
	fmt.Printf("Best BS: %d\n", size)
	if err != nil {
		return nil, errors.New("Could not determine best byte size")
	}
	comp := NewCompressedData(size)
	fmt.Println("Made comp")
	for i := 0; i < len(*uncompressedData); i += size {
		fmt.Printf("Adding slice to comp")
		var slice Memory

		//determine if current slice is end slice (not to go out of bounds)
		if i+size > len(*uncompressedData) {
			slice = (*uncompressedData)[i:]
		} else {
			slice = (*uncompressedData)[i:size]
		}

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
	for i := 1; i <= length/2; i++ {
		fmt.Printf("Running BS: %d\n", i)
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
	number_of_tokens = 0
	for i := 0; i < len(*data); i += bs {
		m := toMemory(toBytes(*data)[i : i+bs])
		if isNewForTokenList(&tokens_list, &m) {
			fmt.Printf("     Detected new Token(%s)\n", m)
			tokens_list = append(tokens_list, m)
			number_of_tokens++
		}
	}
	fmt.Printf("Number of Tokens (BS: %d): %d\n", bs, number_of_tokens)
	return number_of_tokens
}
