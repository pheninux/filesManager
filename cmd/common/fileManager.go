package common

import (
	"fileManager2/pkg/models"
	"os"
	"sync"
)

type IFileManager interface {
	CopyOrMove(src, dst, action string, wgc *sync.WaitGroup)
	CheckExtAndCopy(entry os.FileInfo, param *models.File, wgp *sync.WaitGroup)
	StartProcessing(param *models.File)
	Process(fi []os.FileInfo, param *models.File, wg *sync.WaitGroup)
	MakeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup)
}
