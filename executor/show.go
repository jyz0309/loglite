package executor

import (
	"fmt"
	"os"

	"github.com/jyz0309/loglite/types"
	"github.com/xwb1989/sqlparser"
)

type ShowExecutor struct {
	stmt *sqlparser.Show
	path string
}

func (e *ShowExecutor) Exec() (ResultSet, error) {
	ch := make(chan []types.Item, 10)
	result := BuildBaseResultSet([]string{"name", "is_dir", "size", "mod_time"}, ch)
	go func() {
		defer close(ch)
		switch e.stmt.Type {
		case "files":
			files, err := os.ReadDir(e.path)
			if err != nil {
				return
			}
			for _, file := range files {
				info, _ := file.Info()
				ch <- []types.Item{
					types.BuildStrItem(file.Name()),
					types.BuildStrItem(fmt.Sprintf("%v", info.IsDir())),
					types.BuildStrItem(fmt.Sprintf("%v", info.Size())),
					types.BuildStrItem(info.ModTime().String()),
				}
			}
		}
	}()
	return result, nil
}

func BuildShowExecutor(stmt *sqlparser.Show, path string) *ShowExecutor {
	return &ShowExecutor{
		stmt: stmt,
		path: path,
	}
}
