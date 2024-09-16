package validators

import (
	"io/ioutil"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func ValidateXMLAgainstXSD(xmlData []byte, xsdPath string) error {
	schema, err := ioutil.ReadFile(xsdPath)
	if err != nil {
		return err
	}

	xsdSchema, err := xsd.Parse(schema)
	if err != nil {
		return err
	}
	defer xsdSchema.Free()

	doc, err := libxml2.Parse(xmlData)
	if err != nil {
		return err
	}
	defer doc.Free()

	return xsdSchema.Validate(doc)
}
