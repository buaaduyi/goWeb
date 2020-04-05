package main

import (
	"goweb/db"
	"goweb/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := db.ConnDB("localhost:6379")
	if r == nil {
		utils.ColorPrintf("database is not ready\n", utils.Red)
		os.Exit(1)
	}
	g := gin.Default()
	g.GET("/set", func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("val")
		reply := db.SetDB(r, key, value)
		c.String(http.StatusOK, "%s\n", reply)
	})
	g.GET("/get", func(c *gin.Context) {
		key := c.Query("key")
		reply := db.GetDB(r, key)
		c.String(http.StatusOK, "%s\n", reply)
	})
	g.Run("localhost:8080")

}
