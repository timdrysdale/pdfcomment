/*
 * funtions to get PDF comments, and flatten them
 *
 *
 */

package pdfcomment

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/timdrysdale/geo"
	pdfcore "github.com/unidoc/unipdf/v3/core"
	pdf "github.com/unidoc/unipdf/v3/model"
)

type Comment struct {
	Pos    geo.Point
	Text   string
	Author string
	Page   int
}

type Comments map[int][]Comment

// this doesn't know about pages ...?!!
func GetComments(reader *pdf.PdfReader) (Comments, error) {

	comments := make(map[int][]Comment)

	for p, page := range reader.PageList {

		if annotations, err := page.GetAnnotations(); err == nil {

			for _, annot := range annotations {

				if reflect.TypeOf(annot.GetContext()).String() == "*model.PdfAnnotationText" {

					fmt.Println(annot.Contents)
					fmt.Println(annot.Rect)

					if rect, is := annot.Rect.(*pdfcore.PdfObjectArray); is {

						fmt.Printf("%v %v %v %v\n", rect.Get(0), rect.Get(1), rect.Get(2), rect.Get(3))

						x, err := strconv.ParseFloat(rect.Get(0).String(), 64)
						if err != nil {
							return comments, err
						}
						y, err := strconv.ParseFloat(rect.Get(1).String(), 64)
						if err != nil {
							return comments, err
						}

						newComment := Comment{
							Pos:  geo.Point{X: x, Y: y},
							Text: annot.Contents.String(),
							Page: p,
						}

						comments[p] = append(comments[p], newComment)

					}

				}

			}

		}

	}

	return comments, nil

}
