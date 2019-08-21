package readfile

import (
	"io/ioutil"
	"os"
)

//OpenFile is used to open file from filepath and return []byte data
func OpenFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
