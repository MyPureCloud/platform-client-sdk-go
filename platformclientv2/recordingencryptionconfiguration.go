package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingencryptionconfiguration
type Recordingencryptionconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingencryptionconfiguration) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Recordingencryptionconfiguration) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
	}{ 
		Id: o.Id,
		
		Url: o.Url,
		
		ApiId: o.ApiId,
		
		ApiKey: o.ApiKey,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		LastError: o.LastError,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
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
