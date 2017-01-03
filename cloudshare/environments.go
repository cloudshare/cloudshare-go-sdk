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

type Environments []Environment

type EnvironmentExtended struct {
	Vms               []VMAccessDetails `json:"vms"`
	Description       string            `json:"description"`
	BlueprintID       string            `json:"blueprintId"`
	BlueprintName     string            `json:"blueprintName"`
	PolicyID          string            `json:"policyId"`
	PolicyName        string            `json:"policyName"`
	ExpirationTime    string            `json:"expirationTime"`
	InvitationAllowed bool              `json:"invitationAllowed"`
	Organization      interface{}       `json:"organization"`
	OwnerEmail        string            `json:"ownerEmail"`
	ProjectID         string            `json:"projectId"`
	ProjectName       string            `json:"projectName"`
	SnapshotID        interface{}       `json:"snapshotId"`
	SnapshotName      interface{}       `json:"snapshotName"`
	StatusCode        int               `json:"statusCode"`
	StatusText        string            `json:"statusText"`
	RegionID          string            `json:"regionId"`
	Name              string            `json:"name"`
	ID                string            `json:"id"`
}

type VMAccessDetails struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	StatusText        string      `json:"statusText"`
	Progress          int         `json:"progress"`
	ImageID           string      `json:"imageId"`
	Os                string      `json:"os"`
	WebAccessURL      interface{} `json:"webAccessUrl"`
	Fqdn              string      `json:"fqdn"`
	ExternalAddress   string      `json:"externalAddress"`
	InternalAddresses []string    `json:"internalAddresses"`
	CPUCount          int         `json:"cpuCount"`
	DiskSizeGb        int         `json:"diskSizeGb"`
	MemorySizeMb      int         `json:"memorySizeMb"`
	Username          string      `json:"username"`
	Password          string      `json:"password"`
	ConsoleToken      string      `json:"consoleToken"`
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

type VM struct {
	Type         int         `json:"type"`
	Name         string      `json:"name"`
	Description  interface{} `json:"description"`
	TemplateVMID string      `json:"templateVmId"`
}

type EnvironmentTemplateRequest struct {
	Environment Environment `json:"environment"`
	ItemsCart   []VM        `json:"itemsCart"`
}

type CreateTemplateEnvResponse struct {
	Resources struct {
		CPUCount     int `json:"cpuCount"`
		DiskSizeMB   int `json:"diskSizeMB"`
		MemorySizeMB int `json:"memorySizeMB"`
	} `json:"resources"`
	Vms []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		OsTypeName  string `json:"osTypeName"`
		ImageURL    string `json:"imageUrl"`
		Resources   struct {
			CPUCount     int `json:"cpuCount"`
			DiskSizeMB   int `json:"diskSizeMB"`
			MemorySizeMB int `json:"memorySizeMB"`
		} `json:"resources"`
		DomainName              interface{}   `json:"domainName"`
		InternalIPs             []interface{} `json:"internalIPs"`
		MacAddresses            []interface{} `json:"macAddresses"`
		CanAddMultipleInstances bool          `json:"canAddMultipleInstances"`
		HostName                string        `json:"hostName"`
		VanityName              interface{}   `json:"vanityName"`
		HTTPAccessEnabled       bool          `json:"httpAccessEnabled"`
		StartWithHTTPS          bool          `json:"startWithHttps"`
		User                    string        `json:"user"`
		Password                string        `json:"password"`
		ID                      string        `json:"id"`
	} `json:"vms"`
	EnvironmentID string `json:"environmentId"`
}

/*
	GetTemplateParams allows you to filter templates by various criteria:

	projectID string (optional). "" means don't filter
	regionID string (optional). "" means don't filter
	templateType string (optional). "0" = bluebrint, "1" = VM
	skip int (default 0) - how many to skip.
	take int (default 0) - how many to return. 0 = return all.
*/
type GetTemplateParams struct {
	templateType string
	projectID    string
	regionID     string
	skip         int
	take         int
}
