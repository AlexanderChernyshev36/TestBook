package exportExcel

import (
	"TestBook/models/modelsjson"
	"TestBook/models/modelssql"
	"fmt"
	"github.com/xuri/excelize"
	"gorm.io/gorm"
	"strconv"
)

func ViewQuerySaveBook(Db_Connect *gorm.DB, request_SaveBook modelsjson.QuerySaveBook) string {

	f := excelize.NewFile()

	countlist := 0

	if request_SaveBook.PrintList1 {
		if request_SaveBook.NameList1 == "" {
			return "Добавьте название к NameList_1"
		}

		err := printList_1(Db_Connect, f, request_SaveBook)
		if err != "" {
			return err
		}

		countlist++
	}

	if request_SaveBook.PrintList2 {
		if request_SaveBook.NameList2 == "" {
			return "Добавьте название к NameList_2"
		}
		err := printList_2(Db_Connect, f, request_SaveBook, countlist)
		if err != "" {
			return err
		}
		countlist++
	}

	if request_SaveBook.PrintList3 {
		if request_SaveBook.NameList3 == "" {
			return "Добавьте название к NameList_3"
		}
		err := printList_3(Db_Connect, f, request_SaveBook, countlist)
		if err != "" {
			return err
		}

		countlist++

	}

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	return "Все прошло успешно"

}

func printList_1(Db_Connect *gorm.DB, f *excelize.File, request_SaveBook modelsjson.QuerySaveBook) string {

	sqlZhanrs := []modelssql.Zhanrs{}
	Db_Connect.Find(&sqlZhanrs)

	nameSheet_1 := request_SaveBook.NameList1
	f.SetSheetName("Sheet1", nameSheet_1)

	headerList(nameSheet_1, f, "A1", "id")
	headerList(nameSheet_1, f, "B1", "Наименование")

	StyleBorderLeft, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	StyleBorderCenter, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	StyleBorderRight, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	i := 2
	for _, sqlZhanr := range sqlZhanrs {
		f.SetCellValue(nameSheet_1, "A"+strconv.Itoa(i), sqlZhanr.Id)
		f.SetCellValue(nameSheet_1, "B"+strconv.Itoa(i), sqlZhanr.NameZhanr)
		i++
	}
	_ = f.SetCellStyle(nameSheet_1, "A1", "A"+strconv.Itoa(len(sqlZhanrs)+1), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_1, "B1", "B"+strconv.Itoa(len(sqlZhanrs)+1), StyleBorderLeft)

	_ = f.SetCellStyle(nameSheet_1, "A1", "A"+strconv.Itoa(1), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_1, "B1", "B"+strconv.Itoa(1), StyleBorderCenter)

	f.SetColWidth(nameSheet_1, "A", "A", 20)
	f.SetColWidth(nameSheet_1, "B", "B", 30)

	return ""

}

func printList_2(Db_Connect *gorm.DB, f *excelize.File, request_SaveBook modelsjson.QuerySaveBook, countlist int) string {

	sqlBooks := []modelssql.Books{}
	Db_Connect.Find(&sqlBooks)

	nameSheet_2 := request_SaveBook.NameList2

	if countlist == 0 {
		f.SetSheetName("Sheet1", nameSheet_2)
	} else {
		f.NewSheet(nameSheet_2)
	}

	f.SetColWidth(nameSheet_2, "A", "E", 30)

	f.MergeCell(nameSheet_2, "A1", "A2")
	f.MergeCell(nameSheet_2, "B1", "B2")
	f.MergeCell(nameSheet_2, "C1", "C2")
	f.MergeCell(nameSheet_2, "D1", "D2")
	f.MergeCell(nameSheet_2, "E1", "E2")

	headerList(nameSheet_2, f, "A1", "id книги")
	headerList(nameSheet_2, f, "B1", "id автора")
	headerList(nameSheet_2, f, "C1", "Наименование книги")
	headerList(nameSheet_2, f, "D1", "Кол. Страниц")
	headerList(nameSheet_2, f, "E1", "Жанр")

	StyleBorderLeft, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	StyleBorderCenter, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	StyleBorderRight, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	i := 3
	for _, sqlBook := range sqlBooks {
		f.SetCellValue(nameSheet_2, "A"+strconv.Itoa(i), sqlBook.Id)
		f.SetCellValue(nameSheet_2, "B"+strconv.Itoa(i), sqlBook.IdAuthor)
		f.SetCellValue(nameSheet_2, "C"+strconv.Itoa(i), sqlBook.NameBook)
		f.SetCellValue(nameSheet_2, "D"+strconv.Itoa(i), sqlBook.NumberPages)
		f.SetCellValue(nameSheet_2, "E"+strconv.Itoa(i), sqlBook.Zhanr)
		i++
	}
	_ = f.SetCellStyle(nameSheet_2, "A2", "A"+strconv.Itoa(len(sqlBooks)+2), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_2, "B2", "B"+strconv.Itoa(len(sqlBooks)+2), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_2, "C2", "C"+strconv.Itoa(len(sqlBooks)+2), StyleBorderLeft)
	_ = f.SetCellStyle(nameSheet_2, "D2", "D"+strconv.Itoa(len(sqlBooks)+2), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_2, "E2", "E"+strconv.Itoa(len(sqlBooks)+2), StyleBorderLeft)

	_ = f.SetCellStyle(nameSheet_2, "A1", "A"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_2, "B1", "B"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_2, "C1", "C"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_2, "D1", "D"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_2, "E1", "E"+strconv.Itoa(2), StyleBorderCenter)

	return ""
}

func printList_3(Db_Connect *gorm.DB, f *excelize.File, request_SaveBook modelsjson.QuerySaveBook, countlist int) string {

	sqlAuthorsBooks := []modelssql.AuthorsBook{}
	Db_Connect.Find(&sqlAuthorsBooks)

	nameSheet_3 := request_SaveBook.NameList3

	if countlist == 0 {
		f.SetSheetName("Sheet1", nameSheet_3)
	} else {
		f.NewSheet(nameSheet_3)
	}

	f.SetColWidth(nameSheet_3, "A", "D", 20)

	f.MergeCell(nameSheet_3, "A1", "A2")
	f.MergeCell(nameSheet_3, "B1", "B2")
	f.MergeCell(nameSheet_3, "C1", "C2")
	f.MergeCell(nameSheet_3, "D1", "D2")

	headerList(nameSheet_3, f, "A1", "id автора")
	headerList(nameSheet_3, f, "B1", "Имя")
	headerList(nameSheet_3, f, "C1", "Возраст")
	headerList(nameSheet_3, f, "D1", "Контакты")

	i := 3
	ii := 4
	for _, sqlAuthorBook := range sqlAuthorsBooks {

		f.MergeCell(nameSheet_3, "A"+strconv.Itoa(i), "A"+strconv.Itoa(ii))
		f.SetCellValue(nameSheet_3, "A"+strconv.Itoa(i), sqlAuthorBook.Id)

		f.MergeCell(nameSheet_3, "B"+strconv.Itoa(i), "B"+strconv.Itoa(ii))
		f.SetCellValue(nameSheet_3, "B"+strconv.Itoa(i), sqlAuthorBook.Name)

		f.MergeCell(nameSheet_3, "C"+strconv.Itoa(i), "C"+strconv.Itoa(ii))
		f.SetCellValue(nameSheet_3, "C"+strconv.Itoa(i), sqlAuthorBook.Age)

		f.SetCellValue(nameSheet_3, "D"+strconv.Itoa(i), sqlAuthorBook.Website)

		f.SetCellValue(nameSheet_3, "D"+strconv.Itoa(ii), sqlAuthorBook.Email)

		i = i + 2
		ii = ii + 2
	}

	StyleBorderCenter, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	StyleBorderRight, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})

	_ = f.SetCellStyle(nameSheet_3, "A1", "A"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_3, "B1", "B"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_3, "C1", "C"+strconv.Itoa(2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_3, "D1", "D"+strconv.Itoa(2), StyleBorderCenter)

	_ = f.SetCellStyle(nameSheet_3, "A3", "A"+strconv.Itoa(len(sqlAuthorsBooks)*2+2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_3, "B3", "B"+strconv.Itoa(len(sqlAuthorsBooks)*2+2), StyleBorderCenter)
	_ = f.SetCellStyle(nameSheet_3, "C3", "C"+strconv.Itoa(len(sqlAuthorsBooks)*2+2), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_3, "D3", "D"+strconv.Itoa(len(sqlAuthorsBooks)*2+2), StyleBorderRight)
	_ = f.SetCellStyle(nameSheet_3, "D3", "D"+strconv.Itoa(len(sqlAuthorsBooks)*2+2), StyleBorderRight)

	return ""
}

func headerList(nameSheet string, f *excelize.File, numCol string, textCol string) string {

	if err := f.SetCellRichText(nameSheet, numCol, []excelize.RichTextRun{
		{
			Text: textCol,
			Font: &excelize.Font{
				Bold: true,
			},
		},
	}); err != nil {
		fmt.Println(err)
		return "Не удалось отформатировать текст"
	}
	return ""

}
