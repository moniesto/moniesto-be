package api

import (
	"fmt"

	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) UpdatePostStatus() {

	fmt.Println("CRON JOB RUN")

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		systemError.Log("CRON JOB - Update Post Status: db error while getting active posts")
	}

	for _, post := range activePosts {
		err := server.service.UpdatePostStatus(post)
		_ = err
	}
}
