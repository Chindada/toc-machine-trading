package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"tmt/internal/config"
	"tmt/internal/entity"
	"tmt/internal/usecase/modules/calendar"
	"tmt/internal/usecase/repo"
	"tmt/pkg/eventbus"
	"tmt/pkg/log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FcmUseCase struct {
	repo SystemRepo

	app      *firebase.App
	logger   *log.Log
	bus      *eventbus.Bus
	tradeDay *calendar.Calendar

	pushTokens     []string
	pushTokensLock sync.RWMutex
}

// NewFCM -.
func NewFCM() FCM {
	logger := log.Get()
	fb, err := newFCM()
	if err != nil {
		logger.Fatal(err)
	}

	cfg := config.Get()
	uc := &FcmUseCase{
		repo:     repo.NewSystemRepo(cfg.GetPostgresPool()),
		app:      fb,
		logger:   logger,
		bus:      eventbus.Get(),
		tradeDay: calendar.Get(),
	}

	uc.updatePushToken()

	uc.bus.SubscribeAsync(topicUpdatePushUser, true, uc.updatePushToken)
	uc.bus.SubscribeAsync(topicFetchStockHistory, true, uc.sendTargets)

	return uc
}

type srvAccount struct {
	ProjectID string `json:"project_id"`
}

func newFCM() (*firebase.App, error) {
	serviceAccountFilePath := "configs/service_account.json"
	opt := option.WithCredentialsFile(serviceAccountFilePath)

	data, err := os.ReadFile(serviceAccountFilePath)
	if err != nil {
		return nil, err
	}

	content := srvAccount{}
	if err = json.Unmarshal(data, &content); err != nil {
		return nil, err
	}

	config := &firebase.Config{ProjectID: content.ProjectID}
	fb, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}
	return fb, nil
}

func (uc *FcmUseCase) updatePushToken() {
	uc.pushTokensLock.Lock()
	defer uc.pushTokensLock.Unlock()

	tokens, err := uc.repo.GetAllPushTokens(context.Background())
	if err != nil {
		uc.logger.Error(err)
		return
	}
	uc.pushTokens = tokens
}

func (uc *FcmUseCase) getAllPushToken() []string {
	uc.pushTokensLock.RLock()
	defer uc.pushTokensLock.RUnlock()
	return uc.pushTokens
}

func (uc *FcmUseCase) sendTargets(targetArr []*entity.StockTarget) error {
	tokens := uc.getAllPushToken()
	if len(tokens) == 0 {
		return nil
	}

	ctx := context.Background()
	client, err := uc.app.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Found New Targets",
			Body:  fmt.Sprintf("%s has %d targets", uc.tradeDay.GetStockTradeDay().TradeDay.Format(entity.ShortTimeLayout), len(targetArr)),
		},
		Data: map[string]string{
			"type":              "new_targets",
			"new_targets_count": fmt.Sprintf("%d", len(targetArr)),
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "default",
				},
			},
		},
		Tokens: tokens,
	}

	_, err = client.SendEachForMulticast(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (uc *FcmUseCase) AnnounceMessage(msg string) error {
	ctx := context.Background()
	client, err := uc.app.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Announcement",
			Body:  msg,
		},
		Topic: "announcement",
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (uc *FcmUseCase) PushNotification(title, msg string) error {
	tokens := uc.getAllPushToken()
	if len(tokens) == 0 {
		return nil
	}

	ctx := context.Background()
	client, err := uc.app.Messaging(ctx)
	if err != nil {
		return err
	}

	badge := 1
	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  msg,
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "default",
					Badge: &badge,
				},
			},
		},
		Tokens: tokens,
	}

	_, err = client.SendEachForMulticast(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
