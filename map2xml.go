package chaos

import (
	"encoding/xml"
	"io"
	"fmt"
)

//from:https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go
type ForXmlMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m ForXmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	start.Name.Local="xml"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *ForXmlMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = ForXmlMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func main() {
	// The Map
	m := map[string]string{
		"key_1": "Value One",
		"key_2": "Value Two",
	}
	fmt.Println(m)

	// Encode to XML
	x, _ := xml.MarshalIndent(ForXmlMap(m), "", "  ")
	fmt.Println(string(x))

	// Decode back from XML
	var rm map[string]string
	xml.Unmarshal(x, (*ForXmlMap)(&rm))
	fmt.Println(rm)
}