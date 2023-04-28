package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userinsightstrend
type Userinsightstrend struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PerformanceProfile - The performance profile
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`

	// Division - The division
	Division *Divisionreference `json:"division,omitempty"`

	// Granularity - Granularity
	Granularity *string `json:"granularity,omitempty"`

	// ComparativePeriod - The comparative period work day date range
	ComparativePeriod *Workdayperiod `json:"comparativePeriod,omitempty"`

	// PrimaryPeriod - The primary period work day date range
	PrimaryPeriod *Workdayperiod `json:"primaryPeriod,omitempty"`

	// Entities - The list of insights trend for each metric
	Entities *[]Insightstrendmetricitem `json:"entities,omitempty"`

	// Total - The insights trend in total
	Total *Insightstrendtotalitem `json:"total,omitempty"`

	// User - The query user
	User *Userreference `json:"user,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userinsightstrend) SetField(field string, fieldValue interface{}) {
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

func (o Userinsightstrend) MarshalJSON() ([]byte, error) {
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
	type Alias Userinsightstrend
	
	return json.Marshal(&struct { 
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		ComparativePeriod *Workdayperiod `json:"comparativePeriod,omitempty"`
		
		PrimaryPeriod *Workdayperiod `json:"primaryPeriod,omitempty"`
		
		Entities *[]Insightstrendmetricitem `json:"entities,omitempty"`
		
		Total *Insightstrendtotalitem `json:"total,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		Alias
	}{ 
		PerformanceProfile: o.PerformanceProfile,
		
		Division: o.Division,
		
		Granularity: o.Granularity,
		
		ComparativePeriod: o.ComparativePeriod,
		
		PrimaryPeriod: o.PrimaryPeriod,
		
		Entities: o.Entities,
		
		Total: o.Total,
		
		User: o.User,
		Alias:    (Alias)(o),
	})
}

func (o *Userinsightstrend) UnmarshalJSON(b []byte) error {
	var UserinsightstrendMap map[string]interface{}
	err := json.Unmarshal(b, &UserinsightstrendMap)
	if err != nil {
		return err
	}
	
	if PerformanceProfile, ok := UserinsightstrendMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Division, ok := UserinsightstrendMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Granularity, ok := UserinsightstrendMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if ComparativePeriod, ok := UserinsightstrendMap["comparativePeriod"].(map[string]interface{}); ok {
		ComparativePeriodString, _ := json.Marshal(ComparativePeriod)
		json.Unmarshal(ComparativePeriodString, &o.ComparativePeriod)
	}
	
	if PrimaryPeriod, ok := UserinsightstrendMap["primaryPeriod"].(map[string]interface{}); ok {
		PrimaryPeriodString, _ := json.Marshal(PrimaryPeriod)
		json.Unmarshal(PrimaryPeriodString, &o.PrimaryPeriod)
	}
	
	if Entities, ok := UserinsightstrendMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Total, ok := UserinsightstrendMap["total"].(map[string]interface{}); ok {
		TotalString, _ := json.Marshal(Total)
		json.Unmarshal(TotalString, &o.Total)
	}
	
	if User, ok := UserinsightstrendMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userinsightstrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
