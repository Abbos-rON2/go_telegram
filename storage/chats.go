package storage

import (
	"context"

	"telegram_back/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatI interface {
	CreateChat(ctx context.Context, chat models.CreateChatRequest) error
	GetChat(ctx context.Context, id string, chat *models.Chat) error
	GetAllChats(ctx context.Context, chats *[]models.Chat) error

	GetUserChats(ctx context.Context, userID string, chats *[]models.Chat) error
	GetChatUsers(ctx context.Context, chatID string, users *[]models.UserDTO) error
	AddUserToChat(ctx context.Context, chatID, userID string) error
	RemoveUserFromChat(ctx context.Context, chatID string, userID string) error
}

func (s *storage) GetChat(ctx context.Context, chatID string, chat *models.Chat) error {
	obj, _ := primitive.ObjectIDFromHex(chatID)
	result := s.db.Collection("chats").FindOne(ctx, bson.M{"_id": obj})
	result.Decode(chat)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (s *storage) GetAllChats(ctx context.Context, userID string, chats *[]models.Chat) error {
	cursor, err := s.db.Collection("chats").Find(ctx, map[string]interface{}{
		"type": "direct",
		"members": map[string]interface{}{
			"$in": []string{userID},
		},
	})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var chat models.Chat
		cursor.Decode(&chat)
		*chats = append(*chats, chat)
	}
	return cursor.Err()
}

// func (r *chatRepo) GetUserChats(ctx context.Context, userID string, chats *[]models.Chat) error {
// 	rows, err := r.db.Query(
// 		ctx,
// 		`SELECT id, user_id, content, created_at FROM chats WHERE user_id = $1`,
// 		userID,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var chat models.Chat
// 		if err := rows.Scan(
// 			&chat.ID,
// 			&chat.UserID,
// 			&chat.Content,
// 			&chat.CreatedAt,
// 		); err != nil {
// 			return err
// 		}
// 		*chats = append(*chats, chat)
// 	}
// 	return nil
// }

// func (r *chatRepo) AddUserToChat(ctx context.Context, chatID, userID string) error {
// 	_, err := r.db.Exec(
// 		ctx,
// 		`INSERT INTO chats_users (chat_id, user_id) VALUES ($1, $2)`,
// 		chatID,
// 		userID,
// 	)
// 	return err
// }
// func (r *storage) RemoveUserFromChat(ctx context.Context, userID, chatID string) error {
// 	_, err := r.db.Exec(
// 		ctx,
// 		`DELETE FROM chats_users WHERE chat_id = $1 AND user_id = $2`,
// 		chatID,
// 		userID,
// 	)
// 	return err
// }
// func (r *storage) GetChatUsers(ctx context.Context, chatID string, users *[]models.UserDTO) error {
// 	rows, err := r.db.Query(
// 		ctx,
// 		`SELECT users.id, users.name, users.phone, users.created_at FROM chats_users
// 			INNER JOIN users ON chats_users.user_id = users.id
// 			WHERE chats_users.chat_id = $1`,
// 		chatID,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var user models.UserDTO
// 		if err := rows.Scan(
// 			&user.ID,
// 			&user.Name,
// 			&user.Phone,
// 			&user.CreatedAt,
// 		); err != nil {
// 			return err
// 		}
// 		*users = append(*users, user)
// 	}
// 	return nil
// }
