package addspot

import (
	"context"
	"testing"

	mock_db "github.com/doniacld/outdoorsight/internal/db/mocks"

	"github.com/doniacld/outdoorsight/internal/spot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddSpot(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		request     AddSpotRequest
		expectedRes AddSpotResponse
		expectedErr string
	}{
		"nominal case": {
			request: AddSpotRequest{
				Name:     "LaBagarre",
				Routes:   []spot.Route{{Name: "AlineLaMaline", Level: "5a", Points: 5, Information: "easy one"}},
				Metadata: map[string]interface{}{"Access": "Park at the canoe spot"},
			},
			expectedRes: AddSpotResponse{
				Name:     "LaBagarre",
				Routes:   []spot.Route{{Name: "AlineLaMaline", Level: "5a", Points: 5, Information: "easy one"}},
				Metadata: map[string]interface{}{"Access": "Park at the canoe spot"},
			},
			expectedErr: "",
		},
	}

	// set database mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	odsDB := mock_db.NewMockDB(ctrl)

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			// set mocks calls
			odsDB.EXPECT().GetSpot(ctx, gomock.Any()).Return(nil, nil).Times(1)
			odsDB.EXPECT().AddSpot(ctx, gomock.Any()).Return(nil, nil).Times(1)
			spotDetails := spot.Details(tc.request)
			odsDB.EXPECT().GetSpot(ctx, gomock.Any()).Return(&spotDetails, nil).Times(1)

			// make the call
			response, err := AddSpot(ctx, tc.request, odsDB)
			if tc.expectedErr != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedErr, err.Message)
			} else {
				assert.Equal(t, tc.expectedRes, response)
			}
		})
	}
}
