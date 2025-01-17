// Code generated by goctl. DO NOT EDIT.
package types

type AuthorInfo struct {
	ID              int64  `json:"id"`
	NickName        string `json:"nickname"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type VideoInfo struct {
	VideoID       int64      `json:"video_id"`
	Author        AuthorInfo `json:"author"`
	PlayURL       string     `json:"play_url"`
	CoverURL      string     `json:"cover_url"`
	FavoriteCount int64      `json:"favorite_count"`
	CommentCount  int64      `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
	Partition     int64      `json:"partition"`
	CreateTime    string     `json:"create_time"`
}

type VideoList struct {
	List []*VideoInfo `json:"video_list"`
}

type PublishVideoReq struct {
	VideoUrl  string `json:"video_url"`
	CoverUrl  string `json:"cover_url"`
	Title     string `json:"title"`
	Partition int64  `json:"partition"`
}

type FeedUserReq struct {
	UserId   int64 `form:"user_id"`
	Sort     int64 `form:"sort"  validate:"gte=1,lte=2"`
	MaxValue int64 `form:"max_value"`
}

type FeedResp struct {
	NextMaxValue int64 `json:"next_max_value"`
	IsFinal      bool  `json:"is_final"`
	VideoList
}

type FeedHotReq struct {
	MaxHot int64 `form:"max_hot"`
}

type FeedHotResp struct {
	NextMaxHot int64 `form:"next_max_hot"`
	VideoList
}

type FeedHomeReq struct {
	LatestTime int64 `form:"latest_time"`
}

type FeedHomeResp struct {
	NextTime int64 `json:"next_time"`
	VideoList
}

type FeedFollowReq struct {
	Sort     int64 `form:"sort" validate:"gte=1,lte=2"`
	MaxValue int64 `form:"max_value"`
}

type FeedPartitionReq struct {
	Sort      int64 `form:"sort" validate:"gte=1,lte=2"`
	MaxValue  int64 `form:"max_value"`
	Partition int64 `form:"partition" validate:"gte=1,lte=5"`
}

type DeleteVideoReq struct {
	VideoId   int64 `form:"video_id"`
	Partition int64 `form:"partition" validate:"gte=1,lte=5"`
}

type SearchVideoReq struct {
	KeyWords string `form:"key_words"`
	Page     int64  `form:"page" validate:"gte=1"` // 从1开始
}
