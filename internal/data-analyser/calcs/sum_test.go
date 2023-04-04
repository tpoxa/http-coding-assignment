package calcs

import (
	"testing"
)

func TestSumJsonNumbers(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "simplefloat1",
			args: args{
				data: []byte(`[1,0.5]`),
			},
			want:    1.5,
			wantErr: false,
		},
		{
			name: "simple10",
			args: args{
				data: []byte(`[1,2,3,4]`),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "simple10-2",
			args: args{
				data: []byte(`{"a":6,"b":4}`),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "simple2",
			args: args{
				data: []byte(`{"a":{"b":4},"c":-2}`),
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "simple2-2",
			args: args{
				data: []byte(`[[[2]]]`),
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "simple0",
			args: args{
				data: []byte(`{"a":[-1,1,"dark"]}`),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "simple0-2",
			args: args{
				data: []byte(`[-1,{"a":1, "b":"light"}]`),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "invalidjson1",
			args: args{
				data: []byte(`[`),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "threetimeszero",
			args: args{
				data: []byte(`[[-1,{"a":1, "b":"light"}],[-1,{"a":1, "b":"light"}],[-1,{"a":1, "b":"light"}]]`),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "nestedtwos",
			args: args{
				data: []byte(`{"a":{"b":4},"c":-2, "d": {"a":{"b":4},"c":-2}}`),
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumJsonNumbers(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calcJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalcJSON(b *testing.B) {
	data := []byte(``)
	for n := 0; n < b.N; n++ {
		_, _ = SumJsonNumbers(data)
	}
}
