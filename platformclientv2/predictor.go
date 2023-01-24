package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictor
type Predictor struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Queues - The queue IDs associated with the predictor.
	Queues *[]Addressableentityref `json:"queues,omitempty"`

	// Kpi - The KPI that the predictor attempts to maximize/minimize.
	Kpi *string `json:"kpi,omitempty"`

	// RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`

	// Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *Predictorschedule `json:"schedule,omitempty"`

	// State - The predictor state.
	State *string `json:"state,omitempty"`

	// DateCreated - DateTime indicating when the predictor was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - DateTime indicating when the predictor was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing.
	WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`

	// ErrorCode - Predictor error code - optional details on why the predictor went into error state.
	ErrorCode *string `json:"errorCode,omitempty"`

	// Models - Predictor's models
	Models *[]Predictormodelbrief `json:"models,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictor) SetField(field string, fieldValue interface{}) {
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

func (o Predictor) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Predictor
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Models *[]Predictormodelbrief `json:"models,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Queues: o.Queues,
		
		Kpi: o.Kpi,
		
		RoutingTimeoutSeconds: o.RoutingTimeoutSeconds,
		
		Schedule: o.Schedule,
		
		State: o.State,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		WorkloadBalancingConfig: o.WorkloadBalancingConfig,
		
		ErrorCode: o.ErrorCode,
		
		Models: o.Models,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Predictor) UnmarshalJSON(b []byte) error {
	var PredictorMap map[string]interface{}
	err := json.Unmarshal(b, &PredictorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PredictorMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Queues, ok := PredictorMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Kpi, ok := PredictorMap["kpi"].(string); ok {
		o.Kpi = &Kpi
	}
    
	if RoutingTimeoutSeconds, ok := PredictorMap["routingTimeoutSeconds"].(float64); ok {
		RoutingTimeoutSecondsInt := int(RoutingTimeoutSeconds)
		o.RoutingTimeoutSeconds = &RoutingTimeoutSecondsInt
	}
	
	if Schedule, ok := PredictorMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if State, ok := PredictorMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateCreatedString, ok := PredictorMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := PredictorMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if WorkloadBalancingConfig, ok := PredictorMap["workloadBalancingConfig"].(map[string]interface{}); ok {
		WorkloadBalancingConfigString, _ := json.Marshal(WorkloadBalancingConfig)
		json.Unmarshal(WorkloadBalancingConfigString, &o.WorkloadBalancingConfig)
	}
	
	if ErrorCode, ok := PredictorMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if Models, ok := PredictorMap["models"].([]interface{}); ok {
		ModelsString, _ := json.Marshal(Models)
		json.Unmarshal(ModelsString, &o.Models)
	}
	
	if SelfUri, ok := PredictorMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
