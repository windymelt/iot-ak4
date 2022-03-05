package iotak4

import (
	"bytes"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func Test_prepareUrl(t *testing.T) {
	type args struct {
		endpoint string
		query    map[string]string
	}
	os.Setenv("AK_COOP_ID", "TEST_COOP_ID")
	os.Setenv("AK_TOKEN", "TEST_TOKEN")
	stamps_url, _ := url.Parse("https://atnd.ak4.jp/api/cooperation/TEST_COOP_ID/stamps?token=TEST_TOKEN")
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
		// TODO: Add test cases.
		{
			name: "/stamps",
			args: args{
				endpoint: "/stamps",
				query:    map[string]string{},
			},
			want: stamps_url,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareUrl(tt.args.endpoint, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preparePunchBody(t *testing.T) {
	type args struct {
		punch_type int
		token      string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "punch_type=11",
			args: args{
				punch_type: 11,
				token:      "TEST_TOKEN",
			},
			want: bytes.NewBufferString(`{"token":"TEST_TOKEN","type":11}`).Bytes(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePunchBody(tt.args.punch_type, tt.args.token); !bytes.Equal(got, tt.want) {
				t.Errorf("preparePunchBody() = %v, want %v", bytes.NewBuffer(got).String(), bytes.NewBuffer(tt.want).String())
			}
		})
	}
}
