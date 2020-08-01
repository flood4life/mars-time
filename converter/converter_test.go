package converter

import (
	"testing"
	"time"

	leap_seconds "github.com/flood4life/mars-time/leap-seconds"
)

var sampleLeapSecondsData = leap_seconds.LeapSecondsData{
	LastUpdate: 3803144275,
	ExpiresOn:  3833827200,
	LeapSeconds: []leap_seconds.LeapSecond{
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

func Test_mtcFromMSD(t *testing.T) {
	type args struct {
		msd float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "noon",
			args: args{0.5},
			want: "12:00:00",
		},
		{
			name: "noon for positive date",
			args: args{1.5},
			want: "12:00:00",
		},
		{
			name: "noon for negative date",
			args: args{1.5},
			want: "12:00:00",
		},
		{
			name: "6pm",
			args: args{0.75},
			want: "18:00:00",
		},
		{
			name: "6pm for positive date",
			args: args{1.75},
			want: "18:00:00",
		},
		{
			name: "6pm for negative date",
			args: args{-1.25},
			want: "18:00:00",
		},
		{
			name: "midnight",
			args: args{0},
			want: "00:00:00",
		},
		{
			name: "midnight for positive date",
			args: args{10},
			want: "00:00:00",
		},
		{
			name: "midnight for negative date",
			args: args{-10},
			want: "00:00:00",
		},
		{
			name: "random date from wiki",
			args: args{52095.46858},
			want: "11:14:46",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mtcFromMSD(tt.args.msd); got != tt.want {
				t.Errorf("mtcFromMSD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_EarthTimeToMarsTime(t *testing.T) {
	const delta = 0.0005
	type fields struct {
		LeapSecondsData leap_seconds.LeapSecondsData
	}
	type args struct {
		earth time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MarsTime
	}{
		{
			name:   "19 Jul 2020, 03:22:02 UTC",
			fields: fields{sampleLeapSecondsData},
			args:   args{time.Date(2020, time.July, 19, 3, 22, 2, 0, time.UTC)},
			want: MarsTime{
				MarsSolDate: 52095.46858,
			},
		},
		{
			name:   "30 Jul 2020, 13:54:35 UTC",
			fields: fields{sampleLeapSecondsData},
			args:   args{time.Date(2020, time.July, 30, 13, 54, 35, 0, time.UTC)},
			want: MarsTime{
				MarsSolDate: 52106.60179,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Converter{
				LeapSecondsData: tt.fields.LeapSecondsData,
			}
			got := c.EarthTimeToMarsTime(tt.args.earth)
			if diff := got.MarsSolDate - tt.want.MarsSolDate; diff > delta {
				t.Errorf("EarthTimeToMarsTime() = %v, want %v, delta %v", got.MarsSolDate, tt.want.MarsSolDate, diff)
			}
		})
	}
}
