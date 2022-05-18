package storage

import (
	"context"
	"time"

	"telegram_back/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *storage) CreateChat(ctx context.Context, chatType string, members []string, message models.CreateMessageRequest) (chatID string, err error) {
	doc := models.Message{
		ID:        uuid.NewString(),
		Content:   message.Content,
		UserID:    message.UserID,
		CreatedAt: time.Now(),
	}

	result, err := s.db.Collection("chats").InsertOne(ctx,
		map[string]interface{}{
			"members":  members,
			"type":     chatType,
			"messages": []models.Message{doc},
		},
	)
	if err != nil {
		return
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	chatID = id.Hex()
	return
}

func (s *storage) CreateMessage(ctx context.Context, message models.CreateMessageRequest) error {
	doc := models.Message{
		Content:   message.Content,
		UserID:    message.UserID,
		CreatedAt: time.Now(),
	}
	_, err := s.db.Collection("chats").UpdateOne(ctx,
		map[string]interface{}{"id": message.ChatID},
		map[string]interface{}{
			"$push": map[string]interface{}{
				"messages": doc,
			},
		})
	return err
}

// func (r *messageRepo) GetMessage(ctx context.Context, id string, message *models.Message) error {
// 	return r.db.QueryRow(
// 		ctx,
// 		`SELECT id, chat_id, user_id, content, created_at FROM messages WHERE id = $1`,
// 		id,
// 	).Scan(
// 		&message.ID,
// 		&message.ChatID,
// 		&message.UserID,
// 		&message.Content,
// 		&message.CreatedAt,
// 	)
// }

// func (r *messageRepo) GetAllMessages(ctx context.Context, messages *[]models.Message) error {
// 	rows, err := r.db.Query(
// 		ctx,
// 		`SELECT id, chat_id, user_id, content, created_at FROM messages`,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var message models.Message
// 		if err := rows.Scan(
// 			&message.ID,
// 			&message.ChatID,
// 			&message.UserID,
// 			&message.Content,
// 			&message.CreatedAt,
// 		); err != nil {
// 			return err
// 		}
// 		*messages = append(*messages, message)
// 	}
// 	return nil
// }
