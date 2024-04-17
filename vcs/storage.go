package vcs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Storage struct {
	Path    string
	Commits []Commit
}

func InitStorage(path string) Storage {
	if _, err := os.Open(path); err == os.ErrNotExist {
		os.Mkdir(path, os.ModePerm)
	}

	return Storage{
		path,
		make([]Commit, 0),
	}
}

// добавление коммитов в хранилище storage
func (storage *Storage) AddCommit(commit Commit) {
	storage.Commits = append(storage.Commits, commit)
}

func (storage *Storage) RestoreByCommits(commits []Commit) {
	// Почему 0?
	storage.Commits = make([]Commit, 0)
	// readDirectory, _ := os.Open(storage.Path)
	// ??
	// allFiles, _ := readDirectory.ReadDir(0)

	// for f := range allFiles {
	// 	file := allFiles[f]
	// 	fileName := file.Name()
	// 	filePath := storage.Path + "/" + fileName

	// 	os.Remove(filePath)
	// 	fmt.Println("Deleted file:", filePath)
	// }

	for _, c := range commits {
		storage.AddCommit(c)
		storage.ApplyCommit(c)
	}

}

func (storage *Storage) ApplyCommit(commit Commit) {
	for _, fileChange := range commit.Changes {
		var b []byte
		// OpenFile посмотреть 
		file, err := os.OpenFile(storage.Path + "/" + fileChange.FilePath, 1, os.ModePerm)
		if err != os.ErrNotExist {
			file.Read(b)
		}
		text := string(b)
		lines := strings.Split(text, "\n")
		fmt.Println("Lines: ", lines)
		remains := make([][]int, 0)
		for i, r := range fileChange.RemoveInfo {
			if i == 0 {
				remains = append(remains, []int{0, r.StartIndex})
				continue
			}
			remains = append(remains, []int{fileChange.RemoveInfo[i-1].EndIndex, r.StartIndex})
			if i == len(fileChange.RemoveInfo) - 1 {
				remains = append(remains, []int{r.EndIndex, len(lines)})
			}
		}
		fmt.Println("remains: ", remains)
		 
		maped := make([]string, 0)
		for _, r := range remains {
			// r хранит индекс начала и конца того, что оставляем от lines
			fmt.Println(r)
			maped = append(maped, lines[r[0]:r[1]]...)
		}

		for _, a := range fileChange.AddInfo {
			fmt.Println(a)
			// ??
			maped = append(append(maped[0:a.StartIndex], strings.Split(string(a.Data), "\n")...), maped[a.StartIndex:]...)
		}

		result := strings.Join(maped, "\n")
		file.Write([]byte(result))
		file.Close()
		fmt.Println("\"", result, "\"")

		if result == "" {
			err := os.Remove(storage.Path + "/" + fileChange.FilePath)
			if err != nil {
				log.Panic(err)
			}
		}
	}

}