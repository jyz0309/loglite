package sqlparser

import (
	"fmt"
	"os"

	"github.com/jyz0309/loglite/executor"
	"github.com/xwb1989/sqlparser"
)

func Parse(sql string) error {
	statement, err := sqlparser.Parse(sql)
	if err != nil {
		return err
	}
	var e executor.Executor
	switch stmt := statement.(type) {
	case *sqlparser.Select:
		fmt.Println(stmt)
	case *sqlparser.Show:
		path, _ := os.Getwd()
		e = executor.BuildShowExecutor(stmt, path)
	default:
		return fmt.Errorf("unsupport query type")
	}
	rs, err := e.Exec()
	if err != nil {
		return err
	}

	for {
		set, done, err := rs.Next()
		if err != nil {
			return err
		}
		fmt.Println(set)
		if done {
			break
		}
	}
	return nil
}
