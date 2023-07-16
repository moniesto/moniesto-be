package api

import (
	"fmt"

	"github.com/moniesto/moniesto-be/util/systemError"
)

func (server *Server) UpdatePostStatus() {

	activePosts, err := server.service.GetAllActivePosts()
	if err != nil {
		systemError.Log("JOB ERROR: POST STATUS => db error while getting active posts")
		return
	}

	fmt.Println("# Active posts", len(activePosts))

	for i, post := range activePosts {
		fmt.Printf("post start: %d id: %s\n", i, post.ID)
		err := server.service.UpdatePostStatus(post)
		if err != nil {
			systemError.Log(fmt.Sprintf("JOB ERROR: POST STATUS => %s", err))
		}
	}
}

func (server *Server) PayoutToMoniest() {

	pendingPayouts, err := server.service.GetAllPendingPayouts()
	if err != nil {
		systemError.Log("JOB ERROR: PAYOUT => db error while getting pending payouts", err.Error())
		return
	}

	for _, pendingPayout := range pendingPayouts {
		err := server.service.PayoutToMoniest(pendingPayout)
		if err != nil {
			systemError.Log(err)
		}
	}
}
