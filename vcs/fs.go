package vcs

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

type FileSystem struct {
	path      string
	ROOT_HASH []byte
	treeMap   map[string]*Tree
}

func InitFileSystem(path string) *FileSystem {
	fs := FileSystem{
		path,
		[]byte{},
		make(map[string]*Tree),
	}
	root := fs.CreateTree(path)
	fs.ROOT_HASH = root.Hash
	return &fs
}

func (fs *FileSystem) getData(k []byte) *Tree {
	return fs.treeMap[fmt.Sprintf("%x", k)]
}

func (fs *FileSystem) CreateTree(path string) *Tree {
	// Stat gives us info about file
	stat, err := os.Stat(path)
	if err != nil {
		log.Panic(err)
	}
	if !stat.IsDir() {
		tree := Tree{
			[]byte(stat.Name()),
			[]byte{},
			[][]byte{},
		}
		data, err := os.ReadFile(path)
		if err != nil {
			log.Panic(err)
		}
		hash := sha256.Sum256(data)
		tree.Hash = hash[:]
		// DO NOT UNDERSTAND
		fs.treeMap[fmt.Sprintf("%x", tree.Hash)] = &tree
		return &tree
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}
	children := make([][]byte, 0)
	for _, e := range entries {
		if e.Name() == ".fsdb" {
			continue
		}
		tree := fs.CreateTree(path+"/"+e.Name())
		children = append(children, tree.Hash)
	}
	tree := Tree{
		[]byte(stat.Name()),
		[]byte{},
		children,
	}
	tree.CalculateHash()
	fs.treeMap[fmt.Sprintf("%x", tree.Hash)] = &tree
	return &tree
}

