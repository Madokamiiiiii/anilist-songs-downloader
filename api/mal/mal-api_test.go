package mal

import (
	"testing"
)

func TestGetSongInformationFromMAL(t *testing.T) {
	type args struct {
		malId int
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestMALConnection", args{malId: 10721}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSongInformationFromMAL(tt.args.malId)
		})
	}
}
