package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationencryptionconfiguration
type Conversationencryptionconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Url - keyConfigurationType is always KmsSymmetric, and should be the arn to the key alias for the master key
	Url *string `json:"url,omitempty"`


	// KeyConfigurationType - Type should be 'KmsSymmetric' when create or update Key configurations, 'None' to disable configuration.
	KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`


	// LastError - The error message related to the configuration
	LastError *Errorbody `json:"lastError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Conversationencryptionconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationencryptionconfiguration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`
		
		LastError *Errorbody `json:"lastError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Url: o.Url,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		LastError: o.LastError,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationencryptionconfiguration) UnmarshalJSON(b []byte) error {
	var ConversationencryptionconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationencryptionconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationencryptionconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Url, ok := ConversationencryptionconfigurationMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if KeyConfigurationType, ok := ConversationencryptionconfigurationMap["keyConfigurationType"].(string); ok {
		o.KeyConfigurationType = &KeyConfigurationType
	}
    
	if LastError, ok := ConversationencryptionconfigurationMap["lastError"].(map[string]interface{}); ok {
		LastErrorString, _ := json.Marshal(LastError)
		json.Unmarshal(LastErrorString, &o.LastError)
	}
	
	if SelfUri, ok := ConversationencryptionconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationencryptionconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
