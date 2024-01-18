package jsonmodels

import "github.com/google/uuid"

type FormDataCollection struct {
	UuidFK                  uuid.UUID `gorm:"column:uuid_fk" json:"uuid_fk"`
	PowerPointFileName      string    `gorm:"column:powerpoint_file_name" json:"powerpoint_file_name"`
	IsNewData               string    `gorm:"column:is_this_new_data" json:"is_this_new_data"`
	DatasetName             string    `gorm:"column:dataset_name" json:"dataset_name"`
	DataStorageRequirements string    `gorm:"column:data_storage_requirements" json:"data_storage_requirements"`
}

func (FormDataCollection) TableName() string {
	return "ftf.form_data_collection"
}
