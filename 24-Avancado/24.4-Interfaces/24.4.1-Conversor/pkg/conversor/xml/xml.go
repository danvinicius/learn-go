package xml

import "github.com/clbanning/mxj/v2"

type ConversorXML struct {}

func (c ConversorXML) Converter(dados string) (map[string]any, error) {
	xml, err := mxj.NewMapXml([]byte(dados))
	if err != nil {
		return nil, err
	}
	return xml.Old(), nil
}

func (c ConversorXML) Formato() string {
	return "XML"
}
