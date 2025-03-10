package app

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	mock_storage "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/mock"
)

type FeedServiceTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	t           *testing.T
	mockStorage *mock_storage.MockStorage
	feedService *FeedService
}

func NewFeedServiceTestSuite(t *testing.T) *FeedServiceTestSuite {
	return &FeedServiceTestSuite{t: t}
}

func (suite *FeedServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.t)
	suite.mockStorage = mock_storage.NewMockStorage(suite.ctrl)
	logger := logrus.New()
	suite.feedService = NewFeedService(logger, suite.mockStorage)
}

func (suite *FeedServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *FeedServiceTestSuite) TestGetAllFeed() {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          []domain.Feed
		expectedResponse *responses.GetAllFeed
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetAllFeed{
				Feed: []responses.Feed{
					{
						ID:          1,
						Title:       "testAll",
						Description: "testAbout",
					},
				},
			},
			request: []domain.Feed{
				{
					ID:          1,
					Title:       "testAll",
					Description: "testAbout",
				}},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          []domain.Feed{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetAllFeed(ctx).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.feedService.GetAllFeed(ctx)

		// Compare the expected and actual responses
		assert.Equal(suite.T(), test.expectedError, actualError)
		assert.True(suite.T(), reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *FeedServiceTestSuite) TestGetFeed() {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          domain.Feed
		expectedResponse *responses.GetFeed
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetFeed{
				ID:              1,
				Title:           "testAll",
				RegistrationURL: "testURL",
				Description:     "testAbout",
			},
			request: domain.Feed{
				ID:    1,
				Title: "testAll",

				RegistrationURL: "testURL",
				Description:     "testAbout",
			},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          domain.Feed{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetFeed(ctx, 1).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.feedService.GetFeed(ctx, 1)

		// Compare the expected and actual responses
		assert.Equal(suite.T(), test.expectedError, actualError)
		assert.True(suite.T(), reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func TestFeedService_GetAllFeed(t *testing.T) {
	suite.Run(t, NewFeedServiceTestSuite(t))
}
