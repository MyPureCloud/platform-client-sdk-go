package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Metric
type Metric struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of this metric
	Name *string `json:"name,omitempty"`

	// MetricDefinitionId - The id of associated metric definition
	MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`

	// ExternalMetricDefinitionId - The id of associated external metric definition
	ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`

	// Objective - Associated objective for this metric
	Objective *Objective `json:"objective,omitempty"`

	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`

	// LinkedMetric - The linked metric entity reference
	LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`

	// DateCreated - The created date of this metric. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateUnlinked - The unlinked workday for this metric if this metric was ever unlinked. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateUnlinked *time.Time `json:"dateUnlinked,omitempty"`

	// Precision - The precision of the metric, must be between 0 and 5
	Precision *int `json:"precision,omitempty"`

	// TimeDisplayUnit - The time unit in which the metric should be displayed -- this parameter is ignored when displaying non-time values
	TimeDisplayUnit *string `json:"timeDisplayUnit,omitempty"`

	// SourcePerformanceProfile - The source performance profile when this metric is linked
	SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Metric) SetField(field string, fieldValue interface{}) {
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

func (o Metric) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateUnlinked", }

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
	type Alias Metric
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateUnlinked := new(string)
	if o.DateUnlinked != nil {
		*DateUnlinked = timeutil.Strftime(o.DateUnlinked, "%Y-%m-%d")
	} else {
		DateUnlinked = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`
		
		Objective *Objective `json:"objective,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateUnlinked *string `json:"dateUnlinked,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		TimeDisplayUnit *string `json:"timeDisplayUnit,omitempty"`
		
		SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MetricDefinitionId: o.MetricDefinitionId,
		
		ExternalMetricDefinitionId: o.ExternalMetricDefinitionId,
		
		Objective: o.Objective,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		LinkedMetric: o.LinkedMetric,
		
		DateCreated: DateCreated,
		
		DateUnlinked: DateUnlinked,
		
		Precision: o.Precision,
		
		TimeDisplayUnit: o.TimeDisplayUnit,
		
		SourcePerformanceProfile: o.SourcePerformanceProfile,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Metric) UnmarshalJSON(b []byte) error {
	var MetricMap map[string]interface{}
	err := json.Unmarshal(b, &MetricMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MetricMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MetricMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MetricDefinitionId, ok := MetricMap["metricDefinitionId"].(string); ok {
		o.MetricDefinitionId = &MetricDefinitionId
	}
    
	if ExternalMetricDefinitionId, ok := MetricMap["externalMetricDefinitionId"].(string); ok {
		o.ExternalMetricDefinitionId = &ExternalMetricDefinitionId
	}
    
	if Objective, ok := MetricMap["objective"].(map[string]interface{}); ok {
		ObjectiveString, _ := json.Marshal(Objective)
		json.Unmarshal(ObjectiveString, &o.Objective)
	}
	
	if PerformanceProfileId, ok := MetricMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
    
	if LinkedMetric, ok := MetricMap["linkedMetric"].(map[string]interface{}); ok {
		LinkedMetricString, _ := json.Marshal(LinkedMetric)
		json.Unmarshal(LinkedMetricString, &o.LinkedMetric)
	}
	
	if dateCreatedString, ok := MetricMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateUnlinkedString, ok := MetricMap["dateUnlinked"].(string); ok {
		DateUnlinked, _ := time.Parse("2006-01-02", dateUnlinkedString)
		o.DateUnlinked = &DateUnlinked
	}
	
	if Precision, ok := MetricMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if TimeDisplayUnit, ok := MetricMap["timeDisplayUnit"].(string); ok {
		o.TimeDisplayUnit = &TimeDisplayUnit
	}
    
	if SourcePerformanceProfile, ok := MetricMap["sourcePerformanceProfile"].(map[string]interface{}); ok {
		SourcePerformanceProfileString, _ := json.Marshal(SourcePerformanceProfile)
		json.Unmarshal(SourcePerformanceProfileString, &o.SourcePerformanceProfile)
	}
	
	if SelfUri, ok := MetricMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Metric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
