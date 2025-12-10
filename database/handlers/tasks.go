package handlers

import (
	"time"
	"gorm.io/gorm"
	"database/models"
)

func createTask(db *gorm.DB, title,description string, deadline time.Time) uint { // creating new task - works
  task := models.Task{Title: title,Description: description, Deadline: deadline, Status: "Pending"}
  db.Select("ID","Title", "Description", "Deadline", "Status").Create(&task)
  return task.ID
}

func startTask(db *gorm.DB, taskid uint){ // start task from pending block - works but need to figure out the id
  z := time.Now()
  db.Model(&models.Task{}).Where("id = ?", taskid).Update("started_at",z)
  db.Model(&models.Task{}).Where("id = ?", taskid).Update("status","In Progress")
}

func endTask(db *gorm.DB, taskid uint){ // end started task - !!!!!!! should only work if task has started
  z := time.Now()
  db.Model(&models.Task{}).Where("id = ?", taskid).Update("finished_at",z)
  db.Model(&models.Task{}).Where("id = ?", taskid).Update("status","Finished")
}

func deleteTask(db *gorm.DB, taskid uint){ // remove task = works
  db.Delete(&models.Task{}, taskid)
}

func addSubtaskToTask(){ // adding subtask to the subtaskchecklist in a specific task

}