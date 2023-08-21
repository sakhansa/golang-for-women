package controllers

import (
	"assignment-project-rest-api/database"
	"assignment-project-rest-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create student
//
//	{"name":"Fitri","age":20,"scores":[{"assignment_title":"Assignment Project 1","description":"Create simple API without middleware","score":95}]}
func CreateStudent(ctx *gin.Context) {
	db := database.GetDB()

	var newStudent models.Student

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		Name:   newStudent.Name,
		Age:    newStudent.Age,
		Scores: newStudent.Scores,
	}

	err := db.Create(&student).Error
	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
	})

}

// get all student
func GetAllStudent(ctx *gin.Context) {
	db := database.GetDB()
	var results = []models.Student{}

	res := db.Preload("Scores").Find(&results)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

// update Student by ID
func UpdateStudent(ctx *gin.Context) {
	studentId := ctx.Param("studentID")
	db := database.GetDB()

	var newStudent models.Student

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	student := models.Student{
		Name:   newStudent.Name,
		Age:    newStudent.Age,
		Scores: newStudent.Scores,
	}

	// Fetch the existing Student from the database along with its associated Scores
	existingStudent := models.Student{}
	res := db.Preload("Scores").First(&existingStudent, uint(id))
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the student"})
		return
	}

	// Start a transaction
	tx := db.Begin()

	// Update the fields of the fetched Student
	existingStudent.Name = student.Name
	existingStudent.Age = student.Age

	// Update the fields of the associated Scores
	for i, newScore := range student.Scores {
		if i < len(existingStudent.Scores) {
			existingStudent.Scores[i].AssignmentTitle = newScore.AssignmentTitle
			existingStudent.Scores[i].Description = newScore.Description
			existingStudent.Scores[i].Score = newScore.Score

			// Save the Score update within the transaction
			if err := tx.Save(&existingStudent.Scores[i]).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update a Score"})
				return
			}
		} else {
			//  Assign existing Student ID as FK to new Score
			newScore.StudentId = existingStudent.ID
			// If the newScore is a new Score, add it to the existingStuden.Scores slice
			existingStudent.Scores = append(existingStudent.Scores, newScore)
			// Save the new Score within the transaction
			if err := tx.Create(&existingStudent.Scores[i]).Error; err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add a new Score"})
				return
			}
		}
	}

	// Save the changes to the Student within the transaction
	if err := tx.Save(&existingStudent).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Student"})
		return
	}

	// Commit the transaction
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    existingStudent,
	})
}

// Delete Student by ID
func DeleteStudent(ctx *gin.Context) {
	studentId := ctx.Param("studentID")
	db := database.GetDB()

	var student models.Student

	id, err := strconv.Atoi(studentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Fetch the existing Student from the database along with its associated Scores
	res := db.Preload("Scores").First(&student, uint(id))
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the student"})
		return
	}

	// Start a transaction
	tx := db.Begin()

	// Delete the associated Scores within the transaction
	if err := tx.Delete(&student.Scores).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the associated Scores"})
		return
	}

	// Delete the Student from the database within the transaction
	if err := tx.Delete(&student).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the Student"})
		return
	}

	// Commit the transaction
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
	})
}
