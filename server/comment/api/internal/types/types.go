// Code generated by goctl. DO NOT EDIT.
package types

type CommentDetailRequest struct {
	CommentId int64 `form:"comment_id" vd:"$>0;msg:'comment_id error'"`
}

type CommentDetailResponse struct {
	CommentList []*Comment `json:"comment_list"`
}

type CommentActionRequest struct {
	VideoId    int64  `json:"video_id" vd:"$>0;msg:'video_id error'"`
	ActionType int64  `json:"action_type" vd:"$==1||$==2;msg:'action_type error'"`
	Content    string `json:"content,optional"`
	ParentId   int64  `json:"parent_id,optional"`
	CommentId  int64  `json:"comment_id,optional"`
}

type CommentActionResponse struct {
	Comment *Comment `json:"comment,omitempty"`
}

type CommentListRequest struct {
	VideoId       int64 `form:"video_id" vd:"$>0;msg:'video_id error'"`
	LastCommentId int64 `form:"last_comment_id"`
}

type CommentListResponse struct {
	CommentList []*Comment `json:"comment_list"`
}

type Comment struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	ParentId   int64  `json:"parent_id"`
	SubCount   int64  `json:"sub_count"`
	User       *User  `json:"user"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        *User  `json:"author" copier:"User"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"name"`
	Avatar          string `json:"avatar"`
	FollowCount     int64  `json:"follow_count"`
	TotalFavorited  int64  `json:"total_favorited"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	FollowerCount   int64  `json:"follower_count"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	IsFollow        bool   `json:"is_follow"`
}
