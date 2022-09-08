package storage

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"encoding/csv"
	"finance/internal/model"
	"io"
	"strconv"
	"strings"
	"time"
)

type FileSystem struct {
}

func NewFile() *FileSystem {
	return &FileSystem{}
}

func (f *FileSystem) GetDataFromBankStatement(ctx context.Context, fileBase64 string) ([]*model.TinkoffBankStatement, error) {
	reader, err := decodeBase64(fileBase64)
	if err != nil {
		return nil, err
	}
	csvData, err := readCSV(reader)
	if err != nil {
		return nil, err
	}
	parsedData, err := parseTinkoffBankStatementToStruct(csvData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}
func parseTinkoffBankStatementToStruct(data *[][]string) ([]*model.TinkoffBankStatement, error) {
	result := make([]*model.TinkoffBankStatement, 0, len(*data))
	for rn := 1; rn < len(*data); rn++ {
		row := (*data)[rn]
		// underscore error - if value is empty, set default value
		operationDate, _ := time.Parse("02.01.2006 15:04:05", row[0])
		paymentDate, _ := time.Parse("02.01.2006", row[1])
		operation, _ := strconv.ParseFloat(normalizeCurrency(row[4]), 64)
		payment, _ := strconv.ParseFloat(normalizeCurrency(row[6]), 64)
		cashback, _ := strconv.ParseFloat(normalizeCurrency(row[8]), 64)
		mcc, _ := strconv.Atoi(row[10])
		bonuses, _ := strconv.ParseFloat(normalizeCurrency(row[12]), 64)
		investmentBankRounding, _ := strconv.ParseFloat(normalizeCurrency(row[13]), 64)
		rounding, _ := strconv.ParseFloat(normalizeCurrency(row[14]), 64)
		bankStatement := &model.TinkoffBankStatement{
			OperationDate:          operationDate,
			PaymentDate:            paymentDate,
			CardNumber:             row[2],
			Status:                 model.Status(row[3]),
			Operation:              operation,
			Currency:               model.Currency(row[5]),
			Payment:                payment,
			PaymentCurrency:        model.Currency(row[7]),
			Cashback:               cashback,
			Category:               model.Category(row[9]),
			MCC:                    int64(mcc),
			Description:            row[11],
			Bonuses:                bonuses,
			InvestmentBankRounding: investmentBankRounding,
			Rounding:               rounding,
		}
		result = append(result, bankStatement)
	}
	return result, nil
}
func normalizeCurrency(old string) string {
	s := strings.Replace(old, ",", ".", -1)
	return s
}

func decodeBase64(base64 string) (io.Reader, error) {
	bytesArr, err := b64.StdEncoding.DecodeString(base64)
	if err != nil {
		return nil, err
	}
	byteReader := bytes.NewReader(bytesArr)
	return byteReader, nil
}

func readCSV(reader io.Reader) (*[][]string, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return &records, nil
}
