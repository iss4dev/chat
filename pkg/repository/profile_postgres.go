package repository

import (
	"fmt"
	"strings"
	todo "to-do-list"

	"github.com/jmoiron/sqlx"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (p *ProfilePostgres) GetUserProfile(id int) (todo.UserProfile, error) {
	var user todo.UserProfile
	query := fmt.Sprintf("SELECT avatar_url, bio, updated_at FROM %s WHERE user_id=$1", userProfilesTable)

	err := p.db.Get(&user, query, id)
	if err != nil {
		return todo.UserProfile{}, err
	}

	return user, nil
}
func (p *ProfilePostgres) UpdateUserProfile(userProfile todo.UserProfile, userID int) error {
	query := fmt.Sprintf("UPDATE %s SET ", userProfilesTable)
	args := []interface{}{}
	updates := []string{}
	argCount := 1

	if userProfile.AvatarURL != "" {
		updates = append(updates, fmt.Sprintf("avatar_url=$%d", argCount))
		argCount++
		args = append(args, userProfile.AvatarURL)
	}

	if userProfile.Bio != "" {
		updates = append(updates, fmt.Sprintf("bio=$%d", argCount))
		argCount++
		args = append(args, userProfile.Bio)
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query += fmt.Sprintf("%s WHERE user_id=$%d", strings.Join(updates, ", "), argCount)

	args = append(args, userID)

	_, err := p.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
