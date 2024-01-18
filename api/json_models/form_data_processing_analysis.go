package jsonmodels

import "github.com/google/uuid"

type FormDataProcessingAnalysis struct {
	UuidFK             uuid.UUID `gorm:"column:uuid_fk" json:"uuid_fk"`
	DataProcessingTech string    `gorm:"column:data_processing_tech" json:"data_processing_tech"`
	DataTypes          string    `gorm:"column:data_types" json:"data_types"`
}

func (FormDataProcessingAnalysis) TableName() string {
	return "ftf.form_data_processing_analysis"
}
