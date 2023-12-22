package query

import (
	"context"
	"persona/internal/user/domain/entity"
	"persona/libs/database/ent"
	"persona/libs/database/ent/userschema"

	"time"
)

type UserQueryRepository interface {
	GetTo(id int) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetBy(id int, condition string) (entity.User, error)
	GetListBy(condition string) ([]entity.User, error)
	GetAllList() ([]entity.User, error)
}

type repository struct {
	client *ent.Client
	ctx    context.Context
}

func NewQueryRepository(client *ent.Client) UserQueryRepository {
	return &repository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *repository) GetTo(id int) (entity.User, error) {
	userSchema, err := r.client.UserSchema.Get(r.ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		ID:        userSchema.ID,
		Username:  userSchema.Username,
		Password:  userSchema.Password,
		Email:     userSchema.Email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return user, nil
}

func (r *repository) GetByEmail(email string) (entity.User, error) {
	userSchema, err := r.client.UserSchema.
		Query().
		Where(userschema.Email(email)).
		Only(r.ctx)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		ID:        userSchema.ID,
		Username:  userSchema.Username,
		Password:  userSchema.Password,
		Email:     userSchema.Email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return user, nil
}

func (r *repository) GetBy(id int, condition string) (entity.User, error) {
	userSchema, err := r.client.UserSchema.Get(r.ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		ID:        userSchema.ID,
		UUID:      userSchema.UUID,
		Username:  userSchema.Username,
		Password:  userSchema.Password,
		Email:     userSchema.Email,
		CreatedAt: userSchema.CreatedAt,
		UpdatedAt: userSchema.UpdatedAt,
	}

	return user, nil
}

func (r *repository) GetListBy(condition string) ([]entity.User, error) {
	return []entity.User{}, nil
}

func (r *repository) GetAllList() ([]entity.User, error) {
	userSchemas, err := r.client.UserSchema.Query().All(r.ctx)
	if err != nil {
		return []entity.User{}, err
	}

	users := make([]entity.User, len(userSchemas))

	for i, userSchema := range userSchemas {
		users[i] = entity.User{
			ID:        userSchema.ID,
			Username:  userSchema.Username,
			Password:  userSchema.Password,
			Email:     userSchema.Email,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
	}

	return users, nil
}
