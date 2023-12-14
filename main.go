package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qqMelon/mynotor-addons/internal"
)

func main() {
	dir := "/home/maximen/projects/mynotor-addons/"
	baseDir := "/home/maximen/projects/mynotor-addons/AddOns/"

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		fileList, err := internal.ListFiles(dir, baseDir)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error lors de la recuperation du dossier")
			return
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"files": fileList,
		})

	})

	r.POST("/create", func(ctx *gin.Context) {
		newFolderName := ctx.PostForm("folderName")
		newFolderPath := filepath.Join(dir, "AddOns", newFolderName)

		err := os.Mkdir(newFolderPath, os.ModePerm)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error lors de la création du dossier")
			return
		}

		ctx.Redirect(http.StatusSeeOther, "/")
	})

	r.Static("/static", "./static")
	r.Run(":8082")
}
