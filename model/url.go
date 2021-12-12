package model

import (
	"time"
)

type Url struct {
	ID          uint      `gorm:"primaryKey"`
	Path        string    `gorm:"type:varchar(64);uniqueIndex;comment:短key"`
	OriginalUrl string    `gorm:"type:varchar(255);comment:原链接"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
