package models

import (
	"github.com/jinzhu/gorm"

	"git.codepku.com/examinate/exam/pkg/util"
)

const (
	// 考试发布状态
	STATUS_EXAMINATION = 2
	// 考试发布成绩状态
	STATUS_ACHIEVEMENT = 3
	
	// 考试确认
	EXAMINATION_EXAMINEE_STATUS_OK = 1
)

// type Examination struct {
// 	Model

//     MatchId uint `json:"match_id"`
//     ExaminationCategoryId uint `json:"examination_category_id"`
//     Title string `json:"title"`
//     ExaminationPaperTitle string `json:"examination_paper_title"`
//     StartAt util.JSONTime `json:"start_at"`
//     EndAt util.JSONTime `json:"end_at"`
//     Description string `json:"description"`
//     ExamFileId uint `json:"exam_file_id"`
//     Status uint8 `json:"status"`
// }

type ExaminationExamineeDetail struct {
	Model
	
	ExamineeDeviceProbings []ExamineeDeviceProbling `gorm:"ForeignKey:examination_examinee_id" json:"examinee_device_probings"`
	ExamineeVideos []ExamineeVideo `gorm:"ForeignKey:examination_examinee_id" json:"examinee_videos"`
	ExamineeTencentFaces []ExamineeTencentFace `gorm:"ForeignKey:examination_examinee_id" json:"examinee_tencent_faces"`
	ExamineeOperations []ExamineeOperation `gorm:"ForeignKey:examination_examinee_id" json:"examinee_operations"`

    ExamineeId uint `json:"examinee_id"`
    ExaminationId uint `json:"examination_id"`
	AdmissionTicket string `json:"admission_ticket"`
    IsHand uint8 `json:"is_hand"`
    StartTime util.JSONTime `json:"start_time"`
    HandTime util.JSONTime `json:"hand_time"`
    SubjectiveScore uint16 `json:"subjective_score"`
    ObjectiveScore uint16 `json:"objective_score"`
    Rank uint16 `json:"rank"`
    Status uint8 `json:"status"`
    AchievementStatus uint8 `json:"achievement_status"`
    TestingStatus uint8 `json:"testing_status"`
    IsTestCheck uint8 `json:"is_test_check"`
    IsOnline uint8 `json:"is_online"`
    Title string `json:"title"`
    ExaminationCategoryId uint `json:"examination_category_id"`
    MatchCategory uint8 `json:"match_category"`
    MatchTitle string `json:"match_title"`
    ExaminationCategoryCategory string `json:"examination_category_category"`
    ExaminationCategoryTitle string `json:"examination_category_title"`
    ExaminationPaperTitle string `json:"examination_paper_title"`
    TotalScore uint `json:"total_score"`
    StartAt util.JSONTime `json:"start_at"`
    EndAt util.JSONTime `json:"end_at"`
    ExamFileUrl string `json:"exam_file_url"`
}

func (ExaminationExamineeDetail) TableName() string {
	return "examination_examinees"
}

// 获取考试详情
func GetExaminationExamineeDetail(maps interface {}) (ExaminationExamineeDetail, error) {
	var eexaminee ExaminationExamineeDetail
	err := db.Select([]string{"examinations.id as examination_id",
		"matches.category as match_category",
		"match_id",
		"examination_categories.category as examination_category_category",
		"examination_categories.title as examination_category_title",
		"examinations.title",
		"is_test_check",
		"examination_category_id",
		"examination_paper_title",
		"examination_examinees.id",
		"start_time",
		"examinations.start_at",
		"examinations.end_at",
		"exam_file_id",
		"is_hand",
		"examinations.status",
		"admission_ticket",
		"rank",
		"achievement_status",
		"subjective_score",
		"objective_score",
		"testing_status",
		"origin_filename",
		"driver_baseurl",
		"filename",
		"(select sum(total_score) from major_problems where major_problems.examination_id=examinations.id and deleted_at is null) as total_score"}).
		Joins("left join examinations on examination_examinees.examination_id = examinations.id").
		Joins("left join files on examinations.exam_file_id = files.id").
		Joins("left join matches on examinations.match_id = matches.id").
		Joins("left join examination_categories on examinations.examination_category_id = examination_categories.id").
		Where("examinations.status in (?)", []int8{STATUS_EXAMINATION, STATUS_ACHIEVEMENT}).
		Where(maps).
		Preload("ExamineeDeviceProbings").
		Preload("ExamineeTencentFaces").
		Preload("ExamineeVideos", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,examination_examinee_id,type").
				Where("type = ?", TYPE_BEFORE)
		}).
		Preload("ExamineeOperations", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,examination_examinee_id,category").
				Where("category in (?)", []int8{CATEGORY_OFFLINE, CATEGORY_EXAM_READ})
		}).
		First(&eexaminee).Error

	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return flase, err
	// }

    return eexaminee, err
}