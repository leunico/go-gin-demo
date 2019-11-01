package models

const (
	// 离线
    CATEGORY_OFFLINE = 3
    // ...
	CATEGORY_EXAM_READ = 4
)

type ExamineeOperation struct {
    Model

    ExaminationExamineeId uint `json:"examination_examinee_id"`
    Category uint8 `json:"category"`
    SourceId uint `json:"source_id"`
    Remark string `json:"remark"`
}

func (ExamineeOperation) TableName() string {
    return "examinee_operations"
}