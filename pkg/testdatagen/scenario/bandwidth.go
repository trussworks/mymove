package scenario

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/models/roles"

	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
	"github.com/transcom/mymove/pkg/uploader"
)

// bandwidthScenario builds 1 move with 3 documents of varying sizes for each
// of the User Order uploads, and the Prime's Proof of Service docs
type bandwidthScenario NamedScenario

// BandwidthScenario is the thing
var BandwidthScenario = bandwidthScenario{"bandwidth"}

func createHHGMove(db *pop.Connection, userUploader *uploader.UserUploader, primeUploader *uploader.PrimeUploader) {
	serviceMember := makeServiceMember(db)
	orders := makeOrdersForServiceMember(serviceMember, db, userUploader)
	move := makeMoveForOrders(orders, db)
	shipment := makeShipmentForMove(move, db)
	makePaymentRequestForShipment(move, shipment, db, primeUploader)
}

func makeServiceMember(db *pop.Connection) models.ServiceMember {
	affiliation := models.AffiliationCOASTGUARD
	serviceMember := testdatagen.MakeExtendedServiceMember(db, testdatagen.Assertions{
		ServiceMember: models.ServiceMember{
			Affiliation: &affiliation,
		},
	})

	return serviceMember
}

func makeOrdersForServiceMember(serviceMember models.ServiceMember, db *pop.Connection, userUploader *uploader.UserUploader) models.Order {
	document := testdatagen.MakeDocument(db, testdatagen.Assertions{
		Document: models.Document{
			ServiceMemberID: serviceMember.ID,
			ServiceMember:   serviceMember,
		},
	})

	// Creates order upload documents from the files in this directory:
	// pkg/testdatagen/testdata/bandwidth_test_docs
	files := filesInBandwidthTestDirectory()

	for _, file := range files {
		filePath := fmt.Sprintf("bandwidth_test_docs/%s", file)
		fixture := testdatagen.Fixture(filePath)

		upload := testdatagen.MakeUserUpload(db, testdatagen.Assertions{
			File: fixture,
			UserUpload: models.UserUpload{
				UploaderID: serviceMember.UserID,
				DocumentID: &document.ID,
				Document:   document,
			},
			UserUploader: userUploader,
		})
		document.UserUploads = append(document.UserUploads, upload)
	}

	orders := testdatagen.MakeOrder(db, testdatagen.Assertions{
		Order: models.Order{
			ServiceMemberID:  serviceMember.ID,
			ServiceMember:    serviceMember,
			UploadedOrders:   document,
			UploadedOrdersID: document.ID,
		},
		UserUploader: userUploader,
	})

	return orders
}

func makeMoveForOrders(orders models.Order, db *pop.Connection) models.Move {
	hhgMoveType := models.SelectedMoveTypeHHG
	move := testdatagen.MakeMove(db, testdatagen.Assertions{
		Move: models.Move{
			Status:           models.MoveStatusSUBMITTED,
			OrdersID:         orders.ID,
			Orders:           orders,
			SelectedMoveType: &hhgMoveType,
		},
	})

	return move
}

func makeShipmentForMove(move models.Move, db *pop.Connection) models.MTOShipment {
	estimatedWeight := unit.Pound(1400)
	actualWeight := unit.Pound(2000)
	MTOShipment := testdatagen.MakeMTOShipment(db, testdatagen.Assertions{
		MTOShipment: models.MTOShipment{
			PrimeEstimatedWeight: &estimatedWeight,
			PrimeActualWeight:    &actualWeight,
			ShipmentType:         models.MTOShipmentTypeHHGLongHaulDom,
			ApprovedDate:         swag.Time(time.Now()),
			Status:               models.MTOShipmentStatusSubmitted,
		},
		Move: move,
	})

	testdatagen.MakeMTOAgent(db, testdatagen.Assertions{
		MTOAgent: models.MTOAgent{
			MTOShipment:   MTOShipment,
			MTOShipmentID: MTOShipment.ID,
			FirstName:     swag.String("Test"),
			LastName:      swag.String("Agent"),
			Email:         swag.String("test@test.email.com"),
			MTOAgentType:  models.MTOAgentReleasing,
		},
	})

	return MTOShipment
}

func makePaymentRequestForShipment(move models.Move, shipment models.MTOShipment, db *pop.Connection, primeUploader *uploader.PrimeUploader) {
	paymentRequest := testdatagen.MakePaymentRequest(db, testdatagen.Assertions{
		PaymentRequest: models.PaymentRequest{
			MoveTaskOrder: move,
			IsFinal:       false,
			Status:        models.PaymentRequestStatusPending,
		},
		Move: move,
	})

	dcrtCost := unit.Cents(99999)
	mtoServiceItemDCRT := testdatagen.MakeMTOServiceItem(db, testdatagen.Assertions{
		Move:        move,
		MTOShipment: shipment,
		ReService: models.ReService{
			ID: uuid.FromStringOrNil("68417bd7-4a9d-4472-941e-2ba6aeaf15f4"), // DCRT - Domestic crating
		},
	})

	testdatagen.MakePaymentServiceItem(db, testdatagen.Assertions{
		PaymentServiceItem: models.PaymentServiceItem{
			PriceCents: &dcrtCost,
		},
		PaymentRequest: paymentRequest,
		MTOServiceItem: mtoServiceItemDCRT,
	})

	ducrtCost := unit.Cents(99999)
	mtoServiceItemDUCRT := testdatagen.MakeMTOServiceItem(db, testdatagen.Assertions{
		Move:        move,
		MTOShipment: shipment,
		ReService: models.ReService{
			ID: uuid.FromStringOrNil("fc14935b-ebd3-4df3-940b-f30e71b6a56c"), // DUCRT - Domestic uncrating
		},
	})

	testdatagen.MakePaymentServiceItem(db, testdatagen.Assertions{
		PaymentServiceItem: models.PaymentServiceItem{
			PriceCents: &ducrtCost,
		},
		PaymentRequest: paymentRequest,
		MTOServiceItem: mtoServiceItemDUCRT,
	})

	files := filesInBandwidthTestDirectory()
	// Creates prime upload documents from the files in this directory:
	// pkg/testdatagen/testdata/bandwidth_test_docs
	for _, file := range files {
		filePath := fmt.Sprintf("bandwidth_test_docs/%s", file)
		fixture := testdatagen.Fixture(filePath)
		testdatagen.MakePrimeUpload(db, testdatagen.Assertions{
			File:           fixture,
			PaymentRequest: paymentRequest,
			PrimeUploader:  primeUploader,
		})
	}
}

func createOfficeUser(db *pop.Connection) {
	/* A user with both too and tio roles */
	email := "too_tio_role@office.mil"
	tooRole := roles.Role{}
	err := db.Where("role_type = $1", roles.RoleTypeTOO).First(&tooRole)
	if err != nil {
		log.Panic(fmt.Errorf("Failed to find RoleTypeTOO in the DB: %w", err))
	}

	tioRole := roles.Role{}
	err = db.Where("role_type = $1", roles.RoleTypeTIO).First(&tioRole)
	if err != nil {
		log.Panic(fmt.Errorf("Failed to find RoleTypeTIO in the DB: %w", err))
	}

	tooTioUUID := uuid.Must(uuid.FromString("9bda91d2-7a0c-4de1-ae02-b8cf8b4b858b"))
	loginGovUUID := uuid.Must(uuid.NewV4())
	testdatagen.MakeUser(db, testdatagen.Assertions{
		User: models.User{
			ID:            tooTioUUID,
			LoginGovUUID:  &loginGovUUID,
			LoginGovEmail: email,
			Active:        true,
			Roles:         []roles.Role{tooRole, tioRole},
		},
	})
	testdatagen.MakeOfficeUser(db, testdatagen.Assertions{
		OfficeUser: models.OfficeUser{
			ID:     uuid.FromStringOrNil("dce86235-53d3-43dd-8ee8-54212ae3078f"),
			Email:  email,
			Active: true,
			UserID: &tooTioUUID,
		},
	})
}

func filesInBandwidthTestDirectory() []string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(fmt.Errorf("failed to get current directory: %s", err))
	}

	dirName := path.Join(cwd, "pkg/testdatagen/testdata/bandwidth_test_docs")

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	var docs []string

	for _, file := range files {
		if file.Name() == ".DS_Store" {
			continue
		}
		docs = append(docs, file.Name())
	}

	return docs
}

// Run does that data load thing
func (e bandwidthScenario) Run(db *pop.Connection, userUploader *uploader.UserUploader, primeUploader *uploader.PrimeUploader) {
	createOfficeUser(db)
	createHHGMove(db, userUploader, primeUploader)
}