package requests

type OperationComponent struct {
	MaintenanceOrder             string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation    string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation string `json:"MaintenanceOrderSubOperation"`
	MaintenanceOrderComponent    string `json:"MaintenanceOrderComponent"`
	IsCancelled                  *bool  `json:"IsCancelled"`
}
