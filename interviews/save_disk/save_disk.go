package main

import (
	"crypto/sha256"
	"log"
	"os"
	"path/filepath"
)

func saveDisk(dirPath string) error {
	hashMap := map[[32]byte]string{}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == true {
			return nil
		}
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		oldPath, exists := hashMap[sha256.Sum256(fileBytes)]
		if !exists {
			hashMap[sha256.Sum256(fileBytes)] = path
		} else {
			err = os.Remove(path)
			if err != nil {
				return err
			}
			err = os.Symlink(oldPath, path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := saveDisk(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
