package requests

type Operation struct {
	MaintenanceOrder             string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation    string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation string `json:"MaintenanceOrderSubOperation"`
	IsCancelled                  *bool  `json:"IsCancelled"`
}
