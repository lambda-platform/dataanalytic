package utils

import (
	"encoding/json"
	"fmt"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/config"
	analyticModels "github.com/lambda-platform/dataanalytic/models"
	puzzleModels "github.com/lambda-platform/lambda/DB/DBSchema/models"
	"os"
)

func AutoMigrateSeed() {

	DB.DB.AutoMigrate(
		&analyticModels.Analytic{},
		&analyticModels.AnalyticFilter{},
		&analyticModels.AnalyticRangeFilter{},
		&analyticModels.AnalyticRowsColumn{},
		&analyticModels.AnalyticRangeRowColumn{},
		&analyticModels.AnalyticDateRange{},
	)

	if config.Config.App.Seed == "true" {

		var vbs []puzzleModels.VBSchemaAdmin
		DB.DB.Where("name = ?", "Анализ").Find(&vbs)
		if len(vbs) <= 0 {
			seedData()
		}
	}
}
func seedData() {

	var vbs []puzzleModels.VBSchemaAdmin
	AbsolutePath := AbsolutePath()
	dataFile, err := os.Open(AbsolutePath+"initialData/vb_schemas_admin.json")
	defer dataFile.Close()
	if err != nil {
		fmt.Println("PUZZLE SEED ERROR")
	}
	jsonParser := json.NewDecoder(dataFile)
	err = jsonParser.Decode(&vbs)
	if err != nil {
		fmt.Println(err)
		fmt.Println("PUZZLE SEED DATA ERROR")
	}
	//fmt.Println(len(vbs))

	for _, vb := range vbs {

		DB.DB.Create(&vb)
	}


	var vbs2 []puzzleModels.VBSchema

	dataFile2, err2 := os.Open(AbsolutePath+"initialData/vb_schemas.json")
	defer dataFile2.Close()
	if err2 != nil {
		fmt.Println("PUZZLE SEED ERROR")
	}
	jsonParser2 := json.NewDecoder(dataFile2)
	err = jsonParser2.Decode(&vbs2)
	if err != nil {
		fmt.Println(err)
		fmt.Println("PUZZLE SEED DATA ERROR")
	}
	//fmt.Println(len(vbs))

	for _, vb := range vbs2 {

		DB.DB.Create(&vb)

	}

	
}
