package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
