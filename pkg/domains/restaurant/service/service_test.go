package service

import (
	"reflect"
	"testing"

	"my-restaurant/pkg/domains/restaurant/repository"
)

func TestNewService(t *testing.T) {
	repo := repository.NewRepository()
	type args struct {
		repository repository.RepositoryI
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				repository: repo,
			},
			want: &Service{
				repository: repo,
			},
			wantErr: false,
		},
		{
			name:    "fail test",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
