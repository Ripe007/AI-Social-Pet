package bmodel

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strings"
	"time"
)

const (
	LocalDateTimeFormat string = "2006-01-02 15:04:05"
	DateTimeFormat      string = "2006-01-02"
)

type LocalTime time.Time

//MarshalJSON jsonTime序列化调用的方法
//func (l LocalTime) MarshalJSON() ([]byte, error) {
//	b := make([]byte, 0, len(LocalDateTimeFormat)+2)
//	b = append(b, '"')
//	b = time.Time(l).AppendFormat(b, LocalDateTimeFormat)
//	b = append(b, '"')
//	if string(b) == `"0001-01-01 00:00:00"` {
//		b = []byte(`""`)
//	}
//	return b, nil
//}
//
//func (l *LocalTime) UnmarshalJSON(b []byte) error {
//	if len(b) == 2 {
//		return nil
//	}
//	input := string(b)
//	if input == "null" {
//		return nil
//	}
//	now, err := time.ParseInLocation(`"`+LocalDateTimeFormat+`"`, string(b), time.Local)
//	*l = LocalTime(now)
//	return err
//}

func (l LocalTime) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(LocalDateTimeFormat))
	//b = append(b, '"')
	b = time.Time(l).AppendFormat(b, LocalDateTimeFormat)
	//b = append(b, '"')
	if string(b) == `0001-01-01 00:00:00` {
		b = []byte(``)
	}
	return b, nil
}

func (l *LocalTime) UnmarshalText(b []byte) error {
	if len(b) == 2 || b == nil {
		return nil
	}
	input := string(b)
	if input == "null" || input == "" {
		return nil
	}
	now, err := time.ParseInLocation(LocalDateTimeFormat, string(b), time.Local)
	*l = LocalTime(now)
	return err
}

func (l LocalTime) String() string {
	t := time.Time(l)
	if t.IsZero() {
		return ""
	}
	return t.Format(LocalDateTimeFormat)
}

func (l LocalTime) ToDateTime() DateTime {
	date := time.Time(l).Format("2006-01-02")
	return NewDateTime(date)
}

func NewNowLocalTime() LocalTime {
	return LocalTime(time.Now())
}

//通过日期格式生成
func NewLocalTime(date string) LocalTime {
	var (
		t   time.Time
		err error
	)
	if len(date) == 0 {
		t = time.Time{}
		return LocalTime(t)
	}
	t, err = time.ParseInLocation(LocalDateTimeFormat, date, time.Local)
	if err != nil {
		logs.Error(err)
		t = time.Time{}
	}
	return LocalTime(t)
}

//通过身份证号获取出生年月
func NewLocalTimeWithSocialNum(socialNum string) (LocalTime, error) {
	birth, err := GetSocialNumBirth(socialNum)
	if err != nil {
		return LocalTime{}, err
	}
	return birth.ToLocalTime(), nil
}

func (l LocalTime) GetTime() time.Time {
	return time.Time(l)
}

//添加日期
func (l *LocalTime) AddDate(years int, months int, days int) string {
	*l = NewLocalTime(l.GetTime().AddDate(years, months, days).Format(LocalDateTimeFormat))
	return ""
}

//添加时间戳
func (l *LocalTime) Add(d time.Duration) {
	*l = NewLocalTime(l.GetTime().Add(d).Format(LocalDateTimeFormat))
}

//比较时间教前
func (l LocalTime) Before(t LocalTime) bool {
	return l.GetTime().Before(t.GetTime())
}

//比较时间相等
func (l LocalTime) Equal(t LocalTime) bool {
	return l.GetTime().Equal(t.GetTime())
}

//是否今年
func (l LocalTime) IsNowYear() bool {
	dYear, _, _ := l.GetTime().Date()
	nYear, _, _ := time.Now().Date()
	return dYear == nYear
}

//是否这个月
func (l LocalTime) IsNowMonth() bool {
	dYear, dMonth, _ := l.GetTime().Date()
	nYear, nMonth, _ := time.Now().Date()
	return dYear == nYear &&
		dMonth == nMonth
}

//是否今天
func (l LocalTime) IsToday() bool {
	dYear, dMonth, dDay := l.GetTime().Date()
	nYear, nMonth, nDay := time.Now().Date()
	return dYear == nYear &&
		dMonth == nMonth &&
		dDay == nDay
}

type DateTime time.Time //日期类型

//MarshalJSON jsonTime序列化调用的方法
//func (l DateTime) MarshalJSON() ([]byte, error) {
//	b := make([]byte, 0, len(DateTimeFormat)+2)
//	b = append(b, '"')
//	b = time.Time(l).AppendFormat(b, DateTimeFormat)
//	b = append(b, '"')
//	if string(b) == `"0001-01-01"` {
//		b = []byte(`""`)
//	}
//	return b, nil
//}
//
//func (l *DateTime) UnmarshalJSON(b []byte) error {
//	if len(b) == 2 {
//		return nil
//	}
//	input := string(b)
//	if input == "null" {
//		return nil
//	}
//	now, err := time.ParseInLocation(`"`+DateTimeFormat+`"`, input, time.Local)
//	now.UnixNano()
//	*l = DateTime(now)
//	return err
//}

func (dateTime DateTime) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(DateTimeFormat))
	//b = append(b, '"')
	b = time.Time(dateTime).AppendFormat(b, DateTimeFormat)
	//b = append(b, '"')
	if string(b) == `0001-01-01` {
		b = []byte(``)
	}
	return b, nil
}

func (dateTime *DateTime) UnmarshalText(b []byte) error {
	if len(b) == 2 || b == nil {
		return nil
	}
	input := string(b)
	if input == "null" || input == "" {
		return nil
	}
	now, err := time.ParseInLocation(DateTimeFormat, string(b), time.Local)
	*dateTime = DateTime(now)
	return err
}

func (dateTime DateTime) String() string {
	t := time.Time(dateTime)
	if t.IsZero() {
		return ""
	}
	return t.Format(DateTimeFormat)
}

func (dateTime DateTime) ToLocalTime() LocalTime {
	date := time.Time(dateTime).Format(LocalDateTimeFormat)
	return NewLocalTime(date)
}

func (dateTime DateTime) GetTime() time.Time {
	return time.Time(dateTime)
}

func NewNowDateTime() DateTime {
	return NewDateTime(time.Now().Format(DateTimeFormat))
}

//通过身份证号获取出生年月
func NewDateTimeWithSocialNum(socialNum string) (DateTime, error) {
	return GetSocialNumBirth(socialNum)
}

func NewDateTime(date string) DateTime {
	var (
		t   time.Time
		err error
	)
	if len(date) == 0 {
		t = time.Time{}
		return DateTime(t)
	}
	t, err = time.ParseInLocation(DateTimeFormat, date, time.Local)
	if err != nil {
		logs.Error(err)
		t = time.Time{}
	}
	return DateTime(t)
}

//添加日期
func (dateTime *DateTime) AddDate(years int, months int, days int) {
	*dateTime = NewDateTime(dateTime.GetTime().AddDate(years, months, days).Format(DateTimeFormat))
}

//添加时间戳
func (dateTime *DateTime) Add(d time.Duration) {
	*dateTime = NewDateTime(dateTime.GetTime().Add(d).Format(DateTimeFormat))
}

//比较时间教前
func (dateTime DateTime) Before(t DateTime) bool {
	return dateTime.GetTime().Before(t.GetTime())
}

//比较时间相等
func (dateTime DateTime) Equal(t DateTime) bool {
	return dateTime.GetTime().Equal(t.GetTime())
}

//是否今年
func (dateTime DateTime) IsNowYear() bool {
	dYear, _, _ := dateTime.GetTime().Date()
	nYear, _, _ := time.Now().Date()
	return dYear == nYear
}

//是否这个月
func (dateTime DateTime) IsNowMonth() bool {
	dYear, dMonth, _ := dateTime.GetTime().Date()
	nYear, nMonth, _ := time.Now().Date()
	return dYear == nYear &&
		dMonth == nMonth
}

//是否今天
func (dateTime DateTime) IsToday() bool {
	dYear, dMonth, dDay := dateTime.GetTime().Date()
	nYear, nMonth, nDay := time.Now().Date()
	return dYear == nYear &&
		dMonth == nMonth &&
		dDay == nDay
}

//通过身份证获取出生年月
func GetSocialNumBirth(socialNum string) (date DateTime, err error) {
	socialNum = strings.TrimSpace(socialNum)
	if len(socialNum) == 18 {
		//date = socialNum[6:14]
		year := socialNum[6:10]
		month := socialNum[10:12]
		day := socialNum[12:14]
		date = NewDateTime(fmt.Sprintf("%s-%s-%s", year, month, day))
	} else if len(socialNum) == 15 {
		//date = "19" + socialNum[6:12]
		year := socialNum[6:8]
		month := socialNum[8:10]
		day := socialNum[10:12]
		date = NewDateTime(fmt.Sprintf("19%s-%s-%s", year, month, day))
	} else {
		err = errors.New(fmt.Sprintf("无法识别身份证号:%s", socialNum))
		return
	}
	return
}
