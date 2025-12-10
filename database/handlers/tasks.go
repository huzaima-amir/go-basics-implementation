package handlers

import (
	"time"
	"gorm.io/gorm"
	"database/models"
  "fmt"
)

func createTask(db *gorm.DB, title,description string, deadline time.Time) uint { // creating new task - works
  task := models.Task{Title: title,Description: description, Deadline: deadline, Status: "Pending"}
  db.Select("ID","Title", "Description", "Deadline", "Status").Create(&task)
  return task.ID
}

func startTask(db *gorm.DB, taskid uint) error {
    var task models.Task
    if err := db.First(&task, taskid).Error; err != nil {
        return err // task not found
    }

    // Only allow start if task is Pending
    if task.Status == "Finished" {
        return fmt.Errorf("cannot start a finished task")
    }
    if task.Status != "Pending" {
        return fmt.Errorf("task already started")
    }

    task.StartedAt = time.Now()
    task.Status = "In Progress"
    return db.Save(&task).Error
}

func endTask(db *gorm.DB, taskid uint) error {
    var task models.Task
    if err := db.First(&task, taskid).Error; err != nil {
        return err // task not found
    }

    // Only allow end if task is In Progress
    if task.Status != "In Progress" {
        return fmt.Errorf("cannot end task that hasn't started")
    }

    task.FinishedAt = time.Now()
    task.Status = "Finished"
    return db.Save(&task).Error
}

func deleteTask(db *gorm.DB, taskid uint){ // remove task = works
  db.Delete(&models.Task{}, taskid)
}

func addSubtaskToTask(){ // adding subtask to the subtaskchecklist in a specific task
  
}

func updateTaskOverdueStatus(){ //use hook(callback)

}