package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V3sourcecreaterequest
type V3sourcecreaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the source.
	Name *string `json:"name,omitempty"`

	// VarType - The type of the source. Required if connectionId is not specified, inherits the connection type otherwise.
	VarType *string `json:"type,omitempty"`

	// ConnectionId - The id of the connection related to the source. Required if type is Sharepoint.
	ConnectionId *string `json:"connectionId,omitempty"`

	// TriggerType - The trigger type of the source. Default is Manual.
	TriggerType *string `json:"triggerType,omitempty"`

	// ScheduleSettings - Settings that determine when the source starts a sync. Required if triggerType is Scheduled.
	ScheduleSettings *V3sourceschedulesettings `json:"scheduleSettings,omitempty"`

	// Filters - Filters that determine what documents are synced.
	Filters *V3sourcefilter `json:"filters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V3sourcecreaterequest) SetField(field string, fieldValue interface{}) {
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

func (o V3sourcecreaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias V3sourcecreaterequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ConnectionId *string `json:"connectionId,omitempty"`
		
		TriggerType *string `json:"triggerType,omitempty"`
		
		ScheduleSettings *V3sourceschedulesettings `json:"scheduleSettings,omitempty"`
		
		Filters *V3sourcefilter `json:"filters,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		ConnectionId: o.ConnectionId,
		
		TriggerType: o.TriggerType,
		
		ScheduleSettings: o.ScheduleSettings,
		
		Filters: o.Filters,
		Alias:    (Alias)(o),
	})
}

func (o *V3sourcecreaterequest) UnmarshalJSON(b []byte) error {
	var V3sourcecreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &V3sourcecreaterequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := V3sourcecreaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := V3sourcecreaterequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ConnectionId, ok := V3sourcecreaterequestMap["connectionId"].(string); ok {
		o.ConnectionId = &ConnectionId
	}
    
	if TriggerType, ok := V3sourcecreaterequestMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if ScheduleSettings, ok := V3sourcecreaterequestMap["scheduleSettings"].(map[string]interface{}); ok {
		ScheduleSettingsString, _ := json.Marshal(ScheduleSettings)
		json.Unmarshal(ScheduleSettingsString, &o.ScheduleSettings)
	}
	
	if Filters, ok := V3sourcecreaterequestMap["filters"].(map[string]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V3sourcecreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
