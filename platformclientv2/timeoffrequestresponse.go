package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestresponse
type Timeoffrequestresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// User - The user associated with this time off request
	User *Userreference `json:"user,omitempty"`

	// IsFullDayRequest - Whether this is a full day request (false means partial day)
	IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`

	// MarkedAsRead - Whether this request has been marked as read by the agent
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`

	// ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// Paid - Whether this is a paid time off request
	Paid *bool `json:"paid,omitempty"`

	// Status - The status of this time off request
	Status *string `json:"status,omitempty"`

	// Substatus - The substatus of this time off request
	Substatus *string `json:"substatus,omitempty"`

	// PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests. Will be not empty if isFullDayRequest == false
	PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`

	// FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone. Will be not empty if isFullDayRequest == true
	FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`

	// DailyDurationMinutes - The daily duration of this time off request in minutes
	DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`

	// DurationMinutes - Daily durations for each day of this time off request in minutes
	DurationMinutes *[]int `json:"durationMinutes,omitempty"`

	// PayableMinutes - Payable minutes for each day of this time off request
	PayableMinutes *[]int `json:"payableMinutes,omitempty"`

	// Notes - Notes about the time off request
	Notes *string `json:"notes,omitempty"`

	// SubmittedBy - The user who submitted this time off request. The id may be 'System' if it was an automated process
	SubmittedBy *Userreference `json:"submittedBy,omitempty"`

	// SubmittedDate - The timestamp when this request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SubmittedDate *time.Time `json:"submittedDate,omitempty"`

	// ReviewedBy - The user who reviewed this time off request. The id may be 'System' if it was an automated process
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`

	// ReviewedDate - The timestamp when this request was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// ModifiedBy - The user who last modified this TimeOffRequestResponse
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// ModifiedDate - The timestamp when this request was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// SyncVersion - The sync version of this time off request for which the scheduled activity is associated
	SyncVersion *int `json:"syncVersion,omitempty"`

	// Metadata - The version metadata of the time off request
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffrequestresponse) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffrequestresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "SubmittedDate","ReviewedDate","ModifiedDate", }
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
	type Alias Timeoffrequestresponse
	
	SubmittedDate := new(string)
	if o.SubmittedDate != nil {
		
		*SubmittedDate = timeutil.Strftime(o.SubmittedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SubmittedDate = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		IsFullDayRequest *bool `json:"isFullDayRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Substatus *string `json:"substatus,omitempty"`
		
		PartialDayStartDateTimes *[]time.Time `json:"partialDayStartDateTimes,omitempty"`
		
		FullDayManagementUnitDates *[]string `json:"fullDayManagementUnitDates,omitempty"`
		
		DailyDurationMinutes *int `json:"dailyDurationMinutes,omitempty"`
		
		DurationMinutes *[]int `json:"durationMinutes,omitempty"`
		
		PayableMinutes *[]int `json:"payableMinutes,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		SubmittedBy *Userreference `json:"submittedBy,omitempty"`
		
		SubmittedDate *string `json:"submittedDate,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SyncVersion *int `json:"syncVersion,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		
		DurationMinutes: o.DurationMinutes,
		
		PayableMinutes: o.PayableMinutes,
		
		Notes: o.Notes,
		
		SubmittedBy: o.SubmittedBy,
		
		SubmittedDate: SubmittedDate,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		ModifiedDate: ModifiedDate,
		
		SyncVersion: o.SyncVersion,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffrequestresponse) UnmarshalJSON(b []byte) error {
	var TimeoffrequestresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffrequestresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if User, ok := TimeoffrequestresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IsFullDayRequest, ok := TimeoffrequestresponseMap["isFullDayRequest"].(bool); ok {
		o.IsFullDayRequest = &IsFullDayRequest
	}
    
	if MarkedAsRead, ok := TimeoffrequestresponseMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
    
	if ActivityCodeId, ok := TimeoffrequestresponseMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := TimeoffrequestresponseMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if Status, ok := TimeoffrequestresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Substatus, ok := TimeoffrequestresponseMap["substatus"].(string); ok {
		o.Substatus = &Substatus
	}
    
	if PartialDayStartDateTimes, ok := TimeoffrequestresponseMap["partialDayStartDateTimes"].([]interface{}); ok {
		PartialDayStartDateTimesString, _ := json.Marshal(PartialDayStartDateTimes)
		json.Unmarshal(PartialDayStartDateTimesString, &o.PartialDayStartDateTimes)
	}
	
	if FullDayManagementUnitDates, ok := TimeoffrequestresponseMap["fullDayManagementUnitDates"].([]interface{}); ok {
		FullDayManagementUnitDatesString, _ := json.Marshal(FullDayManagementUnitDates)
		json.Unmarshal(FullDayManagementUnitDatesString, &o.FullDayManagementUnitDates)
	}
	
	if DailyDurationMinutes, ok := TimeoffrequestresponseMap["dailyDurationMinutes"].(float64); ok {
		DailyDurationMinutesInt := int(DailyDurationMinutes)
		o.DailyDurationMinutes = &DailyDurationMinutesInt
	}
	
	if DurationMinutes, ok := TimeoffrequestresponseMap["durationMinutes"].([]interface{}); ok {
		DurationMinutesString, _ := json.Marshal(DurationMinutes)
		json.Unmarshal(DurationMinutesString, &o.DurationMinutes)
	}
	
	if PayableMinutes, ok := TimeoffrequestresponseMap["payableMinutes"].([]interface{}); ok {
		PayableMinutesString, _ := json.Marshal(PayableMinutes)
		json.Unmarshal(PayableMinutesString, &o.PayableMinutes)
	}
	
	if Notes, ok := TimeoffrequestresponseMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if SubmittedBy, ok := TimeoffrequestresponseMap["submittedBy"].(map[string]interface{}); ok {
		SubmittedByString, _ := json.Marshal(SubmittedBy)
		json.Unmarshal(SubmittedByString, &o.SubmittedBy)
	}
	
	if submittedDateString, ok := TimeoffrequestresponseMap["submittedDate"].(string); ok {
		SubmittedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", submittedDateString)
		o.SubmittedDate = &SubmittedDate
	}
	
	if ReviewedBy, ok := TimeoffrequestresponseMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := TimeoffrequestresponseMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if ModifiedBy, ok := TimeoffrequestresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if modifiedDateString, ok := TimeoffrequestresponseMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SyncVersion, ok := TimeoffrequestresponseMap["syncVersion"].(float64); ok {
		SyncVersionInt := int(SyncVersion)
		o.SyncVersion = &SyncVersionInt
	}
	
	if Metadata, ok := TimeoffrequestresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := TimeoffrequestresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
