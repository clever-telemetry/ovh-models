package me

// GENERATED SDK for me API

// Domain tasks
type DomainTask struct {
	CanAccelerate bool                     `json:"canAccelerate,omitempty"`
	CanCancel     bool                     `json:"canCancel,omitempty"`
	CanRelaunch   bool                     `json:"canRelaunch,omitempty"`
	Comment       *string                  `json:"comment,omitempty"`
	CreationDate  string                   `json:"creationDate,omitempty"`
	Domain        string                   `json:"domain,omitempty"`
	DoneDate      *string                  `json:"doneDate,omitempty"`
	Function      NicOperationFunctionEnum `json:"function,omitempty"`
	Id            int64                    `json:"id,omitempty"`
	LastUpdate    string                   `json:"lastUpdate,omitempty"`
	Status        OperationStatusEnum      `json:"status,omitempty"`
	TodoDate      string                   `json:"todoDate,omitempty"`
}
