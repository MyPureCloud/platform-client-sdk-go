package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationupdatedefinition
type Locationupdatedefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the Location. Required for creates, not required for updates
	Name *string `json:"name,omitempty"`

	// Version - Current version of the location
	Version *int `json:"version,omitempty"`

	// State - Current activity status of the location.
	State *string `json:"state,omitempty"`

	// Path - A list of ancestor ids
	Path *[]string `json:"path,omitempty"`

	// Notes - Notes for the location
	Notes *string `json:"notes,omitempty"`

	// ContactUser - The user id of the location contact
	ContactUser *string `json:"contactUser,omitempty"`

	// EmergencyNumber - Emergency number for the location
	EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`

	// Address - Address of the location
	Address *Locationaddress `json:"address,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Locationupdatedefinition) SetField(field string, fieldValue interface{}) {
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

func (o Locationupdatedefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Locationupdatedefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Path *[]string `json:"path,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ContactUser *string `json:"contactUser,omitempty"`
		
		EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`
		
		Address *Locationaddress `json:"address,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Version: o.Version,
		
		State: o.State,
		
		Path: o.Path,
		
		Notes: o.Notes,
		
		ContactUser: o.ContactUser,
		
		EmergencyNumber: o.EmergencyNumber,
		
		Address: o.Address,
		Alias:    (Alias)(o),
	})
}

func (o *Locationupdatedefinition) UnmarshalJSON(b []byte) error {
	var LocationupdatedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &LocationupdatedefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LocationupdatedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := LocationupdatedefinitionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if State, ok := LocationupdatedefinitionMap["state"].(string); ok {
		o.State = &State
	}
    
	if Path, ok := LocationupdatedefinitionMap["path"].([]interface{}); ok {
		PathString, _ := json.Marshal(Path)
		json.Unmarshal(PathString, &o.Path)
	}
	
	if Notes, ok := LocationupdatedefinitionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ContactUser, ok := LocationupdatedefinitionMap["contactUser"].(string); ok {
		o.ContactUser = &ContactUser
	}
    
	if EmergencyNumber, ok := LocationupdatedefinitionMap["emergencyNumber"].(map[string]interface{}); ok {
		EmergencyNumberString, _ := json.Marshal(EmergencyNumber)
		json.Unmarshal(EmergencyNumberString, &o.EmergencyNumber)
	}
	
	if Address, ok := LocationupdatedefinitionMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationupdatedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
