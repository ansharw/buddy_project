package domain

import "task2/model/response"

type User struct {
	id       int
	fullname string
	username string
	password string
	email    string
	point    int
}

// register user (admin)
func (u *User) ToResponseRegisteredUser() response.ResponseRegisteredUser {
	return response.ResponseRegisteredUser{
		UserId:   u.id,
		Fullname: u.fullname,
		Username: u.username,
		Point:    u.point,
	}
}

// reward point (admin)
func (u *User) ToResponseRewardPoint() response.ResponseRewardPoint {
	return response.ResponseRewardPoint{
		UserId:   u.id,
		Fullname: u.fullname,
		Username: u.username,
		Point:    u.point,
	}
}

func (u *User) ToResponseUserPoint() response.ResponseUserPoint {
	return response.ResponseUserPoint{
		Point: u.point,
	}
}

// SETTER
func (c *User) SetUser(id *int, fullname, username, password, email *string, point *int) {
	c.id = *id
	c.fullname = *fullname
	c.username = *username
	c.password = *password
	c.email = *email
	c.point = *point
}

func (c *User) SetId(id *int) {
	c.id = *id
}

func (c *User) SetFullname(fullname *string) {
	c.fullname = *fullname
}

func (c *User) SetUsername(username *string) {
	c.username = *username
}

func (c *User) SetEmail(email *string) {
	c.email = *email
}

func (c *User) SetPassword(password *string) {
	c.password = *password
}

func (c *User) SetPoint(point *int) {
	c.point = *point
}

// GETTER
func (c *User) GetUser() (*int, *string, *string, *string, *int) {
	return &c.id, &c.fullname, &c.username, &c.password, &c.point
}

func (c *User) GetId() *int {
	return &c.id
}

func (c *User) GetFullname() *string {
	return &c.fullname
}

func (c *User) GetUsername() *string {
	return &c.username
}

func (c *User) GetEmail() *string {
	return &c.email
}

func (c *User) GetPassword() *string {
	return &c.password
}

func (c *User) GetPoint() *int {
	return &c.point
}

func NewUserDomain(id int, fullname, username, password, email string, point int) *User {
	return &User{
		id:       id,
		fullname: fullname,
		username: username,
		email:    email,
		password: password,
		point:    point,
	}
}
