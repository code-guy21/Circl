package models

type Friend struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json: "user_id"`
	FriendID uint `json:"friend_id"` 
	Status string `json:"status"`
}