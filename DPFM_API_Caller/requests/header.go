package requests

type Header struct {
	MaintenanceOrder string `json:"MaintenanceOrder"`
	IsCancelled      *bool  `json:"IsCancelled"`
}
