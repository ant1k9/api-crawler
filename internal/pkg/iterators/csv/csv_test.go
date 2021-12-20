package csv

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
			payload: `id,other
123,abc
345,cde`,
			wantItems: []dto.Item{
				{
					ID:      123,
					Plugin:  "some-type",
					Payload: `{"id":"123","other":"abc"}`,
				},
				{
					ID:      345,
					Plugin:  "some-type",
					Payload: `{"id":"345","other":"cde"}`,
				},
			},
		},
		{
			name: "incorrect identificator path",
			payload: `idx,other
123,abc
345,cde`,
			wantErr: true,
		},
		{
			name:    "no items",
			payload: ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := New(
				config.Iterator{
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
