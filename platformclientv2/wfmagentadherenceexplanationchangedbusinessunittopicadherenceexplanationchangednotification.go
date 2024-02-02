package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification
type Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Agent
	Agent *Wfmagentadherenceexplanationchangedbusinessunittopicuserreference `json:"agent,omitempty"`

	// ManagementUnit
	ManagementUnit *Wfmagentadherenceexplanationchangedbusinessunittopicmanagementunit `json:"managementUnit,omitempty"`

	// BusinessUnit
	BusinessUnit *Wfmagentadherenceexplanationchangedbusinessunittopicbusinessunit `json:"businessUnit,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// LengthMinutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// Notes
	Notes *string `json:"notes,omitempty"`

	// ReviewedBy
	ReviewedBy *Wfmagentadherenceexplanationchangedbusinessunittopicuserreference `json:"reviewedBy,omitempty"`

	// ReviewedDate
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","ReviewedDate", }
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
	type Alias Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Agent *Wfmagentadherenceexplanationchangedbusinessunittopicuserreference `json:"agent,omitempty"`
		
		ManagementUnit *Wfmagentadherenceexplanationchangedbusinessunittopicmanagementunit `json:"managementUnit,omitempty"`
		
		BusinessUnit *Wfmagentadherenceexplanationchangedbusinessunittopicbusinessunit `json:"businessUnit,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ReviewedBy *Wfmagentadherenceexplanationchangedbusinessunittopicuserreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Agent: o.Agent,
		
		ManagementUnit: o.ManagementUnit,
		
		BusinessUnit: o.BusinessUnit,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Notes: o.Notes,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification) UnmarshalJSON(b []byte) error {
	var WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Agent, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if ManagementUnit, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if BusinessUnit, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if VarType, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if startDateString, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Notes, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ReviewedBy, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := WfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotificationMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentadherenceexplanationchangedbusinessunittopicadherenceexplanationchangednotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
