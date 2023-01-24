package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification
type Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - The organization's new dependency tracking build status
	Status *string `json:"status,omitempty"`

	// User
	User *Architectdependencytrackingbuildnotificationuser `json:"user,omitempty"`

	// Client
	Client *Architectdependencytrackingbuildnotificationclient `json:"client,omitempty"`

	// StartTime - The time the last build started, in ISO 8601 format
	StartTime *time.Time `json:"startTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) SetField(field string, fieldValue interface{}) {
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

func (o Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		User *Architectdependencytrackingbuildnotificationuser `json:"user,omitempty"`
		
		Client *Architectdependencytrackingbuildnotificationclient `json:"client,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		User: o.User,
		
		Client: o.Client,
		
		StartTime: StartTime,
		Alias:    (Alias)(o),
	})
}

func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) UnmarshalJSON(b []byte) error {
	var ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if User, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if startTimeString, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
