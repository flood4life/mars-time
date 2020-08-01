package leap_seconds

import (
	"testing"
	"time"
)

func TestLeapSecondsData_TAIUTCDiff(t *testing.T) {
	type args struct {
		atUnixEpoch int64
	}
	tests := []struct {
		name   string
		fields LeapSecondsData
		args   args
		want   float64
	}{
		// test cases are based on https://www.ietf.org/timezones/data/leap-seconds.list
		{
			name:   "Jan 1, 1972",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1972, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   42.184,
		},
		{
			name:   "Jul 1, 1972",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1972, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   43.184,
		},
		{
			name:   "Jan 1, 1973",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1973, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   44.184,
		},
		{
			name:   "Jan 1, 1974",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1974, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   45.184,
		},
		{
			name:   "Jan 1, 1975",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1975, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   46.184,
		},
		{
			name:   "Jan 1, 1976",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1976, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   47.184,
		},
		{
			name:   "Jan 1, 1977",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1977, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   48.184,
		},
		{
			name:   "Jan 1, 1978",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1978, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   49.184,
		},
		{
			name:   "Jan 1, 1979",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1979, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   50.184,
		},
		{
			name:   "Jan 1, 1980",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   51.184,
		},
		{
			name:   "Jul 1, 1981",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1981, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   52.184,
		},
		{
			name:   "Jul 1, 1982",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1982, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   53.184,
		},
		{
			name:   "Jul 1, 1983",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1983, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   54.184,
		},
		{
			name:   "Jul 1, 1985",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1985, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   55.184,
		},
		{
			name:   "Jan 1, 1988",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1988, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   56.184,
		},
		{
			name:   "Jan 1, 1990",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   57.184,
		},
		{
			name:   "Jan 1, 1991",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   58.184,
		},
		{
			name:   "Jul 1, 1992",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1992, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   59.184,
		},
		{
			name:   "Jul 1, 1993",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1993, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   60.184,
		},
		{
			name:   "Jul 1, 1994",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1994, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   61.184,
		},
		{
			name:   "Jan 1, 1996",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1996, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   62.184,
		},
		{
			name:   "Jul 1, 1997",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1997, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   63.184,
		},
		{
			name:   "Jan 1, 1999",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   64.184,
		},
		{
			name:   "Jan 1, 2006",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   65.184,
		},
		{
			name:   "Jan 1, 2009",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   66.184,
		},
		{
			name:   "Jul 1, 2012",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   67.184,
		},
		{
			name:   "Jul 1, 2015",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(2015, time.July, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   68.184,
		},
		{
			name:   "Jan 1, 2017",
			fields: sampleLeapSecondsData,
			args:   args{time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()},
			want:   69.184,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.TAIUTCDiff(tt.args.atUnixEpoch); got != tt.want {
				t.Errorf("TAIUTCDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
