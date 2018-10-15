package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

func main() {
	readxml("1.xml")
	writexml("2.xml")
}
