package model

type ObjectModelInterface interface {
	Push() error
	Update() error
	Get() (interface{}, error)
	Delete() error
}
