package graph

import (
	"fmt"
	"gqlgen-app/ent"
)

func tryRollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("Failed to rollback: %s (rollbacked by %s)", rerr, err)
	}
	return err
}
