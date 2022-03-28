package author

import (
	"TestBook/models/modelsjson"
	"TestBook/models/modelssql"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func ViewAuthor(Db_Connect *gorm.DB, id string) modelsjson.Result_AuthorsBook {

	result := modelsjson.Result_AuthorsBook{}

	checkId := 0
	if id != "" {
		checkId, _ = strconv.Atoi(id)
	}

	sqlAuthorBook := modelssql.AuthorsBook{}
	Db_Connect.Where("id = ?", checkId).Find(&sqlAuthorBook)

	result.Data.Age = sqlAuthorBook.Age
	result.Data.Id = sqlAuthorBook.Id
	result.Data.Name = sqlAuthorBook.Name
	result.Data.Website = sqlAuthorBook.Website
	result.Data.Email = sqlAuthorBook.Email

	if sqlAuthorBook.Id == 0 {
		result.Error.Code = 2
		result.Error.TextError = "Ничего не найдено"
		result.Error.Error = true
		fmt.Println("Ничего не найдено")
	}

	return result
}

func SetAuthor(Db_Connect *gorm.DB, request_SetAuthor []modelsjson.AuthorsBook) string {
	result := ""

	for _, jsonAuthor := range request_SetAuthor {
		sqlAuthorsBook := modelssql.AuthorsBook{}
		Db_Connect.Model(&sqlAuthorsBook).Where("id = ?", jsonAuthor.Id).Find(&sqlAuthorsBook)

		if sqlAuthorsBook.Id == 0 {
			//create
			sqlAuthorsBook.Id = jsonAuthor.Id
			sqlAuthorsBook.Name = jsonAuthor.Name
			sqlAuthorsBook.Age = jsonAuthor.Age
			sqlAuthorsBook.Website = jsonAuthor.Website
			sqlAuthorsBook.Email = jsonAuthor.Email
			Db_Connect.Model(sqlAuthorsBook).Create(&sqlAuthorsBook)
			fmt.Println("create" + strconv.Itoa(jsonAuthor.Id))
		} else {
			//update
			Db_Connect.Model(&sqlAuthorsBook).Where("id = ?", jsonAuthor.Id).
				Updates(map[string]interface{}{
					"name":    jsonAuthor.Name,
					"age":     jsonAuthor.Age,
					"website": jsonAuthor.Website,
					"email":   jsonAuthor.Email})
			fmt.Println("update " + strconv.Itoa(jsonAuthor.Id))
		}
	}
	return result
}

func DeleteAuthor(Db_Connect *gorm.DB, id string) string {
	result := ""

	checkId := 0
	if id != "" {
		checkId, _ = strconv.Atoi(id)
	}

	if checkId == 0 {
		result = "Не правильно указан ID"
		return result
	}
	sqlAuthor := modelssql.AuthorsBook{}

	Db_Connect.Where("id = ?", checkId).Find(&sqlAuthor)

	if sqlAuthor.Id > 0 {
		Db_Connect.Where("id = ?", checkId).Delete(&sqlAuthor)
	} else {
		result = "ID не найден"
		return result
	}

	if Db_Connect.Error == nil {
		result = "Автор удален"
	} else {
		result = "Ошибка при удалении"
	}

	return result
}
