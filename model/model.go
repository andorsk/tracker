package model

type ModelInterface interface {
	Push() error
	Update() error
	Get() error
	Gets() error
	Detete() error
}
