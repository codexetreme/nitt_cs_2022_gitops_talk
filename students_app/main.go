package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // add this
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var dbConnStr = os.Getenv("DATABASE_URL")

type student struct {
	ID   string `json:"rollnumber"`
	Name string `json:"name"`
}

//go:embed templates info.json
var res embed.FS

func main() {
	if dbConnStr == "" {
		fmt.Printf("did not find a connection string")
	}
	var db, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "landingpage.html", gin.H{
			"title": "NITT CS registration app",
		})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forms.html", gin.H{
			"title": "NITT CS registration app",
		})
	})
	router.POST("/submit", func(c *gin.Context) {
		details := student{}
		err := c.Bind(&details)
		fmt.Printf("details: %v", details)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("could not bind form request"))
		}

		_, err = db.Exec("INSERT into roster VALUES ($1,$2)", details.ID, details.Name)
		if err != nil {
			c.String(500, "An error occured while persisting to database: %v", err)
		} else {
			c.HTML(http.StatusOK, "forms.html", gin.H{
				"Success": true,
			})
		}
	})

	router.GET("/getRegisteredStudents", func(c *gin.Context) {
		results := []student{}

		rows, err := db.Query("SELECT * FROM roster")
		defer rows.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, "An error occurred")
			log.Println(err)
		}

		for rows.Next() {
			r := student{}
			//fmt.Printf("rows: ")
			err = rows.Scan(&r.ID, &r.Name)
			if err != nil {
				fmt.Printf("error: %s", err)
			}
			results = append(results, r)
		}
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title":   "List of all registered students",
			"Results": results,
		})

	})

	router.GET("/healthcheck", func(c *gin.Context) {
		type healthCheck struct {
			Version   string `json:"version,omitempty"`
			BuildTime string `json:"build_time,omitempty"`
		}
		infoData, _ := ioutil.ReadFile("info.json")
		var h healthCheck
		_ = json.Unmarshal(infoData, &h)
		c.JSON(http.StatusOK, h)
	})

	_ = router.Run(":8080")
}
