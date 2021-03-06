package tracking

import (
	"api/database"
	"api/model/user"
	"time"
)

type TrackedSubmission struct {
	ID                 int                 `gorm:"primaryKey;not null" json:"id"`
	User               user.User           `json:"user" gorm:"foreignKey:UserID"`
	UserID             int                 `json:"user_id" gorm:"foreignKey:UserID;index"`
	IsConfirmed        bool                `json:"is_confirmed"`
	Description        string              `json:"description"`
	Subject            string              `json:"subject"`
	SubmitDate         time.Time           `json:"submit_date"`
	TrackedAttachments []TrackedAttachment `json:"tracked_attachments" gorm:"foreignKey:SubmissionID"`
}

type TrackedAttachment struct {
	ID           int    `gorm:"primaryKey;not null" json:"id"`
	Path         string `json:"path"`
	SubmissionID int    `json:"submission_id" gorm:"foreignKey:SubmissionID;index"`
}

func AddSubmission(submission *TrackedSubmission) interface{} {
	results := database.DBConn.Create(&submission)
	return results.Error
}

func AddTracking(attachment *TrackedAttachment) interface{} {
	results := database.DBConn.Create(&attachment)
	return results.Error
}

func GetYourSubmissions(userId interface{}) []TrackedSubmission {
	var submissions []TrackedSubmission
	database.DBConn.Preload("TrackedAttachments").Where("user_id = ?", userId).Order("id desc").Find(&submissions)
	return submissions
}

func GetSubmissions() []TrackedSubmission {
	var submissions []TrackedSubmission
	database.DBConn.Preload("User").Preload("TrackedAttachments").Order("id desc").Find(&submissions)
	return submissions
}

func OpenSubmissions(id interface{}) interface{} {
	response := database.DBConn.Model(&TrackedSubmission{}).Where("id = ?", id).Update("is_confirmed", false)
	return response.Error
}

func CloseSubmissions(id interface{}) interface{} {
	response := database.DBConn.Model(&TrackedSubmission{}).Where("id = ?", id).Update("is_confirmed", true)
	return response.Error
}
