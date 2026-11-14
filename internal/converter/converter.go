package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"kr1sal.com/avl-converter/internal/utils"
)

func ReadCSVFiles(path string) ([][][]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		filepaths, err := utils.IterFiles(path, "", ".csv", -1)
		if err != nil {
			return nil, err
		}

		tables := make([][][]string, len(filepaths))
		for index, filepath := range filepaths {
			records, err := ReadCSVFile(filepath)
			if err != nil {
				return nil, err
			}
			tables[index] = records
		}

		return tables, nil
	} else if info.Mode().IsRegular() {
		records, err := ReadCSVFile(path)
		if err != nil {
			return nil, err
		}

		return [][][]string{records}, nil
	}

	return nil, fmt.Errorf("invalid path: %s", path)
}

func ReadCSVFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func WriteJSONFile(path string, data [][][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	records := make(map[string][]map[string]string)
	for tableIndex, table := range data {
		tableName := fmt.Sprintf("table_%d", tableIndex+1)
		tableTitles := table[0]
		tableRecords := make([]map[string]string, len(table)-1)
		for rowIndex, row := range table[1:] {
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
