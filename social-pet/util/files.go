package util

//import (
//	"fmt"
//	"io/ioutil"
//	"os"
//	"path"
//	"runtime"
//	"strings"
//)
//
//func GetFileModTime(path string) string {
//	f, err := os.Open(path)
//	if err != nil {
//		return ""
//	}
//	defer f.Close()
//
//	fi, err := f.Stat()
//	if err != nil {
//		return ""
//	}
//
//	LFileTime := fi.ModTime().Format("2006-01-02 15:04:05")
//	LFileTime = strings.Replace(LFileTime, "-", "", -1)
//	LFileTime = strings.Replace(LFileTime, " ", "", -1)
//	LFileTime = strings.Replace(LFileTime, ":", "", -1)
//	return LFileTime
//}
//
//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
//
//func GetDirAllFile(pathname string, extname string, includedir bool) ([]string, error) {
//	var (
//		LFiles, LFiles1 []string
//		LPathChar       string
//	)
//
//	if runtime.GOOS == "windows" {
//		LPathChar = "\\"
//	} else {
//		LPathChar = "//"
//	}
//
//	if includedir {
//		LFiles = append(LFiles, pathname)
//	}
//
//	LFiles = make([]string, 0)
//	LFiles1 = make([]string, 0)
//	rd, err := ioutil.ReadDir(pathname)
//	for _, fi := range rd {
//		if fi.IsDir() {
//			LFiles1, _ = GetDirAllFile(pathname+LPathChar+fi.Name(), extname, includedir)
//
//			for _, str := range LFiles1 {
//				LFiles = append(LFiles, str)
//			}
//		} else {
//			if extname != "" {
//				Lextname := path.Ext(fi.Name())
//				if (strings.ToLower(Lextname) == strings.ToLower(extname)) || ("."+strings.ToLower(Lextname) == strings.ToLower(extname)) {
//					LFiles = append(LFiles, pathname+LPathChar+fi.Name())
//				}
//			} else {
//				LFiles = append(LFiles, pathname+LPathChar+fi.Name())
//			}
//		}
//	}
//	return LFiles, err
//}
//
//func GetDirAllSubDir(pathname string) ([]string, error) {
//	var (
//		LFiles, LFiles1 []string
//		LPathChar       string
//	)
//
//	if runtime.GOOS == "windows" {
//		LPathChar = "\\"
//	} else {
//		LPathChar = "//"
//	}
//
//	LFiles = append(LFiles, pathname)
//
//	LFiles = make([]string, 0)
//	LFiles1 = make([]string, 0)
//	rd, err := ioutil.ReadDir(pathname)
//	for _, fi := range rd {
//		if fi.IsDir() {
//			LFiles1, _ = GetDirAllSubDir(pathname + LPathChar + fi.Name())
//
//			for _, str := range LFiles1 {
//				LFiles = append(LFiles, str)
//			}
//		}
//	}
//	return LFiles, err
//}
//
//func IsNullDir(localPath string) bool {
//	rd, err := ioutil.ReadDir(localPath)
//	if err != nil {
//		fmt.Println("IsNullDir Error=" + err.Error())
//		return true
//	}
//	return len(rd) == 0
//}
//
//func DeleteFileOnDisk(localPath string) error {
//	LFiles, err := GetDirAllFile(localPath, "", false)
//	if err != nil {
//		fmt.Println("DeleteFileOnDisk Error1=" + err.Error())
//		return err
//	}
//
//	for _, LFile := range LFiles {
//		err = os.Remove(LFile)
//		if err != nil {
//			fmt.Println("DeleteFileOnDisk RemoveFile Error=" + err.Error())
//		}
//	}
//
//	fmt.Println("localPath=" + localPath)
//	rd, err := ioutil.ReadDir(localPath)
//	for _, fi := range rd {
//		fmt.Println(fi.Name())
//		if fi.IsDir() {
//			if !IsNullDir(localPath + "\\" + fi.Name()) {
//				fmt.Println("not null dir")
//				for {
//					DeleteFileOnDisk(localPath + "\\" + fi.Name())
//					if IsNullDir(localPath + "\\" + fi.Name()) {
//						os.Remove(localPath + "\\" + fi.Name())
//						break
//					}
//				}
//			} else {
//				fmt.Println("null dir=" + localPath + "\\" + fi.Name())
//				os.Remove(localPath + "\\" + fi.Name())
//			}
//		}
//	}
//
//	os.Remove(localPath)
//	return err
//
//}
