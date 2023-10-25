package requests

type ObjectListItem struct {
	MaintenanceOrder           string `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList int    `json:"MaintenanceOrderObjectList"`
	MaintenanceObjectListItem  int    `json:"MaintenanceObjectListItem"`
	IsCancelled                *bool  `json:"IsCancelled"`
}
