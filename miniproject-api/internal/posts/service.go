package posts

import "miniproject-api/pkg/database"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreatePost(content string, threadID uint) Post {
	post := Post{
		Content:  content,
		ThreadID: threadID,
	}
	database.DB.Create(&post)
	return post
}

func (s *Service) ListPosts(threadID uint) []Post {
	var posts []Post
	database.DB.Where("thread_id = ?", threadID).Find(&posts)
	return posts
}
