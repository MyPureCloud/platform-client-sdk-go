package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationnotification
type Adherenceexplanationnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Agent - The agent for whom the adherence explanation applies
	Agent *Userreference `json:"agent,omitempty"`

	// ManagementUnit - The management unit to which the agent belonged at the time the adherence explanation was submitted
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// BusinessUnit - The business unit to which the agent belonged at the time the adherence explanation was submitted
	BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`

	// StartDate - The start date of the adherence explanation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes - The length of the adherence explanation in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`

	// VarType - The type of the adherence explanation
	VarType *string `json:"type,omitempty"`

	// Notes - Notes about the adherence explanation
	Notes *string `json:"notes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adherenceexplanationnotification) SetField(field string, fieldValue interface{}) {
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

func (o Adherenceexplanationnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate", }
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
	type Alias Adherenceexplanationnotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Userreference `json:"agent,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		BusinessUnit *Businessunitreference `json:"businessUnit,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Agent: o.Agent,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Status: o.Status,
		
		VarType: o.VarType,
		
		Notes: o.Notes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Adherenceexplanationnotification) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceexplanationnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := AdherenceexplanationnotificationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := AdherenceexplanationnotificationMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := AdherenceexplanationnotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if startDateString, ok := AdherenceexplanationnotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := AdherenceexplanationnotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Status, ok := AdherenceexplanationnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarType, ok := AdherenceexplanationnotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Notes, ok := AdherenceexplanationnotificationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if SelfUri, ok := AdherenceexplanationnotificationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
