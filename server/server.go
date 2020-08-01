package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/flood4life/mars-time/converter"
	leap_seconds "github.com/flood4life/mars-time/leap-seconds"
)

type MarsTimeServer struct {
	Converter converter.Converter
}

type MarsTimeRequest struct {
	Earth string `json:"earth"`
}

type MarsTimeResponse struct {
	MSD float64 `json:"msd"`
	MTC string  `json:"mtc"`
}

func NewServer(data leap_seconds.LeapSecondsData) MarsTimeServer {
	cnvtr := converter.Converter{LeapSecondsData: data}
	return MarsTimeServer{
		Converter: cnvtr,
	}
}

func (s MarsTimeServer) convert(request MarsTimeRequest) (MarsTimeResponse, error) {
	date, err := time.Parse(time.RFC3339, request.Earth)
	if err != nil {
		return MarsTimeResponse{}, err
	}
	marsTime := s.Converter.EarthTimeToMarsTime(date)
	return MarsTimeResponse{
		MSD: marsTime.MarsSolDate,
		MTC: marsTime.MartianCoordinatedTime,
	}, nil
}

func (s MarsTimeServer) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get(http.CanonicalHeaderKey("Content-Type")) != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var request MarsTimeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	response, err := s.convert(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`invalid request format`))
		return
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
