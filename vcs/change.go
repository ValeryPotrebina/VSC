package vcs

type Fragment struct {
	StartIndex int
	EndIndex   int
	Data       []byte
}

type FileChange struct {
	FilePath   string
	AddInfo    []Fragment
	RemoveInfo []Fragment
}

type Commit struct {
	Description string
	Author      string
	Timestamp   int64 //change
	Changes     []FileChange
}