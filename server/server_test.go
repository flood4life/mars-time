package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flood4life/mars-time/converter"
)

func postRequestWithContentTypeHeader(body string) *http.Request {
	r := httptest.NewRequest("POST", "/path", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	return r
}

func bodyFromString(body string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(body))
}

func TestMarsTimeServer_ConvertHandler(t *testing.T) {
	type fields struct {
		Converter converter.Converter
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Response
	}{
		{
			name:   "GET results in 405",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/path", nil),
			},
			want: http.Response{
				StatusCode: 405,
				Body:       nil,
			},
		},
		{
			name:   "POST without application/json content-type results in 415",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/path", nil),
			},
			want: http.Response{
				StatusCode: 415,
				Body:       nil,
			},
		},
		{
			name:   "POST with empty body results in 400",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: postRequestWithContentTypeHeader(""),
			},
			want: http.Response{
				StatusCode: 400,
				Body:       nil,
			},
		},
		{
			name:   "POST with body without a date results in 400",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: postRequestWithContentTypeHeader(`{}`),
			},
			want: http.Response{
				StatusCode: 400,
				Body:       nil,
			},
		},
		{
			name:   "POST with a date in non-RFC3339 format results in 400",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: postRequestWithContentTypeHeader(`{"earth":"2020-01-01"}`),
			},
			want: http.Response{
				StatusCode: 400,
				Body:       bodyFromString("invalid request format"),
			},
		},
		{
			name:   "POST with a date in RFC3339 format results in 200",
			fields: fields{Converter: converter.Converter{}},
			args: args{
				w: httptest.NewRecorder(),
				r: postRequestWithContentTypeHeader(`{"earth":"2020-01-01T00:00:00Z"}`),
			},
			want: http.Response{
				StatusCode: 200,
				Body: bodyFromString(`{"msd":51900.68358788009,"mtc":"16:24:22"}
`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MarsTimeServer{
				Converter: tt.fields.Converter,
			}
			s.ConvertHandler(tt.args.w, tt.args.r)
			got := tt.args.w.Result()
			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("Code = %v, want %v", got.StatusCode, tt.want.StatusCode)
			}
			if tt.want.Body != nil {
				gotBody, _ := ioutil.ReadAll(got.Body)
				wantBody, _ := ioutil.ReadAll(tt.want.Body)
				if !bytes.Equal(gotBody, wantBody) {
					t.Errorf("Body = %q, want %q", string(gotBody), string(wantBody))
				}
			}
		})
	}
}
