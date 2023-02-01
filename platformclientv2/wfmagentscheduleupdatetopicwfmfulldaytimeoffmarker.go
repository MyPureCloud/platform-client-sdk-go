package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker
type Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`

	// ManagementUnitDate
	ManagementUnitDate *string `json:"managementUnitDate,omitempty"`

	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// IsPaid
	IsPaid *bool `json:"isPaid,omitempty"`

	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Paid
	Paid *bool `json:"paid,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker
	
	return json.Marshal(&struct { 
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ManagementUnitDate *string `json:"managementUnitDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		Alias
	}{ 
		TimeOffRequestId: o.TimeOffRequestId,
		
		ManagementUnitDate: o.ManagementUnitDate,
		
		ActivityCodeId: o.ActivityCodeId,
		
		IsPaid: o.IsPaid,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Description: o.Description,
		
		Paid: o.Paid,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequestId, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if ManagementUnitDate, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["managementUnitDate"].(string); ok {
		o.ManagementUnitDate = &ManagementUnitDate
	}
    
	if ActivityCodeId, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if IsPaid, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["isPaid"].(bool); ok {
		o.IsPaid = &IsPaid
	}
    
	if LengthInMinutes, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Description, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Paid, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
