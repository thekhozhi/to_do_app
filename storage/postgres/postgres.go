package postgres

import (
	"context"
	"fmt"
	"strings"
	"to_do_app/config"
	"to_do_app/storage"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("error while parsing to postgres config!", err.Error())
		return nil, err
	}

	poolConfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil{
		fmt.Println("error while connecting to postgres db!", err.Error())
		return nil, err
	}

		//migration
		m, err := migrate.New("file://migrations/postgres", url)
		if err != nil {
			fmt.Println("ERROR IN NEW MIGRATE", err.Error())
			return Store{}, err
		}
	
		if err = m.Up(); err != nil {
			if !strings.Contains(err.Error(), "no change") {
				fmt.Println("entered", err)
				version, dirty, err := m.Version()
				if err != nil {
					fmt.Println("ERROR IN VERSION", err.Error())
					return Store{}, err
				}
	
				if dirty {
					version--
					if err = m.Force(int(version)); err != nil {
						fmt.Println("ERROR IN FORCE", err.Error())
						return Store{}, err
					}
				}
				fmt.Println("ERROR IN UP", err.Error())
				return Store{}, err
			}
		}

	return Store{
		Pool: pool,
	},nil
}


func (s Store) Close(){
	s.Pool.Close()
}

func (s Store) User() storage.IUserStorage {
	return NewUserRepo(s.Pool)
}

func (s Store) Task() storage.ITaskStorage {
	return NewTaskRepo(s.Pool)
}

func (s Store) TaskList() storage.ITaskListStorage {
	return NewTaskListRepo(s.Pool)
}

func (s Store) Label() storage.ILabelStorage {
	return NewLabelRepo(s.Pool)
}
