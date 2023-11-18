package main

import (
	sqlFetch "fivem-stats/sql-fetch"
	"net/http"
	"html/template"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
)

func startRouter() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
        "dec": func(i int) int { return i - 1 },
    })

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static/")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "pong",
		})
	  })

	router.GET("/charactersByBank", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, charactersByBank)
	})

	router.GET("/characters", func(c *gin.Context) {
		characters := charactersByBank
		blocksize := 25
		currentshown := c.DefaultQuery("currentshown", "")
		var intcurrentshown int
		var html string
		if currentshown == "" {
			intcurrentshown = blocksize
			html = "characters.html"
			if intcurrentshown > len(characters) {
				intcurrentshown = 0
			} else {
				characters = characters[0:intcurrentshown]
			}
		} else {
			html = "newBlock.html"
			var err error
			intcurrentshown, err = strconv.Atoi(currentshown)
			if err != nil {
				fmt.Printf("htmx more string: %q\n", err)
			}
			if intcurrentshown + blocksize >= len(characters) {		
				characters = characters[intcurrentshown:]
				intcurrentshown = 0
			} else {
				characters = characters[intcurrentshown:intcurrentshown+blocksize]
				intcurrentshown = intcurrentshown + blocksize
			}
		}
		fmt.Printf("final currentshown: %d\n", intcurrentshown)
		c.HTML(http.StatusOK, html, gin.H{
			"characters": characters,
			"currentShown": intcurrentshown,
		})


	})

	router.GET("/character/:citizenID", func(c *gin.Context) {
		citizenID := c.Param("citizenID")

		// Find the character in the list based on the citizenID
		var character sqlFetch.Character
		var found bool
		for _, p := range charactersByBank {
			if p.CitizenID == citizenID {
				character = p
				found = true
				break
			}
		}

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
		}

		// Do something with the character data, for example, render a template or return JSON
		c.HTML(http.StatusOK, "character.html", character)
	})
	router.Run("localhost:80")
}