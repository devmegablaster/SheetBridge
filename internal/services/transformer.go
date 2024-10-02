package services

import (
	"fmt"

	"github.com/devmegablaster/SheetBridge/pb"
)

type TransformerService struct {
	keyOrder []string
}

func NewTransformerService(keyOrder []string) *TransformerService {
	return &TransformerService{
		keyOrder: keyOrder,
	}
}

func (t *TransformerService) TransformToSheetData(data []map[string]interface{}) [][]interface{} {
	if len(data) == 0 {
		return nil
	}

	var finalResult [][]interface{}

	headerRow := make([]interface{}, len(t.keyOrder))
	for i, key := range t.keyOrder {
		headerRow[i] = key
	}
	finalResult = append(finalResult, headerRow)

	for _, row := range data {
		var rowData []interface{}
		for _, key := range t.keyOrder {
			rowData = append(rowData, row[key])
		}
		finalResult = append(finalResult, rowData)
	}

	return finalResult
}

func (t *TransformerService) TransformToWriteMessage(data []map[string]interface{}, writeType pb.WriteType, spreadsheetId, sheetId, accessToken string) *pb.Write {
	transformed := t.TransformToSheetData(data)

	pbValue := []*pb.Value{}

	for _, row := range transformed {
		for _, cell := range row {
			pbValue = append(pbValue, &pb.Value{
				Value: fmt.Sprintf("%v", cell),
			})
		}
	}

	return &pb.Write{
		WriteType: writeType,
		WriteData: &pb.WriteData{
			Values: pbValue,
		},
		SpreadsheetId: spreadsheetId,
		SheetId:       sheetId,
		AccessToken:   accessToken,
		Columns:       t.keyOrder,
	}
}

func (t *TransformerService) TransformToSheetDataFromWrite(write *pb.Write) [][]interface{} {
	var finalResult [][]interface{}

	for i := 0; i < len(write.WriteData.Values); i += len(t.keyOrder) {
		var rowData []interface{}
		for j := 0; j < len(t.keyOrder); j++ {
			rowData = append(rowData, write.WriteData.Values[i+j].Value)
		}
		finalResult = append(finalResult, rowData)
	}

	return finalResult
}
