package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Kpiresult
type Kpiresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// KpiTotalOn - Absolute metric (in which the KPI is based) total for the interactions handled by predictive routing (GPR was on)
	KpiTotalOn *int `json:"kpiTotalOn,omitempty"`

	// KpiTotalOff - Absolute metric (in which the KPI is based) total for the interactions not routed by predictive routing (GPR was off)
	KpiTotalOff *int `json:"kpiTotalOff,omitempty"`

	// InteractionCountOn - Total interactions handled by predictive routing (GPR was on)
	InteractionCountOn *int `json:"interactionCountOn,omitempty"`

	// InteractionCountOff - Total interactions not routed by predictive routing (GPR was off)
	InteractionCountOff *int `json:"interactionCountOff,omitempty"`

	// MediaType - Media type used for the KPI
	MediaType *string `json:"mediaType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Kpiresult) SetField(field string, fieldValue interface{}) {
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

func (o Kpiresult) MarshalJSON() ([]byte, error) {
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
	type Alias Kpiresult
	
	return json.Marshal(&struct { 
		KpiTotalOn *int `json:"kpiTotalOn,omitempty"`
		
		KpiTotalOff *int `json:"kpiTotalOff,omitempty"`
		
		InteractionCountOn *int `json:"interactionCountOn,omitempty"`
		
		InteractionCountOff *int `json:"interactionCountOff,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		Alias
	}{ 
		KpiTotalOn: o.KpiTotalOn,
		
		KpiTotalOff: o.KpiTotalOff,
		
		InteractionCountOn: o.InteractionCountOn,
		
		InteractionCountOff: o.InteractionCountOff,
		
		MediaType: o.MediaType,
		Alias:    (Alias)(o),
	})
}

func (o *Kpiresult) UnmarshalJSON(b []byte) error {
	var KpiresultMap map[string]interface{}
	err := json.Unmarshal(b, &KpiresultMap)
	if err != nil {
		return err
	}
	
	if KpiTotalOn, ok := KpiresultMap["kpiTotalOn"].(float64); ok {
		KpiTotalOnInt := int(KpiTotalOn)
		o.KpiTotalOn = &KpiTotalOnInt
	}
	
	if KpiTotalOff, ok := KpiresultMap["kpiTotalOff"].(float64); ok {
		KpiTotalOffInt := int(KpiTotalOff)
		o.KpiTotalOff = &KpiTotalOffInt
	}
	
	if InteractionCountOn, ok := KpiresultMap["interactionCountOn"].(float64); ok {
		InteractionCountOnInt := int(InteractionCountOn)
		o.InteractionCountOn = &InteractionCountOnInt
	}
	
	if InteractionCountOff, ok := KpiresultMap["interactionCountOff"].(float64); ok {
		InteractionCountOffInt := int(InteractionCountOff)
		o.InteractionCountOff = &InteractionCountOffInt
	}
	
	if MediaType, ok := KpiresultMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Kpiresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
