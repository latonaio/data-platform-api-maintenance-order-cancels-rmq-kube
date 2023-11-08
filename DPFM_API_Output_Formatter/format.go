package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*Header, error) {
	defer rows.Close()
	header := Header{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&header.MaintenanceOrder,
			&header.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil
}

func ConvertToObjectListItem(rows *sql.Rows) (*[]ObjectListItem, error) {
	defer rows.Close()
	items := make([]ObjectListItem, 0)
	i := 0

	for rows.Next() {
		i++
		objectListItem := ObjectListItem{}
		err := rows.Scan(
			&objectListItem.MaintenanceOrder,
			&objectListItem.MaintenanceOrderObjectList,
			&objectListItem.MaintenanceObjectListItem,
			&objectListItem.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &items, err
		}

		items = append(items, objectListItem)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &items, nil
	}

	return &items, nil
}
