package main

/*




 */
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	//readXML()
	writeXML()
}

func writeXML() {
	s := &Recurlyservers{Version: "1.0"}
	s.Svs = append(s.Svs, server{ServerName: "Tokoyo-vpn", ServerIP: "190.1.1.1"})
	s.Svs = append(s.Svs, server{ServerName: "Newyork-vpn", ServerIP: "192.1.1.1"})

	output, err := xml.MarshalIndent(s, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

/**

***为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）
 */
func readXML() {
	file, err := os.Open("server.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	s := Recurlyservers{}
	err = xml.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(s)
}
