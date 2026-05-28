package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activateexternaleventresponse - Response for activation of an external event
type Activateexternaleventresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SchemaId - The schema ID of the external event
	SchemaId *string `json:"schemaId,omitempty"`

	// EventName - The name of the external event
	EventName *string `json:"eventName,omitempty"`

	// DisplayName - The display name of the external event
	DisplayName *string `json:"displayName,omitempty"`

	// Rank - The rank of the external event
	Rank *int `json:"rank,omitempty"`

	// ActivationStatus - The activation status of the external event
	ActivationStatus *string `json:"activationStatus,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activateexternaleventresponse) SetField(field string, fieldValue interface{}) {
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

func (o Activateexternaleventresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Activateexternaleventresponse
	
	return json.Marshal(&struct { 
		SchemaId *string `json:"schemaId,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		
		ActivationStatus *string `json:"activationStatus,omitempty"`
		Alias
	}{ 
		SchemaId: o.SchemaId,
		
		EventName: o.EventName,
		
		DisplayName: o.DisplayName,
		
		Rank: o.Rank,
		
		ActivationStatus: o.ActivationStatus,
		Alias:    (Alias)(o),
	})
}

func (o *Activateexternaleventresponse) UnmarshalJSON(b []byte) error {
	var ActivateexternaleventresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ActivateexternaleventresponseMap)
	if err != nil {
		return err
	}
	
	if SchemaId, ok := ActivateexternaleventresponseMap["schemaId"].(string); ok {
		o.SchemaId = &SchemaId
	}
    
	if EventName, ok := ActivateexternaleventresponseMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if DisplayName, ok := ActivateexternaleventresponseMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Rank, ok := ActivateexternaleventresponseMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	
	if ActivationStatus, ok := ActivateexternaleventresponseMap["activationStatus"].(string); ok {
		o.ActivationStatus = &ActivationStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activateexternaleventresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
