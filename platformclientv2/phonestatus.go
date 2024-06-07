package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonestatus
type Phonestatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// OperationalStatus - The Operational Status of this phone
	OperationalStatus *string `json:"operationalStatus,omitempty"`

	// EdgesStatus - The status of the primary or secondary Edges assigned to the phone lines.
	EdgesStatus *string `json:"edgesStatus,omitempty"`

	// EventCreationTime - Event Creation Time represents an ISO-8601 string. For example: UTC, UTC+01:00, or Europe/London
	EventCreationTime *string `json:"eventCreationTime,omitempty"`

	// Provision - Provision information for this phone
	Provision *Provisioninfo `json:"provision,omitempty"`

	// LineStatuses - A list of LineStatus information for each of the lines of this phone
	LineStatuses *[]Linestatus `json:"lineStatuses,omitempty"`

	// PhoneAssignmentToEdgeType - The phone status's edge assignment type.
	PhoneAssignmentToEdgeType *string `json:"phoneAssignmentToEdgeType,omitempty"`

	// Edge - The URI of the edge that provided this status information.
	Edge *Domainentityref `json:"edge,omitempty"`

	// SelfUri - The URI for this object. Deprecated. Do not use.
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Phonestatus) SetField(field string, fieldValue interface{}) {
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

func (o Phonestatus) MarshalJSON() ([]byte, error) {
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
	type Alias Phonestatus
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		OperationalStatus *string `json:"operationalStatus,omitempty"`
		
		EdgesStatus *string `json:"edgesStatus,omitempty"`
		
		EventCreationTime *string `json:"eventCreationTime,omitempty"`
		
		Provision *Provisioninfo `json:"provision,omitempty"`
		
		LineStatuses *[]Linestatus `json:"lineStatuses,omitempty"`
		
		PhoneAssignmentToEdgeType *string `json:"phoneAssignmentToEdgeType,omitempty"`
		
		Edge *Domainentityref `json:"edge,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		OperationalStatus: o.OperationalStatus,
		
		EdgesStatus: o.EdgesStatus,
		
		EventCreationTime: o.EventCreationTime,
		
		Provision: o.Provision,
		
		LineStatuses: o.LineStatuses,
		
		PhoneAssignmentToEdgeType: o.PhoneAssignmentToEdgeType,
		
		Edge: o.Edge,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Phonestatus) UnmarshalJSON(b []byte) error {
	var PhonestatusMap map[string]interface{}
	err := json.Unmarshal(b, &PhonestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonestatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if OperationalStatus, ok := PhonestatusMap["operationalStatus"].(string); ok {
		o.OperationalStatus = &OperationalStatus
	}
    
	if EdgesStatus, ok := PhonestatusMap["edgesStatus"].(string); ok {
		o.EdgesStatus = &EdgesStatus
	}
    
	if EventCreationTime, ok := PhonestatusMap["eventCreationTime"].(string); ok {
		o.EventCreationTime = &EventCreationTime
	}
    
	if Provision, ok := PhonestatusMap["provision"].(map[string]interface{}); ok {
		ProvisionString, _ := json.Marshal(Provision)
		json.Unmarshal(ProvisionString, &o.Provision)
	}
	
	if LineStatuses, ok := PhonestatusMap["lineStatuses"].([]interface{}); ok {
		LineStatusesString, _ := json.Marshal(LineStatuses)
		json.Unmarshal(LineStatusesString, &o.LineStatuses)
	}
	
	if PhoneAssignmentToEdgeType, ok := PhonestatusMap["phoneAssignmentToEdgeType"].(string); ok {
		o.PhoneAssignmentToEdgeType = &PhoneAssignmentToEdgeType
	}
    
	if Edge, ok := PhonestatusMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if SelfUri, ok := PhonestatusMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
