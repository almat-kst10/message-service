package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/almat-kst10/message-service/internal/models"
)

type IRoomRepo interface {
	RoomList(ctx context.Context, profiles_id int) ([]*models.RoomGeneralInfo, error)
	RoomCreate(ctx context.Context, roomTitle string) (int, error)
	RoomDelete(ctx context.Context, roomId int) error
}

type RoomRepo struct {
	db *sql.DB
}

func NewRoomRepo(db *sql.DB) IRoomRepo {
	return &RoomRepo{
		db: db,
	}
}

func (r *RoomRepo) RoomList(ctx context.Context, profiles_id int) ([]*models.RoomGeneralInfo, error) {
	query := `
		SELECT 
			r.id,
			r.title,
			rc.id,
			rc.profile_id,
			rc.role_id,
			rl.title,
			rc.is_muted,
			rc.is_typing
		FROM room r
		JOIN client_room rc ON r.id = rc.room_id
		JOIN profiles p ON p.id = rc.profile_id
		JOIN room_role rl ON rc.role_id = rl.id
		WHERE rc.profile_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, profiles_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomList []*models.RoomGeneralInfo
	for rows.Next() {
		room := &models.RoomGeneralInfo{}
		err := rows.Scan(
			&room.RoomId,
			&room.RoomTitle,
			&room.ClientId,
			&room.ProfileId,
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

func (r *RoomRepo) RoomCreate(ctx context.Context, roomTitle string) (int, error) {
	query := "INSERT INTO room(title) VALUES($1) RETURNING id"

	var id int
	err := r.db.QueryRowContext(ctx, query, roomTitle).Scan(&id)
	// result, err := r.db.ExecContext(ctx, query, roomTitle)
	if err != nil {
		return 0, fmt.Errorf("error create room: %w", err)
	}

	return id, nil
}

func (r *RoomRepo) RoomDelete(ctx context.Context, roomId int) error {
	query := "DELETE FROM room WHERE id = $1"
	result, err := r.db.ExecContext(ctx, query, roomId)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("error delete room %s", err)
	}

	return nil
}
