package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kr1sal.com/avl-converter/internal/utils"
)

func ReadCSVFiles(path string) (map[string][][]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		filepaths, err := utils.IterFiles(path, "", ".csv", -1)
		if err != nil {
			return nil, err
		}

		tables := make(map[string][][]string, len(filepaths))
		for _, filepath := range filepaths {
			title, records, err := ReadCSVFile(filepath)
			if err != nil {
				return nil, err
			}
			tables[title] = records
		}

		return tables, nil
	} else if info.Mode().IsRegular() {
		title, records, err := ReadCSVFile(path)
		if err != nil {
			return nil, err
		}
		table := map[string][][]string{}
		table[title] = records

		return table, nil
	}

	return nil, fmt.Errorf("invalid path: %s", path)
}

func ReadCSVFile(path string) (string, [][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return "", nil, err
	}

	fileName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	return fileName, records, nil
}

func WriteJSONFile(path string, data map[string][][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	records := make(map[string][]map[string]string)
	for tableTitle, tableData := range data {
		tableName := fmt.Sprintf("%v", tableTitle)
		if len(tableData) == 0 {
			continue
		}
		tableTitles := tableData[0]
		tableRecords := make([]map[string]string, len(tableData)-1)
		for rowIndex, row := range tableData[1:] {
			recordData := make(map[string]string)
			for colIndex, col := range row {
				field := tableTitles[colIndex]
				recordData[field] = col
			}
			tableRecords[rowIndex] = recordData
		}
		records[tableName] = tableRecords
	}

	jsonData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
