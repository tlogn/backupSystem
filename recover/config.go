package recover

import (
	"backupSystem/utils"
)

type recoverHandler func (recoverInfo *utils.RecoverInfo) error

var (
	recoverHandlerMap = map[string]*recoverHandler {
		utils.FILE_TYPE_FILE: &normalFileRecover,
		utils.FILE_TYPE_DIR: &dirRecover,
		utils.FILE_TYPE_HARDLINK: &hardLinkRecover,
		utils.FILE_TYPE_SYMLINK: &symLinkRecover,
		utils.FILE_TYPE_PIPELINE: &pipelineRecover,
	}
	normalFileRecover 	recoverHandler
	dirRecover 			recoverHandler
	hardLinkRecover		recoverHandler
	symLinkRecover		recoverHandler
	pipelineRecover		recoverHandler
)

func init() {
	normalFileRecover = NormalFileRecover
	dirRecover = DirRecover
	hardLinkRecover = HardLinkRecover
	symLinkRecover = SymLinkRecover
	pipelineRecover = PipelineRecover
}