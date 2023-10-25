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

func Operation(rows *sql.Rows) (*[]Operation, error) {
	defer rows.Close()
	items := make([]Operation, 0)
	i := 0

	for rows.Next() {
		i++
		operation := Operation{}
		err := rows.Scan(
			&operation.MaintenanceOrder,
			&operation.MaintenanceOrderOperation,
			&operation.MaintenanceOrderSubOperation,
			&operation.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &items, err
		}

		items = append(items, operation)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &items, nil
	}

	return &items, nil
}

func OperationComponent(rows *sql.Rows) (*[]OperationComponent, error) {
	defer rows.Close()
	items := make([]OperationComponent, 0)
	i := 0

	for rows.Next() {
		i++
		operationComponent := OperationComponent{}
		err := rows.Scan(
			&operationComponent.MaintenanceOrder,
			&operationComponent.MaintenanceOrderOperation,
			&operationComponent.MaintenanceOrderSubOperation,
			&operationComponent.MaintenanceOrderComponent,
			&operationComponent.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &items, err
		}

		items = append(items, operationComponent)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &items, nil
	}

	return &items, nil
}
