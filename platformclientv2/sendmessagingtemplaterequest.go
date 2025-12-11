package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendmessagingtemplaterequest
type Sendmessagingtemplaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ResponseId - Unique identifier for a Response Management response to fetch and apply pre-configured message content when sending outbound responses.
	ResponseId *string `json:"responseId,omitempty"`

	// Parameters - A list of Response Management response substitutions for the response's messaging template. (Deprecated) use bodyParameters instead.
	Parameters *[]Templateparameter `json:"parameters,omitempty"`

	// HeaderParameters - A list of Response Management header parameter substitutions for the response's messaging template
	HeaderParameters *[]Templateparameter `json:"headerParameters,omitempty"`

	// BodyParameters - A list of Response Management body parameter substitutions for the response's messaging template
	BodyParameters *[]Templateparameter `json:"bodyParameters,omitempty"`

	// ButtonUrlParameters - A list of Response Management button url parameter substitutions for the response's messaging template
	ButtonUrlParameters *[]Templateparameter `json:"buttonUrlParameters,omitempty"`

	// CarouselParameters - Template parameters for carousel card components
	CarouselParameters *Carouselparameters `json:"carouselParameters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sendmessagingtemplaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Sendmessagingtemplaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Sendmessagingtemplaterequest
	
	return json.Marshal(&struct { 
		ResponseId *string `json:"responseId,omitempty"`
		
		Parameters *[]Templateparameter `json:"parameters,omitempty"`
		
		HeaderParameters *[]Templateparameter `json:"headerParameters,omitempty"`
		
		BodyParameters *[]Templateparameter `json:"bodyParameters,omitempty"`
		
		ButtonUrlParameters *[]Templateparameter `json:"buttonUrlParameters,omitempty"`
		
		CarouselParameters *Carouselparameters `json:"carouselParameters,omitempty"`
		Alias
	}{ 
		ResponseId: o.ResponseId,
		
		Parameters: o.Parameters,
		
		HeaderParameters: o.HeaderParameters,
		
		BodyParameters: o.BodyParameters,
		
		ButtonUrlParameters: o.ButtonUrlParameters,
		
		CarouselParameters: o.CarouselParameters,
		Alias:    (Alias)(o),
	})
}

func (o *Sendmessagingtemplaterequest) UnmarshalJSON(b []byte) error {
	var SendmessagingtemplaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &SendmessagingtemplaterequestMap)
	if err != nil {
		return err
	}
	
	if ResponseId, ok := SendmessagingtemplaterequestMap["responseId"].(string); ok {
		o.ResponseId = &ResponseId
	}
    
	if Parameters, ok := SendmessagingtemplaterequestMap["parameters"].([]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if HeaderParameters, ok := SendmessagingtemplaterequestMap["headerParameters"].([]interface{}); ok {
		HeaderParametersString, _ := json.Marshal(HeaderParameters)
		json.Unmarshal(HeaderParametersString, &o.HeaderParameters)
	}
	
	if BodyParameters, ok := SendmessagingtemplaterequestMap["bodyParameters"].([]interface{}); ok {
		BodyParametersString, _ := json.Marshal(BodyParameters)
		json.Unmarshal(BodyParametersString, &o.BodyParameters)
	}
	
	if ButtonUrlParameters, ok := SendmessagingtemplaterequestMap["buttonUrlParameters"].([]interface{}); ok {
		ButtonUrlParametersString, _ := json.Marshal(ButtonUrlParameters)
		json.Unmarshal(ButtonUrlParametersString, &o.ButtonUrlParameters)
	}
	
	if CarouselParameters, ok := SendmessagingtemplaterequestMap["carouselParameters"].(map[string]interface{}); ok {
		CarouselParametersString, _ := json.Marshal(CarouselParameters)
		json.Unmarshal(CarouselParametersString, &o.CarouselParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sendmessagingtemplaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
