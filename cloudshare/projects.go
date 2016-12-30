package cloudshare

// Blueprint available for project.
type Blueprint struct {
	ID                    string      `json:"id"`
	Name                  string      `json:"name"`
	Description           string      `json:"description"`
	IsEnvironmentTemplate bool        `json:"isEnvironmentTemplate"`
	Type                  int         `json:"type"`
	ImageURL              string      `json:"imageUrl"`
	RegionID              string      `json:"regionId"`
	Tags                  interface{} `json:"tags"`
	Categories            interface{} `json:"categories"`
	Resources             struct {
		CPUCount     int `json:"cpuCount"`
		DiskSizeMB   int `json:"diskSizeMB"`
		MemorySizeMB int `json:"memorySizeMB"`
	} `json:"resources"`
	NumberOfMachines                       int         `json:"numberOfMachines"`
	HasMultipleVersions                    bool        `json:"hasMultipleVersions"`
	HasDefaultVersion                      bool        `json:"hasDefaultVersion"`
	DisabledForRegularEnvironmentCreation  bool        `json:"disabledForRegularEnvironmentCreation"`
	DisabledForTrainingEnvironmentCreation bool        `json:"disabledForTrainingEnvironmentCreation"`
	CanAddMultipleInstances                bool        `json:"canAddMultipleInstances"`
	EnvTemplateScope                       interface{} `json:"envTemplateScope"`
	CreationDate                           string      `json:"creationDate"`
}

// BlueprintDetails holds blueprint information including snapshots (createFromVersions).
type BlueprintDetails struct {
	CreateFromVersions []struct {
		Machines []struct {
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
			User                    interface{}   `json:"user"`
			Password                interface{}   `json:"password"`
			ID                      string        `json:"id"`
		} `json:"machines"`
		AuthorName string      `json:"authorName"`
		Comment    interface{} `json:"comment"`
		Type       int         `json:"type"`
		Name       string      `json:"name"`
		IsDefault  bool        `json:"isDefault"`
		IsLatest   bool        `json:"isLatest"`
		Number     int         `json:"number"`
		Resources  struct {
			CPUCount     int `json:"cpuCount"`
			DiskSizeMB   int `json:"diskSizeMB"`
			MemorySizeMB int `json:"memorySizeMB"`
		} `json:"resources"`
		CreateTime  string      `json:"createTime"`
		Description interface{} `json:"description"`
		ImageURL    interface{} `json:"imageUrl"`
		Regions     []string    `json:"regions"`
		ID          string      `json:"id"`
	} `json:"createFromVersions"`
	Description           interface{} `json:"description"`
	IsEnvironmentTemplate bool        `json:"isEnvironmentTemplate"`
	Type                  int         `json:"type"`
	ImageURL              string      `json:"imageUrl"`
	Tags                  interface{} `json:"tags"`
	Categories            interface{} `json:"categories"`
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
	ShortID                                interface{} `json:"shortId"`
	EnvTemplateScope                       interface{} `json:"envTemplateScope"`
	CreationDate                           string      `json:"creationDate"`
	Name                                   string      `json:"name"`
	ID                                     string      `json:"id"`
}

// Project name and ID
type Project struct {
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	ID       string `json:"id"`
}

// ProjectDetails of a given projects
type ProjectDetails struct {
	HasNonGenericPolicy      bool `json:"hasNonGenericPolicy"`
	CanAddPolicy             bool `json:"canAddPolicy"`
	CanSeeMultipleRegions    bool `json:"canSeeMultipleRegions"`
	MultipleUserRolesEnabled bool `json:"multipleUserRolesEnabled"`
	EnvironmentResourceQuota struct {
		CPUCount     int `json:"cpuCount"`
		DiskSizeMB   int `json:"diskSizeMB"`
		MemorySizeMB int `json:"memorySizeMB"`
	} `json:"environmentResourceQuota"`
	ProjectResourceQuota struct {
		CPUCount     interface{} `json:"cpuCount"`
		DiskSizeMB   interface{} `json:"diskSizeMB"`
		MemorySizeMB interface{} `json:"memorySizeMB"`
	} `json:"projectResourceQuota"`
	SubscriptionResourceQuota struct {
		CPUCount     interface{} `json:"cpuCount"`
		DiskSizeMB   interface{} `json:"diskSizeMB"`
		MemorySizeMB interface{} `json:"memorySizeMB"`
	} `json:"subscriptionResourceQuota"`
	Regions []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		FriendlyName string `json:"friendlyName"`
		CloudName    string `json:"cloudName"`
	} `json:"regions"`
	CanCreateFromScratch        bool        `json:"canCreateFromScratch"`
	DefaultPolicyForEnvCreation interface{} `json:"defaultPolicyForEnvCreation"`
	Teams                       []struct {
		IsDefaultTeam bool   `json:"isDefaultTeam"`
		Name          string `json:"name"`
		ID            string `json:"id"`
	} `json:"teams"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	ID       string `json:"id"`
}

type Policy struct {
	Name                     string `json:"name"`
	ProjectID                string `json:"projectId"`
	AllowEnvironmentCreation bool   `json:"allowEnvironmentCreation"`
	ID                       string `json:"id"`
}
