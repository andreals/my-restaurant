package repository

import (
	"my-restaurant/pkg/domains/restaurant/model"
	"reflect"
	"testing"
	"time"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name string
		want *RepositoryMemory
	}{
		{
			name: "success test",
			want: &RepositoryMemory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMemory_TakeOrder(t *testing.T) {
	type args struct {
		rest model.Restaurant
	}
	tests := []struct {
		name    string
		r       *RepositoryMemory
		args    args
		want    []model.Table
		wantErr bool
	}{
		{
			name: "success test",
			r:    &RepositoryMemory{},
			args: args{
				rest: model.Restaurant{
					Name:  "Test Rest",
					Chefs: 1,
					Menu: model.Menu{
						Dishes: []model.Dish{
							{
								Name:            "Test1",
								Price:           15.00,
								PreparationTime: time.Second * 5,
							},
							{
								Name:            "Test2",
								Price:           18.00,
								PreparationTime: time.Second * 3,
							},
						},
					},
				},
			},
			want:    []model.Table{{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepositoryMemory{}
			got, err := r.TakeOrder(tt.args.rest)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMemory.TakeOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositoryMemory.TakeOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryMemory_PrepareOrder(t *testing.T) {
	type args struct {
		table model.Table
		chefs int64
	}
	tests := []struct {
		name    string
		r       *RepositoryMemory
		args    args
		wantErr bool
	}{
		{
			name: "success test",
			r:    &RepositoryMemory{},
			args: args{
				table: model.Table{
					Number:    25,
					Customers: 1,
					Orders: []model.Order{
						{
							Customer: 1,
							Dishes: []model.Dish{
								{
									Name:            "Test1",
									Price:           15.00,
									PreparationTime: time.Second * 1,
								},
							},
						},
					},
				},
				chefs: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepositoryMemory{}
			if err := r.PrepareOrder(tt.args.table, tt.args.chefs); (err != nil) != tt.wantErr {
				t.Errorf("RepositoryMemory.PrepareOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generateBill(t *testing.T) {
	type args struct {
		table model.Table
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success test",
			args: args{
				table: model.Table{
					Number:    25,
					Customers: 1,
					Orders: []model.Order{
						{
							Customer: 1,
							Dishes: []model.Dish{
								{
									Name:            "Test1",
									Price:           15.00,
									PreparationTime: time.Second * 20,
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateBill(tt.args.table); (err != nil) != tt.wantErr {
				t.Errorf("generateBill() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
