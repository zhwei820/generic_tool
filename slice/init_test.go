package slice

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSum11(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Sum(tt.args.numbers))
			assert.Equal(t, Avg(tt.args.numbers), 3)
			assert.Equal(t, Max(1, 2), 2)
			assert.Equal(t, Min(1, 2), 1)
		})
	}
}

func TestFloat11(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				numbers: []float64{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Sum(tt.args.numbers))

			assert.Equal(t, Avg(tt.args.numbers), 3.)
			assert.Equal(t, Max(1., 2.), 2.)
			assert.Equal(t, Min(2., 1.), 1.)
		})
	}
}

func TestMinSlice(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
	}{
		{
			name: "", args: args{a: []int{1, 2, 3}}, wantMin: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantMin, MinSlice(tt.args.a))
		})
	}
}

func TestMaxSlice(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			name: "", args: args{a: []int{1, 2, 3}}, wantMax: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantMax, MaxSlice(tt.args.a))
		})
	}
}
