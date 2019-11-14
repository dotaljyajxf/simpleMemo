package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"strings"
)

func LocalStatic(relativePath string, root string, group *gin.Engine) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := createStaticHandler(relativePath, Dir(root, false), group)
	urlPattern := path.Join(relativePath, "/*filepath")

	// Register GET and HEAD handlers
	group.GET(urlPattern, handler)
	group.HEAD(urlPattern, handler)
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyfilesFS{fs}
}

type onlyfilesFS struct {
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

// Readdir overrides the http.File default implementation.
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	// this disables directory listing
	return nil, nil
}

func (fs onlyfilesFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

func createStaticHandler(relativePath string, fs http.FileSystem, group *gin.Engine) gin.HandlerFunc {
	absolutePath := joinPaths(group.BasePath(), relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *gin.Context) {
		if _, nolisting := fs.(*onlyfilesFS); nolisting {
			c.Writer.WriteHeader(http.StatusNotFound)
		}

		file := c.Param("filepath")
		fileExt := path.Ext(file)
		logrus.Infof("fileExt : %s", fileExt)
		if fileExt == ".js" {
			file = file + ".gz"
			c.Writer.Header().Set("Content-Encoding", "gzip")
		}
		logrus.Infof("fileAfter : %s", file)
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			c.AbortWithStatus(http.StatusNotFound)
			//c.handlers = group.engine.noRoute
			//// Reset index
			//c.index = -1
			return
		}
		//c.Request.Header.Set("")
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}
