package updatespot

import (
	"context"
	"testing"

	mock_db "github.com/doniacld/outdoorsight/internal/db/mocks"

	"github.com/doniacld/outdoorsight/internal/spot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateSpot(t *testing.T) {

	tt := map[string]struct {
		request     UpdateSpotRequest
		expectedRes UpdateSpotResponse
		expectedErr string
	}{
		"nominal case": {
			request: UpdateSpotRequest{
				Name:     "LaBagarre",
				Routes:   []spot.Route{{Name: "HenriLeChauve", Level: "7a+", Points: 5, Information: "easy one"}},
				Metadata: map[string]interface{}{"Access": "Park at the canoe spot"},
			},
			expectedRes: UpdateSpotResponse{
				Name:     "LaBagarre",
				Routes:   []spot.Route{{Name: "HenriLeChauve", Level: "7a+", Points: 5, Information: "easy one"}},
				Metadata: map[string]interface{}{"Access": "Park at the canoe spot"},
			},
			expectedErr: "",
		},
	}

	ctx := context.Background()

	// set mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	odsDB := mock_db.NewMockDB(ctrl)

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			// set the calls to the mock DB methods
			spotDetails := spot.Details{Name: tc.request.Name, Routes: tc.request.Routes, Metadata: tc.request.Metadata}
			odsDB.EXPECT().GetSpot(ctx, tc.request.Name).Return( &spotDetails, nil).Times(1)
			odsDB.EXPECT().UpdateSpot(ctx, tc.request.Name, gomock.Any()).Return(int64(1), int64(1), nil).AnyTimes()
			//sd := &spotDetails
			odsDB.EXPECT().GetSpot(ctx, tc.request.Name).Return(&spotDetails, nil).AnyTimes()

			// make the call
			response, err := UpdateSpot(ctx, tc.request, odsDB)
			if tc.expectedErr != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedErr, err.Message)
			} else {
				assert.Equal(t, tc.expectedRes, response)
			}
		})
	}
}
