package cloudshare

// Environment details
type Environment struct {
	ProjectID   string      `json:"projectId"`
	TeamID      interface{} `json:"teamId"`
	PolicyID    interface{} `json:"policyId"`
	Description interface{} `json:"description"`
	Status      string      `json:"status"`
	OwnerEmail  string      `json:"ownerEmail"`
	RegionID    string      `json:"regionId"`
	Name        string      `json:"name"`
	ID          string      `json:"id"`
}

// VMTemplate
type VMTemplate struct {
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	IsEnvironmentTemplate bool     `json:"isEnvironmentTemplate"`
	Type                  int      `json:"type"`
	ImageURL              string   `json:"imageUrl"`
	RegionID              string   `json:"regionId"`
	Tags                  []string `json:"tags"`
	Categories            []string `json:"categories"`
	Resources             struct {
		CPUCount     int `json:"cpuCount"`
		DiskSizeMB   int `json:"diskSizeMB"`
		MemorySizeMB int `json:"memorySizeMB"`
	} `json:"resources"`
	NumberOfMachines                       int         `json:"numberOfMachines"`
	HasMultipleVersions                    bool        `json:"hasMultipleVersions"`
	HasDefaultVersion                      bool        `json:"hasDefaultVersion"`
	DisabledForRegularEnvironmentCreation  interface{} `json:"disabledForRegularEnvironmentCreation"`
	DisabledForTrainingEnvironmentCreation interface{} `json:"disabledForTrainingEnvironmentCreation"`
	CanAddMultipleInstances                bool        `json:"canAddMultipleInstances"`
	EnvTemplateScope                       interface{} `json:"envTemplateScope"`
	CreationDate                           string      `json:"creationDate"`
	ID                                     string      `json:"id"`
}
