package connectors

import (
	"context"
	"fmt"
	"log"
)

func (pc *PostgresConnection) GetTables() []string {
	rows, err := pc.DB.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		log.Fatalf("Failed to get tables: %v", err)
	}
	defer rows.Close()

	tables := []string{}

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Fatalf("Failed to scan table: %v", err)
		}

		tables = append(tables, table)
	}

	return tables
}

func (pc *PostgresConnection) GetTableData(table string) ([]map[string]interface{}, error) {
	rows, err := pc.DB.Query(context.Background(), fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		return nil, fmt.Errorf("Failed to select data from table: %v", err)
	}
	defer rows.Close()

	columns := rows.FieldDescriptions()
	data := []map[string]interface{}{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("Failed to get values from row: %v", err)
		}

		rowData := map[string]interface{}{}

		for i, value := range values {
			rowData[columns[i].Name] = value
		}

		data = append(data, rowData)
	}

	return data, nil
}
