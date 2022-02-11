package main

import (
	"context"
	"fmt"
	"kpidrive/api"
	"kpidrive/sheets"
	"log"
)

const (
	sheetID    = "1NbaNM4bvm24hVXG57iAgasgK0Ef3K-SFG_OefOirVP4"
	cellsRange = "5. RKPI-Карта!A4:F100"
	// path       = "https://testdb.kpi-drive.ru/_api/mo/get_mo"
	path2      = "https://testdb.kpi-drive.ru/_api/indicators/save_indicator_instance_field"
	path3      = "https://testdb.kpi-drive.ru/_api/interpretations/save_interpretation"
	token      = "d0f00715-09ad-4808-b7a7-a7208e90bdec"
)

func main() {
	ctx := context.Background()

	emps, err := sheets.ReadGoogleSheets(ctx, sheetID, cellsRange)
	if err != nil {
		log.Println(err)
		return
	}

	funcs := sheets.GetSalaryFuncs(emps)

	_ = funcs

	f1 := "interpretation_id"
	f2 := "12"
	for _, val := range emps {
		mos, _ := api.EmployeeUpdate(path2, token, val.MoID, f1, f2)
		fmt.Println(val.MoID, mos)
	}
}
