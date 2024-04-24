package vcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Tree struct {

	Id 			[]byte
	Hash 		[]byte
	Children 	[][]byte
}

func (t *Tree) GetHeader() []byte {
	// rootc1c2c3 
	return bytes.Join(append([][]byte{t.Id}, t.Children...,), []byte{})
}

func (t *Tree) CalculateHash() {
	header := t.GetHeader()
	hash := sha256.Sum256(header)
	t.Hash = hash[:]
}

func (t *Tree) String() string {
	var res string
	res += fmt.Sprintf("ID:    %s\n", t.Id)
	res += fmt.Sprintf("HASH:  %x\n", t.Hash)
	res += "CHILDREN: "
	for i, hash := range t.Children {
		if i == 0 {
			res += fmt.Sprintf("%x\n", hash)
		} else {
			res += fmt.Sprintf("          %x\n", hash)
		}
	}
	return res
}

func Serialize(tree *Tree) []byte{
	var b bytes.Buffer 
	encoder := gob.NewEncoder(&b)

	err := encoder.Encode(tree)
	if err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}

func Deserialize(data []byte) *Tree{
	var tree Tree
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&tree)

	if err != nil {
		log.Fatal(err)
	}

	return &tree
}