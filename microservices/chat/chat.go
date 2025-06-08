package chatapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	chat "goa-example/microservices/chat/gen/chat"
	security2 "goa-example/pkg/security"
	"io"
	"time"

	"goa.design/clue/log"
	"goa.design/goa/v3/security"
)

const (
	roomsKey       = "rooms"
	roomPublishKey = "room_publish"
	historyKey     = "history"
)

type chatsrvc struct {
	redis *redis.Client
}

func NewChat(client *redis.Client) chat.Service {
	return &chatsrvc{
		redis: client,
	}
}

func (s *chatsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims, err := security2.ValidToken(token)
	if err != nil {
		return ctx, err
	}

	return security2.HasPermission(ctx, claims, scheme)
}

func (s *chatsrvc) CreateRoom(ctx context.Context, p *chat.CreateRoomPayload) (res string, err error) {
	log.Printf(ctx, "chat.create-room")

	newRoomId := uuid.New().String()
	if err := s.redis.SAdd(ctx, roomsKey, newRoomId).Err(); err != nil {
		log.Printf(ctx, "redis SAdd err: %+v", err)
		return "", chat.Internal("Internal server error")
	}

	return newRoomId, nil
}

func (s *chatsrvc) History(ctx context.Context, p *chat.HistoryPayload) (res []*chat.Chat, err error) {
	log.Printf(ctx, "chat.history")

	histories, err := s.redis.LRange(ctx, fmt.Sprintf("%s:%s", historyKey, p.RoomID), 0, -1).Result()
	if err != nil {
		log.Printf(ctx, "redis LRange err: %+v", err)

		return nil, chat.Internal("Internal server error")
	}

	if len(histories) == 0 {
		return make([]*chat.Chat, 0), nil
	}

	for _, historyJSON := range histories {
		var chatMessage chat.Chat

		if err := json.Unmarshal([]byte(historyJSON), &chatMessage); err != nil {
			log.Printf(ctx, "redis Unmarshal err: %+v", err)

			return nil, chat.Internal("Internal server error")
		}

		res = append(res, &chatMessage)
	}

	return res, nil
}

func (s *chatsrvc) StreamRoom(ctx context.Context, p *chat.StreamRoomPayload, stream chat.StreamRoomServerStream) (err error) {
	log.Printf(ctx, "chat.stream-room")

	pubsubChannel := roomPublishKey + ":" + p.RoomID
	pubsub := s.redis.Subscribe(ctx, pubsubChannel)
	defer pubsub.Close()

	msgCh := make(chan string)
	errCh := make(chan error)
	go func() {
		for {
			str, err := stream.Recv()
			if err == io.EOF {
				errCh <- err

				close(msgCh)
				close(errCh)
				
				return
			}

			msgCh <- str
		}
	}()

	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			var newChat chat.Chat
			if err := json.Unmarshal([]byte(msg.Payload), &newChat); err != nil {
				log.Printf(ctx, "json.Unmarshal from pubsub err: %+v", err)
				continue
			}

			if err := stream.Send(&newChat); err != nil {
				log.Printf(ctx, "stream.Send err: %+v", err)

				errCh <- err
				return
			}
		}
	}()

	for done := false; !done; {
		select {
		case msg := <-msgCh:
			newChat := chat.Chat{
				Username: security2.ContextAuthInfo(ctx).Claims["sub"].(string),
				Message:  msg,
				SentAt:   time.Now().Unix(),
			}

			chatJSON, err := json.Marshal(newChat)
			if err != nil {
				log.Printf(ctx, "json.Marshal err: %+v", err)
				return chat.Internal("Internal server error")
			}

			if err := s.redis.LPush(ctx, fmt.Sprintf("%s:%s", historyKey, p.RoomID), chatJSON).Err(); err != nil {
				log.Printf(ctx, "redis LPush err: %+v", err)
				return chat.Internal("Internal server error")
			}

			if err := s.redis.Publish(ctx, pubsubChannel, chatJSON).Err(); err != nil {
				log.Printf(ctx, "redis Publish err: %+v", err)
				return chat.Internal("Internal server error")
			}

		case err := <-errCh:
			if err != nil {
				return chat.Internal("Internal server error")
			}
		case <-ctx.Done():
			done = true
		}
	}

	return stream.Close()
}
