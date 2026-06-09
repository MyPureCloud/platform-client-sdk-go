package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationlinkresponse
type Organizationlinkresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SourceOrganizationId - Organization Id for the login organization.
	SourceOrganizationId *string `json:"sourceOrganizationId,omitempty"`

	// TargetOrganizationId - Organization Id for the linking organization.
	TargetOrganizationId *string `json:"targetOrganizationId,omitempty"`

	// SourceRegion - Region where context organization is hosted, ie. us-east-1
	SourceRegion *string `json:"sourceRegion,omitempty"`

	// TargetRegion - Region where linking organization is hosted, ie. us-east-2
	TargetRegion *string `json:"targetRegion,omitempty"`

	// TargetName - Name for the linking organization.
	TargetName *string `json:"targetName,omitempty"`

	// Status - Status of the linking.
	Status *string `json:"status,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationlinkresponse) SetField(field string, fieldValue interface{}) {
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

func (o Organizationlinkresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationlinkresponse
	
	return json.Marshal(&struct { 
		SourceOrganizationId *string `json:"sourceOrganizationId,omitempty"`
		
		TargetOrganizationId *string `json:"targetOrganizationId,omitempty"`
		
		SourceRegion *string `json:"sourceRegion,omitempty"`
		
		TargetRegion *string `json:"targetRegion,omitempty"`
		
		TargetName *string `json:"targetName,omitempty"`
		
		Status *string `json:"status,omitempty"`
		Alias
	}{ 
		SourceOrganizationId: o.SourceOrganizationId,
		
		TargetOrganizationId: o.TargetOrganizationId,
		
		SourceRegion: o.SourceRegion,
		
		TargetRegion: o.TargetRegion,
		
		TargetName: o.TargetName,
		
		Status: o.Status,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationlinkresponse) UnmarshalJSON(b []byte) error {
	var OrganizationlinkresponseMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationlinkresponseMap)
	if err != nil {
		return err
	}
	
	if SourceOrganizationId, ok := OrganizationlinkresponseMap["sourceOrganizationId"].(string); ok {
		o.SourceOrganizationId = &SourceOrganizationId
	}
    
	if TargetOrganizationId, ok := OrganizationlinkresponseMap["targetOrganizationId"].(string); ok {
		o.TargetOrganizationId = &TargetOrganizationId
	}
    
	if SourceRegion, ok := OrganizationlinkresponseMap["sourceRegion"].(string); ok {
		o.SourceRegion = &SourceRegion
	}
    
	if TargetRegion, ok := OrganizationlinkresponseMap["targetRegion"].(string); ok {
		o.TargetRegion = &TargetRegion
	}
    
	if TargetName, ok := OrganizationlinkresponseMap["targetName"].(string); ok {
		o.TargetName = &TargetName
	}
    
	if Status, ok := OrganizationlinkresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationlinkresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
