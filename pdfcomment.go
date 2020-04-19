/*
 * funtions to get PDF comments, and flatten them
 *
 *
 */

package pdfcomment

import (
	"fmt"

	"github.com/timdrysdale/geo"
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

	comments := make(Comments)

	//numPages, err := reader.GetNumPages()
	//if err != nil {
	//	return comments, err
	//}

	for _, page := range reader.PageList {

		if annotations, err := page.GetAnnotations(); err == nil {

			for _, annot := range annotations {
				fmt.Printf("%v\n", annot)
				/*name := fmt.Sprintf("%s", annot.Get("Name"))
				if strings.Compare(name, "Comment") == 0 {
					newComment = &Comment{
						Pos:    geo.Point{X: 0, Y: 0},
						Text:   dict.Get("Contents"),
						Author: dict.Get("T"),
						Page:   p,
					}
				*/
			}

		}

	}

	return comments, nil

}
