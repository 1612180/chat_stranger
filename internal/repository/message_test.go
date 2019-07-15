package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMessageGorm_FetchByTime(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	messageGorm := messageGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.User{}).Error; err != nil {
		t.Error(err)
	}

	if err := db.Create(&model.Message{
		RoomID: 1,
		UserID: 1,
	}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		roomID       int
		fromTime     time.Time
		wantMessages []*model.Message
		wantOK       bool
	}{
		{
			roomID:       0,
			fromTime:     time.Time{},
			wantMessages: []*model.Message{},
			wantOK:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("roomID=%d", tc.roomID), func(t *testing.T) {
			messages, ok := messageGorm.FetchByTime(tc.roomID, tc.fromTime)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantMessages, messages)
		})
	}
}

func TestMessageGorm_Create(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	messageGorm := messageGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		message *model.Message
		wantOK  bool
	}{
		{
			message: nil,
			wantOK:  false,
		},
		{
			message: &model.Message{},
			wantOK:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("message=%v", tc.message), func(t *testing.T) {
			ok := messageGorm.Create(tc.message)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestMessageGorm_Delete(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	messageGorm := messageGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Message{
		RoomID: 1,
	}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		roomID int
		wantOK bool
	}{
		{
			roomID: 1,
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("roomID=%d", tc.roomID), func(t *testing.T) {
			ok := messageGorm.Delete(tc.roomID)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}
