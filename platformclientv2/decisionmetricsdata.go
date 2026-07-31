package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Decisionmetricsdata
type Decisionmetricsdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User - The user associated with the decision metrics
	User *Userreference `json:"user,omitempty"`

	// PerformanceRank - The performance ranking value of the user for decision metrics. The value ranges from 0 to 9999, with the highest value indicating the best performer
	PerformanceRank *int `json:"performanceRank,omitempty"`

	// TieBreakerValue - A numeric tie-breaker value used to resolve ties in performance rankings. Values are sorted in ascending order, with lower values taking precedence
	TieBreakerValue *int `json:"tieBreakerValue,omitempty"`

	// Metadata - The metadata associated to the users decision metric, which will be null if the user has no associated decision metrics
	Metadata *Wfmentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Decisionmetricsdata) SetField(field string, fieldValue interface{}) {
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

func (o Decisionmetricsdata) MarshalJSON() ([]byte, error) {
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
	type Alias Decisionmetricsdata
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		PerformanceRank *int `json:"performanceRank,omitempty"`
		
		TieBreakerValue *int `json:"tieBreakerValue,omitempty"`
		
		Metadata *Wfmentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		PerformanceRank: o.PerformanceRank,
		
		TieBreakerValue: o.TieBreakerValue,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Decisionmetricsdata) UnmarshalJSON(b []byte) error {
	var DecisionmetricsdataMap map[string]interface{}
	err := json.Unmarshal(b, &DecisionmetricsdataMap)
	if err != nil {
		return err
	}
	
	if User, ok := DecisionmetricsdataMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if PerformanceRank, ok := DecisionmetricsdataMap["performanceRank"].(float64); ok {
		PerformanceRankInt := int(PerformanceRank)
		o.PerformanceRank = &PerformanceRankInt
	}
	
	if TieBreakerValue, ok := DecisionmetricsdataMap["tieBreakerValue"].(float64); ok {
		TieBreakerValueInt := int(TieBreakerValue)
		o.TieBreakerValue = &TieBreakerValueInt
	}
	
	if Metadata, ok := DecisionmetricsdataMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Decisionmetricsdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
