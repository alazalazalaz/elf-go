package appfile

import (
	"elf-go/components/applogs"
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

func TestDownloadFile(t *testing.T) {

	//imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	imgUrl := "https://api.paypal.com/v1/notifications/certs/CERT-360caa42-fca2a594-2f1e8d33"
	data, err := DownloadFile(imgUrl)
	if err != nil {
		applogs.Errorf("download file error, err:%v", err)
		return
	}

	applogs.Infof("download success, data:%s", string(data))
}
