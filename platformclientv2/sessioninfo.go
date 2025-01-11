package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sessioninfo
type Sessioninfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - Version of the continuous forecast session
	Version *int `json:"version,omitempty"`

	// SessionId - Session ID of the continuous forecast session
	SessionId *string `json:"sessionId,omitempty"`

	// BusinessUnitId - Business unit ID of the continuous forecast session
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// PlanningGroupsVersion - Version of the planning groups
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`

	// DateOfSession - Date and Time of the Session. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateOfSession *time.Time `json:"dateOfSession,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sessioninfo) SetField(field string, fieldValue interface{}) {
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

func (o Sessioninfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateOfSession", }
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
	type Alias Sessioninfo
	
	DateOfSession := new(string)
	if o.DateOfSession != nil {
		
		*DateOfSession = timeutil.Strftime(o.DateOfSession, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateOfSession = nil
	}
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		DateOfSession *string `json:"dateOfSession,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		SessionId: o.SessionId,
		
		BusinessUnitId: o.BusinessUnitId,
		
		PlanningGroupsVersion: o.PlanningGroupsVersion,
		
		DateOfSession: DateOfSession,
		Alias:    (Alias)(o),
	})
}

func (o *Sessioninfo) UnmarshalJSON(b []byte) error {
	var SessioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &SessioninfoMap)
	if err != nil {
		return err
	}
	
	if Version, ok := SessioninfoMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if SessionId, ok := SessioninfoMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if BusinessUnitId, ok := SessioninfoMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if PlanningGroupsVersion, ok := SessioninfoMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	
	if dateOfSessionString, ok := SessioninfoMap["dateOfSession"].(string); ok {
		DateOfSession, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateOfSessionString)
		o.DateOfSession = &DateOfSession
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sessioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
