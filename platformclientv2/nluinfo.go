package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluinfo
type Nluinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Domain
	Domain *Addressableentityref `json:"domain,omitempty"`

	// Version
	Version *Nludomainversion `json:"version,omitempty"`

	// Intents
	Intents *[]Intent `json:"intents,omitempty"`

	// EngineVersion
	EngineVersion *string `json:"engineVersion,omitempty"`

	// NluData
	NluData *Nludomainversion `json:"nluData,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nluinfo) SetField(field string, fieldValue interface{}) {
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

func (o Nluinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Nluinfo
	
	return json.Marshal(&struct { 
		Domain *Addressableentityref `json:"domain,omitempty"`
		
		Version *Nludomainversion `json:"version,omitempty"`
		
		Intents *[]Intent `json:"intents,omitempty"`
		
		EngineVersion *string `json:"engineVersion,omitempty"`
		
		NluData *Nludomainversion `json:"nluData,omitempty"`
		Alias
	}{ 
		Domain: o.Domain,
		
		Version: o.Version,
		
		Intents: o.Intents,
		
		EngineVersion: o.EngineVersion,
		
		NluData: o.NluData,
		Alias:    (Alias)(o),
	})
}

func (o *Nluinfo) UnmarshalJSON(b []byte) error {
	var NluinfoMap map[string]interface{}
	err := json.Unmarshal(b, &NluinfoMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := NluinfoMap["domain"].(map[string]interface{}); ok {
		DomainString, _ := json.Marshal(Domain)
		json.Unmarshal(DomainString, &o.Domain)
	}
	
	if Version, ok := NluinfoMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if Intents, ok := NluinfoMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if EngineVersion, ok := NluinfoMap["engineVersion"].(string); ok {
		o.EngineVersion = &EngineVersion
	}
    
	if NluData, ok := NluinfoMap["nluData"].(map[string]interface{}); ok {
		NluDataString, _ := json.Marshal(NluData)
		json.Unmarshal(NluDataString, &o.NluData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
