package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffrequestupdatemutopictimeoffrequestupdate
type Wfmtimeoffrequestupdatemutopictimeoffrequestupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// MuId
	MuId *string `json:"muId,omitempty"`

	// User
	User *Wfmtimeoffrequestupdatemutopicuserreference `json:"user,omitempty"`

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

	// FullDayEarliestStartOffsetMinutes
	FullDayEarliestStartOffsetMinutes *[]int `json:"fullDayEarliestStartOffsetMinutes,omitempty"`

	// FullDayLatestEndOffsetMinutes
	FullDayLatestEndOffsetMinutes *[]int `json:"fullDayLatestEndOffsetMinutes,omitempty"`

	// DailyDurationMinutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`

	// DurationMinutes
	DurationMinutes *[]int `json:"durationMinutes,omitempty"`

	// PayableMinutes
	PayableMinutes *[]int `json:"payableMinutes,omitempty"`

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
func (o *Wfmtimeoffrequestupdatemutopictimeoffrequestupdate) SetField(field string, fieldValue interface{}) {
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

func (o Wfmtimeoffrequestupdatemutopictimeoffrequestupdate) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmtimeoffrequestupdatemutopictimeoffrequestupdate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MuId *string `json:"muId,omitempty"`
		
		User *Wfmtimeoffrequestupdatemutopicuserreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Substatus *string `json:"substatus,omitempty"`
		
		PartialDayStartDateTimes *[]string `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		FullDayEarliestStartOffsetMinutes *[]int `json:"fullDayEarliestStartOffsetMinutes,omitempty"`
		
		FullDayLatestEndOffsetMinutes *[]int `json:"fullDayLatestEndOffsetMinutes,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		DurationMinutes *[]int `json:"durationMinutes,omitempty"`
		
		PayableMinutes *[]int `json:"payableMinutes,omitempty"`
		
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
		
		MuId: o.MuId,
		
		User: o.User,
		
		IsFullDayRequest: o.IsFullDayRequest,
		
		MarkedAsRead: o.MarkedAsRead,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		Status: o.Status,
		
		Substatus: o.Substatus,
		
		PartialDayStartDateTimes: o.PartialDayStartDateTimes,
		
		FullDayManagementUnitDates: o.FullDayManagementUnitDates,
		
		FullDayEarliestStartOffsetMinutes: o.FullDayEarliestStartOffsetMinutes,
		
		FullDayLatestEndOffsetMinutes: o.FullDayLatestEndOffsetMinutes,
		
		DailyDurationMinutes: o.DailyDurationMinutes,
		
		DurationMinutes: o.DurationMinutes,
		
		PayableMinutes: o.PayableMinutes,
		
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

func (o *Wfmtimeoffrequestupdatemutopictimeoffrequestupdate) UnmarshalJSON(b []byte) error {
	var WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MuId, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["muId"].(string); ok {
		o.MuId = &MuId
	}
    
	if User, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
    
	if MarkedAsRead, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    
	if ActivityCodeId, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if Status, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Substatus, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["substatus"].(string); ok {
		o.Substatus = &Substatus
	}
    
	if PartialDayStartDateTimes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if FullDayEarliestStartOffsetMinutes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["fullDayEarliestStartOffsetMinutes"].([]interface{}); ok {
		FullDayEarliestStartOffsetMinutesString, _ := json.Marshal(FullDayEarliestStartOffsetMinutes)
		json.Unmarshal(FullDayEarliestStartOffsetMinutesString, &o.FullDayEarliestStartOffsetMinutes)
	}
	
	if FullDayLatestEndOffsetMinutes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["fullDayLatestEndOffsetMinutes"].([]interface{}); ok {
		FullDayLatestEndOffsetMinutesString, _ := json.Marshal(FullDayLatestEndOffsetMinutes)
		json.Unmarshal(FullDayLatestEndOffsetMinutesString, &o.FullDayLatestEndOffsetMinutes)
	}
	
	if DailyDurationMinutes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if DurationMinutes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["durationMinutes"].([]interface{}); ok {
		DurationMinutesString, _ := json.Marshal(DurationMinutes)
		json.Unmarshal(DurationMinutesString, &o.DurationMinutes)
	}
	
	if PayableMinutes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["payableMinutes"].([]interface{}); ok {
		PayableMinutesString, _ := json.Marshal(PayableMinutes)
		json.Unmarshal(PayableMinutesString, &o.PayableMinutes)
	}
	
	if Notes, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ReviewedDate, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["reviewedDate"].(string); ok {
		o.ReviewedDate = &ReviewedDate
	}
    
	if ReviewedBy, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["reviewedBy"].(string); ok {
		o.ReviewedBy = &ReviewedBy
	}
    
	if SubmittedDate, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["submittedDate"].(string); ok {
		o.SubmittedDate = &SubmittedDate
	}
    
	if SubmittedBy, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["submittedBy"].(string); ok {
		o.SubmittedBy = &SubmittedBy
	}
    
	if ModifiedDate, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["modifiedDate"].(string); ok {
		o.ModifiedDate = &ModifiedDate
	}
    
	if ModifiedBy, ok := WfmtimeoffrequestupdatemutopictimeoffrequestupdateMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffrequestupdatemutopictimeoffrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
