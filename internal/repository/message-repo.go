package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/almat-kst10/message-service/configs"
	"github.com/almat-kst10/message-service/internal/models"
)

type IMessageRepo interface {
	SaveMessage(ctx context.Context, message *models.Message) (bool, error)
	GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error)
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

func (r *MessageRepo) SaveMessage(ctx context.Context, message *models.Message) (bool, error) {
	query := "INSERT INTO message (sender_id, receiver_id, text, created_at) VALUES ($1, $2, $3, NOW())"
	_, err := r.db.Exec(query, message.SenderId, message.ReceiverId, message.Text)
	return err == nil, err
}

func (r *MessageRepo) GetMessage(ctx context.Context, user1Id, user2Id int) ([]models.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, text, created
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