package domain

import (
	"task2/model/response"
	"time"
)

type Content struct {
	id         int
	title      string
	image_name string
	body       string
	status     int
	created    time.Time
	modified   time.Time
}

func (c *Content) ToResponseUserContentAll() response.ResponseUserContentAll {
	return response.ResponseUserContentAll{
		Id:         c.id,
		Title:      c.title,
		Image_name: c.image_name,
		Body:       c.body,
		Status:     c.status,
		Created:    c.created,
		Modified:   c.modified,
	}
}

func (c *Content) ToResponseUserContentCreated() response.ResponseUserContentCreated {
	return response.ResponseUserContentCreated{
		Id:         c.id,
		Title:      c.title,
		Image_name: c.image_name,
		Body:       c.body,
		Status:     c.status,
		Created:    c.created,
	}
}

func (c *Content) ToResponseUserContentByContentId() response.ResponseUserContentByContentId {
	return response.ResponseUserContentByContentId{
		Id:         c.id,
		Title:      c.title,
		Image_name: c.image_name,
		Body:       c.body,
		Status:     c.status,
		Created:   	c.created,
		Modified:   c.modified,
	}
}

func (c *Content) ToResponseUserContentUpdateByContentId() response.ResponseUserContentUpdateByContentId {
	return response.ResponseUserContentUpdateByContentId{
		Id:         c.id,
		Title:      c.title,
		Image_name: c.image_name,
		Body:       c.body,
		Status:     c.status,
		Modified:   c.modified,
	}
}

// SETTER
func (c *Content) SetContent(id *int, title, image_name, body *string, status *int, created, modified *time.Time) {
	c.id = *id
	c.title = *title
	c.image_name = *image_name
	c.body = *body
	c.status = *status
	c.created = *created
	c.modified = *modified
}

func (c *Content) SetId(id *int) {
	c.id = *id
}

func (c *Content) SetTitle(title *string) {
	c.title = *title
}

func (c *Content) SetImageName(image_name *string) {
	c.image_name = *image_name
}

func (c *Content) SetBody(body *string) {
	c.body = *body
}

func (c *Content) SetStatus(status *int) {
	c.status = *status
}

func (c *Content) SetCreated(created *time.Time) {
	c.created = *created
}

func (c *Content) SetModified(modified *time.Time) {
	c.modified = *modified
}

// GETTER
func (c *Content) GetContent() (*int, *string, *string, *string, *int, *time.Time, *time.Time) {
	return &c.id, &c.title, &c.image_name, &c.body, &c.status, &c.created, &c.modified
}

func (c *Content) GetId() *int {
	return &c.id
}

func (c *Content) GetTitle() *string {
	return &c.title
}

func (c *Content) GetImageName() *string {
	return &c.image_name
}

func (c *Content) GetBody() *string {
	return &c.body
}

func (c *Content) GetStatus() *int {
	return &c.status
}

func (c *Content) GetCreated() *time.Time {
	return &c.created
}

func (c *Content) GetModified() *time.Time {
	return &c.created
}

func NewContentDomain(id int, title, image_name, body string, status int, created, modified time.Time) *Content {
	return &Content{
		id:         id,
		title:      title,
		image_name: image_name,
		body:       body,
		status:     status,
		created:    created,
		modified:   modified,
	}
}
