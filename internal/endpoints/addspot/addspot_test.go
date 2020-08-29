package addspot
/*
import (
	"context"
	"github.com/doniacld/outdoorsight/internal/db"
	"github.com/doniacld/outdoorsight/internal/db/core"
	mock_core "github.com/doniacld/outdoorsight/internal/db/core/mocks"
	mock_db "github.com/doniacld/outdoorsight/internal/db/mocks"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/doniacld/outdoorsight/internal/spot"

	"github.com/stretchr/testify/assert"
)

type mockDB struct {
	*mock_core.MockDB
}

func TestAddSpot(t *testing.T) {
	ctx := context.Background()

	// NewClient creates the OutdoorsightDB structure
	func New() CoreDB {
		mongoDB := core.NewDB()
		osDB := OutdoorsightDB{mongoDB}
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_db.NewMockDB(ctrl)
	osDB := db.New(core.MongoDB{})

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

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			mockDB.EXPECT().Find(ctx, gomock.Any(), gomock.Any()).Return(spot.Details{}, nil).Times(1)
			mockDB.EXPECT().Insert(ctx, gomock.Any(), gomock.Any()).Return(gomock.Any(), nil).Times(1)
			mockDB.EXPECT().Find(ctx, gomock.Any(), gomock.Any()).Return(spot.Details{}, nil).Times(1)

			response, err := AddSpot(ctx, tc.request, osDB)

			if tc.expectedErr == "" {
				assert.Equal(t, tc.expectedErr, err.Message)
			} else {
				assert.Equal(t, tc.expectedRes, response)
			}
		})
	}
}
*/