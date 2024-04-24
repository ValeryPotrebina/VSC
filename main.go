package main

import (
	"blockchain/vcs"
)

func main() {
	storage := vcs.InitStorage("tmp")
	defer storage.CloseStorage()
	storage.FindDiffs()
}
