package services

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/devmegablaster/SheetBridge/pb"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: t.AccessToken}, nil
}

type WriterService struct {
	SpreadsheetId string
	SheetId       string
	SheetName     string
	AccessToken   string
	SheetSvc      *sheets.Service
	KeyOrder      []string
}

func NewWriterService(spreadsheetId, sheetId, accessToken string, keyOrder []string) *WriterService {
	ctx := context.Background()
	sheetSvc, err := sheets.NewService(ctx, option.WithTokenSource(&TokenSource{AccessToken: accessToken}))
	if err != nil {
		slog.Error("Unable to create Sheets client", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &WriterService{
		SpreadsheetId: spreadsheetId,
		SheetId:       sheetId,
		AccessToken:   accessToken,
		SheetSvc:      sheetSvc,
		KeyOrder:      keyOrder,
	}
}

func (s *WriterService) IdtoName() {
	spreadSheet, err := s.SheetSvc.Spreadsheets.Get(s.SpreadsheetId).IncludeGridData(false).Do()
	if err != nil {
		slog.Error("Unable to retrieve spreadsheet", slog.String("error", err.Error()))
		os.Exit(1)
	}

	sheetId, err := strconv.Atoi(s.SheetId)
	if err != nil {
		slog.Error("SheetId must be an integer", slog.String("error", err.Error()))
		return
	}

	for _, sheet := range spreadSheet.Sheets {
		if sheet.Properties.SheetId == int64(sheetId) {
			s.SheetName = sheet.Properties.Title
			break
		}
	}
}

func (s *WriterService) WriteToSheet(writeRange string, valueRange *sheets.ValueRange) error {
	_, err := s.SheetSvc.Spreadsheets.Values.Update(s.SpreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		slog.Error("Unable to write data", slog.String("error", err.Error()))
		return err
	}

	slog.Info("Data writter")

	return nil
}

func (s *WriterService) WriteFull(data *pb.Write) {
	rng := fmt.Sprintf("%s!A1", s.SheetName)

	TransformerService := NewTransformerService(s.KeyOrder)
	finalResult := TransformerService.TransformToSheetDataFromWrite(data)

	valueRange := sheets.ValueRange{
		Values: finalResult,
		Range:  rng,
	}

	err := s.WriteToSheet(rng, &valueRange)
	if err != nil {
		slog.Error("Unable to write data", slog.String("error", err.Error()))
		return
	}

	slog.Info("Data written", slog.String("Sheet", s.SheetName))
}
