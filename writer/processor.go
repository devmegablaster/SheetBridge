package writer

import (
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/devmegablaster/SheetBridge/pb"
)

type WriteProcessor struct {
	writerSvc *services.WriterService
}

func NewWriteProcessor() *WriteProcessor {
	return &WriteProcessor{}
}

func (w *WriteProcessor) Handle(write *pb.Write) {
	w.writerSvc = services.NewWriterService(write.SpreadsheetId, write.SheetId, write.AccessToken, write.Columns)
	w.writerSvc.IdtoName()
	w.writerSvc.WriteFull(write)
}
