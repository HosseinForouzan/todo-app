package psql

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
    Username string `koanf:"username"`
    Password string `koanf:"password"`
    Port     int    `koanf:"port"`
    Host     string `koanf:"host"`
    DBName   string `koanf:"db_name"`
}

type PsqlDB struct {
    db *pgxpool.Pool
}

func NewPgxPool(ctx context.Context, config Config) (*PsqlDB, error) {
    connString := fmt.Sprintf(
        "postgres://%s:%s@%s:%d/%s",
        config.Username,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
    )

    poolConfig, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, fmt.Errorf("parse postgres config: %w", err)
    }

    poolConfig.MaxConns = 20
    poolConfig.MinConns = 5

    db, err := pgxpool.NewWithConfig(ctx, poolConfig)
    if err != nil {
        return nil, fmt.Errorf("create postgres pool: %w", err)
    }

    if err := db.Ping(ctx); err != nil {
        db.Close()
        return nil, fmt.Errorf("ping postgres: %w", err)
    }

    return &PsqlDB{
        db: db,
    }, nil
}

func (p *PsqlDB) Pool() *pgxpool.Pool {
    return p.db
}

func (p *PsqlDB) Close() {
    p.db.Close()
}