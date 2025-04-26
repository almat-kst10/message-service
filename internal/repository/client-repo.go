package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/almat-kst10/message-service/internal/models"
)

type IClientRoomRepo interface {
	ClientJoin(ctx context.Context, client *models.ClientRoom) error
	ClientDelete(ctx context.Context, client *models.ClientRoom) error
	ClientEdit(ctx context.Context, client *models.ClientRoom) error
}

type ClientRoomRepo struct {
	db *sql.DB
}

func NewClientRoomRepo(db *sql.DB) IClientRoomRepo {
	return &ClientRoomRepo{
		db: db,
	}
}

func (r *ClientRoomRepo) ClientJoin(ctx context.Context, client *models.ClientRoom) error {
	query := "INSERT INTO client_room(room_id, profile_id, role_id, is_muted, is_typing) VALUES($1, $2, $3, $4, $5)"
	result, err := r.db.ExecContext(ctx, query, client.RoomId, client.ProfileId, client.RoleId, client.IsMuted, client.IsTyping)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("error join room %s", err)
	}

	return nil
}

func (r *ClientRoomRepo) ClientDelete(ctx context.Context, client *models.ClientRoom) error {
	query := "DELETE FROM client_room WHERE room_id = $1 AND profile_id = $2"
	result, err := r.db.ExecContext(ctx, query, client.RoomId, client.ProfileId)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("error delete client room %s", err)
	}

	return nil
}

func (r *ClientRoomRepo) ClientEdit(ctx context.Context, client *models.ClientRoom) error {
	var query strings.Builder
	var args []interface{}
	argIndex := 1

	query.WriteString("UPDATE client_room SET ")
	if client.RoleId > 0 {
		query.WriteString(fmt.Sprintf("role_id=$%d, ", argIndex))
		argIndex++
	}

	if client.IsMuted {
		query.WriteString(fmt.Sprintf("is_muted=$%d, ", argIndex))
		argIndex++
	}

	if client.IsTyping {
		query.WriteString(fmt.Sprintf("is_typing=$%d", argIndex))
		argIndex++
	}

	queryRequest := strings.TrimSuffix(query.String(), ", ")
	queryRequest += fmt.Sprintf(" WHERE profile_id=$%d", argIndex)
	args = append(args, client.ProfileId)

	result, err := r.db.ExecContext(ctx, queryRequest, args...)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("error update client-room %s", err)
	}

	return nil
}
