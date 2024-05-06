package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/entity"
	"github.com/LeoAntunesBrombilla/readspacev2/internal/repository/interfaces"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user *entity.UserEntity) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err // handle the error appropriately
	}

	query := `INSERT INTO users (email, username, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	conn, err := u.db.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer conn.Release()

	row := conn.QueryRow(context.Background(), query, user.Email, user.Username, string(hashedPassword), user.CreatedAt)
	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UpdateUser(id *int64, user *entity.UserUpdateDetails) error {
	updateFields := []string{}
	args := []interface{}{}
	argCounter := 1

	if user.Email != "" {
		updateFields = append(updateFields, "email = $"+strconv.Itoa(argCounter))
		args = append(args, user.Email)
		argCounter++
	}

	if user.Username != "" {
		updateFields = append(updateFields, "username = $"+strconv.Itoa(argCounter))
		args = append(args, user.Username)
		argCounter++
	}

	args = append(args, id)

	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE id = $%d",
		strings.Join(updateFields, ", "),
		argCounter,
	)

	_, err := u.db.Exec(context.Background(), query, args...)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (u *userRepository) DeleteUserById(id *int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) ListAllUsers() ([]*entity.UserEntity, error) {
	query := `SELECT id, email, username, password, created_at FROM users`
	rows, err := u.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.UserEntity

	for rows.Next() {
		var user entity.UserEntity

		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) FindByUserName(username string) (*entity.UserEntity, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE username = $1`
	row := u.db.QueryRow(context.Background(), query, username)

	var user entity.UserEntity

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindPasswordById(id *int64) (*string, error) {
	query := `SELECT password FROM users WHERE id = $1`
	row := u.db.QueryRow(context.Background(), query, *id)

	var password string

	if err := row.Scan(&password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &password, nil
}

func (u *userRepository) UpdateUserPassword(id *int64, password string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := u.db.Exec(context.Background(), query, password, id)

	if err != nil {
		return err
	}

	return nil
}
