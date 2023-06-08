package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomeattributionrequest
type Outcomeattributionrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OutcomeId - ID of Outcome.
	OutcomeId *string `json:"outcomeId,omitempty"`

	// ExternalContactId - The external contact ID of the customer who achieved the outcome.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// AssociatedValue - The total value associated with the customer's outcome.
	AssociatedValue *float32 `json:"associatedValue,omitempty"`

	// Touchpoints - List of interactions that led to this outcome being achieved.
	Touchpoints *[]Touchpoint `json:"touchpoints,omitempty"`

	// CreatedDate - Date outcome was achieved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outcomeattributionrequest) SetField(field string, fieldValue interface{}) {
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

func (o Outcomeattributionrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Outcomeattributionrequest
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		AssociatedValue *float32 `json:"associatedValue,omitempty"`
		
		Touchpoints *[]Touchpoint `json:"touchpoints,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		OutcomeId: o.OutcomeId,
		
		ExternalContactId: o.ExternalContactId,
		
		AssociatedValue: o.AssociatedValue,
		
		Touchpoints: o.Touchpoints,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Outcomeattributionrequest) UnmarshalJSON(b []byte) error {
	var OutcomeattributionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomeattributionrequestMap)
	if err != nil {
		return err
	}
	
	if OutcomeId, ok := OutcomeattributionrequestMap["outcomeId"].(string); ok {
		o.OutcomeId = &OutcomeId
	}
    
	if ExternalContactId, ok := OutcomeattributionrequestMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if AssociatedValue, ok := OutcomeattributionrequestMap["associatedValue"].(float64); ok {
		AssociatedValueFloat32 := float32(AssociatedValue)
		o.AssociatedValue = &AssociatedValueFloat32
	}
    
	if Touchpoints, ok := OutcomeattributionrequestMap["touchpoints"].([]interface{}); ok {
		TouchpointsString, _ := json.Marshal(Touchpoints)
		json.Unmarshal(TouchpointsString, &o.Touchpoints)
	}
	
	if createdDateString, ok := OutcomeattributionrequestMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomeattributionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
