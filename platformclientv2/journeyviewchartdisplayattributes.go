package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewchartdisplayattributes - Display attributes for the chart, such as type, labels and legends
type Journeyviewchartdisplayattributes struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of chart to display
	VarType *string `json:"type,omitempty"`

	// GroupByTitle - A title for the grouped by attributes (aka the x axis)
	GroupByTitle *string `json:"groupByTitle,omitempty"`

	// MetricsTitle - A title for the metrics (aka the y axis)
	MetricsTitle *string `json:"metricsTitle,omitempty"`

	// ShowLegend - Whether to show a legend
	ShowLegend *bool `json:"showLegend,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewchartdisplayattributes) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewchartdisplayattributes) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewchartdisplayattributes
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		GroupByTitle *string `json:"groupByTitle,omitempty"`
		
		MetricsTitle *string `json:"metricsTitle,omitempty"`
		
		ShowLegend *bool `json:"showLegend,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		GroupByTitle: o.GroupByTitle,
		
		MetricsTitle: o.MetricsTitle,
		
		ShowLegend: o.ShowLegend,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewchartdisplayattributes) UnmarshalJSON(b []byte) error {
	var JourneyviewchartdisplayattributesMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewchartdisplayattributesMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneyviewchartdisplayattributesMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if GroupByTitle, ok := JourneyviewchartdisplayattributesMap["groupByTitle"].(string); ok {
		o.GroupByTitle = &GroupByTitle
	}
    
	if MetricsTitle, ok := JourneyviewchartdisplayattributesMap["metricsTitle"].(string); ok {
		o.MetricsTitle = &MetricsTitle
	}
    
	if ShowLegend, ok := JourneyviewchartdisplayattributesMap["showLegend"].(bool); ok {
		o.ShowLegend = &ShowLegend
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewchartdisplayattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
