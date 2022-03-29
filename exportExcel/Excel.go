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
		sqlZhanrs := []modelssql.Zhanrs{}
		Db_Connect.Find(&sqlZhanrs)

		nameSheet_1 := request_SaveBook.NameList1
		f.SetSheetName("Sheet1", nameSheet_1)

		f.SetCellValue(nameSheet_1, "A1", "id")
		f.SetCellValue(nameSheet_1, "B1", "Наименование")

		if err := f.SetCellRichText(nameSheet_1, "B1", []excelize.RichTextRun{
			{
				Text: "Наименование",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"

		}

		if err := f.SetCellRichText(nameSheet_1, "A1", []excelize.RichTextRun{
			{
				Text: "id",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"
		}

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

		countlist++
	}

	if request_SaveBook.PrintList2 {
		if request_SaveBook.NameList2 == "" {
			return "Добавьте название к NameList_2"
		}

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

		if err := f.SetCellRichText(nameSheet_2, "A1", []excelize.RichTextRun{
			{
				Text: "id книги",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"
		}

		if err := f.SetCellRichText(nameSheet_2, "B1", []excelize.RichTextRun{
			{
				Text: "id автора",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"

		}

		if err := f.SetCellRichText(nameSheet_2, "C1", []excelize.RichTextRun{
			{
				Text: "Наименование книги",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"

		}

		if err := f.SetCellRichText(nameSheet_2, "D1", []excelize.RichTextRun{
			{
				Text: "Кол. Страниц",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"

		}

		if err := f.SetCellRichText(nameSheet_2, "E1", []excelize.RichTextRun{
			{
				Text: "Жанр",
				Font: &excelize.Font{
					Bold: true,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return "Не удалось отформатировать текст"

		}

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

		countlist++
	}

	return "Все прошло успешно"
}
