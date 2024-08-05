package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alternativeshifttraderesponse
type Alternativeshifttraderesponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ShiftOfferJobId - The job ID of the alternative shift offer listing, from which the trade was chosen
	ShiftOfferJobId *string `json:"shiftOfferJobId,omitempty"`

	// ExistingShifts - The existing shifts from the offer, may be empty
	ExistingShifts *[]Alternativeshiftagentscheduledshift `json:"existingShifts,omitempty"`

	// OfferedShifts - The offered shifts from the offer, may be empty
	OfferedShifts *[]Alternativeshiftagentscheduledshift `json:"offeredShifts,omitempty"`

	// Schedule - The existing schedule information associated with the trade
	Schedule *Alternativeshiftschedulelookup `json:"schedule,omitempty"`

	// ManagementUnit - The management unit of this alternative shift trade request
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// User - The user who submitted the trade request
	User *Userreference `json:"user,omitempty"`

	// WeekDate - The start week date of the associated schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

	// ExpirationDate - The date when the trade will expire in ISO-8601 format. The trade cannot be approved after expiration
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// State - The state of this alternative shift trade
	State *string `json:"state,omitempty"`

	// ProcessingStatus - The processing status of this alternative shift trade
	ProcessingStatus *string `json:"processingStatus,omitempty"`

	// SystemDateReviewed - The timestamp of when the trade request was reviewed by the system in ISO-8601 format
	SystemDateReviewed *time.Time `json:"systemDateReviewed,omitempty"`

	// AdminDateReviewed - The timestamp of when the trade request was reviewed by an admin in ISO-8601 format
	AdminDateReviewed *time.Time `json:"adminDateReviewed,omitempty"`

	// AdminReviewedBy - The admin who reviewed this alternative shift trade after system denial
	AdminReviewedBy *Userreference `json:"adminReviewedBy,omitempty"`

	// Violations - A list of trade match violations
	Violations *[]string `json:"violations,omitempty"`

	// Metadata - Version metadata for this alternative shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alternativeshifttraderesponse) SetField(field string, fieldValue interface{}) {
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

func (o Alternativeshifttraderesponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ExpirationDate","SystemDateReviewed","AdminDateReviewed", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "WeekDate", }

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
	type Alias Alternativeshifttraderesponse
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	ExpirationDate := new(string)
	if o.ExpirationDate != nil {
		
		*ExpirationDate = timeutil.Strftime(o.ExpirationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExpirationDate = nil
	}
	
	SystemDateReviewed := new(string)
	if o.SystemDateReviewed != nil {
		
		*SystemDateReviewed = timeutil.Strftime(o.SystemDateReviewed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SystemDateReviewed = nil
	}
	
	AdminDateReviewed := new(string)
	if o.AdminDateReviewed != nil {
		
		*AdminDateReviewed = timeutil.Strftime(o.AdminDateReviewed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AdminDateReviewed = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ShiftOfferJobId *string `json:"shiftOfferJobId,omitempty"`
		
		ExistingShifts *[]Alternativeshiftagentscheduledshift `json:"existingShifts,omitempty"`
		
		OfferedShifts *[]Alternativeshiftagentscheduledshift `json:"offeredShifts,omitempty"`
		
		Schedule *Alternativeshiftschedulelookup `json:"schedule,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		ExpirationDate *string `json:"expirationDate,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ProcessingStatus *string `json:"processingStatus,omitempty"`
		
		SystemDateReviewed *string `json:"systemDateReviewed,omitempty"`
		
		AdminDateReviewed *string `json:"adminDateReviewed,omitempty"`
		
		AdminReviewedBy *Userreference `json:"adminReviewedBy,omitempty"`
		
		Violations *[]string `json:"violations,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ShiftOfferJobId: o.ShiftOfferJobId,
		
		ExistingShifts: o.ExistingShifts,
		
		OfferedShifts: o.OfferedShifts,
		
		Schedule: o.Schedule,
		
		ManagementUnit: o.ManagementUnit,
		
		User: o.User,
		
		WeekDate: WeekDate,
		
		ExpirationDate: ExpirationDate,
		
		State: o.State,
		
		ProcessingStatus: o.ProcessingStatus,
		
		SystemDateReviewed: SystemDateReviewed,
		
		AdminDateReviewed: AdminDateReviewed,
		
		AdminReviewedBy: o.AdminReviewedBy,
		
		Violations: o.Violations,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Alternativeshifttraderesponse) UnmarshalJSON(b []byte) error {
	var AlternativeshifttraderesponseMap map[string]interface{}
	err := json.Unmarshal(b, &AlternativeshifttraderesponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AlternativeshifttraderesponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ShiftOfferJobId, ok := AlternativeshifttraderesponseMap["shiftOfferJobId"].(string); ok {
		o.ShiftOfferJobId = &ShiftOfferJobId
	}
    
	if ExistingShifts, ok := AlternativeshifttraderesponseMap["existingShifts"].([]interface{}); ok {
		ExistingShiftsString, _ := json.Marshal(ExistingShifts)
		json.Unmarshal(ExistingShiftsString, &o.ExistingShifts)
	}
	
	if OfferedShifts, ok := AlternativeshifttraderesponseMap["offeredShifts"].([]interface{}); ok {
		OfferedShiftsString, _ := json.Marshal(OfferedShifts)
		json.Unmarshal(OfferedShiftsString, &o.OfferedShifts)
	}
	
	if Schedule, ok := AlternativeshifttraderesponseMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ManagementUnit, ok := AlternativeshifttraderesponseMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if User, ok := AlternativeshifttraderesponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if weekDateString, ok := AlternativeshifttraderesponseMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if expirationDateString, ok := AlternativeshifttraderesponseMap["expirationDate"].(string); ok {
		ExpirationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationDateString)
		o.ExpirationDate = &ExpirationDate
	}
	
	if State, ok := AlternativeshifttraderesponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if ProcessingStatus, ok := AlternativeshifttraderesponseMap["processingStatus"].(string); ok {
		o.ProcessingStatus = &ProcessingStatus
	}
    
	if systemDateReviewedString, ok := AlternativeshifttraderesponseMap["systemDateReviewed"].(string); ok {
		SystemDateReviewed, _ := time.Parse("2006-01-02T15:04:05.999999Z", systemDateReviewedString)
		o.SystemDateReviewed = &SystemDateReviewed
	}
	
	if adminDateReviewedString, ok := AlternativeshifttraderesponseMap["adminDateReviewed"].(string); ok {
		AdminDateReviewed, _ := time.Parse("2006-01-02T15:04:05.999999Z", adminDateReviewedString)
		o.AdminDateReviewed = &AdminDateReviewed
	}
	
	if AdminReviewedBy, ok := AlternativeshifttraderesponseMap["adminReviewedBy"].(map[string]interface{}); ok {
		AdminReviewedByString, _ := json.Marshal(AdminReviewedBy)
		json.Unmarshal(AdminReviewedByString, &o.AdminReviewedBy)
	}
	
	if Violations, ok := AlternativeshifttraderesponseMap["violations"].([]interface{}); ok {
		ViolationsString, _ := json.Marshal(Violations)
		json.Unmarshal(ViolationsString, &o.Violations)
	}
	
	if Metadata, ok := AlternativeshifttraderesponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := AlternativeshifttraderesponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Alternativeshifttraderesponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
