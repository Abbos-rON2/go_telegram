package storage

import (
	"context"
	"errors"
	"time"

	"telegram_back/errs"
	"telegram_back/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserI interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) error
	GetUser(ctx context.Context, id string, user *models.User) error
	GetUserByPhone(ctx context.Context, phone string, user *models.User) error
	GetAllUsers(ctx context.Context, users *[]models.User) error
	UpdateUserUsername(ctx context.Context, user models.User) error
	UpdateUserInfo(ctx context.Context, user models.User) error
	CreateChat(ctx context.Context, members []string, message models.CreateMessageRequest) error
}

func (s *storage) CreateUser(ctx context.Context, user models.CreateUserRequest) error {
	_, err := s.db.Collection("users").UpdateOne(ctx, bson.M{
		"phone": user.Phone,
	}, bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"phone":      user.Phone,
			"password":   user.Password,
			"created_at": time.Now(),
		},
	}, options.Update().SetUpsert(true))
	return err
}

func (s *storage) UpdateUserUsername(ctx context.Context, user models.User) error {
	resp := s.db.Collection("users").FindOne(ctx, bson.M{
		"username": user.Username,
	})
	if !errors.Is(mongo.ErrNoDocuments, resp.Err()) {
		return errs.ErrUsernameTaken
	}

	_, err := s.db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{"username": user.Username}})
	return err
}

func (s *storage) UpdateUserInfo(ctx context.Context, user models.User) error {
	_, err := s.db.Collection("users").UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"bio":        user.Bio,
			"updated_at": time.Now(),
		},
	})
	return err
}

func (s *storage) GetUser(ctx context.Context, id string, user *models.User) error {
	return s.db.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(user)
}

func (s *storage) GetUserByPhone(ctx context.Context, phone string, user *models.User) error {
	return s.db.Collection("users").FindOne(ctx, bson.M{"phone": phone}).Decode(user)
}

func (s *storage) GetAllUsers(ctx context.Context, users *[]models.User) error {
	result, err := s.db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	err = result.All(ctx, users)
	return err
}
