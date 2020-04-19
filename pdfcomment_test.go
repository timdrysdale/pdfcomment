package pdfcomment

import (
	"fmt"
	"os"
	"testing"

	pdf "github.com/unidoc/unipdf/v3/model"
)

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
		fmt.Printf("%v\n", comments)
	}

}

//func printPdfFieldData(inputPath, targetFieldName string) error {
//	f, err := os.Open(inputPath)
//	if err != nil {
//		return err
//	}
//	fmt.Printf("Input file: %s\n", inputPath)
//
//	defer f.Close()
//
//	pdfReader, err := pdf.NewPdfReader(f)
//	if err != nil {
//		return err
//	}
//
//	acroForm := pdfReader.AcroForm
//	if acroForm == nil {
//		fmt.Printf(" No formdata present\n")
//		return nil
//	}
//
//	match := false
//	fields := acroForm.AllFields()
//	for _, field := range fields {
//		fullname, err := field.FullName()
//		if err != nil {
//			return err
//		}
//		if fullname == targetFieldName || targetFieldName == "" {
//			match = true
//			if field.V != nil {
//				fmt.Printf("Field '%s': '%v' (%T)\n", fullname, field.V, field.V)
//			} else {
//				fmt.Printf("Field '%s': not filled\n", fullname)
//			}
//		}
//	}
//
//	if !match {
//		return errors.New("field not found")
//	}
//	return nil
//}
//
//func inspectPdf(inputPath string, opt cmdOptions) error {
//	f, err := os.Open(inputPath)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	pdfReader, err := pdf.NewPdfReader(f)
//	if err != nil {
//		return err
//	}
//
//	isEncrypted, err := pdfReader.IsEncrypted()
//	if err != nil {
//		return err
//	}
//
//	// Try decrypting with an empty one.
//	if isEncrypted {
//		auth, err := pdfReader.Decrypt([]byte(opt.pdfPassword))
//		if err != nil {
//			return err
//		}
//
//		if !auth {
//			return errors.New("Unable to decrypt password protected file - need to specify pass to Decrypt")
//		}
//	}
//
//	numPages, err := pdfReader.GetNumPages()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("PDF Num Pages: %d\n", numPages)
//
//	objNums := pdfReader.GetObjectNums()
//
//	// Output.
//	fmt.Printf("%d PDF objects:\n", len(objNums))
//	for _, objNum := range objNums {
//		obj, err := pdfReader.GetIndirectObjectByNumber(objNum)
//		if err != nil {
//			return err
//		}
//		//fmt.Printf("%3d: %d 0 %T\n", i, objNum, obj)
//		/*if stream, is := obj.(*pdfcore.PdfObjectStream); is {
//			//decoded, err := pdfcore.DecodeStream(stream)
//			if err != nil {
//				return err
//			}
//			//fmt.Printf("Decoded:\n%s\n", decoded)
//		} else
//		*/
//		if indObj, is := obj.(*pdfcore.PdfIndirectObject); is {
//
//			//fmt.Println(indObj.PdfObject.String())
//
//			/*
//				switch v := foo.Get("Name").(type) {
//				default:
//				case string:
//					fmt.Println(foo.Get("Contents"))
//				}*/
//
//			if foo, is := indObj.PdfObject.(*pdfcore.PdfObjectDictionary); is {
//
//				v := fmt.Sprintf("%s", foo.Get("Name"))
//				if strings.Compare(v, "Comment") == 0 {
//					fmt.Println("=========================================================")
//					fmt.Printf("%s:%s\n", foo.Get("T"), foo.Get("Contents"))
//					fmt.Println(foo.Get("Rect"))
//					//for _, k := range foo.Keys() {
//					//	fmt.Printf("%s:%s\n", k, foo.Get(k))
//					//}
//				}
//				//	fmt.Println(v)
//
//				/*
//					switch v := foo.Get("Name").(type) {
//									default:
//									case string:
//										if strings.Equal(v, "Contents") {
//											fmt.Println(foo.Get("Contents"))
//										}
//									}*/
//
//				//fmt.Println(foo.Get("Contents"))
//				if false {
//					for _, k := range foo.Keys() {
//						fmt.Printf("%s:%s\n", k, foo.Get(k))
//					}
//				}
//			}
//
//			//if objDict, is := obj.(*pdfcore.PdfObjectDictionary); is {
//			//	fmt.Printf("FOUND ONE: %T \n", objDict)
//			//}
//			/*
//				for _, k := range indObj.PdfObject.PdfObjectDictionary.keys {
//					fmt.Printf("%s:%s\n", k, indObj.PdfObject[k])
//				}*/
//
//			//if objDict, is := obj.(*pdfcore.PdfObjectDictionary); is {
//
//			//fmt.Printf("FOUND ONE: %T \n", indObj.PdfObject) //.PdfObjectDictionary)
//			//fmt.Printf("%T\n", indObj.PdfObject)
//
//			//fmt.Printf("%s\n", indObj.PdfObject.String())
//
//			//contents := indObj.PdfObject.String()
//
//			//}
//
//		}
//
//	}
//
//	return nil
//}
