package tracking

import (
	"api/database"
	"api/model/user"
	"time"
)

type TrackedSubmission struct {
	ID          int       `gorm:"primaryKey;not null" json:"id"`
	User        user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID      int       `json:"user_id" gorm:"foreignKey:UserID;index"`
	IsConfirmed bool      `json:"is_confirmed"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	SubmitDate  time.Time `json:"submit_date"`
}

type TrackedAttachment struct {
	ID           int               `gorm:"primaryKey;not null" json:"id"`
	Path         string            `json:"path"`
	Submission   TrackedSubmission `json:"submission" gorm:"foreignKey:SubmissionID"`
	SubmissionID int               `json:"submission_id" gorm:"foreignKey:SubmissionID;index"`
}

func AddSubmission(submission *TrackedSubmission) interface{} {
	results := database.DBConn.Create(&submission)
	return results.Error
}
