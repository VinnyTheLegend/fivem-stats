package main

import (
	sqlFetch "fivem-stats/sql-fetch"
	"net/http"
	"html/template"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"os"
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
		var characters []sqlFetch.Character
		blocksize := 25
		descending := false
		ascending := true
		sortby := "firstname"
		firstname, lastname, bank := true, false, false
		scrollable := true
		intcurrentshown := blocksize
		characters = charactersByFirstName
		if intcurrentshown > len(characters) {
			scrollable = false
		} else {
			characters = characters[0:intcurrentshown]
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"characters": characters,
			"currentShown": intcurrentshown,
			"descending": descending,
			"ascending": ascending,
			"scrollable": scrollable,
			"sortby": gin.H{"value": sortby, "firstname": firstname, "lastname": lastname, "bank": bank},
		})
	})

	router.GET("/characters/updatelist", func(c *gin.Context) {
		var characters []sqlFetch.Character
		blocksize := 25
		currentshown := c.DefaultQuery("currentshown", "0")
		scrollable := true
		var ascending string
		descending := c.DefaultQuery("descending", "")
		sortby := c.DefaultQuery("sortby", "firstname")
		firstname, lastname, bank := false, false, false
		search := c.DefaultQuery("search", "")

		switch sortby {
		case "firstname":
			characters = charactersByFirstName
			firstname = true
		case "lastname":
			characters = charactersByLastName
			lastname = true
		case "bank":
			characters = charactersByBank
			bank = true
		}

		if descending == "" || descending == "false" {
			ascending = "true"
			descending = ""
		} else {
			ascending = ""
			s := make([]sqlFetch.Character, len(characters))
			copy(s, characters)
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			characters = s
		}

		if search != "" {
			var s []sqlFetch.Character
			for _, character := range characters {
				if strings.Contains(strings.ToLower(character.CharInfo.FirstName) + " " + strings.ToLower(character.CharInfo.LastName), strings.ToLower(search)) {
					s = append(s, character)
				}
			}
			characters = s
		}

		var intcurrentshown int
		var err error
		intcurrentshown, err = strconv.Atoi(currentshown)
		if err != nil {
			fmt.Printf("htmx more string: %q\n", err)
		}
		if intcurrentshown + blocksize >= len(characters) {		
			characters = characters[intcurrentshown:]
			scrollable = false
		} else {
			characters = characters[intcurrentshown:intcurrentshown+blocksize]
			intcurrentshown = intcurrentshown + blocksize
		}

		c.HTML(http.StatusOK, "newBlock.html", gin.H{
			"characters": characters,
			"currentShown": intcurrentshown,
			"descending": descending,
			"ascending": ascending,
			"scrollable": scrollable,
			"sortby": gin.H{"value": sortby, "firstname": firstname, "lastname": lastname, "bank": bank},
			"search": search,
		})

	})
	
	router.GET("/characters/updatefilter", func(c *gin.Context) {
		descending := c.DefaultQuery("descending", "")
		var ascending string
		sortby := c.DefaultQuery("sortby", "firstname")
		firstname, lastname, bank := sortby=="firstname", sortby=="lastname", sortby=="bank"
		search := c.DefaultQuery("search", "")


		if descending == "" || descending == "false" {
			ascending = "true"
			descending = ""
		} else {
			ascending = ""
		}

		c.HTML(http.StatusOK, "listFilter.html", gin.H{
			"descending": descending,
			"ascending": ascending,
			"sortby": gin.H{"value": sortby, "firstname": firstname, "lastname": lastname, "bank": bank},
			"search": search,
		})

	})

	router.GET("/character/:citizenID", func(c *gin.Context) {
		citizenID := c.Param("citizenID")

		// Find the character in the list based on the citizenID
		var character sqlFetch.Character
		var found bool
		for _, c := range charactersByBank {
			if c.CitizenID == citizenID {
				character = c
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
		//c.IndentedJSON(http.StatusOK, character)
	})

	router.Run("localhost:80" + os.Getenv("PORT"))
}