package main

import (
    // "fmt"
    // "os"
    // "io"
	// "github.com/sergi/go-diff/diffmatchpatch"
	"blockchain/vcs"
)



func main() {

	s := vcs.InitStorage("files")
	commits := make([]vcs.Commit, 0)
	commits = append(commits, vcs.Commit{
		Description: "test",
		Author:      "a",
		Timestamp:        1000,
		Changes: []vcs.FileChange{
			{
				FilePath: "a.txt",
				AddInfo: []vcs.Fragment{
					{
						StartIndex: 0,
						EndIndex:   1,
						Data:  []byte("aboba"),
					},
				},
				RemoveInfo: []vcs.Fragment{},
			},
		},
	})
	commits = append(commits, vcs.Commit{
		Description: "test",
		Author:      "a",
		Timestamp:        1000,
		Changes: []vcs.FileChange{
			{
				FilePath: "a.txt",
				AddInfo: []vcs.Fragment{
					{
						StartIndex: 0,
						EndIndex:   2,
						Data:  []byte("abyba\nabuba"),
					},
				},
				RemoveInfo: []vcs.Fragment{
					{
						StartIndex: 0,
						EndIndex:   1,
						Data:  []byte("aboba"),
					},
				},
			},
		},
	})
	s.RestoreByCommits(commits)
	// file1, err := os.Open("files/a.txt")
	
    // if err != nil{
    //     fmt.Println(err) 
    //     os.Exit(1) 
    // }
    // defer file1.Close() 
     
    // data1 := make([]byte, 64)
     
    // for{
    //     n, err := file1.Read(data1)
    //     if err == io.EOF{   // если конец файла
    //         break           // выходим из цикла
    //     }
    //     fmt.Print(string(data1[:n]))
    // }
	// file2, err := os.Open("files/b.txt")
	
    // if err != nil{
    //     fmt.Println(err) 
    //     os.Exit(1) 
    // }
    // defer file2.Close() 
     
    // data2 := make([]byte, 64)
     
    // for{
    //     n, err := file2.Read(data2)
    //     if err == io.EOF{   // если конец файла
    //         break           // выходим из цикла
    //     }
    //     fmt.Print(string(data2[:n]))
    // }

	// fmt.Println()
	// fmt.Println("---------")
	// dmp := diffmatchpatch.New()

    // diffs := dmp.DiffMain(string(data1), string(data2), true)
    // fmt.Println(dmp.DiffCleanupSemantic(diffs))
	// fmt.Println(dmp.DiffPrettyText(diffs))
	// fmt.Println(dmp.DiffLinesToChars(string(data1), string(data2)))

}
