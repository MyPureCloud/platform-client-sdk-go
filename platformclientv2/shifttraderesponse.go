package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttraderesponse
type Shifttraderesponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of this shift trade
	Id *string `json:"id,omitempty"`

	// WeekDate - The start week date of the initiating shift in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

	// Schedule - A reference to the associated schedule
	Schedule *Buschedulereferenceformuroute `json:"schedule,omitempty"`

	// State - The state of this shift trade
	State *string `json:"state,omitempty"`

	// InitiatingUser - The user who initiated this trade
	InitiatingUser *Userreference `json:"initiatingUser,omitempty"`

	// InitiatingShiftId - The ID of the shift offered for trade by the initiating user
	InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`

	// InitiatingShiftStart - The start date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InitiatingShiftStart *time.Time `json:"initiatingShiftStart,omitempty"`

	// InitiatingShiftEnd - The end date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InitiatingShiftEnd *time.Time `json:"initiatingShiftEnd,omitempty"`

	// ReceivingWeekDate - The start week date of the receiving shift in yyyy-MM-dd format for a cross-week shift trade or null otherwise. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReceivingWeekDate *time.Time `json:"receivingWeekDate,omitempty"`

	// ReceivingUser - The user matching the trade, or if the state is not 'Matched', the user to whom the trade request was sent
	ReceivingUser *Userreference `json:"receivingUser,omitempty"`

	// ReceivingShiftId - The ID of the shift being exchanged for the initiating shift, null if the receiving user is picking up a shift
	ReceivingShiftId *string `json:"receivingShiftId,omitempty"`

	// ReceivingShiftStart - The start date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReceivingShiftStart *time.Time `json:"receivingShiftStart,omitempty"`

	// ReceivingShiftEnd - The end date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReceivingShiftEnd *time.Time `json:"receivingShiftEnd,omitempty"`

	// Expiration - When this shift trade offer will expire if not matched or approved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`

	// OneSided - Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
	OneSided *bool `json:"oneSided,omitempty"`

	// AcceptableIntervals - Time frames when the initiating user is willing to accept trades. Empty means giving up the shift. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`

	// ReviewedBy - The user who reviewed this shift trade. The id may be 'System' if it was an automated process
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`

	// ReviewedDate - The timestamp when this shift trade was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`

	// Metadata - Version data for this trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Shifttraderesponse) SetField(field string, fieldValue interface{}) {
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

func (o Shifttraderesponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "InitiatingShiftStart","InitiatingShiftEnd","ReceivingShiftStart","ReceivingShiftEnd","Expiration","ReviewedDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "WeekDate","ReceivingWeekDate", }

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
	type Alias Shifttraderesponse
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	InitiatingShiftStart := new(string)
	if o.InitiatingShiftStart != nil {
		
		*InitiatingShiftStart = timeutil.Strftime(o.InitiatingShiftStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftStart = nil
	}
	
	InitiatingShiftEnd := new(string)
	if o.InitiatingShiftEnd != nil {
		
		*InitiatingShiftEnd = timeutil.Strftime(o.InitiatingShiftEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InitiatingShiftEnd = nil
	}
	
	ReceivingWeekDate := new(string)
	if o.ReceivingWeekDate != nil {
		*ReceivingWeekDate = timeutil.Strftime(o.ReceivingWeekDate, "%Y-%m-%d")
	} else {
		ReceivingWeekDate = nil
	}
	
	ReceivingShiftStart := new(string)
	if o.ReceivingShiftStart != nil {
		
		*ReceivingShiftStart = timeutil.Strftime(o.ReceivingShiftStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReceivingShiftStart = nil
	}
	
	ReceivingShiftEnd := new(string)
	if o.ReceivingShiftEnd != nil {
		
		*ReceivingShiftEnd = timeutil.Strftime(o.ReceivingShiftEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReceivingShiftEnd = nil
	}
	
	Expiration := new(string)
	if o.Expiration != nil {
		
		*Expiration = timeutil.Strftime(o.Expiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expiration = nil
	}
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Schedule *Buschedulereferenceformuroute `json:"schedule,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitiatingUser *Userreference `json:"initiatingUser,omitempty"`
		
		InitiatingShiftId *string `json:"initiatingShiftId,omitempty"`
		
		InitiatingShiftStart *string `json:"initiatingShiftStart,omitempty"`
		
		InitiatingShiftEnd *string `json:"initiatingShiftEnd,omitempty"`
		
		ReceivingWeekDate *string `json:"receivingWeekDate,omitempty"`
		
		ReceivingUser *Userreference `json:"receivingUser,omitempty"`
		
		ReceivingShiftId *string `json:"receivingShiftId,omitempty"`
		
		ReceivingShiftStart *string `json:"receivingShiftStart,omitempty"`
		
		ReceivingShiftEnd *string `json:"receivingShiftEnd,omitempty"`
		
		Expiration *string `json:"expiration,omitempty"`
		
		OneSided *bool `json:"oneSided,omitempty"`
		
		AcceptableIntervals *[]string `json:"acceptableIntervals,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		Schedule: o.Schedule,
		
		State: o.State,
		
		InitiatingUser: o.InitiatingUser,
		
		InitiatingShiftId: o.InitiatingShiftId,
		
		InitiatingShiftStart: InitiatingShiftStart,
		
		InitiatingShiftEnd: InitiatingShiftEnd,
		
		ReceivingWeekDate: ReceivingWeekDate,
		
		ReceivingUser: o.ReceivingUser,
		
		ReceivingShiftId: o.ReceivingShiftId,
		
		ReceivingShiftStart: ReceivingShiftStart,
		
		ReceivingShiftEnd: ReceivingShiftEnd,
		
		Expiration: Expiration,
		
		OneSided: o.OneSided,
		
		AcceptableIntervals: o.AcceptableIntervals,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Shifttraderesponse) UnmarshalJSON(b []byte) error {
	var ShifttraderesponseMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttraderesponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ShifttraderesponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := ShifttraderesponseMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if Schedule, ok := ShifttraderesponseMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if State, ok := ShifttraderesponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitiatingUser, ok := ShifttraderesponseMap["initiatingUser"].(map[string]interface{}); ok {
		InitiatingUserString, _ := json.Marshal(InitiatingUser)
		json.Unmarshal(InitiatingUserString, &o.InitiatingUser)
	}
	
	if InitiatingShiftId, ok := ShifttraderesponseMap["initiatingShiftId"].(string); ok {
		o.InitiatingShiftId = &InitiatingShiftId
	}
    
	if initiatingShiftStartString, ok := ShifttraderesponseMap["initiatingShiftStart"].(string); ok {
		InitiatingShiftStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftStartString)
		o.InitiatingShiftStart = &InitiatingShiftStart
	}
	
	if initiatingShiftEndString, ok := ShifttraderesponseMap["initiatingShiftEnd"].(string); ok {
		InitiatingShiftEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", initiatingShiftEndString)
		o.InitiatingShiftEnd = &InitiatingShiftEnd
	}
	
	if receivingWeekDateString, ok := ShifttraderesponseMap["receivingWeekDate"].(string); ok {
		ReceivingWeekDate, _ := time.Parse("2006-01-02", receivingWeekDateString)
		o.ReceivingWeekDate = &ReceivingWeekDate
	}
	
	if ReceivingUser, ok := ShifttraderesponseMap["receivingUser"].(map[string]interface{}); ok {
		ReceivingUserString, _ := json.Marshal(ReceivingUser)
		json.Unmarshal(ReceivingUserString, &o.ReceivingUser)
	}
	
	if ReceivingShiftId, ok := ShifttraderesponseMap["receivingShiftId"].(string); ok {
		o.ReceivingShiftId = &ReceivingShiftId
	}
    
	if receivingShiftStartString, ok := ShifttraderesponseMap["receivingShiftStart"].(string); ok {
		ReceivingShiftStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftStartString)
		o.ReceivingShiftStart = &ReceivingShiftStart
	}
	
	if receivingShiftEndString, ok := ShifttraderesponseMap["receivingShiftEnd"].(string); ok {
		ReceivingShiftEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", receivingShiftEndString)
		o.ReceivingShiftEnd = &ReceivingShiftEnd
	}
	
	if expirationString, ok := ShifttraderesponseMap["expiration"].(string); ok {
		Expiration, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationString)
		o.Expiration = &Expiration
	}
	
	if OneSided, ok := ShifttraderesponseMap["oneSided"].(bool); ok {
		o.OneSided = &OneSided
	}
    
	if AcceptableIntervals, ok := ShifttraderesponseMap["acceptableIntervals"].([]interface{}); ok {
		AcceptableIntervalsString, _ := json.Marshal(AcceptableIntervals)
		json.Unmarshal(AcceptableIntervalsString, &o.AcceptableIntervals)
	}
	
	if ReviewedBy, ok := ShifttraderesponseMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := ShifttraderesponseMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if Metadata, ok := ShifttraderesponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
