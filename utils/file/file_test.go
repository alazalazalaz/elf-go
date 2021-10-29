package file

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	dirname := "../../"
	files, err := ScanDir(dirname)
	for _, v := range files {
		fmt.Println(v)
	}
	fmt.Println(err)
	Exist("f")
	files2, err := ScanDirRecursion(dirname)
	for _, v := range files2 {
		fmt.Println(v)
	}
	fmt.Println(err)
}
