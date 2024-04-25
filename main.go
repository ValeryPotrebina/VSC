package main

import (
	"blockchain/vcs"
	"fmt"
)

func main() {
	storage := vcs.InitStorage("tmp")
	defer storage.CloseStorage()
	diffs := storage.FindDiffs()
	for _, df := range diffs {
		fmt.Println("df: ", df)
	}

	// set1 := set.New()
	// set1.Insert(1)
	// set1.Insert(2)
/// 1 2 3
/// 2
	// set2 := set.New()
	// set2.Insert(2)
	// set2.Insert(3)


	// inter := set1.Intersection(set2)
	// set1.Union(set2).Difference(inter).Do(func(i interface{}) {
	// 	fmt.Println(i)
	// })
	
}
