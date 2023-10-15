package bmodel

import (
	"encoding/xml"
	"fmt"
	"social-pet/util"
	"testing"
	"time"
)

const (
	inputJson    = `{"LocalTime":"2020-04-03 00:00:00","DateTime":"2020-04-03"}`
	inputEmpJson = `{"LocalTime":"","DateTime":""}`
	inputXml     = `<Test><LocalTime>2020-04-03 00:00:00</LocalTime><DateTime>2020-04-03</DateTime></Test>`
	inputEmpXml  = `<Test><LocalTime></LocalTime><DateTime></DateTime></Test>`
)

type Test struct {
	LocalTime LocalTime
	DateTime  DateTime
}

func TestJson(t *testing.T) {
	//1.正常值
	if err := LocalTimeT(inputJson); err != nil {
		t.Error(err)
	}
	////1.空值
	if err := LocalTimeT(inputEmpJson); err != nil {
		t.Error(err)
	}
	//if err := XML(inputXml); err != nil {
	//	t.Error(err)
	//}
	//if err := XML(inputEmpXml); err != nil {
	//	t.Error(err)
	//}
	before := NewNowLocalTime().GetTime()
	new := NewNowLocalTime().GetTime()
	if new.Sub(before) >= 24*time.Hour {

	}
}

func XML(input string) (err error) {
	var obj Test
	if err = xml.Unmarshal([]byte(input), &obj); err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%+v", obj))
	indent, er := xml.MarshalIndent(obj, "", "")
	if er != nil {
		err = er
		return
	}
	fmt.Println("marshal:", string(indent))
	return
}

func LocalTimeT(input string) (err error) {
	var obj Test

	err = util.JsonStrToStruct(input, &obj)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%+v", obj))
	fmt.Println("marshal:", util.StructToJsonString(obj))
	return
}

func TestAddDate(t *testing.T) {
	d := NewNowDateTime()
	fmt.Println(d.String())
	d.AddDate(0, 0, 1)
	fmt.Println(d.String())

	l := NewNowLocalTime()
	fmt.Println(l.String())
	l.AddDate(0, 0, 1)
	fmt.Println(l.String())
}

func TestIsToday(t *testing.T) {
	d := NewNowLocalTime()
	fmt.Println(d.IsToday())
	d.AddDate(0, 0, 1)
	fmt.Println(d.IsToday())

	date := NewNowDateTime()
	fmt.Println(date.IsToday())
	date.AddDate(0, 0, 1)
	fmt.Println(date.IsToday())
}
