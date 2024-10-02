package connectors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/devmegablaster/SheetBridge/internal/models"
)

func (pc *PostgresConnection) GetTables() []string {
	rows, err := pc.DB.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		slog.Error("Failed to get tables", slog.String("error", err.Error()))
	}
	defer rows.Close()

	tables := []string{}

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			slog.Error("Failed to scan table", slog.String("error", err.Error()))
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

func (pc *PostgresConnection) GetTableSchema(table string) *models.Schema {
	rows, err := pc.DB.Query(context.Background(), fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", table))
	if err != nil {
		slog.Error("Failed to get table schema", slog.String("error", err.Error()))
	}
	defer rows.Close()

	schema := &models.Schema{}

	for rows.Next() {
		var column string
		var dataType string
		if err := rows.Scan(&column, &dataType); err != nil {
			slog.Error("Failed to scan schema", slog.String("error", err.Error()))
		}

		schema.Col = append(schema.Col, column)
		schema.Type = append(schema.Type, dataType)
	}

	return schema
}
