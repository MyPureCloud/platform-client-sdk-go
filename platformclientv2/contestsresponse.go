package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestsresponse
type Contestsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Division - The division for this performance profile associate to
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

	// Version - The Contest Version
	Version *int `json:"version,omitempty"`

	// CreatedBy - The creator of the contest
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// Profile - The performance profile
	Profile *Contestprofile `json:"profile,omitempty"`

	// Participants - The Contest's participants
	Participants *[]Userreference `json:"participants,omitempty"`

	// Status - The Contest status
	Status *string `json:"status,omitempty"`

	// ParticipantCount - The Number of participants in the contest
	ParticipantCount *int `json:"participantCount,omitempty"`

	// DateFinalized - The Contest's finalize datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinalized *time.Time `json:"dateFinalized,omitempty"`

	// DateCancelled - The Contest's cancelled datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCancelled *time.Time `json:"dateCancelled,omitempty"`

	// DateModified - The Contest's last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateScoresModified - The datetime the contest scores were last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateScoresModified *time.Time `json:"dateScoresModified,omitempty"`

	// Winners - The Contest Winners
	Winners *[]Contestwinners `json:"winners,omitempty"`

	// DisqualifiedAgents - The Contest's disqualified agents, returned when a contest is complete
	DisqualifiedAgents *[]Contestdisqualifiedagents `json:"disqualifiedAgents,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Contestsresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAnnounced","DateFinalized","DateCancelled","DateModified","DateScoresModified", }
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
	type Alias Contestsresponse
	
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
	
	DateFinalized := new(string)
	if o.DateFinalized != nil {
		
		*DateFinalized = timeutil.Strftime(o.DateFinalized, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateFinalized = nil
	}
	
	DateCancelled := new(string)
	if o.DateCancelled != nil {
		
		*DateCancelled = timeutil.Strftime(o.DateCancelled, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCancelled = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateScoresModified := new(string)
	if o.DateScoresModified != nil {
		
		*DateScoresModified = timeutil.Strftime(o.DateScoresModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateScoresModified = nil
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
		
		Version *int `json:"version,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		Profile *Contestprofile `json:"profile,omitempty"`
		
		Participants *[]Userreference `json:"participants,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ParticipantCount *int `json:"participantCount,omitempty"`
		
		DateFinalized *string `json:"dateFinalized,omitempty"`
		
		DateCancelled *string `json:"dateCancelled,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateScoresModified *string `json:"dateScoresModified,omitempty"`
		
		Winners *[]Contestwinners `json:"winners,omitempty"`
		
		DisqualifiedAgents *[]Contestdisqualifiedagents `json:"disqualifiedAgents,omitempty"`
		
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
		
		Version: o.Version,
		
		CreatedBy: o.CreatedBy,
		
		Profile: o.Profile,
		
		Participants: o.Participants,
		
		Status: o.Status,
		
		ParticipantCount: o.ParticipantCount,
		
		DateFinalized: DateFinalized,
		
		DateCancelled: DateCancelled,
		
		DateModified: DateModified,
		
		DateScoresModified: DateScoresModified,
		
		Winners: o.Winners,
		
		DisqualifiedAgents: o.DisqualifiedAgents,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contestsresponse) UnmarshalJSON(b []byte) error {
	var ContestsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ContestsresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContestsresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ContestsresponseMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Title, ok := ContestsresponseMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ContestsresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateStartString, ok := ContestsresponseMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := ContestsresponseMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if WinningCriteria, ok := ContestsresponseMap["winningCriteria"].(string); ok {
		o.WinningCriteria = &WinningCriteria
	}
    
	if dateAnnouncedString, ok := ContestsresponseMap["dateAnnounced"].(string); ok {
		DateAnnounced, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAnnouncedString)
		o.DateAnnounced = &DateAnnounced
	}
	
	if AnnouncementTimezone, ok := ContestsresponseMap["announcementTimezone"].(string); ok {
		o.AnnouncementTimezone = &AnnouncementTimezone
	}
    
	if Anonymization, ok := ContestsresponseMap["anonymization"].(string); ok {
		o.Anonymization = &Anonymization
	}
    
	if Metrics, ok := ContestsresponseMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Prizes, ok := ContestsresponseMap["prizes"].([]interface{}); ok {
		PrizesString, _ := json.Marshal(Prizes)
		json.Unmarshal(PrizesString, &o.Prizes)
	}
	
	if Version, ok := ContestsresponseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if CreatedBy, ok := ContestsresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Profile, ok := ContestsresponseMap["profile"].(map[string]interface{}); ok {
		ProfileString, _ := json.Marshal(Profile)
		json.Unmarshal(ProfileString, &o.Profile)
	}
	
	if Participants, ok := ContestsresponseMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if Status, ok := ContestsresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ParticipantCount, ok := ContestsresponseMap["participantCount"].(float64); ok {
		ParticipantCountInt := int(ParticipantCount)
		o.ParticipantCount = &ParticipantCountInt
	}
	
	if dateFinalizedString, ok := ContestsresponseMap["dateFinalized"].(string); ok {
		DateFinalized, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateFinalizedString)
		o.DateFinalized = &DateFinalized
	}
	
	if dateCancelledString, ok := ContestsresponseMap["dateCancelled"].(string); ok {
		DateCancelled, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCancelledString)
		o.DateCancelled = &DateCancelled
	}
	
	if dateModifiedString, ok := ContestsresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateScoresModifiedString, ok := ContestsresponseMap["dateScoresModified"].(string); ok {
		DateScoresModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateScoresModifiedString)
		o.DateScoresModified = &DateScoresModified
	}
	
	if Winners, ok := ContestsresponseMap["winners"].([]interface{}); ok {
		WinnersString, _ := json.Marshal(Winners)
		json.Unmarshal(WinnersString, &o.Winners)
	}
	
	if DisqualifiedAgents, ok := ContestsresponseMap["disqualifiedAgents"].([]interface{}); ok {
		DisqualifiedAgentsString, _ := json.Marshal(DisqualifiedAgents)
		json.Unmarshal(DisqualifiedAgentsString, &o.DisqualifiedAgents)
	}
	
	if SelfUri, ok := ContestsresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contestsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
