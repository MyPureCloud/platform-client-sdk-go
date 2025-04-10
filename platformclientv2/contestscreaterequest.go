package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestscreaterequest
type Contestscreaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Division - The division for this performance profile associate to. Only set for DEFAULT profile.
	Division *Writabledivision `json:"division,omitempty"`

	// Title - The Contest title
	Title *string `json:"title,omitempty"`

	// Description - The Contest description
	Description *string `json:"description,omitempty"`

	// DateStart - Start date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - End date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// WinningCriteria - The Contest winning criteria
	WinningCriteria *string `json:"winningCriteria,omitempty"`

	// DateAnnounced - The Contest's Announcement Datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAnnounced *time.Time `json:"dateAnnounced,omitempty"`

	// AnnouncementTimezone - The Contest's Announcement Timezone. Valid values are strings of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	AnnouncementTimezone *string `json:"announcementTimezone,omitempty"`

	// Anonymization - The Contest anonymization
	Anonymization *string `json:"anonymization,omitempty"`

	// Metrics - The Contest's Metrics
	Metrics *[]Contestmetrics `json:"metrics,omitempty"`

	// Prizes - The Contest Prizes
	Prizes *[]Contestprizes `json:"prizes,omitempty"`

	// ProfileId - The Contest profile
	ProfileId *string `json:"profileId,omitempty"`

	// ParticipantIds - The Contest's participants
	ParticipantIds *[]string `json:"participantIds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestscreaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Contestscreaterequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAnnounced", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStart","DateEnd", }

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
	type Alias Contestscreaterequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	
	DateAnnounced := new(string)
	if o.DateAnnounced != nil {
		
		*DateAnnounced = timeutil.Strftime(o.DateAnnounced, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAnnounced = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		WinningCriteria *string `json:"winningCriteria,omitempty"`
		
		DateAnnounced *string `json:"dateAnnounced,omitempty"`
		
		AnnouncementTimezone *string `json:"announcementTimezone,omitempty"`
		
		Anonymization *string `json:"anonymization,omitempty"`
		
		Metrics *[]Contestmetrics `json:"metrics,omitempty"`
		
		Prizes *[]Contestprizes `json:"prizes,omitempty"`
		
		ProfileId *string `json:"profileId,omitempty"`
		
		ParticipantIds *[]string `json:"participantIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		Title: o.Title,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		WinningCriteria: o.WinningCriteria,
		
		DateAnnounced: DateAnnounced,
		
		AnnouncementTimezone: o.AnnouncementTimezone,
		
		Anonymization: o.Anonymization,
		
		Metrics: o.Metrics,
		
		Prizes: o.Prizes,
		
		ProfileId: o.ProfileId,
		
		ParticipantIds: o.ParticipantIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contestscreaterequest) UnmarshalJSON(b []byte) error {
	var ContestscreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContestscreaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContestscreaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ContestscreaterequestMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Title, ok := ContestscreaterequestMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ContestscreaterequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateStartString, ok := ContestscreaterequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := ContestscreaterequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if WinningCriteria, ok := ContestscreaterequestMap["winningCriteria"].(string); ok {
		o.WinningCriteria = &WinningCriteria
	}
    
	if dateAnnouncedString, ok := ContestscreaterequestMap["dateAnnounced"].(string); ok {
		DateAnnounced, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAnnouncedString)
		o.DateAnnounced = &DateAnnounced
	}
	
	if AnnouncementTimezone, ok := ContestscreaterequestMap["announcementTimezone"].(string); ok {
		o.AnnouncementTimezone = &AnnouncementTimezone
	}
    
	if Anonymization, ok := ContestscreaterequestMap["anonymization"].(string); ok {
		o.Anonymization = &Anonymization
	}
    
	if Metrics, ok := ContestscreaterequestMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Prizes, ok := ContestscreaterequestMap["prizes"].([]interface{}); ok {
		PrizesString, _ := json.Marshal(Prizes)
		json.Unmarshal(PrizesString, &o.Prizes)
	}
	
	if ProfileId, ok := ContestscreaterequestMap["profileId"].(string); ok {
		o.ProfileId = &ProfileId
	}
    
	if ParticipantIds, ok := ContestscreaterequestMap["participantIds"].([]interface{}); ok {
		ParticipantIdsString, _ := json.Marshal(ParticipantIds)
		json.Unmarshal(ParticipantIdsString, &o.ParticipantIds)
	}
	
	if SelfUri, ok := ContestscreaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contestscreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
