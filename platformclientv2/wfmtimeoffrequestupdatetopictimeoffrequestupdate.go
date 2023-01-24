package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffrequestupdatetopictimeoffrequestupdate
type Wfmtimeoffrequestupdatetopictimeoffrequestupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// User
	User *Wfmtimeoffrequestupdatetopicuserreference `json:"user,omitempty"`

	// IsFullDayRequest
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`

	// MarkedAsRead
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`

	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Paid
	Paid *bool `json:"paid,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// Substatus
	Substatus *string `json:"substatus,omitempty"`

	// PartialDayStartDateTimes
	PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`

	// FullDayManagementUnitDates
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`

	// DailyDurationMinutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`

	// Notes
	Notes *string `json:"notes,omitempty"`

	// ReviewedDate
	ReviewedDate *string `json:"reviewedDate,omitempty"`

	// ReviewedBy
	ReviewedBy *string `json:"reviewedBy,omitempty"`

	// SubmittedDate
	SubmittedDate *string `json:"submittedDate,omitempty"`

	// SubmittedBy
	SubmittedBy *string `json:"submittedBy,omitempty"`

	// ModifiedDate
	ModifiedDate *string `json:"modifiedDate,omitempty"`

	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) SetField(field string, fieldValue interface{}) {
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

func (o Wfmtimeoffrequestupdatetopictimeoffrequestupdate) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmtimeoffrequestupdatetopictimeoffrequestupdate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Wfmtimeoffrequestupdatetopicuserreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Substatus *string `json:"substatus,omitempty"`
		
		PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		ReviewedBy *string `json:"reviewedBy,omitempty"`
		
		SubmittedDate *string `json:"submittedDate,omitempty"`
		
		SubmittedBy *string `json:"submittedBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		IsFullDayRequest: o.IsFullDayRequest,
		
		MarkedAsRead: o.MarkedAsRead,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		Status: o.Status,
		
		Substatus: o.Substatus,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		
		Notes: o.Notes,
		
		ReviewedDate: o.ReviewedDate,
		
		ReviewedBy: o.ReviewedBy,
		
		SubmittedDate: o.SubmittedDate,
		
		SubmittedBy: o.SubmittedBy,
		
		ModifiedDate: o.ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) UnmarshalJSON(b []byte) error {
	var WfmtimeoffrequestupdatetopictimeoffrequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffrequestupdatetopictimeoffrequestupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if User, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
    
	if MarkedAsRead, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    
	if ActivityCodeId, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if Status, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Substatus, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["substatus"].(string); ok {
		o.Substatus = &Substatus
	}
    
	if PartialDayStartDateTimes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if DailyDurationMinutes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if Notes, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ReviewedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["reviewedDate"].(string); ok {
		o.ReviewedDate = &ReviewedDate
	}
    
	if ReviewedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["reviewedBy"].(string); ok {
		o.ReviewedBy = &ReviewedBy
	}
    
	if SubmittedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["submittedDate"].(string); ok {
		o.SubmittedDate = &SubmittedDate
	}
    
	if SubmittedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["submittedBy"].(string); ok {
		o.SubmittedBy = &SubmittedBy
	}
    
	if ModifiedDate, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["modifiedDate"].(string); ok {
		o.ModifiedDate = &ModifiedDate
	}
    
	if ModifiedBy, ok := WfmtimeoffrequestupdatetopictimeoffrequestupdateMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatetopictimeoffrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
