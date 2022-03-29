package genre

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
	}

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return "Все прошло успешно"
}

func ViewGenre(Db_Connect *gorm.DB, name string) modelsjson.Result_Zhanrs {

	result := modelsjson.Result_Zhanrs{}

	checkName := ""
	if name != "" {
		checkName = name
	}

	sqlZhanrs := []modelssql.Zhanrs{}
	Db_Connect.Where("name_zhanr Like ?", "%"+checkName+"%").Find(&sqlZhanrs)

	jsonZhanrs := []modelsjson.Zhanrs{}
	for _, sqlZhanr := range sqlZhanrs {
		jsonZhanr := modelsjson.Zhanrs{}
		jsonZhanr.NameZhanr = sqlZhanr.NameZhanr
		jsonZhanr.Id = sqlZhanr.Id
		jsonZhanrs = append(jsonZhanrs, jsonZhanr)
	}

	result.Data = jsonZhanrs

	if len(result.Data) == 0 {
		result.Error.Code = 2
		result.Error.TextError = "Ничего не найдено"
		result.Error.Error = true
		fmt.Println("Ничего не найдено")
	}

	return result
}

func DeleteGenre(Db_Connect *gorm.DB, id string) string {
	result := ""

	checkId := 0
	if id != "" {
		checkId, _ = strconv.Atoi(id)
	}

	if checkId == 0 {
		result = " Не правильно указан ID"
		return result
	}

	sqlGenre := modelssql.Zhanrs{}
	Db_Connect.Where("id = ?", checkId).Find(&sqlGenre)

	if sqlGenre.Id > 0 {
		Db_Connect.Where("id = ?", checkId).Delete(&sqlGenre)
	} else {
		result = "ID не найден"
		return result
	}

	if Db_Connect.Error == nil {
		result = "Жанр удален"
	} else {
		result = "Ошибка при удалении"
	}

	return result
}

func SetGenre(Db_Connect *gorm.DB, request_SetGenre []modelsjson.Zhanrs) string {
	result := ""

	for _, jsonGenre := range request_SetGenre {
		sqlGenre := modelssql.Zhanrs{}
		Db_Connect.Model(&sqlGenre).Where("id = ?", jsonGenre.Id).Find(&sqlGenre)

		if sqlGenre.Id == 0 {
			//create
			sqlGenre.Id = jsonGenre.Id
			sqlGenre.NameZhanr = jsonGenre.NameZhanr
			Db_Connect.Model(sqlGenre).Create(&sqlGenre)
			fmt.Println("create" + strconv.Itoa(jsonGenre.Id))
		} else {
			//update
			Db_Connect.Model(&jsonGenre).Where("id = ?", jsonGenre.Id).
				Updates(map[string]interface{}{
					"name_zhanr": jsonGenre.NameZhanr,
				})
			fmt.Println("update " + strconv.Itoa(jsonGenre.Id))
		}
	}

	return result
}
