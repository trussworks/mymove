package ghcimport

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/testingsuite"
)

type GHCRateEngineImportSuite struct {
	testingsuite.PopTestSuite
	logger *zap.Logger
}

func (suite *GHCRateEngineImportSuite) SetupTest() {
	suite.DB().TruncateAll()
	suite.helperSetupStagingTables()
}

func (suite *GHCRateEngineImportSuite) TearDownSuite() {
	//suite.PopTestSuite.TearDown()
}

func (suite *GHCRateEngineImportSuite) helperSetupStagingTables() {
	path := filepath.Join("fixtures", "stage_ghc_pricing.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		suite.T().Fatal(ioErr)
	}

	sql := string(c)
	err := suite.DB().RawQuery(sql).Exec()
	if err != nil {
		suite.T().Fatal(err)
	}
}

func TestGHCRateEngineImportSuite(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panic(err)
	}

	hs := &GHCRateEngineImportSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       logger,
	}

	suite.Run(t, hs)
}

func (suite *GHCRateEngineImportSuite) TestGHCRateEngineImporter_Import() {
	tests := []struct {
		name    string
		gre     *GHCRateEngineImporter
		wantErr bool
	}{
		{
			name: "Run GHC Rate Engine Importer",
			gre: &GHCRateEngineImporter{
				Logger: suite.logger,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			if err := tt.gre.Import(suite.DB()); (err != nil) != tt.wantErr {
				t.Errorf("GHCRateEngineImporter.Import() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}