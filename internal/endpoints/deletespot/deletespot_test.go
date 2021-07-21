package deletespot

import (
	"context"
	"testing"

	mock_db "github.com/doniacld/outdoorsight/internal/db/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteSpot(t *testing.T) {

	tt := map[string]struct {
		request     DeleteSpotRequest
		expectedRes DeleteSpotResponse
		expectedErr string
	}{
		"nominal case": {
			request: DeleteSpotRequest{
				SpotName: "LaBagarre",
			},
			expectedRes: DeleteSpotResponse{},
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
			odsDB.EXPECT().DeleteSpot(ctx, gomock.Any()).Return(int64(0), nil).Times(1)
			response, err := DeleteSpot(ctx, tc.request, odsDB)
			if tc.expectedErr != "" {
				require.NotNil(t, err)
				assert.Equal(t, tc.expectedErr, err.Message)
			} else {
				assert.Equal(t, tc.expectedRes, response)
			}
		})
	}
}
