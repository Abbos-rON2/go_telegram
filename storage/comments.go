package storage

// import (
// 	"context"

// 	"telegram_back/errs"
// 	"telegram_back/models"

// 	"github.com/jackc/pgx/v4"
// )

// type CommentI interface {
// 	CreateComment(ctx context.Context, comment models.CreateCommentRequest) error
// }

// func (r *storage) CreateComment(ctx context.Context, comment models.CreateCommentRequest) error {
// 	var replyID *int
// 	if comment.ReplyID != 0 {
// 		replyID = &comment.ReplyID
// 	}
// 	_, err := r.db.Exec(ctx, `
// 		INSERT INTO comments (
// 			post_id,
// 			author_id,
// 			content,
// 			reply_id
// 		) VALUES (
// 			$1,
// 			$2,
// 			$3,
// 			$4
// 		)
// 	`, comment.PostID, comment.AuthorID, comment.Content, replyID)
// 	if err != nil {
// 		return errs.Errf(errs.ErrDatabaseError, err)
// 	}
// 	return err
// }
