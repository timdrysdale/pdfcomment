/*
 * funtions to get PDF comments, and flatten them
 *
 *
 */

package pdfcomment

import (
	"fmt"
	"reflect"

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

	comments := make(Comments)

	//numPages, err := reader.GetNumPages()
	//if err != nil {
	//	return comments, err
	//}

	for _, page := range reader.PageList {

		if annotations, err := page.GetAnnotations(); err == nil {

			for _, annot := range annotations {

				if reflect.TypeOf(annot.GetContext()).String() == "*model.PdfAnnotationText" {
					//					fmt.Println(annot.GetContext())

					fmt.Println(annot.Contents)
					fmt.Println(annot.Rect)
					if rect, is := annot.Rect.(*pdfcore.PdfObjectArray); is {

						fmt.Printf("%v %v %v %v\n", rect.Get(0), rect.Get(1), rect.Get(2), rect.Get(3))
					}

					//
					//newComment := &Comment{
					//	Pos:    geo.Point{X: annot.Rect.Get(0), Y: annot.Rect.Get(1)},
					//	Text:   annot.Contents.String(),
					//	Author: dict.Get("T"),
					//	Page:   p,
					//}

				}

				//	fmt.Println(reflect.TypeOf(annot.GetContext()).String())
				//
				//	fmt.Printf("Parent: %v\n", annot.StructParent)
				//	fmt.Printf("Context: %v\n", annot.GetContext())
				//	//if at, is := annot.(*pdf.PdfAnnotationText); is {
				//	//	fmt.Println(reflect.TypeOf(at).String())
				//	//}
				//	//var v *pdf.PdfAnnotationText
				//	//fmt.Println(reflect.TypeOf(v).String())
				//	fmt.Println(reflect.TypeOf(annot).String())
				//
				//	// see pdfcore/primitives
				//	//annot.Contents is *core.PdfObjectString
				//	//annot.Rect is *core.PdfObjectArray
				//
				//	fmt.Printf("%v\n", annot)
				//
				//	if rect, is := annot.Rect.(*pdfcore.PdfObjectArray); is {
				//
				//		fmt.Printf("%v %v %v %v\n", rect.Get(0), rect.Get(1), rect.Get(2), rect.Get(3))
				//	}
				//
				//	fmt.Printf("%T\n\n", annot.Contents)

				//fmt.Printf("%T\n", annot.Rect)
				//if text, is := annot.(*pdf.PdfAnnotationText); is {
				//	fmt.Printf("%v\n", annot)
				//}

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
