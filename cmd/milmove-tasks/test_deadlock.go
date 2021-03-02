package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/cli"
	"github.com/transcom/mymove/pkg/logging"
)

// func checkTestDeadlockConfig(v *viper.Viper) error {
//
// 	if transactionName := v.GetString("transaction-name"); transactionName == "" {
// 		return errors.New("transaction-name is missing")
// 	}
//
// 	return nil
// }

func initTestDeadlockFlags(flag *pflag.FlagSet) {
	// Logging Levels
	cli.InitLoggingFlags(flag)

	// Don't sort flags
	flag.SortFlags = false
}

func forceDeadlock() {
	fmt.Println("vim-go")
	c := make(chan bool)
	fmt.Println("vim-go2")
	<-c
	fmt.Println("vim-go3")
}

func main1() {
	forceDeadlock()
}

// go run ./cmd/milmove-tasks test-deadlock
func testDeadlock(cmd *cobra.Command, args []string) error {
	// Create the logger
	v := viper.New()

	logger, err := logging.Config(
		logging.WithEnvironment(v.GetString(cli.LoggingEnvFlag)),
		logging.WithLoggingLevel(v.GetString(cli.LoggingLevelFlag)),
		logging.WithStacktraceLength(v.GetInt(cli.StacktraceLengthFlag)),
	)
	if err != nil {
		logger.Fatal("Failed to initialize Zap logging", zap.Error(err))
	}

	flag := pflag.CommandLine
	initTestDeadlockFlags(flag)
	err = flag.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal("invalid configuration", zap.Error(err))
	}

	pflagsErr := v.BindPFlags(flag)
	if pflagsErr != nil {
		logger.Fatal("invalid configuration", zap.Error(err))
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	// checkConfigErr := checkTestDeadlockConfig(v)
	// if checkConfigErr != nil {
	// 	logger.Fatal("invalid configuration", zap.Error(err))
	// }

	logger.Info(
		"force a deadlock")

	forceDeadlock()

	logger.Info(
		"somehow i got here after deadlock")

	return nil
}
