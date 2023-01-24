package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createpredictorrequest
type Createpredictorrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueIds - The queue IDs associated with the predictor.
	QueueIds *[]string `json:"queueIds,omitempty"`

	// Kpi - The KPI that the predictor attempts to maximize/minimize.
	Kpi *string `json:"kpi,omitempty"`

	// RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`

	// Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *Predictorschedule `json:"schedule,omitempty"`

	// WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing
	WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createpredictorrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createpredictorrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createpredictorrequest
	
	return json.Marshal(&struct { 
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		Alias
	}{ 
		QueueIds: o.QueueIds,
		
		Kpi: o.Kpi,
		
		RoutingTimeoutSeconds: o.RoutingTimeoutSeconds,
		
		Schedule: o.Schedule,
		
		WorkloadBalancingConfig: o.WorkloadBalancingConfig,
		Alias:    (Alias)(o),
	})
}

func (o *Createpredictorrequest) UnmarshalJSON(b []byte) error {
	var CreatepredictorrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatepredictorrequestMap)
	if err != nil {
		return err
	}
	
	if QueueIds, ok := CreatepredictorrequestMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if Kpi, ok := CreatepredictorrequestMap["kpi"].(string); ok {
		o.Kpi = &Kpi
	}
    
	if RoutingTimeoutSeconds, ok := CreatepredictorrequestMap["routingTimeoutSeconds"].(float64); ok {
		RoutingTimeoutSecondsInt := int(RoutingTimeoutSeconds)
		o.RoutingTimeoutSeconds = &RoutingTimeoutSecondsInt
	}
	
	if Schedule, ok := CreatepredictorrequestMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if WorkloadBalancingConfig, ok := CreatepredictorrequestMap["workloadBalancingConfig"].(map[string]interface{}); ok {
		WorkloadBalancingConfigString, _ := json.Marshal(WorkloadBalancingConfig)
		json.Unmarshal(WorkloadBalancingConfigString, &o.WorkloadBalancingConfig)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createpredictorrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
