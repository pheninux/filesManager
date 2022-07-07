package common

import (
	"fileManager2/pkg/models"
	"os"
	"sync"
)

type IFileManager interface {
	CopyOrMove(src, dst, action string, wgc *sync.WaitGroup)
	CheckExtAndCopy(entry os.FileInfo, param *models.DataTemplate, wgp *sync.WaitGroup)
	StartProcessing(param *models.DataTemplate)
	Process(fi []os.FileInfo, param *models.DataTemplate, wg *sync.WaitGroup)
	MakeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup)
}
