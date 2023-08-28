package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Grammarlanguage
type Grammarlanguage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// GrammarId
	GrammarId *string `json:"grammarId,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`

	// VoiceFileUrl - The URL to the voice mode file associated with this grammar language
	VoiceFileUrl *string `json:"voiceFileUrl,omitempty"`

	// DtmfFileUrl - The URL to the DTMF mode file associated with this grammar language
	DtmfFileUrl *string `json:"dtmfFileUrl,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Grammarlanguage) SetField(field string, fieldValue interface{}) {
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

func (o Grammarlanguage) MarshalJSON() ([]byte, error) {
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
	type Alias Grammarlanguage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		GrammarId *string `json:"grammarId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		VoiceFileUrl *string `json:"voiceFileUrl,omitempty"`
		
		DtmfFileUrl *string `json:"dtmfFileUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		GrammarId: o.GrammarId,
		
		Language: o.Language,
		
		VoiceFileUrl: o.VoiceFileUrl,
		
		DtmfFileUrl: o.DtmfFileUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Grammarlanguage) UnmarshalJSON(b []byte) error {
	var GrammarlanguageMap map[string]interface{}
	err := json.Unmarshal(b, &GrammarlanguageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GrammarlanguageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if GrammarId, ok := GrammarlanguageMap["grammarId"].(string); ok {
		o.GrammarId = &GrammarId
	}
    
	if Language, ok := GrammarlanguageMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if VoiceFileUrl, ok := GrammarlanguageMap["voiceFileUrl"].(string); ok {
		o.VoiceFileUrl = &VoiceFileUrl
	}
    
	if DtmfFileUrl, ok := GrammarlanguageMap["dtmfFileUrl"].(string); ok {
		o.DtmfFileUrl = &DtmfFileUrl
	}
    
	if SelfUri, ok := GrammarlanguageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Grammarlanguage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
