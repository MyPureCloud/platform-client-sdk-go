package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulebid
type Schedulebid struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the schedule bid
	Id *string `json:"id,omitempty"`

	// Name - The name of the schedule bid
	Name *string `json:"name,omitempty"`

	// BidWindowStartDate - The bid start date when agents can start participating in schedule bidding relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowStartDate *time.Time `json:"bidWindowStartDate,omitempty"`

	// BidWindowEndDate - The bid end date relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowEndDate *time.Time `json:"bidWindowEndDate,omitempty"`

	// EffectiveDate - The date when schedule sets would be effective for schedule generation relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// Status - The state of the bid
	Status *string `json:"status,omitempty"`

	// BidType - The type of the bid
	BidType *string `json:"bidType,omitempty"`

	// Forecast - The selected forecast used for schedule set generation for this bid
	Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`

	// WeeksToSchedule - The number of weeks to generate schedule sets through this bid
	WeeksToSchedule *int `json:"weeksToSchedule,omitempty"`

	// EndOverridesAndRotations - If true, all existing overrides, work plan rotations will be ended one day before effective date of this bid
	EndOverridesAndRotations *bool `json:"endOverridesAndRotations,omitempty"`

	// AgentRankingType - The type of agent ranking selected for this bid
	AgentRankingType *string `json:"agentRankingType,omitempty"`

	// RankingTiebreakerType - Ranking tiebreaker
	RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`

	// PublishedDate - The date the schedule bid is published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`

	// EndDate - The end date until which schedule sets can be used for schedule generation. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndDate *time.Time `json:"endDate,omitempty"`

	// Metadata - The metadata of this bid
	Metadata *Workplanbidmetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulebid) SetField(field string, fieldValue interface{}) {
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

func (o Schedulebid) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "PublishedDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BidWindowStartDate","BidWindowEndDate","EffectiveDate","EndDate", }

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
	type Alias Schedulebid
	
	BidWindowStartDate := new(string)
	if o.BidWindowStartDate != nil {
		*BidWindowStartDate = timeutil.Strftime(o.BidWindowStartDate, "%Y-%m-%d")
	} else {
		BidWindowStartDate = nil
	}
	
	BidWindowEndDate := new(string)
	if o.BidWindowEndDate != nil {
		*BidWindowEndDate = timeutil.Strftime(o.BidWindowEndDate, "%Y-%m-%d")
	} else {
		BidWindowEndDate = nil
	}
	
	EffectiveDate := new(string)
	if o.EffectiveDate != nil {
		*EffectiveDate = timeutil.Strftime(o.EffectiveDate, "%Y-%m-%d")
	} else {
		EffectiveDate = nil
	}
	
	PublishedDate := new(string)
	if o.PublishedDate != nil {
		
		*PublishedDate = timeutil.Strftime(o.PublishedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishedDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%d")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		BidWindowStartDate *string `json:"bidWindowStartDate,omitempty"`
		
		BidWindowEndDate *string `json:"bidWindowEndDate,omitempty"`
		
		EffectiveDate *string `json:"effectiveDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		BidType *string `json:"bidType,omitempty"`
		
		Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`
		
		WeeksToSchedule *int `json:"weeksToSchedule,omitempty"`
		
		EndOverridesAndRotations *bool `json:"endOverridesAndRotations,omitempty"`
		
		AgentRankingType *string `json:"agentRankingType,omitempty"`
		
		RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Metadata *Workplanbidmetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		BidWindowStartDate: BidWindowStartDate,
		
		BidWindowEndDate: BidWindowEndDate,
		
		EffectiveDate: EffectiveDate,
		
		Status: o.Status,
		
		BidType: o.BidType,
		
		Forecast: o.Forecast,
		
		WeeksToSchedule: o.WeeksToSchedule,
		
		EndOverridesAndRotations: o.EndOverridesAndRotations,
		
		AgentRankingType: o.AgentRankingType,
		
		RankingTiebreakerType: o.RankingTiebreakerType,
		
		PublishedDate: PublishedDate,
		
		EndDate: EndDate,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulebid) UnmarshalJSON(b []byte) error {
	var SchedulebidMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulebidMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SchedulebidMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SchedulebidMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if bidWindowStartDateString, ok := SchedulebidMap["bidWindowStartDate"].(string); ok {
		BidWindowStartDate, _ := time.Parse("2006-01-02", bidWindowStartDateString)
		o.BidWindowStartDate = &BidWindowStartDate
	}
	
	if bidWindowEndDateString, ok := SchedulebidMap["bidWindowEndDate"].(string); ok {
		BidWindowEndDate, _ := time.Parse("2006-01-02", bidWindowEndDateString)
		o.BidWindowEndDate = &BidWindowEndDate
	}
	
	if effectiveDateString, ok := SchedulebidMap["effectiveDate"].(string); ok {
		EffectiveDate, _ := time.Parse("2006-01-02", effectiveDateString)
		o.EffectiveDate = &EffectiveDate
	}
	
	if Status, ok := SchedulebidMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if BidType, ok := SchedulebidMap["bidType"].(string); ok {
		o.BidType = &BidType
	}
    
	if Forecast, ok := SchedulebidMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if WeeksToSchedule, ok := SchedulebidMap["weeksToSchedule"].(float64); ok {
		WeeksToScheduleInt := int(WeeksToSchedule)
		o.WeeksToSchedule = &WeeksToScheduleInt
	}
	
	if EndOverridesAndRotations, ok := SchedulebidMap["endOverridesAndRotations"].(bool); ok {
		o.EndOverridesAndRotations = &EndOverridesAndRotations
	}
    
	if AgentRankingType, ok := SchedulebidMap["agentRankingType"].(string); ok {
		o.AgentRankingType = &AgentRankingType
	}
    
	if RankingTiebreakerType, ok := SchedulebidMap["rankingTiebreakerType"].(string); ok {
		o.RankingTiebreakerType = &RankingTiebreakerType
	}
    
	if publishedDateString, ok := SchedulebidMap["publishedDate"].(string); ok {
		PublishedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishedDateString)
		o.PublishedDate = &PublishedDate
	}
	
	if endDateString, ok := SchedulebidMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02", endDateString)
		o.EndDate = &EndDate
	}
	
	if Metadata, ok := SchedulebidMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := SchedulebidMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulebid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
