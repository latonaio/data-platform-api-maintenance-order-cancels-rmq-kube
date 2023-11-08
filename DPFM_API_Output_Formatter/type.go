package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header             *Header             `json:"Header"`
	ObjectListItem     *[]ObjectListItem     `json:"ObjectListItem"`
	Operation          *[]Operation          `json:"Operation"`
	OperationComponent *[]OperationComponent `json:"OperationComponent"`
}

type Header struct {
	MaintenanceOrder string `json:"MaintenanceOrder"`
	IsCancelled      *bool  `json:"IsCancelled"`
}

type ObjectListItem struct {
	MaintenanceOrder           string `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList int    `json:"MaintenanceOrderObjectList"`
	MaintenanceObjectListItem  int    `json:"MaintenanceObjectListItem"`
	IsCancelled                *bool  `json:"IsCancelled"`
}

type Operation struct {
	MaintenanceOrder             string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation    string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation string `json:"MaintenanceOrderSubOperation"`
	IsCancelled                  *bool  `json:"IsCancelled"`
}

type OperationComponent struct {
	MaintenanceOrder             string `json:"MaintenanceOrder"`
	MaintenanceOrderOperation    string `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation string `json:"MaintenanceOrderSubOperation"`
	MaintenanceOrderComponent    string `json:"MaintenanceOrderComponent"`
	IsCancelled                  *bool  `json:"IsCancelled"`
}
