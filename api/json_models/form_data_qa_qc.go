package jsonmodels

import "github.com/google/uuid"

type FormDataQaQc struct {
	UuidFK                       uuid.UUID `gorm:"column:uuid_fk" json:"uuid_fk"`
	QaQcMethods                  string    `gorm:"column:qa_qc_methods" json:"qa_qc_methods"`
	QaQcStrategies               string    `gorm:"column:qa_qc_strategies" json:"qa_qc_strategies"`
	ValidationCalibrationMethods string    `gorm:"column:validation_calibration_methods" json:"validation_calibration_methods"`
}

func (FormDataQaQc) TableName() string {
	return "ftf.form_data_qa_qc"
}
