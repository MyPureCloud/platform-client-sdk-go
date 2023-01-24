package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotdisconnectaction - Settings for a next-action of disconnecting, including the reason code for the disconnect.
type Textbotdisconnectaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Reason - The reason for the disconnect.
	Reason *string `json:"reason,omitempty"`

	// ReasonExtendedInfo - Extended information related to the reason, if available.
	ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`

	// FlowLocation - Describes where in the Bot Flow the user was when the disconnect occurred.
	FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`

	// FlowOutcomes - The list of Flow Outcomes for the bot flow and their details.
	FlowOutcomes *[]Textbotflowoutcome `json:"flowOutcomes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotdisconnectaction) SetField(field string, fieldValue interface{}) {
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

func (o Textbotdisconnectaction) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotdisconnectaction
	
	return json.Marshal(&struct { 
		Reason *string `json:"reason,omitempty"`
		
		ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`
		
		FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`
		
		FlowOutcomes *[]Textbotflowoutcome `json:"flowOutcomes,omitempty"`
		Alias
	}{ 
		Reason: o.Reason,
		
		ReasonExtendedInfo: o.ReasonExtendedInfo,
		
		FlowLocation: o.FlowLocation,
		
		FlowOutcomes: o.FlowOutcomes,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotdisconnectaction) UnmarshalJSON(b []byte) error {
	var TextbotdisconnectactionMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotdisconnectactionMap)
	if err != nil {
		return err
	}
	
	if Reason, ok := TextbotdisconnectactionMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    
	if ReasonExtendedInfo, ok := TextbotdisconnectactionMap["reasonExtendedInfo"].(string); ok {
		o.ReasonExtendedInfo = &ReasonExtendedInfo
	}
    
	if FlowLocation, ok := TextbotdisconnectactionMap["flowLocation"].(map[string]interface{}); ok {
		FlowLocationString, _ := json.Marshal(FlowLocation)
		json.Unmarshal(FlowLocationString, &o.FlowLocation)
	}
	
	if FlowOutcomes, ok := TextbotdisconnectactionMap["flowOutcomes"].([]interface{}); ok {
		FlowOutcomesString, _ := json.Marshal(FlowOutcomes)
		json.Unmarshal(FlowOutcomesString, &o.FlowOutcomes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotdisconnectaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
