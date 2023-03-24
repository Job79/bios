package store

import (
	"bios/store/model"
	"fmt"
)

// attachGeneric detects the ids of given objects based on the uid's
func attachGeneric[T model.Model](s Store, table string, models []T) ([]T, error) {
	uids := mapSlice(filter(models, func(m T) bool { return m.GetUID() != "" }), func(m T) any { return m.GetUID() })
	if len(uids) == 0 {
		return models, nil
	}

	rows, err := s.db.Query(s.db.Rebind("SELECT id, uid FROM "+table+" WHERE uid IN ("+inQuery(len(uids))+")"), uids...)
	if err != nil {
		return nil, err
	}

	idxMap := mapToIdx(models, func(m T) string { return m.GetUID() })
	for rows.Next() {
		id, uid := 0, ""
		if err := rows.Scan(&id, &uid); err != nil {
			return nil, err
		} else if idx, ok := idxMap[uid]; ok {
			models[idx].SetID(id)
		}
	}
	return models, nil
}

// deleteGeneric removes given models from the database
func deleteGeneric[T model.Model](s Store, table string, models []T) error {
	uids := mapSlice(models, func(m T) any { return m.GetUID() })
	_, err := s.db.Exec(s.db.Rebind("DELETE FROM "+table+" WHERE uid IN ("+inQuery(len(uids))+")"), uids...)
	return err
}

// fetchRelationsMap fetches a map of relations for the given ids as map[keyColumn][]valColumn
func fetchRelationsMap(s Store, table, keyCol, valCol string, ids []int) (map[int][]int, error) {
	// Fetch relations
	rows, err := s.db.Query(
		fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s IN (%s)", keyCol, valCol, table, keyCol, inQuery(len(ids))),
		ids,
	)
	if err != nil {
		return nil, err
	}

	// Create map[keyCol][]valCol
	relations := make(map[int][]int)
	for rows.Next() {
		var keyID, valID int
		if err := rows.Scan(&keyID, &valID); err != nil {
			return relations, err
		}
		relations[keyID] = append(relations[keyID], valID)
	}

	return relations, rows.Err()
}
