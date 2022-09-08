package storage

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"encoding/csv"
	"finance/internal/model"
	"fmt"
	"io"
	"strconv"
	"time"
)

type FileSystem struct {
}

func NewFile() *FileSystem {
	return &FileSystem{}
}

func (f *FileSystem) GetDataFromBankStatement(ctx context.Context, fileBase64 string) {
	reader, err := decodeBase64(fileBase64)
	if err != nil {

	}
	data, err := readCSV(reader)
	fmt.Println(data)
}
func parseTinkoffBankStatementToStruct(data *[][]string) ([]*model.TinkoffBankStatement, error) {
	result := make([]*model.TinkoffBankStatement, len(*data))
	for rn := 1; rn < len(*data); rn++ {
		row := (*data)[rn]
		operationDate, err := time.Parse(time.RFC3339, row[0])
		if err != nil {
			return nil, err
		}
		paymentDate, err := time.Parse(time.RFC3339, row[1])
		if err != nil {
			return nil, err
		}
		operation, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return nil, err
		}
		payment, err := strconv.ParseFloat(row[6], 64)
		if err != nil {
			return nil, err
		}
		cashback, err := strconv.ParseFloat(row[8], 64)
		if err != nil {
			return nil, err
		}
		mcc, err := strconv.Atoi(row[10])
		if err != nil {
			return nil, err
		}
		bonuses, err := strconv.ParseFloat(row[12], 64)
		if err != nil {
			return nil, err
		}
		investmentBankRounding, err := strconv.ParseFloat(row[13], 64)
		if err != nil {
			return nil, err
		}
		rounding, err := strconv.ParseFloat(row[14], 64)
		if err != nil {
			return nil, err
		}
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
