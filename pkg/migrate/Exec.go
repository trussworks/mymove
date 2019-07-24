package migrate

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

func Exec(inputReader io.Reader, tx *pop.Connection, wait time.Duration) error {

	in := NewBuffer()

	go ReadInSQL(inputReader, in, true, true, true) // read in lines as a separate thread

	lines := make(chan string, 1000)
	// read values out of the buffer
	go func() {
		formattedSQL := in.String()
		scanner := bufio.NewScanner(strings.NewReader(formattedSQL))
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	statements := make(chan string, 1000)
	go SplitStatements(lines, statements)
	for stmt := range statements {
		//if it is COPY statement then assume rest of the statements are part of copy and execute then as part of stdin
		match := copyStdinPattern.FindStringSubmatch(stmt)
		if match != nil {

			dataRows, err := extractCopyDataRows(statements)
			if err != nil {
				return errors.Wrapf(err, "error extracting data from")
			}

			// prepare statement so we can insert data rows
			stmt, err := prepareCopyFromStdin(match[4], parseColumns(match[6]), tx)
			if err != nil {
				return errors.Wrap(err, "error preparing copy from stdin statement")
			}
			for _, dataRow := range dataRows {
				values := lineToValues(dataRow)
				_, err = stmt.Exec(values...)
				if err != nil {
					return errors.Wrapf(err, "error executing copy from stdin with values %q", values)
				}
			}
			//exit loop after encountering the first COPY statement
			break
		}

		//if it is a regular statement then execute it as raw sql
		errExec := tx.RawQuery(stmt).Exec()
		if errExec != nil {
			return errors.Wrapf(errExec, "error executing statement: %q", stmt)
		}
	}
	return nil
}

// function to find and return data statements from COPY migration
func extractCopyDataRows(statements chan string) ([]string, error) {
	//create buffer of remaining statements
	var copyStmts []string
	var isCopy []string

	//range over the channel again
	for stmt := range statements {
		//only do this once, we are looking for first instance of copy
		//with an assumption that copy is always the last statement
		if isCopy == nil {
			isCopy = copyStdinPattern.FindStringSubmatch(stmt)
		}

		//skip all non copy statements
		if isCopy == nil {
			continue
		} else {
			// data statements end with newline instead of ; let's split them, by newline
			copyStmts = strings.Split(stmt, "\n")
			// remove trailing \. from slice
			if copyStmts[len(copyStmts)-1] == "\\." {
				//drop element from slice
				copyStmts = copyStmts[:len(copyStmts)-1]
			}
		}
	}
	return copyStmts, nil
}