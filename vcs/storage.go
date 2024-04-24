package vcs

import (
	"bytes"
	"github.com/dgraph-io/badger/v4"
	"log"
	"os"
)

type Storage struct {
	DB        *badger.DB
	ROOT_HASH []byte
	Path      string
	Commits   []Commit
}

func InitStorage(path string) Storage {
	if _, err := os.Open(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
	}

	// Открываем бдшку
	opt := badger.DefaultOptions(path + "/.fsdb")
	db, err := badger.Open(opt)
	if err != nil {
		log.Panic(err)
	}

	storage := Storage{
		db,
		[]byte{},
		path,
		[]Commit{},
	}

	rootHash, err := storage.GetData([]byte("ROOT_HASH"))
	if err == badger.ErrKeyNotFound {
		log.Println("ROOT_HASH NOT FOUND. INITIALIZING ROOH_HASH...")
		rootObj := Tree{
			[]byte("ROOT"),
			[]byte(""),
			[][]byte{},
		}
		rootObj.CalculateHash()
		err := storage.SetData(rootObj.Hash, Serialize(&rootObj))
		if err != nil {
			log.Panic(err)
		}
		//
		err = storage.SetData([]byte("ROOT_HASH"), rootObj.Hash)
		if err != nil {
			log.Panic(err)
		}
		storage.ROOT_HASH = rootObj.Hash
		return storage
	}
	log.Printf("ROOT_HASH found: %x", rootHash)
	storage.ROOT_HASH = rootHash
	return storage
}

func (s *Storage) GetData(k []byte) ([]byte, error) {
	var hash []byte
	err := s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(k)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			hash = val
			return nil
		})
		return err
	})
	return hash, err
}

func (s *Storage) SetData(k []byte, v []byte) error {
	err := s.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		return err
	})
	return err
}

func (s *Storage) CloseStorage() {
	s.DB.Close()
}


func (s *Storage) FindDiffs() {
	fs := InitFileSystem(s.Path)
	if !bytes.Equal(fs.ROOT_HASH, s.ROOT_HASH) {
		log.Println("CHANGES DETECTED")
	}
}



