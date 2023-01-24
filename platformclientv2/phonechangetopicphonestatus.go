package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicphonestatus
type Phonechangetopicphonestatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// OperationalStatus
	OperationalStatus *string `json:"operationalStatus,omitempty"`

	// Edge
	Edge *Phonechangetopicedgereference `json:"edge,omitempty"`

	// Provision
	Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`

	// LineStatuses
	LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`

	// EventCreationTime
	EventCreationTime *time.Time `json:"eventCreationTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Phonechangetopicphonestatus) SetField(field string, fieldValue interface{}) {
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

func (o Phonechangetopicphonestatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventCreationTime", }
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
	type Alias Phonechangetopicphonestatus
	
	EventCreationTime := new(string)
	if o.EventCreationTime != nil {
		
		*EventCreationTime = timeutil.Strftime(o.EventCreationTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventCreationTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		OperationalStatus *string `json:"operationalStatus,omitempty"`
		
		Edge *Phonechangetopicedgereference `json:"edge,omitempty"`
		
		Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`
		
		LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`
		
		EventCreationTime *string `json:"eventCreationTime,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		OperationalStatus: o.OperationalStatus,
		
		Edge: o.Edge,
		
		Provision: o.Provision,
		
		LineStatuses: o.LineStatuses,
		
		EventCreationTime: EventCreationTime,
		Alias:    (Alias)(o),
	})
}

func (o *Phonechangetopicphonestatus) UnmarshalJSON(b []byte) error {
	var PhonechangetopicphonestatusMap map[string]interface{}
	err := json.Unmarshal(b, &PhonechangetopicphonestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonechangetopicphonestatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if OperationalStatus, ok := PhonechangetopicphonestatusMap["operationalStatus"].(string); ok {
		o.OperationalStatus = &OperationalStatus
	}
    
	if Edge, ok := PhonechangetopicphonestatusMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if Provision, ok := PhonechangetopicphonestatusMap["provision"].(map[string]interface{}); ok {
		ProvisionString, _ := json.Marshal(Provision)
		json.Unmarshal(ProvisionString, &o.Provision)
	}
	
	if LineStatuses, ok := PhonechangetopicphonestatusMap["lineStatuses"].([]interface{}); ok {
		LineStatusesString, _ := json.Marshal(LineStatuses)
		json.Unmarshal(LineStatusesString, &o.LineStatuses)
	}
	
	if eventCreationTimeString, ok := PhonechangetopicphonestatusMap["eventCreationTime"].(string); ok {
		EventCreationTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventCreationTimeString)
		o.EventCreationTime = &EventCreationTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phonechangetopicphonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
