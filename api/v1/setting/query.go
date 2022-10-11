package setting

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type settingQuery struct {
	conn *sqlx.DB
}

func NewSettingQuery(connection *sqlx.DB) *settingQuery {
	return &settingQuery{
		conn: connection,
	}
}
func (r *settingQuery) GetCatCount(filter CatFilter) (int, error) {
	where, args := []string{"status > 0"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count
		FROM c_categories
		WHERE `+strings.Join(where, " AND "), args...)
	return count, err
}

func (r *settingQuery) GetCatList(filter CatFilter) (*[]CatResponse, error) {
	where, args := []string{"status > 0"}, []interface{}{}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageID*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var cats []CatResponse
	err := r.conn.Select(&cats, `
		SELECT id, name, status
		FROM c_categories
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	return &cats, err
}
