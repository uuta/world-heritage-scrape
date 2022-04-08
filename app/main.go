package main

import (
	"encoding/xml"
	"fmt"
)

type Data struct {
	Families []Family `xml:"family"`
}

type Family struct {
	Lastname string    `xml:"lastname,attr"`
	Members  []Members `xml:"member"`
}

type Members struct {
	Firstname string `xml:"firstname"`
	Mail      string `xml:"mail"`
}

func main() {
	xmlStr := `
<?xml version="1.0" encoding="UTF-8"?>
<data>
    <family lastname="山田">
        <member>
            <firstname>太郎</firstname>
            <mail>tao.yamada@aaa.bbb.cc</mail>
        </member>
        <member>
            <firstname>花子</firstname>
            <mail>hanako.yamada@aaa.bbb.cc</mail>
        </member>
    </family>
    <family lastname="花形">
        <member>
            <firstname>みつる</firstname>
            <mail>mitsuru.hanagata@aaa.bbb.cc</mail>
        </member>
        <member>
            <firstname>かほる</firstname>
            <mail>kahoru.hanagata@aaa.bbb.cc</mail>
        </member>
    </family>
</data>
`

	data := new(Data)
	if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
		fmt.Println("XML Unmarshal error:", err)
		return
	}

	// 家族1
	fmt.Println(data.Families[0].Lastname)             // 苗字
	fmt.Println(data.Families[0].Members[0].Firstname) // 名前1
	fmt.Println(data.Families[0].Members[0].Mail)      // メールアドレス1
	fmt.Println(data.Families[0].Members[1].Firstname) // 名前2
	fmt.Println(data.Families[0].Members[1].Mail)      // メールアドレス2

	// 区切り
	fmt.Println("")

	// 家族2
	fmt.Println(data.Families[1].Lastname)             // 苗字
	fmt.Println(data.Families[1].Members[0].Firstname) // 名前1
	fmt.Println(data.Families[1].Members[0].Mail)      // メールアドレス1
	fmt.Println(data.Families[1].Members[1].Firstname) // 名前2
	fmt.Println(data.Families[1].Members[1].Mail)      // メールアドレス2

}
