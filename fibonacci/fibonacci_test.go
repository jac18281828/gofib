package fibonacci

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCalculator_Fib(t *testing.T) {
	type fields struct {
		fibn map[int]*big.Int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *big.Int
	}{
		{
			name:   "ZeroCero",
			fields: fields{make(map[int]*big.Int)},
			args:   args{0},
			want:   big.NewInt(0),
		},
		{
			name:   "OneUne",
			fields: fields{make(map[int]*big.Int)},
			args:   args{1},
			want:   big.NewInt(1),
		},
		{
			name:   "SimpleExample13",
			fields: fields{make(map[int]*big.Int)},
			args:   args{13},
			want:   big.NewInt(233),
		},
		{
			name:   "SimpleExample84",
			fields: fields{make(map[int]*big.Int)},
			args:   args{84},
			want:   big.NewInt(160500643816367088),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := Calculator{
				fibn: tt.fields.fibn,
			}
			if got := calc.Fib(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculator.Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCalc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "CreateOne",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCalc(); got == nil {
				t.Errorf("NewCalc() = %v, want not nil", got)
			}
		})
	}
}

func Test_fib_fast(t *testing.T) {
	type args struct {
		n    int
		fibn map[int]*big.Int
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 *big.Int
	}{
		{
			name:  "Fast0",
			args:  args{0, make(map[int]*big.Int)},
			want:  big.NewInt(0),
			want1: big.NewInt(1),
		},
		{
			name:  "Fast1",
			args:  args{1, make(map[int]*big.Int)},
			want:  big.NewInt(1),
			want1: big.NewInt(1),
		},
		{
			name:  "Fast13",
			args:  args{13, make(map[int]*big.Int)},
			want:  big.NewInt(233),
			want1: big.NewInt(377),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fib_fast(tt.args.n, tt.args.fibn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fib_fast() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("fib_fast() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 *big.Int
	}{
		{
			name:  "fib_zero",
			args:  args{0},
			want:  big.NewInt(0),
			want1: big.NewInt(1),
		},
		{
			name:  "fib_one",
			args:  args{1},
			want:  big.NewInt(1),
			want1: big.NewInt(1),
		},
		{
			name:  "fib_exact100",
			args:  args{90},
			want:  big.NewInt(2880067194370816120),
			want1: big.NewInt(4660046610375530309),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fib(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fib() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("fib() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
