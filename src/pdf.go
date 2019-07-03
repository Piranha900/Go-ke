package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"

)

func createTable(pdf *gofpdf.Fpdf, header []string, data [][]string, columnSize []float64) {
	if pdf == nil {
		return
	}
	fmt.Println(header)
	fmt.Println("---")
	fmt.Println(data)
	fmt.Println("---")
	fmt.Println(columnSize)
	fmt.Println("---")

	cLen := len(columnSize)
	if cLen == 0 {
		return
	}
	fmt.Println("cLen > 0")
	if len(header) != 0 && cLen != len(header) {
		return
	}
	fmt.Println("len header 0 or eq to cLen")
	for _, v := range data {
		if len(v) != cLen {
			return
		}
	}
	fmt.Println("Check done")
	pdf.SetFillColor(128, 128, 128)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(128, 0, 0)
	pdf.SetLineWidth(.3)
	pdf.SetFont("Arial", "", 12)
	// display header fields
	for i, str := range header {
		pdf.CellFormat(columnSize[i], 7, str, "1", 0, "C", true, 0, "")
	}
	// only advance a line if we wrote a header
	if len(header) > 0 {
		pdf.Ln(-1)
	}
	// display data /w alternating background
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("", "", 0)
	//  Data
	fill := false
	for _, c := range data {
		for i, str := range c {
			pdf.CellFormat(columnSize[i], 6, str, "LR", 0, "", fill, 0, "")
		}
		pdf.Ln(-1)
		fill = !fill
	}
	return
}

func main()  {

	pdf := gofpdf.New("P", "mm", "A4", "") //crea il pdf
	pdf.AddPage()
	ciao := []string{"ciao"}//crea la pagina
	data := [][]string{{"22"},{"06"},{"mao"}}
	cs := []float64{10.0}
	createTable(pdf,ciao, data,cs)
	pdf.OutputFileAndClose("La scheda di " + "prova" + ".pdf")
}