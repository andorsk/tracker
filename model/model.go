package model

type ObjectModelInterface interface {
	Push() error
	Update() error
	Get() error
	Delete() error
}
