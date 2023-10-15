package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"github.com/bwmarrin/snowflake"
	"github.com/cnbattle/wubi"
	"github.com/satori/go.uuid"
	"regexp"
	"strings"
	"time"
)

//生成his密码
func NewHisMd5Pwd(input string) (pwd string) {
	var (
		fixByte []byte
		err     error
	)
	for _, b := range []byte(input) {
		fixByte = append(fixByte, b)
		fixByte = append(fixByte, 0)
	}
	//MD5算法. 
	hash := md5.New()
	if _, err = hash.Write(fixByte); err != nil {
		fmt.Println(err)
		return
	}
	for _, b := range hash.Sum([]byte("")) {
		pwd += fmt.Sprintf("%02x", b)
	}
	return
}

//登录token
func NewUserToken(userId, password string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%s%d", userId, password, time.Now().Unix())))
	return hex.EncodeToString(h.Sum(nil))
}

//生成随机串
func NewShortUUId() string {
	uuid := uuid.NewV4()
	escaper := strings.NewReplacer("9", "99", "-", "90", "_", "91")
	ret := escaper.Replace(base64.RawURLEncoding.EncodeToString(uuid.Bytes()))
	return ret
}

//生成数字流水号
func NewSnowFlakeID() (id string) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	id = node.Generate().String()
	return
}

//打码：身份证、姓名、手机
func EncryptField(input string) (output string) {
	output = input
	regsfz := regexp.MustCompile(`^(^[1-9]\d{5}(18|19|([23]\d))\d{2}(((0[13578]|10|12)(0[1-9]|[12]\d|3[01]))|((0[469]|11)(0[1-9]|[12]\d|30))|(02(0[1-9]|[12]\d)))\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}(((0[13578]|10|12)(0[1-9]|[12]\d|3[01]))|((0[469]|11)(0[1-9]|[12]\d|30))|(02(0[1-9]|[12]\d)))\d{3}$)`)
	regxm := regexp.MustCompile("[\u4e00-\u9fa5]")
	regdhhm := regexp.MustCompile(`^1([38]\d|5[0-35-9]|7[3678])\d{8}$`)
	if result := string(regsfz.Find([]byte(string(input)))); len(result) > 0 {
		output = result[0:5] + `******` + result[11:]
		return

	} else if result := regxm.FindAll([]byte(input), -1); len(result) > 1 {
		var m = []byte{}
		if len(result) == 2 {
			m = append(m, []byte("*")...)
			m = append(m, result[1]...)
			output = string(m)
			return
		} else if len(result) > 2 {
			m = append(m, result[0]...)
			m = append(m, []byte("*")...)
			for i := 2; i < len(result); i++ {
				m = append(m, result[i]...)
			}
			output = string(m)
			return
		} else {
			output = string(result[0])
			return
		}
	} else if result := string(regdhhm.Find([]byte(string(input)))); len(result) > 0 {
		m := result[0:5] + `***` + result[8:]
		output = m
		return
	} else {
		output = input
		return
	}
	return
}

//生成拼音码
func Pycode(in string) (out string, err error) {
	in = DBCtoSBC(in)
	str, err := pinyin.New(in).Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`[A-Z0-9]`)
	result := reg.FindAll([]byte(str), -1)
	for _, i := range result {
		out += string(i)
	}
	return
}

//生成五笔码
func Wucode(in string) (out string, err error) {
	in = DBCtoSBC(in)
	var strbuf []string
	strbuf, err = wubi.Get(in)
	if err != nil {
		return
	}
	for _, i := range strbuf {
		out += strings.ToUpper(i[0:1])
	}
	return
}

//全角转换半角
func DBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}
	return retstr
}
