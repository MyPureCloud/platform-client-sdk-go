package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Snapshotinfo
type Snapshotinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - Version of the snapshot
	Version *int `json:"version,omitempty"`

	// SnapshotId - Snapshot Id of the continuous forecast session
	SnapshotId *string `json:"snapshotId,omitempty"`

	// SessionId - Session Id of the continuous forecast session
	SessionId *string `json:"sessionId,omitempty"`

	// BusinessUnitId - Business unit ID of the continuous forecast session
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// PlanningGroupsVersion - Version of the planning groups
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Snapshotinfo) SetField(field string, fieldValue interface{}) {
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

func (o Snapshotinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Snapshotinfo
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		SnapshotId *string `json:"snapshotId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		SnapshotId: o.SnapshotId,
		
		SessionId: o.SessionId,
		
		BusinessUnitId: o.BusinessUnitId,
		
		PlanningGroupsVersion: o.PlanningGroupsVersion,
		Alias:    (Alias)(o),
	})
}

func (o *Snapshotinfo) UnmarshalJSON(b []byte) error {
	var SnapshotinfoMap map[string]interface{}
	err := json.Unmarshal(b, &SnapshotinfoMap)
	if err != nil {
		return err
	}
	
	if Version, ok := SnapshotinfoMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if SnapshotId, ok := SnapshotinfoMap["snapshotId"].(string); ok {
		o.SnapshotId = &SnapshotId
	}
    
	if SessionId, ok := SnapshotinfoMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if BusinessUnitId, ok := SnapshotinfoMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if PlanningGroupsVersion, ok := SnapshotinfoMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Snapshotinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
