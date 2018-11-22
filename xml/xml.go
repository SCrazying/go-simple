package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type XMLPerson struct {
	XMLName xml.Name           `xml:"person"`
	Para    XMLPersonParamNode `xml:"para"`
	Desc    XMLPersonDescNode  `xml:"desc"`
	Id      XMLPersonIdNode    `xml:"id"`
}

type XMLPersonParamNode struct {
	XMLName xml.Name `xml:"para"`
	Name    string   `xml:"name,attr"`
	Age     string   `xml:"age,attr"`
}
type XMLPersonDescNode struct {
	XMLName    xml.Name `xml:"desc"`
	Desc_Value string   `xml:",innerxml"`
}
type XMLPersonIdNode struct {
	XMLName xml.Name `xml:"id"`
	Attr1   string   `xml:"attr1,attr"`
	Attr2   string   `xml:"attr2,attr"`
	Id      string   `xml:",innerxml"`
}

func writexml(filepath string) {
	person := XMLPerson{
		Para: XMLPersonParamNode{Name: "sss", Age: "18"},
		Desc: XMLPersonDescNode{Desc_Value: "emmm"},
		Id:   XMLPersonIdNode{Attr1: "test1", Attr2: "test2", Id: "11"},
	}
	content, err := xml.MarshalIndent(&person, "", "    ")
	if err != nil {
		fmt.Println("xml write error")
	}
	err = ioutil.WriteFile(filepath, content, 0666)
	if err != nil {
		fmt.Println("save xmlfile error")
	}
}

func readxml(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read xml error")
	}
	var person XMLPerson
	err = xml.Unmarshal(content, &person)
	if err != nil {
		fmt.Println("unmarshal []byte error")
	}
	fmt.Println(person)

	return ""
}

type Ui struct {
	Name   string  `xml:"name,attr"`
	Groups []Group `xml:"group"`
}
type Group struct {
	XMLName xml.Name `xml:"xxxroup"`
	Name    string   `xml:"name,attr"`
	Id      string   `xml:"id,attr"`
}

func main() {
	readxml("1.xml")
	writexml("2.xml")

	var ui Ui
	ui.Name = "test"
	ui.Groups = append(ui.Groups, Group{Name: "test", Id: "123"})
	ui.Groups = append(ui.Groups, Group{Name: "11111", Id: "3213"})
	bufer, err := xml.Marshal(ui)
	if err != nil {

		log.Fatalln("Marshal err", err)
	}
	fmt.Println(string(bufer))
}
