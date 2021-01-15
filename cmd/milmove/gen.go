package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

const (
	// tempMigrationPath is the temporary path for generated migrations
	tempMigrationPath string = "./tmp"

	// localMigrationTemplate is the template for local secure migration files
	localMigrationTemplate string = `-- Local test migration.
-- This will be run on development environments.
-- It should mirror what you intend to apply on prod/staging/experimental
-- DO NOT include any sensitive data.
`
)

// Close an open file or exit
func closeFile(outfile *os.File) {
	err := outfile.Close()
	if err != nil {
		log.Printf("error closing %s: %v\n", outfile.Name(), err)
		os.Exit(1)
	}
}

func createMigration(path string, filename string, t *template.Template, templateData interface{}) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if mkdirErr := os.Mkdir(path, 0750); mkdirErr != nil {
			return errors.Wrapf(mkdirErr, "error creating path %q", path)
		}
	}
	migrationPath := filepath.Join(path, filename)
	migrationFile, err := os.Create(migrationPath)
	if err != nil {
		return errors.Wrapf(err, "error creating %s", migrationPath)
	}
	defer closeFile(migrationFile)
	err = t.Execute(migrationFile, templateData)
	if err != nil {
		log.Println("error executing template: ", err)
	}
	log.Printf("new migration file created at: %q\n", migrationPath)
	return nil
}

func addMigrationToManifest(migrationManifest string, filename string) error {
	mmf, err := os.OpenFile(filepath.Clean(migrationManifest), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "could not open migration manifest")
	}
	//RA Summary: gosec - errcheck - Unchecked return value
	//RA: Linter flags errcheck error: Ignoring a method's return value can cause the program to overlook unexpected states and conditions.
	//RA: Functions with unchecked return values in the file are used to end an asynchronous connection pertaining to file formatting
	//RA: Given the functions causing the lint errors are used to end a running asynchronous connection, it does not present a risk
	//RA Developer Status: Mitigated
	//RA Validator Status: {RA Accepted, Return to Developer, Known Issue, Mitigated, False Positive, Bad Practice}
	//RA Validator: jneuner@mitre.org
	//RA Modified Severity:
	defer mmf.Close() // nolint:errcheck

	_, err = mmf.WriteString(filename + "\n")
	if err != nil {
		return errors.Wrap(err, "could not append to the migration manifest")
	}

	log.Printf("new migration appended to manifest at: %q\n", migrationManifest)
	return nil
}

func writeEmptyFile(migrationPath, filename string) error {
	if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
		if mkdirErr := os.Mkdir(migrationPath, 0750); mkdirErr != nil {
			return errors.Wrapf(mkdirErr, "error creating path %q", migrationPath)
		}
	}
	path := filepath.Join(migrationPath, filename)

	err := ioutil.WriteFile(path, []byte{}, 0600)
	if err != nil {
		return errors.Wrap(err, "could not write new migration file")
	}

	log.Printf("new migration file created at: %q\n", path)
	return nil
}
