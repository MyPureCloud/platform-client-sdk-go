package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingencryptionconfiguration
type Recordingencryptionconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Url - When keyConfigurationType is LocalKeyManager, this should be the url for decryption and must specify the path to where GenesysCloud can requests decryption. When keyConfigurationType is KmsSymmetric, this should be the arn to the key alias for the master key
	Url *string `json:"url,omitempty"`


	// ApiId - The api id for Hawk Authentication. Null if keyConfigurationType is KmsSymmetric
	ApiId *string `json:"apiId,omitempty"`


	// ApiKey - The api shared symmetric key used for hawk authentication. Null if keyConfigurationType is KmsSymmetric
	ApiKey *string `json:"apiKey,omitempty"`


	// KeyConfigurationType - Type should be LocalKeyManager or KmsSymmetric when create or update Key configurations; 'Native' for disabling configuration.
	KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`


	// LastError - The error message related to the configuration
	LastError *Errorbody `json:"lastError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Recordingencryptionconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingencryptionconfiguration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		ApiId *string `json:"apiId,omitempty"`
		
		ApiKey *string `json:"apiKey,omitempty"`
		
		KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`
		
		LastError *Errorbody `json:"lastError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Url: o.Url,
		
		ApiId: o.ApiId,
		
		ApiKey: o.ApiKey,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		LastError: o.LastError,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingencryptionconfiguration) UnmarshalJSON(b []byte) error {
	var RecordingencryptionconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingencryptionconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingencryptionconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Url, ok := RecordingencryptionconfigurationMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if ApiId, ok := RecordingencryptionconfigurationMap["apiId"].(string); ok {
		o.ApiId = &ApiId
	}
	
	if ApiKey, ok := RecordingencryptionconfigurationMap["apiKey"].(string); ok {
		o.ApiKey = &ApiKey
	}
	
	if KeyConfigurationType, ok := RecordingencryptionconfigurationMap["keyConfigurationType"].(string); ok {
		o.KeyConfigurationType = &KeyConfigurationType
	}
	
	if LastError, ok := RecordingencryptionconfigurationMap["lastError"].(map[string]interface{}); ok {
		LastErrorString, _ := json.Marshal(LastError)
		json.Unmarshal(LastErrorString, &o.LastError)
	}
	
	if SelfUri, ok := RecordingencryptionconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingencryptionconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
