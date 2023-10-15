package util

import (
	"syscall"
	"unsafe"

	"github.com/axgle/mahonia"
)

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}

func ConvertToUTF8String(src string, srcCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	return srcResult
}

//func GetDllResult(ADllFile, AServiceName, AComName, AInput string) (string, error) {
//	user32 := syscall.NewLazyDLL(ADllFile)
//	MessageBoxW := user32.NewProc("GetHisData")
//
//	r1, _, _ := MessageBoxW.Call(StrPtr(ConvertToUTF8String(AInput, "utf8")), StrPtr(AServiceName), StrPtr(AComName))
//	//	if err != nil {
//	//		fmt.Println("GetDllResult " + AComName + " Error=" + err.Error())
//	//		return "", err
//	//	}
//
//	p := (*byte)(unsafe.Pointer(r1))
//	// 定义一个[]byte切片，用来存储C返回的字符串
//	data := make([]byte, 0)
//	// 遍历C返回的char指针，直到 '\0' 为止
//	for *p != 0 {
//		data = append(data, *p)         // 将得到的byte追加到末尾
//		r1 += unsafe.Sizeof(byte(0))    // 移动指针，指向下一个char
//		p = (*byte)(unsafe.Pointer(r1)) // 获取指针的值，此时指针已经指向下一个char
//	}
//	name := string(data) // 将data转换为字符串
//
//	LHandle := user32.Handle()
//	if err := syscall.FreeLibrary(syscall.Handle(LHandle)); err != nil {
//		fmt.Println(err)
//	}
//
//	return name, nil
//
//}
