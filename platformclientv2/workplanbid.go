package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanbid - Work plan bid response
type Workplanbid struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the work plan bid
	Id *string `json:"id,omitempty"`

	// Name - The name of the work plan bid
	Name *string `json:"name,omitempty"`

	// Forecast - The selected forecast in this work plan bid
	Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`

	// BidWindowStartDate - The bid start date where agents start participate in work plan bidding. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowStartDate *time.Time `json:"bidWindowStartDate,omitempty"`

	// BidWindowEndDate - The bid end date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BidWindowEndDate *time.Time `json:"bidWindowEndDate,omitempty"`

	// EffectiveDate - The date when agents will be assigned to the new work plan. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// Status - The state of the bid
	Status *string `json:"status,omitempty"`

	// AgentRankingType - The type of agent ranking selected for this bid
	AgentRankingType *string `json:"agentRankingType,omitempty"`

	// RankingTiebreakerType - Ranking tiebreaker
	RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`

	// PublishedDate - The date the work plan bid published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`

	// WorkPlanFieldsVisibleToAgents - The work plan fields visible to agents whenever work plan preferences are made
	WorkPlanFieldsVisibleToAgents *[]string `json:"workPlanFieldsVisibleToAgents,omitempty"`

	// Metadata - The meta data of this bid
	Metadata *Workplanbidmetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanbid) SetField(field string, fieldValue interface{}) {
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

func (o Workplanbid) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "PublishedDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BidWindowStartDate","BidWindowEndDate","EffectiveDate", }

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
	type Alias Workplanbid
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Forecast *Bushorttermforecastweekreference `json:"forecast,omitempty"`
		
		BidWindowStartDate *string `json:"bidWindowStartDate,omitempty"`
		
		BidWindowEndDate *string `json:"bidWindowEndDate,omitempty"`
		
		EffectiveDate *string `json:"effectiveDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		AgentRankingType *string `json:"agentRankingType,omitempty"`
		
		RankingTiebreakerType *string `json:"rankingTiebreakerType,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		WorkPlanFieldsVisibleToAgents *[]string `json:"workPlanFieldsVisibleToAgents,omitempty"`
		
		Metadata *Workplanbidmetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Forecast: o.Forecast,
		
		BidWindowStartDate: BidWindowStartDate,
		
		BidWindowEndDate: BidWindowEndDate,
		
		EffectiveDate: EffectiveDate,
		
		Status: o.Status,
		
		AgentRankingType: o.AgentRankingType,
		
		RankingTiebreakerType: o.RankingTiebreakerType,
		
		PublishedDate: PublishedDate,
		
		WorkPlanFieldsVisibleToAgents: o.WorkPlanFieldsVisibleToAgents,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanbid) UnmarshalJSON(b []byte) error {
	var WorkplanbidMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanbidMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanbidMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkplanbidMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Forecast, ok := WorkplanbidMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if bidWindowStartDateString, ok := WorkplanbidMap["bidWindowStartDate"].(string); ok {
		BidWindowStartDate, _ := time.Parse("2006-01-02", bidWindowStartDateString)
		o.BidWindowStartDate = &BidWindowStartDate
	}
	
	if bidWindowEndDateString, ok := WorkplanbidMap["bidWindowEndDate"].(string); ok {
		BidWindowEndDate, _ := time.Parse("2006-01-02", bidWindowEndDateString)
		o.BidWindowEndDate = &BidWindowEndDate
	}
	
	if effectiveDateString, ok := WorkplanbidMap["effectiveDate"].(string); ok {
		EffectiveDate, _ := time.Parse("2006-01-02", effectiveDateString)
		o.EffectiveDate = &EffectiveDate
	}
	
	if Status, ok := WorkplanbidMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if AgentRankingType, ok := WorkplanbidMap["agentRankingType"].(string); ok {
		o.AgentRankingType = &AgentRankingType
	}
    
	if RankingTiebreakerType, ok := WorkplanbidMap["rankingTiebreakerType"].(string); ok {
		o.RankingTiebreakerType = &RankingTiebreakerType
	}
    
	if publishedDateString, ok := WorkplanbidMap["publishedDate"].(string); ok {
		PublishedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishedDateString)
		o.PublishedDate = &PublishedDate
	}
	
	if WorkPlanFieldsVisibleToAgents, ok := WorkplanbidMap["workPlanFieldsVisibleToAgents"].([]interface{}); ok {
		WorkPlanFieldsVisibleToAgentsString, _ := json.Marshal(WorkPlanFieldsVisibleToAgents)
		json.Unmarshal(WorkPlanFieldsVisibleToAgentsString, &o.WorkPlanFieldsVisibleToAgents)
	}
	
	if Metadata, ok := WorkplanbidMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := WorkplanbidMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanbid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
