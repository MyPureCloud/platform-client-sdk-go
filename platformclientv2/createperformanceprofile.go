package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createperformanceprofile
type Createperformanceprofile struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - A name for this performance profile
	Name *string `json:"name,omitempty"`

	// Division - The associated division for this Performance Profile
	Division *Writabledivision `json:"division,omitempty"`

	// Description - A description about this performance profile
	Description *string `json:"description,omitempty"`

	// MetricOrders - Order of the associated metrics. The list should contain valid ids for metrics
	MetricOrders *[]string `json:"metricOrders,omitempty"`

	// DateCreated - Creation date for this performance profile. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ReportingIntervals - The reporting interval periods for this performance profile
	ReportingIntervals *[]Reportinginterval `json:"reportingIntervals,omitempty"`

	// Active - The flag for active profiles
	Active *bool `json:"active,omitempty"`

	// MemberCount - The number of members in this performance profile
	MemberCount *int `json:"memberCount,omitempty"`

	// MaxLeaderboardRankSize - The maximum rank size for the leaderboard. This counts the number of ranks can be retrieved in a leaderboard queries
	MaxLeaderboardRankSize *int `json:"maxLeaderboardRankSize,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createperformanceprofile) SetField(field string, fieldValue interface{}) {
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

func (o Createperformanceprofile) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Createperformanceprofile
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MetricOrders *[]string `json:"metricOrders,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ReportingIntervals *[]Reportinginterval `json:"reportingIntervals,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		MaxLeaderboardRankSize *int `json:"maxLeaderboardRankSize,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		MetricOrders: o.MetricOrders,
		
		DateCreated: DateCreated,
		
		ReportingIntervals: o.ReportingIntervals,
		
		Active: o.Active,
		
		MemberCount: o.MemberCount,
		
		MaxLeaderboardRankSize: o.MaxLeaderboardRankSize,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Createperformanceprofile) UnmarshalJSON(b []byte) error {
	var CreateperformanceprofileMap map[string]interface{}
	err := json.Unmarshal(b, &CreateperformanceprofileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateperformanceprofileMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CreateperformanceprofileMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := CreateperformanceprofileMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := CreateperformanceprofileMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if MetricOrders, ok := CreateperformanceprofileMap["metricOrders"].([]interface{}); ok {
		MetricOrdersString, _ := json.Marshal(MetricOrders)
		json.Unmarshal(MetricOrdersString, &o.MetricOrders)
	}
	
	if dateCreatedString, ok := CreateperformanceprofileMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ReportingIntervals, ok := CreateperformanceprofileMap["reportingIntervals"].([]interface{}); ok {
		ReportingIntervalsString, _ := json.Marshal(ReportingIntervals)
		json.Unmarshal(ReportingIntervalsString, &o.ReportingIntervals)
	}
	
	if Active, ok := CreateperformanceprofileMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if MemberCount, ok := CreateperformanceprofileMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if MaxLeaderboardRankSize, ok := CreateperformanceprofileMap["maxLeaderboardRankSize"].(float64); ok {
		MaxLeaderboardRankSizeInt := int(MaxLeaderboardRankSize)
		o.MaxLeaderboardRankSize = &MaxLeaderboardRankSizeInt
	}
	
	if SelfUri, ok := CreateperformanceprofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createperformanceprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
