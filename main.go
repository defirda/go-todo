package main

// notes
// time.Now()

import (
	"database/sql"
	"go-todo/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3030)/go_todo?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/activity-groups", getAll)
	// router.GET("/activity-groups/:id", getOne)
	// router.POST("/books", createBook)
	// router.PUT("/books/:id", updateBook)
	// router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

func getAll(c *gin.Context) {
	rows, err := db.Query("SELECT activity_id, title, email, created_at, updated_at FROM activities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	activities := []models.Activity{}

	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.activity_id, &activity.title, &activity.email, &activity.created_at, &activity.updated_at); err != nil {
			log.Fatal(err)
		}
		activities = append(activities, activity)
	}

	data := gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    activities,
	}

	c.JSON(http.StatusOK, data)
}

// func getOne(c *gin.Context) {
// 	activity_id := c.Param("id")

// 	var activity Activity

// 	err := db.QueryRow("SELECT activity_id, title, email, created_at, updated_at FROM activities WHERE activity_id = ?", activity_id).Scan(&activity.activity_id, &activity.title, &activity.email, &activity.created_at, &activity.updated_at)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			c.JSON(http.StatusNotFound, gin.H{"message": "Activity not found"})
// 			return
// 		}
// 		log.Fatal(err)
// 	}

// 	data := gin.H{
// 		"status":  "Success",
// 		"message": "Success",
// 		"data":    activity,
// 	}

// 	c.JSON(http.StatusOK, data)
// }

// func createBook(c *gin.Context) {
// 	var book Book
// 	if err := c.ShouldBindJSON(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	book.ID = int(id)

// 	c.JSON(http.StatusOK, book)
// }

// func updateBook(c *gin.Context) {
// 	id := c.Param("id")

// 	var book Book
// 	if err := c.ShouldBindJSON(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result, err := db.Exec("UPDATE books SET title=?, author=? WHERE id=?", book.Title, book.Author, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rowsUpdated, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if rowsUpdated == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Book with ID %s successfully updated", id)})
// }
