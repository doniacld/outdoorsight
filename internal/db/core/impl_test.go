package core

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type spotDetails struct {
	Name string
}

func TestMongoDB_Insert(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		collection  string
		doc         interface{}
		expectedErr string
	}{
		"nominal case": {
			collection:  "spot",
			doc:         spotDetails{Name: "LaBagarre"},
			expectedErr: "",
		},
	}

	mongoDB := NewDB()
	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			err := mongoDB.Insert(ctx, tc.collection, tc.doc)
			if tc.expectedErr != "" {
				assert.Contains(t, tc.expectedErr, err)
			}
		})
	}
}

// TODO DONIA add mocks ?
func TestMongoDB_Find(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		collection  string
		filter      map[string]interface{}
		expectedRes spotDetails
		expectedErr string
	}{
		"nominal case": {
			collection:  "spot",
			filter:      map[string]interface{}{"name": "LesSurplombs"},
			expectedRes: spotDetails{Name: "LesSurplombs"},
			expectedErr: "",
		},
	}

	mongoDB := NewDB()
	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			cursor, err := mongoDB.Find(ctx, tc.collection, tc.filter)
			if tc.expectedErr != "" {
				assert.Contains(t, tc.expectedErr, err)
			} else {
				var sd spotDetails
				fmt.Println(cursor.Decode(&sd))
				assert.Equal(t, tc.expectedRes, sd)
			}
		})
	}
}

func TestMongoDB_Update(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		collection  string
		filter      map[string]interface{}
		doc         bson.D
		expectedErr string
	}{
		"nominal case": {
			collection:  "spot",
			filter:      map[string]interface{}{"name": "LesSurplombs"},
			doc:         bson.D{{"$set", bson.D{{"name", "LesSurplombs"}}}},
			expectedErr: "",
		},
	}

	mongoDB := NewDB()
	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			err := mongoDB.Update(ctx, tc.collection, tc.filter, tc.doc)
			if tc.expectedErr != "" {
				assert.Contains(t, tc.expectedErr, err)
			}
		})
	}
}

func TestMongoDB_Delete(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		collection  string
		filter      map[string]interface{}
		expectedRes spotDetails
		expectedErr string
	}{
		"nominal case": {
			collection:  "spot",
			filter:      map[string]interface{}{"name": "LesSurplombs"},
			expectedErr: "",
		},
	}

	mongoDB := NewDB()
	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			err := mongoDB.Delete(ctx, tc.collection, tc.filter)
			if tc.expectedErr != "" {
				assert.Contains(t, tc.expectedErr, err)
			}
		})
	}
}
