package db

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/doniacld/outdoorsight/internal/db/core/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestOutdoorsightDB_AddSpot(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		spotD       SpotDetails
		expectedErr string
	}{
		"nominal case": {
			spotD:       SpotDetails{Name: "LaBagarre"},
			expectedErr: "",
		},
		"error case": {
			spotD:       SpotDetails{Name: "LaBagarre"},
			expectedErr: "unable to insert spot",
		},
	}
	fmt.Println(tt)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	osDB := mock_core.NewMockDB(ctrl)

	osDB.EXPECT().Insert(ctx, "spot", gomock.Any()).Return(errors.New("unable to insert spot")).AnyTimes()
	db := New()
	err := db.AddSpot(ctx, SpotDetails{})
	assert.Error(t, err)
	/*
		for n, tc := range tt {
			t.Run(n, func(t *testing.T) {
				if tc.expectedErr != "" {
					mongoDB.EXPECT().Insert(ctx, "spot", gomock.Any()).Return(errors.New("unable to insert spot")).AnyTimes()
					err := db.AddSpot(ctx, tc.spotD)
					assert.Error(t, err)
					//	assert.Contains(t, tc.expectedErr, err.Error())
				} else {
					mongoDB.EXPECT().Insert(ctx, "spot", tc.spotD).Return(nil).AnyTimes()
					err := db.AddSpot(ctx, tc.spotD)
					assert.NoError(t, err)
				}
			})
		}
	*/
}

const (
	spotLaBagarre = "LaBagarre"
)

// TODO DONIA return the right cursor
func TestOutdoorsightDB_FindSpot(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		filter      map[string]interface{}
		spotName    string
		expectedRes SpotDetails
		expectedErr string
	}{
		"nominal case": {
			filter:      map[string]interface{}{"name": spotLaBagarre},
			spotName:    spotLaBagarre,
			expectedRes: SpotDetails{Name: "LesSurplombs"},
			expectedErr: "",
		},
		"error case": {
			filter:      map[string]interface{}{"name": spotLaBagarre},
			spotName:    spotLaBagarre,
			expectedErr: "unable to find spot",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongoDB := mock_core.NewMockDB(ctrl)

	db := New()

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			if tc.expectedErr != "" {
				mongoDB.EXPECT().Find(ctx, "spot", tc.filter).Return(errors.New("unable to find spot"))
				_, err := db.GetSpot(ctx, tc.spotName)
				assert.Contains(t, tc.expectedErr, err.Error())
			} else {
				var cursor *mongo.Cursor
				mongoDB.EXPECT().Find(ctx, "spot", tc.filter).Return(cursor, nil).AnyTimes()
				res, _ := db.GetSpot(ctx, tc.spotName)
				assert.Equal(t, tc.expectedRes, res)
			}
		})
	}
}

func TestOutdoorsightDB_DeleteSpot(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		filter      map[string]interface{}
		spotName    string
		expectedErr string
	}{
		"nominal case": {
			filter:      map[string]interface{}{"name": spotLaBagarre},
			spotName:    spotLaBagarre,
			expectedErr: "",
		},
		"error case": {
			filter:      map[string]interface{}{"name": spotLaBagarre},
			spotName:    spotLaBagarre,
			expectedErr: "unable to delete spot",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongoDB := mock_core.NewMockDB(ctrl)

	db := New()

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			if tc.expectedErr != "" {
				mongoDB.EXPECT().Delete(ctx, "spot", tc.filter).Return(errors.New("unable to delete spot"))
				err := db.DeleteSpot(ctx, tc.spotName)
				assert.Contains(t, tc.expectedErr, err.Error())
			}
		})
	}
}

func TestOutdoorsightDB_UpdateSpot(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		spotName    string
		spotD       SpotDetails
		filter      map[string]interface{}
		doc         bson.D
		expectedErr string
	}{
		"nominal case": {
			spotName:    spotLaBagarre,
			spotD:       SpotDetails{Name: "LesSurplombs"},
			filter:      map[string]interface{}{"name": spotLaBagarre},
			doc:         bson.D{{"$set", bson.D{{"name", "LesSurplombs"}}}},
			expectedErr: "",
		},
		"error case": {
			spotName:    spotLaBagarre,
			spotD:       SpotDetails{Name: "LesSurplombs"},
			filter:      map[string]interface{}{"name": spotLaBagarre},
			doc:         bson.D{{"$set", bson.D{{"name", "LesSurplombs"}}}},
			expectedErr: "unable to update spot",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongoDB := mock_core.NewMockDB(ctrl)

	db := New()

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			if tc.expectedErr != "" {
				mongoDB.EXPECT().Update(ctx, "spot", tc.filter, tc.doc).Return(errors.New("unable to update spot"))
				err := db.UpdateSpot(ctx, tc.spotName, tc.spotD)
				assert.Contains(t, tc.expectedErr, err.Error())
			} else {
				mongoDB.EXPECT().Update(ctx, "spot", tc.filter, tc.doc).Return(nil)
				err := db.UpdateSpot(ctx, tc.spotName, tc.spotD)
				assert.NoError(t, err)
			}
		})
	}
}
