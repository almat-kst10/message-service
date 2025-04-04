package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/almat-kst10/message-service/configs"
	"github.com/almat-kst10/message-service/internal/models"
	_ "github.com/lib/pq"
)

type IMessageRepo interface {
	SaveMessage(ctx context.Context, message *models.Message) (bool, error)
	GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error)
	ChatsList(ctx context.Context, profiles_id int) ([]models.Chat, error)
	Close()
}

type MessageRepo struct {
	db *sql.DB
}

func NewRepositoryMessage(cfg *configs.Configs) (IMessageRepo, error) {
	d := `
		host=%s 
		port=%s 
		user=%s 
		dbname=%s 
		password=%s 
		sslmode=%s
		client_encoding=%s
	`
	dsn := fmt.Sprintf(d, cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Name, cfg.DB.Psw, cfg.DB.SllMode, cfg.DB.Encoding)

	db, err := sql.Open(cfg.DB.Driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MessageRepo{db: db}, nil
}

func (r *MessageRepo) Close() {
	r.db.Close()
}

func (r *MessageRepo) ChatsList(ctx context.Context, profiles_id int) ([]models.Chat, error) {
	query := "SELECT id, profiles_id, user2_profile_id, last_message, is_read, count_new_msg, is_visible FROM chat_list WHERE profiles_id = $1 AND is_visible = true"
	rows, err := r.db.QueryContext(ctx, query, profiles_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatsList []models.Chat
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(
			&chat.Id,
			&chat.MyProfileId,
			&chat.User2ProfileId,
			&chat.LastMessage,
			&chat.IsRead,
			&chat.CountNewMsg,
			&chat.IsVisible,
		)

		if err != nil {
			return nil, err
		}

		chatsList = append(chatsList, chat)
	}
	
	return chatsList, nil
}

func (r *MessageRepo) SaveMessage(ctx context.Context, message *models.Message) (bool, error) {
	query := "INSERT INTO messages (sender_id, receiver_id, message, created_at) VALUES ($1, $2, $3, NOW())"
	_, err := r.db.Exec(query, message.SenderId, message.ReceiverId, message.Text)
	return err == nil, err
}

func (r *MessageRepo) GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, message, created_at
		FROM messages
		WHERE (sender_id=$1 AND receiver_id=$2) OR (sender_id=$2 AND receiver_id=$1)
		ORDER BY created_at
	`
	rows, err := r.db.QueryContext(ctx, query, user1Id, user2Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.Id, &msg.SenderId, &msg.ReceiverId, &msg.Text, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	
	return messages, nil
}