package sheets

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	timeout = 10
	APIkey  = "AIzaSyCmQCft5_2pjoLvTnKVvcLWREQnh1haUCA"
)

var (
	errNotString = errors.New("interface is not string")
	errConvert   = errors.New("interface is not string")
)

func ReadGoogleSheets(ctx context.Context, spreadsheetID, cells string) ([]*Employee, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	sheetsService, err := sheets.NewService(ctxWithTimeout, option.WithAPIKey(APIkey))
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve sheets client: %w", err)
	}

	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, cells).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %w", err)
	}

	if len(resp.Values) == 0 {
		return nil, errors.New("Table is empty")
	}
	var emps []*Employee
	for i, row := range resp.Values {
		emp, err := unpackRow(row)
		if err != nil {
			log.Println(fmt.Sprintf("Row %d unpacking error: %v", i, err))
		}

		emps = append(emps, emp)
	}

	return emps, nil
}

func unpackRow(row []interface{}) (*Employee, error) {
	emp := &Employee{}

	ID, ok := row[0].(string)
	if !ok {
		return nil, errNotString
	}
	emp.MoID = strings.TrimSpace(ID)

	fiostring, ok := row[1].(string)
	if !ok {
		return nil, errNotString
	}
	fio := strings.Split(strings.TrimSpace(fiostring), " ")

	for i := range fio {
		switch i {
		case 0:
			emp.FirstName = fio[i]
		case 1:
			emp.MiddleName = fio[i]
		case 2:
			emp.LastName = fio[i]
		}
	}

	for i := 2; i < len(row); i++ {
		salaryValueString, ok := row[i].(string)
		if !ok {
			return nil, errNotString
		}

		num, err := strconv.Atoi(strings.ReplaceAll(salaryValueString, "\u00a0", ""))
		if err != nil {
			return nil, errConvert
		}
		emp.Salary = append(emp.Salary, num)
	}

	return emp, nil
}
