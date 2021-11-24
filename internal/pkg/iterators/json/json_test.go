package json

import (
	"reflect"
	"testing"

	"github.com/ant1k9/api-crawler/config"
	"github.com/ant1k9/api-crawler/internal/pkg/iterators/dto"
)

func Test_iterator_GetCollection(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		wantItems []dto.Item
		wantErr   bool
	}{
		{
			name: "get items",
			payload: `{"items": [
				{"id": "abc-123", "other": 200}
			]}`,
			wantItems: []dto.Item{
				{
					ID:      123,
					Plugin:  "some-type",
					Payload: `{"id": "abc-123", "other": 200}`,
				},
			},
		},
		{
			name:      "no items",
			payload:   `{"items": null}`,
			wantItems: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := New(
				config.Iterator{
					Regexp:            `\w+-(\d+)`,
					CollectionPath:    "items",
					IdentificatorPath: "id",
				},
				"some-type",
			)
			gotItems, err := i.GetCollection(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("iterator.GetCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("iterator.GetCollection() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}
