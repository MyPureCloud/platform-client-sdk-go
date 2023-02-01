package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodel
type Predictormodel struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Kpi - The key performance indicator used in the model.
	Kpi *string `json:"kpi,omitempty"`

	// Queues - The List of Queues that are assessed for Predictive Routing.
	Queues *[]Addressableentityref `json:"queues,omitempty"`

	// DateCreated - DateTime indicating when the model was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateTrained - DateTime indicating when the model was last trained. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTrained *time.Time `json:"dateTrained,omitempty"`

	// MediaType - The media type of the model.
	MediaType *string `json:"mediaType,omitempty"`

	// Features
	Features *[]Predictormodelfeature `json:"features,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictormodel) SetField(field string, fieldValue interface{}) {
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

func (o Predictormodel) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateTrained", }
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
	type Alias Predictormodel
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateTrained := new(string)
	if o.DateTrained != nil {
		
		*DateTrained = timeutil.Strftime(o.DateTrained, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateTrained = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateTrained *string `json:"dateTrained,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Features *[]Predictormodelfeature `json:"features,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Kpi: o.Kpi,
		
		Queues: o.Queues,
		
		DateCreated: DateCreated,
		
		DateTrained: DateTrained,
		
		MediaType: o.MediaType,
		
		Features: o.Features,
		Alias:    (Alias)(o),
	})
}

func (o *Predictormodel) UnmarshalJSON(b []byte) error {
	var PredictormodelMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PredictormodelMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Kpi, ok := PredictormodelMap["kpi"].(string); ok {
		o.Kpi = &Kpi
	}
    
	if Queues, ok := PredictormodelMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if dateCreatedString, ok := PredictormodelMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateTrainedString, ok := PredictormodelMap["dateTrained"].(string); ok {
		DateTrained, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateTrainedString)
		o.DateTrained = &DateTrained
	}
	
	if MediaType, ok := PredictormodelMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Features, ok := PredictormodelMap["features"].([]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
