package leap_seconds

import (
	"reflect"
	"testing"
)

var sampleLeapSecondsData = LeapSecondsData{
	LastUpdate: 3803144275,
	ExpiresOn:  3833827200,
	LeapSeconds: []LeapSecond{
		{AddedOn: 2272060800, TotalCount: 10},
		{AddedOn: 2287785600, TotalCount: 11},
		{AddedOn: 2303683200, TotalCount: 12},
		{AddedOn: 2335219200, TotalCount: 13},
		{AddedOn: 2366755200, TotalCount: 14},
		{AddedOn: 2398291200, TotalCount: 15},
		{AddedOn: 2429913600, TotalCount: 16},
		{AddedOn: 2461449600, TotalCount: 17},
		{AddedOn: 2492985600, TotalCount: 18},
		{AddedOn: 2524521600, TotalCount: 19},
		{AddedOn: 2571782400, TotalCount: 20},
		{AddedOn: 2603318400, TotalCount: 21},
		{AddedOn: 2634854400, TotalCount: 22},
		{AddedOn: 2698012800, TotalCount: 23},
		{AddedOn: 2776982400, TotalCount: 24},
		{AddedOn: 2840140800, TotalCount: 25},
		{AddedOn: 2871676800, TotalCount: 26},
		{AddedOn: 2918937600, TotalCount: 27},
		{AddedOn: 2950473600, TotalCount: 28},
		{AddedOn: 2982009600, TotalCount: 29},
		{AddedOn: 3029443200, TotalCount: 30},
		{AddedOn: 3076704000, TotalCount: 31},
		{AddedOn: 3124137600, TotalCount: 32},
		{AddedOn: 3345062400, TotalCount: 33},
		{AddedOn: 3439756800, TotalCount: 34},
		{AddedOn: 3550089600, TotalCount: 35},
		{AddedOn: 3644697600, TotalCount: 36},
		{AddedOn: 3692217600, TotalCount: 37},
	},
}

func TestParseFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want LeapSecondsData
	}{
		{
			name: "included sample",
			args: args{path: "./sample/leap-seconds.list"},
			want: sampleLeapSecondsData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
