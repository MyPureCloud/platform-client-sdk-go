package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Autostatustransitiondetail
type Autostatustransitiondetail struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// NextStatus - Next status of auto status transition.
	NextStatus *Workitemstatusreference `json:"nextStatus,omitempty"`

	// DateOfTransition - Date at which auto status transition occurs. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateOfTransition *time.Time `json:"dateOfTransition,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Autostatustransitiondetail) SetField(field string, fieldValue interface{}) {
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

func (o Autostatustransitiondetail) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateOfTransition", }
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
	type Alias Autostatustransitiondetail
	
	DateOfTransition := new(string)
	if o.DateOfTransition != nil {
		
		*DateOfTransition = timeutil.Strftime(o.DateOfTransition, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateOfTransition = nil
	}
	
	return json.Marshal(&struct { 
		NextStatus *Workitemstatusreference `json:"nextStatus,omitempty"`
		
		DateOfTransition *string `json:"dateOfTransition,omitempty"`
		Alias
	}{ 
		NextStatus: o.NextStatus,
		
		DateOfTransition: DateOfTransition,
		Alias:    (Alias)(o),
	})
}

func (o *Autostatustransitiondetail) UnmarshalJSON(b []byte) error {
	var AutostatustransitiondetailMap map[string]interface{}
	err := json.Unmarshal(b, &AutostatustransitiondetailMap)
	if err != nil {
		return err
	}
	
	if NextStatus, ok := AutostatustransitiondetailMap["nextStatus"].(map[string]interface{}); ok {
		NextStatusString, _ := json.Marshal(NextStatus)
		json.Unmarshal(NextStatusString, &o.NextStatus)
	}
	
	if dateOfTransitionString, ok := AutostatustransitiondetailMap["dateOfTransition"].(string); ok {
		DateOfTransition, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateOfTransitionString)
		o.DateOfTransition = &DateOfTransition
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Autostatustransitiondetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
