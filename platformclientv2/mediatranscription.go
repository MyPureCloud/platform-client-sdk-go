package platformclientv2
import (
	"encoding/json"
)

// Mediatranscription
type Mediatranscription struct { 
	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// TranscriptionProvider
	TranscriptionProvider *string `json:"transcriptionProvider,omitempty"`


	// IntegrationId
	IntegrationId *string `json:"integrationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediatranscription) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
