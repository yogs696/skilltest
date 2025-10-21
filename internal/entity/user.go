package entity

type (

	// User entity.
	User struct {
		ID           uint64 `gorm:"primaryKey" json:"id"`
		Username     string `gorm:"not null;size:25;index;index:user_idx_username_created_at_str" json:"username"`
		Email        string `gorm:"not null;size:50;unique;index" json:"email"`
		Password     string `json:"password" binding:"required"`
		CreatedAtStr string `gorm:"not null;index;index:user_idx_username_created_at_str;type:date;" json:"createdAt"`
	}
)
