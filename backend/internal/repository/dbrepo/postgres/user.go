package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"readspacev2/internal/entity"
	"readspacev2/internal/repository"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user *entity.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err // handle the error appropriately
	}

	// Aqui a query e preparada para inserir os dados do usuarios
	query := `INSERT INTO users (email, username, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	conn, err := u.db.Acquire(context.Background())

	if err != nil {
		return err
	}

	// Ele vai ser executado logo apos essa funcao retornar
	// garantindo que fechamos a conexao
	defer conn.Release()

	// Executamos a query
	row := conn.QueryRow(context.Background(), query, user.Email, user.Username, string(hashedPassword), user.CreatedAt)
	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetByID(id int64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) ListAll() ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindByUserName(username string) (*entity.User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE username = $1`
	row := u.db.QueryRow(context.Background(), query, username)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
