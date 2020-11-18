package model

import "time"

// UserWeightHistory model
type UserWeightHistory struct {
	ID     uint8     `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Weight float32   `gorm:"column:weight;type:float unsigned"`
	Date   time.Time `gorm:"column:date;type:date"`
	UserID uint8     `gorm:"column:user_id;type:mediumint unsigned"`
	User   User      `gorm:"foreignKey:UserID;references:ID"`
}

// TableName Function to change the name of a table. In this case of UserWeightHistory model
func (uwh *UserWeightHistory) TableName() string {
	return "user_weight_history"
}

// CreateUserWeightHistory Creates a new weight history data point
func (uwh *UserWeightHistory) CreateUserWeightHistory() string {
	return "CreateUserWeightHistory"
}

// UpdateUserWeightHistory Updates an existing weight history data point
func (uwh *UserWeightHistory) UpdateUserWeightHistory() string {
	return "UpdateUserWeightHistory"
}