package domain

import (
	"task2/model/response"
	"time"
)

type UserContent struct {
	userId       int
	contentId    int
	point        int
	updated_date time.Time
}

type UserContentJoin struct {
	// userId     int
	contentId  int
	title      string
	image_name string
	body       string
	status     int
	created    time.Time
	modified   time.Time
}

func (c *UserContent) ToResponseUserContentPoint() response.ResponseUserContentPoint {
	return response.ResponseUserContentPoint{
		UserId:    c.userId,
		ContentId: c.userId,
		Point:     c.point,
	}
}

func (c *UserContentJoin) ToResponseUserContentByContentId() response.ResponseUserContentByContentId {
	return response.ResponseUserContentByContentId{
		Id:         c.contentId,
		Title:      c.title,
		Image_name: c.image_name,
		Body:       c.body,
		Status:     c.status,
		Created:    c.created,
		Modified:   c.modified,
	}
}

// SETTER
func (c *UserContent) SetUserContent(userId, contentId, point *int, updated_date *time.Time) {
	c.userId = *userId
	c.contentId = *contentId
	c.point = *point
	c.updated_date = *updated_date
}

func (c *UserContent) SetUserId(userId *int) {
	c.userId = *userId
}

func (c *UserContent) SetContentId(contentId *int) {
	c.contentId = *contentId
}

func (c *UserContent) SetPoint(point *int) {
	c.point = *point
}

func (c *UserContent) SetUpdatedDate(updated_date *time.Time) {
	c.updated_date = *updated_date
}

// GETTER
func (c *UserContent) GetUserContent() (*int, *int, *int, *time.Time) {
	return &c.userId, &c.contentId, &c.point, &c.updated_date
}

func (c *UserContent) GetUserId() *int {
	return &c.userId
}

func (c *UserContent) GetContentId() *int {
	return &c.contentId
}

func (c *UserContent) GetPoint() *int {
	return &c.point
}

func (c *UserContent) GetUpdatedDate() *time.Time {
	return &c.updated_date
}

func NewUserContentDomain(userId, contentId, point int, updated_date time.Time) *UserContent {
	return &UserContent{
		userId:       userId,
		contentId:    contentId,
		point:        point,
		updated_date: updated_date,
	}
}
