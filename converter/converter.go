package converter

import (
	"fmt"
	"math"
	"time"

	leap_seconds "github.com/flood4life/mars-time/leap-seconds"
)

const (
	secondsInMinute = 60
	minutesInHour   = 60
	hoursInDay      = 24
	secondsInDay    = secondsInMinute * minutesInHour * hoursInDay
)

type MarsTime struct {
	MarsSolDate            float64
	MartianCoordinatedTime string
}

type Converter struct {
	LeapSecondsData leap_seconds.LeapSecondsData
}

func (c Converter) EarthTimeToMarsTime(earth time.Time) MarsTime {
	msd := c.MSDFromUnixEpoch(earth.Unix())
	return MarsTime{
		MarsSolDate:            msd,
		MartianCoordinatedTime: MTCFromMSD(msd),
	}
}

func (c Converter) MSDFromUnixEpoch(epoch int64) float64 {
	// taken from https://en.wikipedia.org/wiki/Timekeeping_on_Mars#Formulas_to_compute_MSD_and_MTC
	return (float64(epoch)+c.LeapSecondsData.TAIUTCDiff(epoch))/88775.244147 + 34127.2954262
}

func MTCFromMSD(msd float64) string {
	fraction := math.Mod(msd, 1)
	if fraction < 0 {
		fraction = 1 + fraction
	}

	totalSeconds := int(math.Ceil(secondsInDay * fraction))
	fullSeconds := totalSeconds % secondsInMinute

	totalMinutes := totalSeconds / secondsInMinute
	fullMinutes := totalMinutes % minutesInHour

	fullHours := totalMinutes / minutesInHour

	return fmt.Sprintf("%02d:%02d:%02d", fullHours, fullMinutes, fullSeconds)
}
