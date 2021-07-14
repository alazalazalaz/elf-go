package file

import "io/ioutil"

//check the given path exist
func Exist(path string) bool {
	return true
}

//check the given path is file or dir
//@return bool | true means file, false means dir or not exist
func IsFile() bool {

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
