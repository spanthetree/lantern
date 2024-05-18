package shared

// this struct defines the local service data, for use in registering itself to consul
type LocalServiceData struct {
	ID         string
	Name       string
	Address    string
	FQDN       string
	Tags       []string
	DataCenter string
}
