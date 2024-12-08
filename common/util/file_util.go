package util

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Ext(path string) string {
	return filepath.Ext(path)
}

func ValidPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, os.ErrNotExist
	}
	return false, err
}


func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}


func IsFile(path string) bool {
	return !IsDir(path)
}


func WalkPath(path string, suffixes []string) ([]string, error) {
	var files = make([]string, 0)
	if !filepath.IsAbs(path) {
		panic(fmt.Sprintf("Not absolute path:%v", path))
	}
	if valid, err := ValidPath(path); !valid {
		return files, err
	}

	if IsDir(path) {
		err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				fmt.Printf("read file:%v\n", d)
				ext := filepath.Ext(p)
				if contains(suffixes, ext) {
					files = append(files, p)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		ext := filepath.Ext(path)
		if contains(suffixes, ext) {
			files = append(files, path)
		}
	}
	
	return files, nil
} 


func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func GetProjPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	projPath := filepath.Join(pwd, "../..")
	return projPath
}

func GetProjName() string {
	_, projName := filepath.Split(GetProjPath())
	return projName
}
