package setting

import (
	"subs/core/database"
)

type settingService struct {
}

func NewSettingService() *settingService {
	return &settingService{}
}

func (s *settingService) GetCatList(filter CatFilter) (int, *[]CatResponse, error) {
	db := database.RDB()
	query := NewSettingQuery(db)
	count, err := query.GetCatCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := query.GetCatList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}
