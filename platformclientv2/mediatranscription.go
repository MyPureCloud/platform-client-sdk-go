package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Mediatranscription) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatranscription

	

	return json.Marshal(&struct { 
		DisplayName *string `json:"displayName,omitempty"`
		
		TranscriptionProvider *string `json:"transcriptionProvider,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		*Alias
	}{ 
		DisplayName: u.DisplayName,
		
		TranscriptionProvider: u.TranscriptionProvider,
		
		IntegrationId: u.IntegrationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediatranscription) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
