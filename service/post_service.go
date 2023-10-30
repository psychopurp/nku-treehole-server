package service

import (
	"fmt"

	"nku-treehole-server/dto"
	"nku-treehole-server/model"
	"nku-treehole-server/pkg/logger"
)

type PostService struct{}

func (s *PostService) GetPosts(page, limit int) (*dto.GetPostsResponse, error) {
	post := &model.Post{}
	res := &dto.GetPostsResponse{
		QueryListResponse: dto.QueryListResponse{
			Page:  page,
			Limit: limit,
		},
		List: make([]*dto.PostData, 0, 32),
	}
	posts, totalCount, err := post.GetPosts(page, limit)
	if err != nil {
		logger.Errorf("GetPosts err=%v ", err)
		return nil, err
	}
	res.Total = (int(totalCount) + limit - 1) / limit
	if len(posts) == 0 {
		return res, err
	}
	var userIds []int64
	userIdMap := make(map[int64]*model.User)
	for _, i := range posts {
		if _, ok := userIdMap[i.UserId]; !ok {
			userIds = append(userIds, i.UserId)
			userIdMap[i.UserId] = nil
		}
	}

	user := &model.User{}
	users, err := user.FindInBatches(userIds)
	if err != nil {
		logger.Errorf("GetPosts err=%v ", err)
		return nil, err
	}
	for _, u := range users {
		userIdMap[u.ID] = u
	}

	for _, i := range posts {
		// 用户不存在则跳过
		if u, exist := userIdMap[i.UserId]; !exist || u == nil {
			totalCount--
			continue
		} else {
			res.List = append(res.List, &dto.PostData{
				PostId:    i.ID,
				Avatar:    u.Avatar,
				Username:  u.Name,
				UserId:    fmt.Sprint(u.ID),
				CreatedAt: i.CreatedAt.String(),
				Content:   i.Content,
			})
		}
	}
	return res, nil
}
