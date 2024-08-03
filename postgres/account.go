package postgres

import (
	"context"
	"database/sql"

	social "github.com/darmonlyone/my-social"
	"github.com/darmonlyone/my-social/postgres/boilentity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) social.AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r *accountRepo) Find(ctx context.Context, id string) (*social.Account, error) {
	boil, err := boilentity.Accounts(qm.Where("id=?", id)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, social.ErrNotFound
		}
		return nil, err
	}
	return &social.Account{
		ID:             boil.ID,
		Username:       boil.Username,
		HashedPassword: boil.HashedPassword,
		FirstName:      boil.Firstname,
		LastName:       boil.Lastname,
		CreatedAt:      boil.CreatedAt,
		UpdatedAt:      boil.UpdatedAt,
	}, nil
}

func (r *accountRepo) Store(ctx context.Context, account *social.Account) error {
	boilAccount := &boilentity.Account{
		ID:             account.ID,
		Username:       account.Username,
		HashedPassword: account.HashedPassword,
		Firstname:      account.FirstName,
		Lastname:       account.LastName,
	}
	return boilAccount.Insert(ctx, r.db, boil.Infer())
}

// FindByUsername implements social.AccountRepo.
func (r *accountRepo) FindByUsername(ctx context.Context, username string) (*social.Account, error) {
	boil, err := boilentity.Accounts(qm.Where("username=?", username)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, social.ErrNotFound
		}
		return nil, err
	}
	return &social.Account{
		ID:             boil.ID,
		Username:       boil.Username,
		HashedPassword: boil.HashedPassword,
		FirstName:      boil.Firstname,
		LastName:       boil.Lastname,
		CreatedAt:      boil.CreatedAt,
		UpdatedAt:      boil.UpdatedAt,
	}, nil
}
