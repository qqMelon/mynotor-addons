package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qqMelon/mynotor-addons/internal"
)

func main() {
	dir := "/home/maximen/project/perso/mynotor-addons/"
	baseDir := "/home/maximen/project/perso/mynotor-addons/AddOns/"

  dlState := ""

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

	r.GET("/browse", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "browse.html", gin.H{
      "state": dlState,
			// "files": fileList,
		})
	})

	r.POST("/create", func(ctx *gin.Context) {
    dlState = "Downloading ..."
		newFolderName := ctx.PostForm("folderName")
		newFolderPath := filepath.Join(dir, "AddOns", newFolderName)

		err := os.Mkdir(newFolderPath, os.ModePerm)
		if err != nil {
      dlState = "Error in downloading"
			ctx.String(http.StatusInternalServerError, "Error lors de la création du dossier")
			return
		}

    dlState = "Successfuly installed"
		ctx.Redirect(http.StatusSeeOther, "/browse")
	})

	r.Static("/static", "./static")
	r.Run(":8082")
}
