package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/transcom/mymove/cmd/webhook-client/utils"
	"github.com/transcom/mymove/pkg/testingsuite"

	"github.com/spf13/cobra"

	"github.com/transcom/mymove/pkg/logging"
)

// WebhookClientTestingSuite is a suite for testing the webhook client
type WebhookClientTestingSuite struct {
	testingsuite.PopTestSuite
	logger   utils.Logger
	certPath string
	keyPath  string
}

func TestWebhookClientTestingSuite(t *testing.T) {
	logger, _ := logging.Config("development", true)

	ts := &WebhookClientTestingSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage()),
		logger:       logger,
		certPath:     "../../config/tls/devlocal-mtls.cer",
		keyPath:      "../../config/tls/devlocal-mtls.key",
	}
	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use: "webhook-client [flags]",
	}
	initRootFlags(root.PersistentFlags())
	return root
}

func NewDbConnectionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "db-connection-test",
		RunE:         dbConnection,
		SilenceUsage: true,
	}
	initDbConnectionFlags(cmd.Flags())
	return cmd
}

func NewPostWebhookNotifyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "post-webhook-notify",
		RunE:         postWebhookNotify,
		SilenceUsage: true,
	}
	initPostWebhookNotifyFlags(cmd.Flags())
	return cmd
}

func (suite *WebhookClientTestingSuite) Test_DBConnection() {

	rootCmd := NewRootCmd()
	cmd := NewDbConnectionCommand()
	rootCmd.AddCommand(cmd)
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)

	suite.T().Run("db-connection: Success", func(t *testing.T) {
		rootCmd.SetArgs([]string{
			"db-connection-test",
		})
		err := rootCmd.Execute()
		suite.Nil(err)
	})

}

func (suite *WebhookClientTestingSuite) Test_PostWebhookNotify() {

	rootCmd := NewRootCmd()
	cmd := NewPostWebhookNotifyCommand()
	rootCmd.AddCommand(cmd)
	b := bytes.NewBufferString("")
	filename := "../../pkg/testdatagen/testdata/webhook_test_data.json"
	rootCmd.SetOut(b)
	suite.T().Run("post-webhook-notify: Send a request to server", func(t *testing.T) {
		rootCmd.SetArgs([]string{
			"--filename", filename,
			"--certpath", suite.certPath,
			"--keypath", suite.keyPath,
			"--insecure",
			"post-webhook-notify"})
		err := rootCmd.Execute()
		suite.Nil(err)
	})

}