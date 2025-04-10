package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestsessentials
type Contestsessentials struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Title - The Contest title
	Title *string `json:"title,omitempty"`

	// Status - The Contest status
	Status *string `json:"status,omitempty"`

	// DateStart - Start date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - End date of the contest. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// Profile - The performance profile
	Profile *Contestprofile `json:"profile,omitempty"`

	// ParticipantCount - The Number of participants in the contest
	ParticipantCount *int `json:"participantCount,omitempty"`

	// DateAnnounced - The Contest's Announcement datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAnnounced *time.Time `json:"dateAnnounced,omitempty"`

	// DateFinalized - The Contest's finalize datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateFinalized *time.Time `json:"dateFinalized,omitempty"`

	// DateCancelled - The Contest's cancelled datetime, returned when a contest is complete. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCancelled *time.Time `json:"dateCancelled,omitempty"`

	// DateModified - The Contest's last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateScoresModified - The datetime the contest scores were last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateScoresModified *time.Time `json:"dateScoresModified,omitempty"`

	// Metrics - The Contest's Metrics
	Metrics *[]Contestmetrics `json:"metrics,omitempty"`

	// RequestingParticipantContestInfo - The Most Recent Contest Info for the requesting participant
	RequestingParticipantContestInfo *Contestrequesingparticipantdailyinfo `json:"requestingParticipantContestInfo,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestsessentials) SetField(field string, fieldValue interface{}) {
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

func (o Contestsessentials) MarshalJSON() ([]byte, error) {
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
	type Alias Contestsessentials
	
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
		
		Title *string `json:"title,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		Profile *Contestprofile `json:"profile,omitempty"`
		
		ParticipantCount *int `json:"participantCount,omitempty"`
		
		DateAnnounced *string `json:"dateAnnounced,omitempty"`
		
		DateFinalized *string `json:"dateFinalized,omitempty"`
		
		DateCancelled *string `json:"dateCancelled,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateScoresModified *string `json:"dateScoresModified,omitempty"`
		
		Metrics *[]Contestmetrics `json:"metrics,omitempty"`
		
		RequestingParticipantContestInfo *Contestrequesingparticipantdailyinfo `json:"requestingParticipantContestInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Status: o.Status,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		Profile: o.Profile,
		
		ParticipantCount: o.ParticipantCount,
		
		DateAnnounced: DateAnnounced,
		
		DateFinalized: DateFinalized,
		
		DateCancelled: DateCancelled,
		
		DateModified: DateModified,
		
		DateScoresModified: DateScoresModified,
		
		Metrics: o.Metrics,
		
		RequestingParticipantContestInfo: o.RequestingParticipantContestInfo,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contestsessentials) UnmarshalJSON(b []byte) error {
	var ContestsessentialsMap map[string]interface{}
	err := json.Unmarshal(b, &ContestsessentialsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContestsessentialsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := ContestsessentialsMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Status, ok := ContestsessentialsMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateStartString, ok := ContestsessentialsMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := ContestsessentialsMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if Profile, ok := ContestsessentialsMap["profile"].(map[string]interface{}); ok {
		ProfileString, _ := json.Marshal(Profile)
		json.Unmarshal(ProfileString, &o.Profile)
	}
	
	if ParticipantCount, ok := ContestsessentialsMap["participantCount"].(float64); ok {
		ParticipantCountInt := int(ParticipantCount)
		o.ParticipantCount = &ParticipantCountInt
	}
	
	if dateAnnouncedString, ok := ContestsessentialsMap["dateAnnounced"].(string); ok {
		DateAnnounced, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAnnouncedString)
		o.DateAnnounced = &DateAnnounced
	}
	
	if dateFinalizedString, ok := ContestsessentialsMap["dateFinalized"].(string); ok {
		DateFinalized, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateFinalizedString)
		o.DateFinalized = &DateFinalized
	}
	
	if dateCancelledString, ok := ContestsessentialsMap["dateCancelled"].(string); ok {
		DateCancelled, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCancelledString)
		o.DateCancelled = &DateCancelled
	}
	
	if dateModifiedString, ok := ContestsessentialsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateScoresModifiedString, ok := ContestsessentialsMap["dateScoresModified"].(string); ok {
		DateScoresModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateScoresModifiedString)
		o.DateScoresModified = &DateScoresModified
	}
	
	if Metrics, ok := ContestsessentialsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if RequestingParticipantContestInfo, ok := ContestsessentialsMap["requestingParticipantContestInfo"].(map[string]interface{}); ok {
		RequestingParticipantContestInfoString, _ := json.Marshal(RequestingParticipantContestInfo)
		json.Unmarshal(RequestingParticipantContestInfoString, &o.RequestingParticipantContestInfo)
	}
	
	if SelfUri, ok := ContestsessentialsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contestsessentials) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
