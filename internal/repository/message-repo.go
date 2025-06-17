package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/almat-kst10/message-service/internal/models"
)

type IMessageClientRepo interface {
	SetMessage(ctx context.Context, message models.MessageClientRoom) error
	GetMessage(ctx context.Context, message models.MessageClientRoom) ([]*models.MessageClientRoom, error)
}

type MessageClientRepo struct {
	db *sql.DB
}

func NewMessageClientRepo(db *sql.DB) IMessageClientRepo {
	return &MessageClientRepo{
		db: db,
	}
}

func (r *MessageClientRepo) SetMessage(ctx context.Context, message models.MessageClientRoom) error {
	query := "INSERT INTO message_client_room(room_id, profile_id, text) VALUES($1, $2, $3)"
	result, err := r.db.ExecContext(ctx, query, message.RoomId, message.ProfileId, message.Text)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("error update client-room %s", err)
	}

	return nil
}

func (r *MessageClientRepo) GetMessage(ctx context.Context, message models.MessageClientRoom) ([]*models.MessageClientRoom, error) {
	query := `
		SELECT 
			mcr.id, mcr.room_id, mcr.profile_id, mcr.text, mcr.created_at, p.name, p.surname
		FROM message_client_room mcr
		JOIN profiles p
		ON p.id = mcr.profile_id
		WHERE room_id = $1 AND profile_id != $2`
	rows, err := r.db.QueryContext(ctx, query, message.RoomId, message.ProfileId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.MessageClientRoom
	for rows.Next() {
		message := models.MessageClientRoom{}
		var name string
		var surname string

		err := rows.Scan(
			&message.Id,
			&message.RoomId,
			&message.ProfileId,
			&message.Text,
			&message.CreatedAt,
			&name,
			&surname,
		)

		message.FullName = fmt.Sprintf("%s %s", name, surname)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}
