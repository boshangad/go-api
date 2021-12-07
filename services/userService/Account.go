package userService

import "time"

type Account struct {
	// 昵称
	Nickname string `json:"nickname,omitempty" binding:""`
	// 姓名
	Name string `json:"name,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty"`
	// 生日
	Birthday time.Time `json:"birthday,omitempty"`
	// 头像
	Avatar    string `json:"avatar,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty" binding:""`
	// 年龄
	Age int `json:"age,omitempty"`
}

// 更新用户身份
func (that Account) Update() {

}

// 更新用户昵称
func (that Account) UpdateNickname() {

}

// 更新用户姓名
func (that Account) UpdateName() {

}

// 更新生日
func (that Account) UpdateBirthday() {

}

// 更新年龄
func (that Account) UpdateAge() {

}

func (that Account) UpdateAvatar() {

}
