package tinkoff

type (
	UpdateBankTransactionsCommand struct {
		FileBase64 string `json:"file,omitempty"`
	}
)
