package dto

type PostData struct {
	PostId    int    `json:"postId"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	UserId    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	Content   string `json:"content"`
}

type CreatePostRequest struct {
	Content string `json:"content"`
}

type GetPostsResponse struct {
	QueryListResponse
	List []*PostData `json:"list"`
}
