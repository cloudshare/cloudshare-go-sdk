package cloudshare

// Region in which environments are created
type Region struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CloudName    string `json:"cloudName"`
	FriendlyName string `json:"friendlyName"`
}
