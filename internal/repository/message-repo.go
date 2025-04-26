package repo

import (
	"context"
	"database/sql"
	"fmt"

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
	query := "SELECT id, room_id, profile_id, text, created_at FROM message_client_room WHERE room_id = $1 AND profile_id = $2"
	rows, err := r.db.QueryContext(ctx, query, message.RoomId, message.ProfileId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.MessageClientRoom
	for rows.Next() {
		message := models.MessageClientRoom{}

		err := rows.Scan(
			&message.Id,
			&message.RoomId,
			&message.ProfileId,
			&message.Text,
			&message.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}