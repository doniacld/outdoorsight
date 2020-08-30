package getspot

import (
	"context"
	"testing"

	"github.com/doniacld/outdoorsight/internal/db/mocks"

	"github.com/doniacld/outdoorsight/internal/spot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSpot(t *testing.T) {
	tt := map[string]struct {
		request     GetSpotRequest
		expectedRes GetSpotResponse
		expectedErr string
	}{
		"nominal case": {
			request: GetSpotRequest{
				SpotName:     "LaBagarre",
			},
			expectedRes: GetSpotResponse{
				Name:     "LaBagarre",
				Routes:   []spot.Route{{Name: "AlineLaMaline", Level: "5a", Points: 5, Information: "easy one"}},
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
			spotDetails := spot.Details(tc.expectedRes)
			odsDB.EXPECT().GetSpot(ctx, tc.request.SpotName).Return(&spotDetails, nil).Times(1)
			response, err := GetSpot(ctx, tc.request, odsDB)

			if tc.expectedErr != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedErr, err.Message)
			} else {
				assert.Equal(t, tc.expectedRes, response)
			}
		})
	}
}
