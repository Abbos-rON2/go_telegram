package storage

import (
	"context"
	"time"

	"telegram_back/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *storage) CreateChat(ctx context.Context, chatType string, members []string, message models.CreateMessageRequest) (chatID string, err error) {
	doc := models.Message{
		Content:   message.Content,
		UserID:    message.UserID,
		CreatedAt: time.Now(),
	}
	members = append(members, message.UserID)

	result, err := s.db.Collection("chats").InsertOne(ctx,
		models.Chat{
			Type:      chatType,
			Members:   members,
			Messages:  []models.Message{doc},
			CreatedAt: time.Now().String(),
		})
	if err != nil {
		return
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	chatID = id.Hex()
	return
}

func (s *storage) CreateMessage(ctx context.Context, message models.CreateMessageRequest) error {
	obj, _ := primitive.ObjectIDFromHex(message.ChatID)
	doc := models.Message{
		Content:   message.Content,
		UserID:    message.UserID,
		CreatedAt: time.Now(),
	}

	_, err := s.db.Collection("chats").UpdateOne(ctx,
		bson.M{"_id": obj},
		bson.M{"$push": bson.M{"messages": doc}},
	)
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
