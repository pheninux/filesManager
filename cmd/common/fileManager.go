package common

import (
	"fileManager2/pkg/models"
	"os"
	"sync"
)

type IFileManager interface {
	CopyOrMove(src, dst string, dt *models.DataTemplate, wgc *sync.WaitGroup)
	CheckExtAndCopy(entry os.FileInfo, dt *models.DataTemplate, wgp *sync.WaitGroup, cc chan int)
	StartProcessing(dt *models.DataTemplate, cc chan int)
	Process(fi []os.FileInfo, dt *models.DataTemplate, wg *sync.WaitGroup, cc chan int)
	MakeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup)
}
