package book

import (
	"TestBook/models/modelsjson"
	"TestBook/models/modelssql"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func SetBook(Db_Connect *gorm.DB, request_SetBook []modelsjson.Books) string {
	result := ""

	for _, jsonBook := range request_SetBook {
		sqlBook := modelssql.Books{}

		Db_Connect.Model(&sqlBook).Where("id = ?", jsonBook.Id).Find(&sqlBook)

		if sqlBook.Id == 0 {
			//create
			sqlBook.Id = jsonBook.Id
			sqlBook.IdAuthor = jsonBook.IdAuthor
			sqlBook.NameBook = jsonBook.NameBook
			sqlBook.NumberPages = jsonBook.NumberPages
			sqlBook.Zhanr = jsonBook.Zhanr
			Db_Connect.Model(sqlBook).Create(&sqlBook)
			fmt.Println("Create " + strconv.Itoa(jsonBook.Id))
		} else {
			//update
			Db_Connect.Model(&jsonBook).Where("id = ?", jsonBook.Id).
				Updates(map[string]interface{}{
					"id_author":    jsonBook.IdAuthor,
					"name_book":    jsonBook.NameBook,
					"number_pages": jsonBook.NumberPages,
					"zhanr":        jsonBook.Zhanr,
				})
			fmt.Println("Update " + strconv.Itoa(jsonBook.Id))
		}
	}
	return result
}

func ViewBook(Db_Connect *gorm.DB, id string) modelsjson.Result_Books {

	result := modelsjson.Result_Books{}

	checkId := 0
	if id != "" {
		checkId, _ = strconv.Atoi(id)
	}

	sqlBook := modelssql.Books{}
	Db_Connect.Where("id = ?", checkId).Find(&sqlBook)

	result.Data.Id = sqlBook.Id
	result.Data.Zhanr = sqlBook.Zhanr
	result.Data.NumberPages = sqlBook.NumberPages
	result.Data.NameBook = sqlBook.NameBook
	result.Data.IdAuthor = sqlBook.IdAuthor

	if sqlBook.Id == 0 {
		result.Error.Code = 2
		result.Error.TextError = "Ничего не найдено"
		result.Error.Error = true
		fmt.Println("Ничего не найдено")
	}

	return result
}

func DeleteBook(Db_Connect *gorm.DB, id string) string {
	result := ""

	checkId := 0
	if id != "" {
		checkId, _ = strconv.Atoi(id)
	}

	if checkId == 0 {
		result = "Не правильно указан ID"
		return result
	}

	sqlBook := modelssql.Books{}
	Db_Connect.Where("id = ?", checkId).Find(&sqlBook)

	if sqlBook.Id > 0 {
		Db_Connect.Where("id = ?", checkId).Delete(&sqlBook)
	} else {
		result = "ID не найден"
		return result
	}

	if Db_Connect.Error == nil {
		result = "Книга удалена"
	} else {
		result = "Ошибка при удалении"
	}

	return result
}
