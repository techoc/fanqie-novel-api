package models

import "gorm.io/gorm"

type Label struct {
	LabelID uint //
	Name    string
	gorm.Model
}
