package models

const (
	// 考前
    TENCENT_FACE_TYPE_BEFORE = 2
)

type ExamineeTencentFace struct {
    Model

    ExaminationExamineeId uint `json:"examination_examinee_id"`
    Category uint8 `json:"category"`
    Sim float64 `json:"sim"`
    Result string `json:"result"`
    Description string `json:"description"`
    Type uint8 `json:"type"`
}

func (ExamineeTencentFace) TableName() string {
    return "examinee_tencent_faces"
}