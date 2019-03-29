package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

type dataStr struct {
	ID        string
	Priority  string
	Discount  string
	UnitP     string
	ShippingC string
	CID       string
}

func main() {

	excelFileName := "/Users/yychenaa//Documents/go/src/github.com/EddieYY/test.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}
	cc := make(chan dataStr)
	done := make(chan bool)
	//	for _, sheet := range xlFile.Sheets {
	rowN := len(xlFile.Sheet["Orders"].Rows)
	for i, row := range xlFile.Sheet["Orders"].Rows {
		var c dataStr
		go func(i int, row *xlsx.Row) {
			if i != 0 {
				for j, cell := range row.Cells {
					text := cell.String()
					switch j {
					case 0:
						c.ID = text
					case 1:
						c.Priority = text
					case 2:
						c.Discount = text
					case 3:
						c.UnitP = text
					case 4:
						c.ShippingC = text
					case 5:
						c.CID = text
					}
					//fmt.Printf("%s\n", i)
				}

				//fmt.Printf("%s\n", c)
				cc <- c
			}
			done <- true
		}(i, row)

	}
	go func() {
		for i := 0; i < rowN; i++ {
			<-done
		}
		close(cc)
	}()

	for n := range cc {
		fmt.Println(n)
	}

}
