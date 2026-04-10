package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"tl/db"
)

func printJSON(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(os.Stdout, string(data))
	return err
}

func printError(err error) {
	data, _ := json.Marshal(map[string]string{"error": err.Error()})
	fmt.Fprintln(os.Stderr, string(data))
}

func taskOut(t *db.Task) map[string]any {
	out := map[string]any{
		"id":     t.ID,
		"title":  t.Title,
		"status": t.Status,
	}
	if t.Meta != nil {
		var meta map[string]any
		if json.Unmarshal([]byte(*t.Meta), &meta) == nil {
			for k, v := range meta {
				out[k] = v
			}
		}
	}
	return out
}

func printTask(t *db.Task) error {
	return printJSON(taskOut(t))
}

func printTasks(tasks []*db.Task) error {
	views := make([]map[string]any, len(tasks))
	for i, t := range tasks {
		views[i] = taskOut(t)
	}
	return printJSON(views)
}
