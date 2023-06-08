package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemwrapup
type Workitemwrapup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Workitem - Workitem that the wrapup code has been added to.
	Workitem *Workitemreference `json:"workitem,omitempty"`

	// WrapupCode - The wrapup code used in the workitem.
	WrapupCode *Wrapupidreference `json:"wrapupCode,omitempty"`

	// ModifiedBy - The user who added the wrapup code to the workitem.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// User - The user for whom wrapup code was added. This may be the same as modifiedBy.
	User *Userreference `json:"user,omitempty"`

	// DateModified - The modified date of the Workitem when the wrapup code was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemwrapup) SetField(field string, fieldValue interface{}) {
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

func (o Workitemwrapup) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified", }
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
	type Alias Workitemwrapup
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Workitem *Workitemreference `json:"workitem,omitempty"`
		
		WrapupCode *Wrapupidreference `json:"wrapupCode,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		Alias
	}{ 
		Workitem: o.Workitem,
		
		WrapupCode: o.WrapupCode,
		
		ModifiedBy: o.ModifiedBy,
		
		User: o.User,
		
		DateModified: DateModified,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemwrapup) UnmarshalJSON(b []byte) error {
	var WorkitemwrapupMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemwrapupMap)
	if err != nil {
		return err
	}
	
	if Workitem, ok := WorkitemwrapupMap["workitem"].(map[string]interface{}); ok {
		WorkitemString, _ := json.Marshal(Workitem)
		json.Unmarshal(WorkitemString, &o.Workitem)
	}
	
	if WrapupCode, ok := WorkitemwrapupMap["wrapupCode"].(map[string]interface{}); ok {
		WrapupCodeString, _ := json.Marshal(WrapupCode)
		json.Unmarshal(WrapupCodeString, &o.WrapupCode)
	}
	
	if ModifiedBy, ok := WorkitemwrapupMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if User, ok := WorkitemwrapupMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if dateModifiedString, ok := WorkitemwrapupMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemwrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
