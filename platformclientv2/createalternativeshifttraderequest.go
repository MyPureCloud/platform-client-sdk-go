package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createalternativeshifttraderequest
type Createalternativeshifttraderequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// JobId - The ID of this alternative shift job
	JobId *string `json:"jobId,omitempty"`

	// DropShiftReferenceKeys - A list of offered shift reference keys an agent wants to drop
	DropShiftReferenceKeys *[]string `json:"dropShiftReferenceKeys,omitempty"`

	// PickupShiftReferenceKeys - A list of offered shift reference keys an agent wants to pick up
	PickupShiftReferenceKeys *[]string `json:"pickupShiftReferenceKeys,omitempty"`

	// AlternativeShiftTradeGranularity - The granularity of alternative shifts to be traded
	AlternativeShiftTradeGranularity *string `json:"alternativeShiftTradeGranularity,omitempty"`

	// ExpirationDate - The date when the trade will expire in ISO-8601 format. The trade cannot be approved after expiration
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createalternativeshifttraderequest) SetField(field string, fieldValue interface{}) {
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

func (o Createalternativeshifttraderequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate", }
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
	type Alias Createalternativeshifttraderequest
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	return json.Marshal(&struct { 
		JobId *string `json:"jobId,omitempty"`
		
		DropShiftReferenceKeys *[]string `json:"dropShiftReferenceKeys,omitempty"`
		
		PickupShiftReferenceKeys *[]string `json:"pickupShiftReferenceKeys,omitempty"`
		
		AlternativeShiftTradeGranularity *string `json:"alternativeShiftTradeGranularity,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		Alias
	}{ 
		JobId: o.JobId,
		
		DropShiftReferenceKeys: o.DropShiftReferenceKeys,
		
		PickupShiftReferenceKeys: o.PickupShiftReferenceKeys,
		
		AlternativeShiftTradeGranularity: o.AlternativeShiftTradeGranularity,
		
		ExpirationDate: ExpirationDate,
		Alias:    (Alias)(o),
	})
}

func (o *Createalternativeshifttraderequest) UnmarshalJSON(b []byte) error {
	var CreatealternativeshifttraderequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatealternativeshifttraderequestMap)
	if err != nil {
		return err
	}
	
	if JobId, ok := CreatealternativeshifttraderequestMap["jobId"].(string); ok {
		o.JobId = &JobId
	}
    
	if DropShiftReferenceKeys, ok := CreatealternativeshifttraderequestMap["dropShiftReferenceKeys"].([]interface{}); ok {
		DropShiftReferenceKeysString, _ := json.Marshal(DropShiftReferenceKeys)
		json.Unmarshal(DropShiftReferenceKeysString, &o.DropShiftReferenceKeys)
	}
	
	if PickupShiftReferenceKeys, ok := CreatealternativeshifttraderequestMap["pickupShiftReferenceKeys"].([]interface{}); ok {
		PickupShiftReferenceKeysString, _ := json.Marshal(PickupShiftReferenceKeys)
		json.Unmarshal(PickupShiftReferenceKeysString, &o.PickupShiftReferenceKeys)
	}
	
	if AlternativeShiftTradeGranularity, ok := CreatealternativeshifttraderequestMap["alternativeShiftTradeGranularity"].(string); ok {
		o.AlternativeShiftTradeGranularity = &AlternativeShiftTradeGranularity
	}
    
	if expirationDateString, ok := CreatealternativeshifttraderequestMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createalternativeshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
