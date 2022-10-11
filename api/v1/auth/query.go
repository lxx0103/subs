package auth

import (
	"github.com/jmoiron/sqlx"
)

type authQuery struct {
	conn *sqlx.DB
}

func NewAuthQuery(connection *sqlx.DB) *authQuery {
	return &authQuery{
		conn: connection,
	}
}

func (r *authQuery) GetWxUserByOpenID(openID string) (*WxUserResponse, error) {
	var user WxUserResponse
	err := r.conn.Get(&user, `
		SELECT id, open_id, name, avatar, status
		FROM u_wx_users
		WHERE open_id = ? AND status > 0
	`, openID)
	return &user, err
}
