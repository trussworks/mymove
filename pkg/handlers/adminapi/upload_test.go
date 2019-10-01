package adminapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/gofrs/uuid"

	uploadop "github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/upload"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/mocks"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/services/upload"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestGetUploadHandler() {
	sm := testdatagen.MakeDefaultServiceMember(suite.DB())
	suite.MustSave(&sm)

	orders := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
		Order: models.Order{
			ServiceMemberID: sm.ID,
			ServiceMember:   sm,
		},
	})
	suite.MustSave(&orders)

	move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
		Move: models.Move{
			Orders:   orders,
			OrdersID: orders.ID,
		},
	})
	suite.MustSave(&move)

	document := testdatagen.MakeDocument(suite.DB(), testdatagen.Assertions{
		Document: models.Document{
			ServiceMember:   sm,
			ServiceMemberID: sm.ID,
		},
	})
	suite.MustSave(&document)

	uploadID, _ := uuid.NewV4()
	uploadAssertions := models.Upload{
		ID:          uploadID,
		Document:    document,
		DocumentID:  &document.ID,
		Filename:    "FileName",
		Bytes:       int64(15),
		ContentType: "application/pdf",
		CreatedAt:   time.Now(),
		UploaderID:  sm.UserID,
	}

	uploadInstance := testdatagen.MakeUpload(suite.DB(), testdatagen.Assertions{Upload: uploadAssertions})
	suite.MustSave(&uploadInstance)

	requestUser := testdatagen.MakeDefaultUser(suite.DB())
	req := httptest.NewRequest("GET", fmt.Sprintf("/uploads/%s", uploadID.String()), nil)
	req = suite.AuthenticateAdminRequest(req, requestUser)

	queryBuilder := query.NewQueryBuilder(suite.DB())

	// test that everything is wired up
	suite.T().Run("integration test ok response", func(t *testing.T) {
		params := uploadop.GetUploadParams{
			HTTPRequest: req,
			UploadID:    *handlers.FmtUUID(uploadID),
		}

		handler := GetUploadHandler{
			HandlerContext: handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			NewQueryFilter: query.NewQueryFilter,
			UploadFetcher:  upload.NewUploadFetcher(queryBuilder),
		}

		response := handler.Handle(params)

		suite.IsType(&uploadop.GetUploadOK{}, response)
		okResponse := response.(*uploadop.GetUploadOK)
		suite.Equal(*handlers.FmtUUID(uploadID), okResponse.Payload.ID)
		suite.Equal(move.Locator, okResponse.Payload.MoveLocator)
	})

	queryFilter := mocks.QueryFilter{}
	newQueryFilter := newMockQueryFilterBuilder(&queryFilter)

	suite.T().Run("successful response", func(t *testing.T) {
		uploaded := models.Upload{ID: uploadID}
		params := uploadop.GetUploadParams{
			HTTPRequest: req,
			UploadID:    *handlers.FmtUUID(uploadID),
		}
		uploadFetcher := &mocks.UploadFetcher{}
		uploadFetcher.On("FetchUploads",
			mock.Anything,
			mock.Anything,
		).Return(uploaded, nil).Once()
		handler := GetUploadHandler{
			HandlerContext: handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			NewQueryFilter: newQueryFilter,
			UploadFetcher:  upload.NewUploadFetcher(queryBuilder),
		}

		response := handler.Handle(params)

		suite.IsType(&uploadop.GetUploadOK{}, response)
		okResponse := response.(*uploadop.GetUploadOK)
		suite.Equal(*handlers.FmtUUID(uploadID), okResponse.Payload.ID)
	})

	suite.T().Run("unsuccessful response when fetch fails", func(t *testing.T) {
		params := uploadop.GetUploadParams{
			HTTPRequest: req,
			UploadID:    *handlers.FmtUUID(uploadID),
		}
		expectedError := models.ErrFetchNotFound
		uploadFetcher := &mocks.UploadFetcher{}
		uploadFetcher.On("FetchUploads",
			mock.Anything,
			mock.Anything,
		).Return(nil, expectedError).Once()
		handler := GetUploadHandler{
			HandlerContext: handlers.NewHandlerContext(suite.DB(), suite.TestLogger()),
			NewQueryFilter: newQueryFilter,
			UploadFetcher:  uploadFetcher,
		}

		response := handler.Handle(params)

		expectedResponse := &handlers.ErrResponse{
			Code: http.StatusNotFound,
			Err:  expectedError,
		}
		suite.Equal(expectedResponse, response)
	})
}