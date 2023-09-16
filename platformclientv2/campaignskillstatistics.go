package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignskillstatistics
type Campaignskillstatistics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SkillCombinations - Number of available skill combinations for the campaign
	SkillCombinations *int `json:"skillCombinations,omitempty"`

	// EligibleSkilledAgents - Number of eligible agents with skills for campaign
	EligibleSkilledAgents *int `json:"eligibleSkilledAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignskillstatistics) SetField(field string, fieldValue interface{}) {
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

func (o Campaignskillstatistics) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignskillstatistics
	
	return json.Marshal(&struct { 
		SkillCombinations *int `json:"skillCombinations,omitempty"`
		
		EligibleSkilledAgents *int `json:"eligibleSkilledAgents,omitempty"`
		Alias
	}{ 
		SkillCombinations: o.SkillCombinations,
		
		EligibleSkilledAgents: o.EligibleSkilledAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignskillstatistics) UnmarshalJSON(b []byte) error {
	var CampaignskillstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignskillstatisticsMap)
	if err != nil {
		return err
	}
	
	if SkillCombinations, ok := CampaignskillstatisticsMap["skillCombinations"].(float64); ok {
		SkillCombinationsInt := int(SkillCombinations)
		o.SkillCombinations = &SkillCombinationsInt
	}
	
	if EligibleSkilledAgents, ok := CampaignskillstatisticsMap["eligibleSkilledAgents"].(float64); ok {
		EligibleSkilledAgentsInt := int(EligibleSkilledAgents)
		o.EligibleSkilledAgents = &EligibleSkilledAgentsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignskillstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
