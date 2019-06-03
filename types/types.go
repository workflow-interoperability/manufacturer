package types

// SubscriberInformation is blockchain concept
type SubscriberInformation struct {
	Roles []string `json:"roles"`
	ID    string   `json:"id"`
}

// FromToData is blockchain concept
type FromToData struct {
	ProcessID         string
	ProcessInstanceID string
	IESMID            string
}

// WorkflowRelevantData is blockchain concept
type WorkflowRelevantData struct {
	From FromToData `json:"from"`
	To   FromToData `json:"to"`
}

// ApplicationData is blockchain concept
type ApplicationData struct {
	URL string `json:"url"`
}

// Payload is blockchain concept
type Payload struct {
	ApplicationData      ApplicationData      `json:"applicationData"`
	WorkflowRelevantData WorkflowRelevantData `json:"workflowRelevantData"`
	WorkflowControlData  string               `json:"workflowControlData"`
}

// IM is blockchain asset
type IM struct {
	ID                    string                `json:"id"`
	Payload               Payload               `json:"payload"`
	Owner                 string                `json:"owner"`
	SubscriberInformation SubscriberInformation `json:"subscriberInformation"`
}

// PIIS is blockchain asset
type PIIS struct {
	ID                    string     `json:"id"`
	From                  FromToData `json:"from"`
	To                    FromToData `json:"to"`
	IMID                  string
	Owner                 string                `json:"owner"`
	SubscriberInformation SubscriberInformation `json:"subscriberInformation"`
}

// PublishIM is blockchain operation
type PublishIM struct {
	IM IM `json:"im"`
}

// PublishPIIS is blockchain operation
type PublishPIIS struct {
	PIIS PIIS `json:"piis"`
}
