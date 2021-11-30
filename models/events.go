package models

type TriggerEvent struct {
	ID string
	Batch bool
	Product *Product
}
