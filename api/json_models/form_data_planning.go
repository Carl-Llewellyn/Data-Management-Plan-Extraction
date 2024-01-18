package jsonmodels

import "github.com/google/uuid"

type ProjectStatus string

const (
	Proposed ProjectStatus = "proposed"
	Submited ProjectStatus = "submitted"
	Complete ProjectStatus = "complete"
)

type FormDataPlanning struct {
	UuidFK                uuid.UUID     `gorm:"column:uuid_fk" json:"uuid_fk"`
	ProjectTitle          string        `gorm:"column:project_title" json:"project_title"`
	PublicationTitle      string        `gorm:"column:publication_title" json:"publication_title"`
	ProjectDescription    string        `gorm:"column:project_description" json:"project_description"`
	DataOwnership         string        `gorm:"column:data_ownership" json:"data_ownership"`
	ProjectStatus         ProjectStatus `gorm:"column:project_status" json:"project_status"`
	ProjectTimeFrameStart string        `gorm:"column:project_time_frame_start" json:"project_time_frame_start"`
	ProjectTimeFrameEnd   string        `gorm:"column:project_time_frame_end" json:"project_time_frame_end"`
	ShapeFileName         string        `gorm:"column:shapefile_name" json:"shapefile_name"`
}

func (FormDataPlanning) TableName() string {
	return "ftf.form_data_planning"
}
