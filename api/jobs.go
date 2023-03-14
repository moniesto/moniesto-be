package api

import (
	"fmt"

	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) UpdatePostStatus() {

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		// TODO: better system error
		systemError.Log("CRON JOB - Update Post Status: db error while getting active posts")
		return
	}

	for _, post := range activePosts {
		err := server.service.UpdatePostStatus(post)
		if err != nil {
			systemError.Log(fmt.Sprintf("CRON JOB - Update Post Status: %s", err))
		}
	}
}
