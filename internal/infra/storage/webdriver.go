package storage

import "context"

type WebDriver struct {
}

func (w *WebDriver) DownloadBankStatement(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewWebDriver() *WebDriver {
	return &WebDriver{}
}
