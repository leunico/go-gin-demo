package models

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	// 教务类型
	EXAMINEE_CERTIFICATE_CGK = 5
)

type Examinees struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Certificates string `json:"certificates"`
	AdmissionTicket string `json:"admission_ticket"`
	CertificateType uint8 `json:"certificate_type"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(certificates, admission_ticket, password string, category uint8) (bool, error) {
	var auth Examinees
	if (category == EXAMINEE_CERTIFICATE_CGK) {
		return false, nil
	}

	err := db.Select("examinees.id,examinees.password,examination_examinees.id as examination_examinee_id").
		Joins("left join examination_examinees on examination_examinees.examinee_id = examinees.id").
		Where(Examinees{Certificates: certificates, CertificateType: category}).
		Where("examination_examinees.admission_ticket = ?", admission_ticket).
		Where("examination_examinees.status = ?", EXAMINATION_EXAMINEE_STATUS_OK).
		First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

    err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if auth.ID > 0 && err == nil {
		return true, nil
	}

	return false, nil
}