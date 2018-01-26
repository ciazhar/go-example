package main

import (
	"github.com/jung-kurt/gofpdf"
	// "github.com/jung-kurt/gofpdf/internal/example"
)

func main (){

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello World!")
	// fileStr := example.Filename("basic")
	// err := pdf.OutputFileAndClose("coba")
	// example.Summary(err, fileStr)

}
