package gmap

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestContains(t *testing.T) {
	type args struct {
		m   map[string]string
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				m:   map[string]string{"key": "value"},
				key: "key",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Contains(tt.args.m, tt.args.key))
		})
	}
}

func TestContains2(t *testing.T) {
	type args struct {
		m   map[string]interface{}
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				m:   map[string]interface{}{"key": "value"},
				key: "key",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Contains(tt.args.m, tt.args.key))
		})
	}
}
