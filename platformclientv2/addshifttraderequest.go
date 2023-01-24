package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Addshifttraderequest
type Addshifttraderequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ScheduleId - The ID of the schedule to which the initiating and receiving shifts belong
	ScheduleId *string `json:"scheduleId,omitempty"`

	// InitiatingShiftId - The ID of the shift that the initiating user wants to give up
	InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`

	// ReceivingUserId - The ID of the user to whom to send the request (for use in direct trade requests)
	ReceivingUserId *string `json:"receivingUserId,omitempty"`

	// Expiration - When this shift trade request should expire. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`

	// AcceptableIntervals
	AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Addshifttraderequest) SetField(field string, fieldValue interface{}) {
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

func (o Addshifttraderequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Expiration", }
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
	type Alias Addshifttraderequest
	
	Expiration := new(string)
	if o.Expiration != nil {
		
		*Expiration = timeutil.Strftime(o.Expiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expiration = nil
	}
	
	return json.Marshal(&struct { 
		ScheduleId *string `json:"scheduleId,omitempty"`
		
		InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`
		
		ReceivingUserId *string `json:"receivingUserId,omitempty"`
		
		Expiration *string `json:"expiration,omitempty"`
		
		AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`
		Alias
	}{ 
		ScheduleId: o.ScheduleId,
		
		InitiatingShiftId: o.InitiatingShiftId,
		
		ReceivingUserId: o.ReceivingUserId,
		
		Expiration: Expiration,
		
		AcceptableIntervals: o.AcceptableIntervals,
		Alias:    (Alias)(o),
	})
}

func (o *Addshifttraderequest) UnmarshalJSON(b []byte) error {
	var AddshifttraderequestMap map[string]interface{}
	err := json.Unmarshal(b, &AddshifttraderequestMap)
	if err != nil {
		return err
	}
	
	if ScheduleId, ok := AddshifttraderequestMap["scheduleId"].(string); ok {
		o.ScheduleId = &ScheduleId
	}
    
	if InitiatingShiftId, ok := AddshifttraderequestMap["initiatingShiftId"].(string); ok {
		o.InitiatingShiftId = &InitiatingShiftId
	}
    
	if ReceivingUserId, ok := AddshifttraderequestMap["receivingUserId"].(string); ok {
		o.ReceivingUserId = &ReceivingUserId
	}
    
	if expirationString, ok := AddshifttraderequestMap["expiration"].(string); ok {
		Expiration, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationString)
		o.Expiration = &Expiration
	}
	
	if AcceptableIntervals, ok := AddshifttraderequestMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Addshifttraderequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
