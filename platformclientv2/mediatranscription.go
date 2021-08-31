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

func (o *Mediatranscription) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatranscription
	
	return json.Marshal(&struct { 
		DisplayName *string `json:"displayName,omitempty"`
		
		TranscriptionProvider *string `json:"transcriptionProvider,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		*Alias
	}{ 
		DisplayName: o.DisplayName,
		
		TranscriptionProvider: o.TranscriptionProvider,
		
		IntegrationId: o.IntegrationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediatranscription) UnmarshalJSON(b []byte) error {
	var MediatranscriptionMap map[string]interface{}
	err := json.Unmarshal(b, &MediatranscriptionMap)
	if err != nil {
		return err
	}
	
	if DisplayName, ok := MediatranscriptionMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	
	if TranscriptionProvider, ok := MediatranscriptionMap["transcriptionProvider"].(string); ok {
		o.TranscriptionProvider = &TranscriptionProvider
	}
	
	if IntegrationId, ok := MediatranscriptionMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediatranscription) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
