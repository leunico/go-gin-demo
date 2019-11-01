package models

const (
	// 考前检测
	TYPE_BEFORE = 0
)

type ExamineeVideo struct {
    Model

    ExaminationExamineeId uint `json:"examination_examinee_id"`
    VideoUrl string `json:"video_url"`
    Type uint8 `json:"type"`
}

func (ExamineeVideo) TableName() string {
    return "examinee_videos"
}