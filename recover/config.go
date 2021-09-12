package recover

import (
	"backupSystem/utils"
)

type recoverHandler func (recoverInfo *utils.RecoverInfo) error

var (
	recoverHandlerMap = map[string]*recoverHandler {
		"File": &normalFileRecover,
		"Dir": &dirRecover,
		"HardLink": &hardLinkRecover,
		"SymLink": &symLinkRecover,
		"Pipeline": &pipelineRecover,
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