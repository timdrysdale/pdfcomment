/*
 * funtions to get PDF comments, and flatten them
 *
 *
 */

package pdfcomment

import (
	"reflect"
	"strconv"

	"github.com/timdrysdale/geo"
	pdfcore "github.com/unidoc/unipdf/v3/core"
	pdf "github.com/unidoc/unipdf/v3/model"
)

type Comment struct {
	Pos  geo.Point
	Text string
	Page int
}

type Comments map[int][]Comment

func GetComments(reader *pdf.PdfReader) (Comments, error) {

	comments := make(map[int][]Comment)

	for p, page := range reader.PageList {

		if annotations, err := page.GetAnnotations(); err == nil {

			for _, annot := range annotations {

				if reflect.TypeOf(annot.GetContext()).String() == "*model.PdfAnnotationText" {

					if rect, is := annot.Rect.(*pdfcore.PdfObjectArray); is {

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

func GetCommentsForPage(comments Comments, page int) []Comment {

	return comments[page]

}
