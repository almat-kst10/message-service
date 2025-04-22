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
	RoomList(ctx context.Context, profiles_id int) ([]models.RoomGeneralInfo, error)
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

func (r *MessageRepo) RoomList(ctx context.Context, profiles_id int) ([]models.RoomGeneralInfo, error) {
	query := `
		SELECT 
			r.id,
			r.title,
			rc.id,
			rc.profile_id,
			p.name,
			p.surname,
			rc.role_id,
			rl.title,
			rc.is_muted,
			rc.is_typing
		FROM room r
		JOIN room_client rc ON r.id = rc.room_id
		JOIN profiles p ON p.id = rc.profile_id
		JOIN room_role rl ON rc.role_id = rl.id
		WHERE rc.profile_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, profiles_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomList []models.RoomGeneralInfo
	for rows.Next() {
		var room models.RoomGeneralInfo
		err := rows.Scan(
			&room.RoomId,
			&room.RoomTitle,
			&room.ClientId,
			&room.ProfileId,
			&room.ProfileName,
			&room.ProfileSurname,
			&room.RoleId,
			&room.RoleName,
			&room.IsMuted,
			&room.IsTyping,
		)
		if err != nil {
			return nil, err
		}

		roomList = append(roomList, room)
	}

	return roomList, nil
}