package services

import (
	"bytes"

	"github.com/AbdullahAlzariqi/pdf"
	"github.com/labstack/echo/v5"
)

type PDFOperations interface {
	ExtractText(c *echo.Context, path string) (string, error)
}

type PDFService struct {}

func NewPDFService() PDFOperations {
	return &PDFService{}

}

func (p *PDFService) ExtractText(c *echo.Context, filePath string) (string, error) {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		c.Logger().Error("Failed to process pdf " + err.Error())
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		c.Logger().Error("Failed to read buffer of pdf " + err.Error())
		return "", err

	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
