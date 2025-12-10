package models

import (
	"time"
)

type Task struct {  // User has many Tasks
  ID uint  `gorm:"primaryKey"` //unique by default
  Title string
  Description string
  SubTaskCheckList []TaskSubTask  `gorm:"foreignKey:TaskID"`
  Status string //pending,in progress or finished
  Deadline time.Time
  StartedAt time.Time
  FinishedAt time.Time
  Overdue bool  //set to true if deadline has passed and status isnt set to finished?
  Tags []Tag  `gorm:"many2many:task_tags;"`
}

// All tasks and events have checklist that has multiple subTasks
type TaskSubTask struct {
  ID uint `gorm:"primaryKey"`
  Title string 
  Checked bool
  TaskID uint 
}