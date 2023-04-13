package common

import (
	"reflect"
	"testing"

	"github.com/bcaglaraydin/go-scoreboard/models"
	"github.com/go-redis/redis/v8"
)

func TestGetUserFromUserID(t *testing.T) {
	type args struct {
		rdb    *redis.Client
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserFromUserID(tt.args.rdb, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFromUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFromUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUserPointAndScore(t *testing.T) {
	type args struct {
		rdb  *redis.Client
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserPointAndScore(tt.args.rdb, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserPointAndScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
