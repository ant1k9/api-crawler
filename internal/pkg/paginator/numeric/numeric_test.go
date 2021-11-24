package numeric

import (
	"testing"

	"github.com/ant1k9/api-crawler/config"
)

func Test_paginator_NextPage(t *testing.T) {
	p := New(config.Paginator{
		Start: 1,
		End:   3,
		Key:   "<PAGE>",
	})

	tests := []struct {
		name    string
		payload string
		want    string
		wantErr bool
	}{
		{
			name:    "first page",
			payload: "page=<PAGE>",
			want:    "page=1",
		},
		{
			name:    "second page",
			payload: "page=<PAGE>",
			want:    "page=2",
		},
		{
			name:    "last page",
			payload: "page=<PAGE>",
			want:    "page=3",
		},
		{
			name:    "pagination finished",
			payload: "page=<PAGE>",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.NextPage(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("paginator.NextPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("paginator.NextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
