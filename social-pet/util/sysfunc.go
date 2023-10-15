package util

import (
	"bufio"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const DateTimeFormat string = "2006-01-02 15:04:05"

var (
	base64Table        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	coder              = base64.NewEncoding(base64Table)
	XorKey      []byte = []byte{0xC2, 0xD9, 0xBB, 0x55, 0x13, 0x6D, 0x44, 0x47}
)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func Base64UnCode(src string) (string, error) {
	Lret, err := base64Decode([]byte(src))
	if err != nil {
		return "", err
	}

	return string(Lret), nil

}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId1() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(GetMd5String(base64.URLEncoding.EncodeToString(b)))
}

func GetValidByte(src []byte) []byte {
	var str_buf []byte
	for _, v := range src {
		if v != 0 {
			str_buf = append(str_buf, v)
		}
	}
	return str_buf
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("checkErr:" + err.Error())
		//		os.Exit(-1)
	}
}

func GetWxPaySign(AParamStr string, AKey string) string {
	LRet := []byte(AParamStr + "&key=" + AKey)
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum(LRet)))
}

func GetAPPRootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ExecCommand(commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return err
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return nil
}

func XorEncode(src string) string {
	var result string
	j := 0
	s := ""

	ret := base64.StdEncoding.EncodeToString([]byte(src))

	bt := []rune(ret)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}

	return strings.ToUpper(result)

}

func XorDecode(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	//	fmt.Println(bt)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}

	ret, _ := base64.StdEncoding.DecodeString(result)

	return string(ret)
}

func InterfaceToStr(AInter interface{}) string {
	if AInter == nil {
		return ""
	} else {
		return AInter.(string)
	}
}

func InterfaceToInt(AInter interface{}) int {
	if AInter == nil {
		return 0
	} else {
		return int(AInter.(float64))
	}
}

func ReadWithIoutil(name string, encode bool) string {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("Use ioutil.ReadFile to read a file:", result)
		if encode {
			result = XorDecode(result)
		}
		//fmt.Println(name + "=" + result)
		return result
	}

	return ""

}

func ReadObject(FileName string, Object interface{}, encode bool) bool {
	ret := ReadWithIoutil(FileName, encode)
	if strings.TrimSpace(ret) == "" {
		return false
	}

	//	fmt.Println("ReadObject Ret=" + ret)
	err := json.Unmarshal([]byte(ret), Object)
	if err != nil {
		fmt.Printf("ReadObject Error=%s\n", err.Error())
		return false
	}
	return true

}

func WriteWithIoutil(name, content string, encode bool) bool {
	if encode {
		content = XorEncode(content)
	}
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		return true
	} else {
		return false
	}

}

func WriteObject(FileName string, Object interface{}, encode bool) bool {
	ret, err := json.Marshal(Object)
	if err != nil {
		fmt.Printf("WriteObject Error=%s\n", err.Error())
		return false
	}

	return WriteWithIoutil(FileName, string(ret), encode)
}

func GetCommandRet(cmdstr string) (string, error) {
	//如果为空需要安装yum install e2fsprogs-devel.x86_64
	cmd := exec.Command("/bin/sh", "-c", cmdstr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("GetCommandRet error=" + err.Error())
		return "", err
	}
	return string(out), nil

}

func GetCpuId() string {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return ""
		}

		str := string(out)
		reg := regexp.MustCompile("\\s+")
		str = reg.ReplaceAllString(str, "")
		return str[11:]
	} else {
		//如果为空需要安装yum install e2fsprogs-devel.x86_64
		out, err := GetCommandRet("cat /sys/class/dmi/id/product_uuid")
		if err != nil {
			out, err = GetCommandRet("ls /dev/disk/by-uuid")
			if err != nil {
				return ""
			}

			LPos := strings.Index(out, "\n")
			if LPos > 0 {
				out = out[:LPos]
				out = strings.TrimSpace(out)
				out = strings.ReplaceAll(out, "-", "")
				return strings.ToLower(out)
			}
			return ""

		}
		return out

	}
}

func StrFirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func StrFirstToLower(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 { // 后文有介绍
				vv[i] += 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
