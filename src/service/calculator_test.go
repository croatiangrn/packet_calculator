package service

import (
	"reflect"
	"testing"
)

func TestPack_CalculatePacks(t *testing.T) {
	type args struct {
		order    int
		packages []int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "Test with 200 packs",
			args: args{
				order:    200,
				packages: []int{250, 500, 1000, 2000, 5000},
			},
			want: map[int]int{
				250: 1,
			},
			wantErr: false,
		},
		{
			name: "Test with valid order and packs",
			args: args{
				order:    251,
				packages: []int{250, 500, 1000, 2000, 5000},
			},
			want: map[int]int{
				500: 1,
			},
			wantErr: false,
		},
		{
			name: "Test with edge case from email",
			args: args{
				order:    500000,
				packages: []int{23, 31, 53},
			},
			want: map[int]int{
				23: 2,
				31: 7,
				53: 9429,
			},
			wantErr: false,
		},
		{
			name: "Amount less than zero",
			args: args{
				order:    -2,
				packages: []int{23, 31, 53},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "No pack sizes provided",
			args: args{
				order:    200,
				packages: []int{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Calculator{}
			got, err := p.CalculatePacks(tt.args.order, tt.args.packages)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePacks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePacks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
