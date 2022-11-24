package models

import "time"

type Post struct {
	Status      int32     `json:"status" db:"status" `
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthuorName      string `json:"authuor_name"`
	Votes            int64  `json:"votes"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community_detail"`
}
