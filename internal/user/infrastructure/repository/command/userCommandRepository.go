package command

import (
	"context"
	"fmt"
	"persona/internal/user/domain/entity"
	"persona/libs/database/ent"
)

type UserCommandRepository interface {
	SaveTo(user *entity.User) error
	UpdateTo(user *entity.User) error
	DeleteTo(userId int) error
}

type repository struct {
	client *ent.Client
	ctx    context.Context
}

func NewCommandRepository(client *ent.Client) UserCommandRepository {
	return &repository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *repository) SaveTo(user *entity.User) error {
	_, err := r.client.UserSchema.
		Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(r.ctx)

	fmt.Println(err)
	return err
}

func (r *repository) UpdateTo(user *entity.User) error {
	//_, err := r.client.UserSchema.
	//	UpdateOneID(int(user.ID)).
	//	SetUsername(user.Username).
	//	SetPassword(user.Password).
	//	Save(r.ctx)

	return nil
}

func (r *repository) DeleteTo(userId int) error {
	err := r.client.UserSchema.
		DeleteOneID(userId).
		Exec(r.ctx)

	return err
}
