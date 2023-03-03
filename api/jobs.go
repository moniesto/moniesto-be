package api

import "fmt"

func (server *Server) UpdatePostStatus() {

	fmt.Println("CRON JOB RUN")

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		// show error
	}

	_ = activePosts
}
