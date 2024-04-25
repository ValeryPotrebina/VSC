package vcs

import (
	"bytes"
	"fmt"
	"log"

	"github.com/golang-collections/collections/set"
	"golang.org/x/exp/slices"
	// "google.golang.org/genproto/googleapis/cloud/functions/v1"
)

type Comparator struct {
	Storage    *Storage
	FileSystem *FileSystem
}

func (cmp *Comparator) Compare() []string {
	s, err := cmp.Storage.GetData(cmp.Storage.ROOT_HASH)
	if err != nil {
		log.Fatal(err)
	}
	root1 := Deserialize(s)
	root2 := cmp.FileSystem.GetData(cmp.FileSystem.ROOT_HASH)
	return cmp.compareTrees(root1, root2)
}

func (cmp *Comparator) compareTrees(tree1 *Tree, tree2 *Tree) []string {
	changedFiles := make([]string, 0)
	if (!bytes.Equal(tree1.Hash, tree2.Hash)) {
		children1 := make([]*Tree, 0)
		children1Names := set.New()
		for _, child := range tree1.Children {
			data, err := cmp.Storage.GetData(child)
			if err != nil {
				log.Panic(data)
			}
			tree := Deserialize(data)
			children1 = append(children1, tree)
			children1Names.Insert(string(tree.Id))
		}
		fmt.Println("children1Names: ", children1Names)

		children2 := make([]*Tree, 0)
		children2Names := set.New()
		for _, child := range tree2.Children {
			tree := cmp.FileSystem.GetData(child)
			children2 = append(children2, tree)
			children2Names.Insert(string(tree.Id))
		}
		fmt.Println("children2Names: ", children2Names)

		intersection := children1Names.Intersection(children2Names)

		// Все что различатся между 1 и 2
		children1Names.Union(children2Names).Difference(intersection).Do(func(i interface{}) {
			changedFiles = append(changedFiles, i.(string))
		})

		intersection.Do(func(i interface{}) {
			fileName := i.(string)
			//WTF?
			t1 := children1[slices.IndexFunc(children1, func(t *Tree) bool {
				return bytes.Equal([]byte(fileName), t.Id)
			})]
			t2 := children2[slices.IndexFunc(children2, func(t *Tree) bool {
				return bytes.Equal([]byte(fileName), t.Id)
			})]
			res := cmp.compareTrees(t1, t2)

			for _, f := range res {
				changedFiles = append(changedFiles, string(tree1.Id) + "/" + f)
			}
		})
	}
	return changedFiles
}