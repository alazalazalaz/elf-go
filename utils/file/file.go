package file

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func TarDeCompress(sourceFile, targetPath string, needCover bool) error {
	if []rune(targetPath)[len([]rune(targetPath)) - 1] != '/' {
		targetPath = targetPath + "/"
	}

	f, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer func() {_ = f.Close()}()

	g, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer func() {_ = g.Close()}()

	tr := tar.NewReader(g)
	for {
		header, err := tr.Next()
		if err != nil{
			if err == io.EOF{
				break
			}else{
				return err
			}
		}

		filename := targetPath + header.Name
		isExist := Exist(filename)
		if isExist == false {
			if !IsFile(filename) {
				if err := os.MkdirAll(filename, 0755); err != nil{
					return err
				}
				log.Printf("decompress add folder : %s", filename)
			}else{
				if err := createFile(filename, tr); err != nil{
					return err
				}
				log.Printf("decompress add file : %s", filename)
			}
		}else{
			if needCover == true {
				if !IsFile(filename) {
					//清空文件夹
					if err := clearFolder(filename); err != nil{
						log.Printf("decompress clear folder's file : %s", filename)
						return err
					}
				}else{
					if err := createFile(filename, tr); err != nil{
						return err
					}
					log.Printf("decompress cover file : %s", filename)
				}
			}else{
				//不需要覆盖的情况无需操作
			}
		}
	}
	return nil
}

func createFile(filename string, tr io.Reader) error {
	file, err := os.Create(filename)
	if err != nil{
		return err
	}
	if _, err := io.Copy(file, tr); err != nil {
		return err
	}
	return nil
}

func clearFolder(folder string)error {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, d := range dir{
		os.RemoveAll(path.Join([]string{folder, d.Name()}...))
	}
	return nil
}

// Exist check the given path exist or not
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsFile check the given path is a file or not
func IsFile(path string) bool{
	if path == filepath.Dir(path) || path == filepath.Dir(path) + "/"{
		return false
	}
	return true
}

//scan all files without recursion
//@param dirname | eg:"aa/bb/cc/" please keep the right slash(/)
func ScanDir(dirname string)([]string, error){
	var files []string
	fdList , err := ioutil.ReadDir(dirname)
	if err != nil {
		return files, err
	}
	for _, fd := range fdList{
		if fd.IsDir() {
			files = append(files, dirname + fd.Name() + "/")
		}else{
			files = append(files, dirname + fd.Name())
		}
	}
	return files, nil
}

//scan all files recursion
//@param dirname | eg:"aa/bb/cc/" please keep the right slash(/)
func ScanDirRecursion(dirname string) ([]string, error){
	var files []string
	fdList , err := ioutil.ReadDir(dirname)
	if err != nil {
		return files, err
	}
	for _, fd := range fdList{
		if fd.IsDir() {
			temp, _ := ScanDirRecursion(dirname + fd.Name() + "/")
			files = append(files, temp...)
		}else{
			files = append(files, dirname + fd.Name())
		}
	}
	return files, nil
}
