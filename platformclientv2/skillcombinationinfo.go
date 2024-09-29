package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillcombinationinfo
type Skillcombinationinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SkillCombination - A skill combination in the contact queue
	SkillCombination *[]string `json:"skillCombination,omitempty"`

	// TotalCount - Number of total contacts in the contact queue for this skill combination
	TotalCount *int `json:"totalCount,omitempty"`

	// RemainingCount - Number of remaining contacts in the contact queue for this skill combination
	RemainingCount *int `json:"remainingCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Skillcombinationinfo) SetField(field string, fieldValue interface{}) {
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

func (o Skillcombinationinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Skillcombinationinfo
	
	return json.Marshal(&struct { 
		SkillCombination *[]string `json:"skillCombination,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		
		RemainingCount *int `json:"remainingCount,omitempty"`
		Alias
	}{ 
		SkillCombination: o.SkillCombination,
		
		TotalCount: o.TotalCount,
		
		RemainingCount: o.RemainingCount,
		Alias:    (Alias)(o),
	})
}

func (o *Skillcombinationinfo) UnmarshalJSON(b []byte) error {
	var SkillcombinationinfoMap map[string]interface{}
	err := json.Unmarshal(b, &SkillcombinationinfoMap)
	if err != nil {
		return err
	}
	
	if SkillCombination, ok := SkillcombinationinfoMap["skillCombination"].([]interface{}); ok {
		SkillCombinationString, _ := json.Marshal(SkillCombination)
		json.Unmarshal(SkillCombinationString, &o.SkillCombination)
	}
	
	if TotalCount, ok := SkillcombinationinfoMap["totalCount"].(float64); ok {
		TotalCountInt := int(TotalCount)
		o.TotalCount = &TotalCountInt
	}
	
	if RemainingCount, ok := SkillcombinationinfoMap["remainingCount"].(float64); ok {
		RemainingCountInt := int(RemainingCount)
		o.RemainingCount = &RemainingCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Skillcombinationinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
