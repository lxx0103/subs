package auth

import (
	"database/sql"
)

type authRepository struct {
	tx *sql.Tx
}

func NewAuthRepository(transaction *sql.Tx) *authRepository {
	return &authRepository{
		tx: transaction,
	}
}

func (r *authRepository) CheckWxUserConfict(openID string) (bool, error) {
	var existed int
	row := r.tx.QueryRow("SELECT count(1) FROM u_wx_users WHERE open_id = ?", openID)
	err := row.Scan(&existed)
	if err != nil {
		return true, err
	}
	return existed != 0, nil
}

func (r *authRepository) CreateWxUser(info WxUser) error {
	_, err := r.tx.Exec(`
		INSERT INTO u_wx_users
		(
			open_id,
			status,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, info.OpenID, info.Status, info.Created, info.CreatedBy, info.Updated, info.UpdatedBy)
	return err
}
