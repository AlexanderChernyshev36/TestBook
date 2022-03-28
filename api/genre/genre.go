package genre

import (
	"TestBook/models/modelsjson"
	"TestBook/models/modelssql"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

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
