package utils

import (
	"strconv"
	"zumar-school/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPaginationParams(c *gin.Context) (page int, limit int) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limitStr := c.DefaultQuery("limit", strconv.Itoa(config.Cfg.PageLimit))
	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = config.Cfg.PageLimit
	}

	if limit > 100 {
		limit = 100
	}

	return page, limit
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if limit <= 0 {
			limit = config.Cfg.PageLimit
		} else if limit > 100 {
			limit = 100
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
