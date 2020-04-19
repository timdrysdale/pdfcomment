package pdfcomment

import (
	"os"
	"reflect"
	"testing"

	"github.com/timdrysdale/geo"
	pdf "github.com/unidoc/unipdf/v3/model"
)

var c00 = Comment{Pos: geo.Point{X: 117.819, Y: 681.924}, Text: "This is a comment on page 1"}
var c10 = Comment{Pos: geo.Point{X: 326.501, Y: 593.954}, Text: "this is a comment on page 2", Page: 1}
var c11 = Comment{Pos: geo.Point{X: 141.883, Y: 685.869}, Text: "this is a second comment on page 2", Page: 1}
var c20 = Comment{Pos: geo.Point{X: 387.252, Y: 696.52}, Text: "this is a comment on page 3", Page: 2}
var c21 = Comment{Pos: geo.Point{X: 184.487, Y: 659.439}, Text: "this is a second comment on page 3", Page: 2}

func TestPDFExtract(t *testing.T) {
	f, err := os.Open("./test/3page-comments.pdf")
	if err != nil {
		t.Error("Can't open test pdf")
	}

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		t.Error("Can't read test pdf")
	}

	comments, err := GetComments(pdfReader)

	if err != nil {
		t.Error(err)
	}

	var expectedComments = make(map[int][]Comment)
	expectedComments[0] = []Comment{c00}
	expectedComments[1] = []Comment{c10, c11}
	expectedComments[2] = []Comment{c20, c21}

	for i := 0; i < 3; i++ {
		if !reflect.DeepEqual(comments[i], expectedComments[i]) {
			t.Errorf("Comments wrong")
		}
	}

}
