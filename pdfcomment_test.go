package pdfcomment

import (
	"os"
	"reflect"
	"testing"

	"github.com/timdrysdale/geo"
	"github.com/unidoc/unipdf/v3/creator"
	pdf "github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/model/optimize"
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
		if !reflect.DeepEqual(comments.GetByPage(i), expectedComments[i]) {
			t.Errorf("Comments wrong")
		}
	}

}

func TestPDFFlatten(t *testing.T) {

	c := creator.New()
	c.SetPageMargins(0, 0, 0, 0) // we're not printing
	c.SetPageSize(creator.PageSizeA4)
	c.NewPage()
	makeMarker(c, c00, "1")
	c.NewPage()
	makeMarker(c, c10, "1")
	makeMarker(c, c11, "2")
	c.NewPage()
	makeMarker(c, c20, "1")
	makeMarker(c, c21, "2")

	c.SetOptimizer(optimize.New(optimize.Options{
		CombineDuplicateDirectObjects:   true,
		CombineIdenticalIndirectObjects: true,
		CombineDuplicateStreams:         true,
		CompressStreams:                 true,
		UseObjectStreams:                true,
		ImageQuality:                    80,
		ImageUpperPPI:                   100,
	}))

	c.WriteToFile("./test/flattened-comments.pdf")
}
