package models

type Scannable interface {
	ScanJson(JsonStr string) error
}
