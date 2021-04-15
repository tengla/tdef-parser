package serializer

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

// ToBin
func ToBin(filename string, v interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

// FromBin
func FromBin(filename string, v interface{}) error {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	dec := gob.NewDecoder(reader)
	defer file.Close()
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}
