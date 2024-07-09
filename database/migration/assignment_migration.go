package migration


type AssignmentType string

const (
	Exam AssignmentType = "Exam"
	Task AssignmentType = "Task"
)

type Assignment struct {
	AssignmentID int64 `gorm:"primaryKey;autoIncrement"`
	TypeAssignment AssignmentType `gorm:"type:assignment_type;not null"`
	Grade []Grade `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
}
