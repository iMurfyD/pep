package compress

import (
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
		compress(&data)
	}
}

func compress(uncompressedData *[]byte) (*CompressedData, error) {
	size, err := determineBestByteSize(uncompressedData)
	if err != nil {
		return nil, errors.New("Could not determine best byte size")
	}
	comp := NewCompressedData(size)
	for i := 0; i < len(*uncompressedData); i += size {
		slice := (*uncompressedData)[i:size]
		fmt.Println(slice)
		//will add bit to "comp"
		//fill in later
		//keep Println to check to make sure that it works
	}
	return comp, nil
}

func determineBestByteSize(uncompressedData *[]byte) (int, error) {
	//Just filler for now, will fill in with actual algorthyms later
	if len(*uncompressedData) == 0 {
		return -1, errors.New("No Data")
	} else {
		return len(*uncompressedData), nil
	}
}
