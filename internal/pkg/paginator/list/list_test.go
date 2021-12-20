package list

import (
	"testing"

	"github.com/ant1k9/api-crawler/config"
)

func Test_paginator_NextPage(t *testing.T) {
	p := New(config.Paginator{
		Key: "<ID>",
		Items: []string{
			"first",
			"second",
			"third",
		},
	})

	tests := []struct {
		name    string
		payload string
		want    string
		adjust  string
		wantErr bool
	}{
		{
			name:    "first page",
			payload: "id=<ID>",
			adjust:  "plug_first",
			want:    "id=first",
		},
		{
			name:    "second page",
			payload: "id=<ID>",
			adjust:  "plug_second",
			want:    "id=second",
		},
		{
			name:    "last page",
			payload: "id=<ID>",
			adjust:  "plug_third",
			want:    "id=third",
		},
		{
			name:    "pagination finished",
			payload: "id=<ID>",
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
			if err != nil {
				return
			}

			if got != tt.want {
				t.Errorf("paginator.NextPage() = %v, want %v", got, tt.want)
				return
			}
			if p.AdjustPlugin("plug") != tt.adjust {
				t.Errorf(
					"paginator.AdjustPlugin() = %v, want %v",
					p.AdjustPlugin("plug"), tt.adjust,
				)
			}
		})
	}
}
