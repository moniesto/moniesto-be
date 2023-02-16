package model

type GetContentPostRequest struct {
	Subscribed bool `form:"subscribed" json:"subscribed"`
	Limit      int  `form:"limit" json:"limit"`
	Offset     int  `form:"offset" json:"offset"`
}
