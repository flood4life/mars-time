package leap_seconds

import "math"

const (
	unixEpochNTPOffset    = 2208988800
	secondsInDay          = 86400
	terrestrialTimeOffset = 32.184
)

type LeapSecondsData struct {
	LastUpdate  int64
	ExpiresOn   int64
	LeapSeconds []LeapSecond
}

type LeapSecond struct {
	AddedOn    int64
	TotalCount int
}

func (d LeapSecondsData) TAIUTCDiff(atUnixEpoch int64) float64 {
	floatUnix := float64(atUnixEpoch)
	if len(d.LeapSeconds) == 0 {
		return empiricalDifference(floatUnix)
	}
	ntp := ntpEpochFromUnix(atUnixEpoch)
	first := d.LeapSeconds[0]
	// if there's no accurate data in table, use the empirical formula
	if ntp < first.AddedOn || ntp > d.ExpiresOn {
		return empiricalDifference(floatUnix)
	}

	leapSecond := d.earliestFittingLeapSecond(ntp)
	if leapSecond == nil {
		return empiricalDifference(floatUnix)
	}
	return float64(leapSecond.TotalCount) + terrestrialTimeOffset
}

func (d LeapSecondsData) earliestFittingLeapSecond(ntpEpoch int64) *LeapSecond {
	for i := 0; i < len(d.LeapSeconds); i++ {
		if d.LeapSeconds[i].AddedOn >= ntpEpoch {
			return &d.LeapSeconds[i]
		}
	}
	return nil
}

func ntpEpochFromUnix(unixEpoch int64) int64 {
	return unixEpoch + unixEpochNTPOffset
}

func empiricalDifference(epoch float64) float64 {
	// T = (JDut - 2451545.0) / 36525
	// TT - UTC = 64.184s + 59 s × T - 51.2 s × T^2 - 67.1 s × T^3 - 16.4 s × T^4
	// taken from https://www.giss.nasa.gov/tools/mars24/help/algorithm.html, A-3, A-4
	t := (EpochToJulianDate(epoch) - 2451545.0) / 36525

	return 64.184 + 59*t - 51.2*math.Pow(t, 2) - 67.1*math.Pow(t, 3) - 16.4*math.Pow(t, 4)
}

func EpochToJulianDate(unixEpoch float64) float64 {
	return 2440587.5 + (unixEpoch / secondsInDay)
}
