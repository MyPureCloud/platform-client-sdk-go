package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyviewchart - A chart within the context of the elements of the the journey view
type Journeyviewchart struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Version - The version of the journey view chart
	Version *int `json:"version,omitempty"`

	// GroupByTime - A time unit to group the metrics by. There is a limit on the number of groupBy properties which can be specified.
	GroupByTime *string `json:"groupByTime,omitempty"`

	// GroupByAttributes - A list of attributes to group the metrics by. There is a limit on the number of groupBy properties which can be specified.
	GroupByAttributes *[]Journeyviewchartgroupbyattribute `json:"groupByAttributes,omitempty"`

	// Metrics - A list of metrics to calculate within the chart by (aka the y axis)
	Metrics *[]Journeyviewchartmetric `json:"metrics,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyviewchart) SetField(field string, fieldValue interface{}) {
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

func (o Journeyviewchart) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyviewchart
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		GroupByTime *string `json:"groupByTime,omitempty"`
		
		GroupByAttributes *[]Journeyviewchartgroupbyattribute `json:"groupByAttributes,omitempty"`
		
		Metrics *[]Journeyviewchartmetric `json:"metrics,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		GroupByTime: o.GroupByTime,
		
		GroupByAttributes: o.GroupByAttributes,
		
		Metrics: o.Metrics,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyviewchart) UnmarshalJSON(b []byte) error {
	var JourneyviewchartMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyviewchartMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyviewchartMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := JourneyviewchartMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := JourneyviewchartMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if GroupByTime, ok := JourneyviewchartMap["groupByTime"].(string); ok {
		o.GroupByTime = &GroupByTime
	}
    
	if GroupByAttributes, ok := JourneyviewchartMap["groupByAttributes"].([]interface{}); ok {
		GroupByAttributesString, _ := json.Marshal(GroupByAttributes)
		json.Unmarshal(GroupByAttributesString, &o.GroupByAttributes)
	}
	
	if Metrics, ok := JourneyviewchartMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if SelfUri, ok := JourneyviewchartMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyviewchart) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
