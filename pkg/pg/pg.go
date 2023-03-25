package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/vmoltaemcrkonrgcechd/online-store/config"
)

type PG struct {
	Sq sq.StatementBuilderType
	DB *sql.DB
}

func New(cfg *config.Config) (*PG, error) {
	db, err := sql.Open("postgres", cfg.PgURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PG{
		Sq: sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			RunWith(db),
		DB: db,
	}, nil
}
