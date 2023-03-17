package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Authorizationsettings
type Authorizationsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// AnalysisEnabled - Boolean showing if organization is opted in or not to unused role/perm analysis
	AnalysisEnabled *bool `json:"analysisEnabled,omitempty"`

	// AnalysisDays - Integer number of days to analyze user usage
	AnalysisDays *int `json:"analysisDays,omitempty"`

	// DateLastCalculated - The date and time of the most recent unused role calculation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLastCalculated *time.Time `json:"dateLastCalculated,omitempty"`

	// DateLastActive - The date of the most recent org activity used for analysis. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateLastActive *time.Time `json:"dateLastActive,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Authorizationsettings) SetField(field string, fieldValue interface{}) {
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

func (o Authorizationsettings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastCalculated", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateLastActive", }

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
	type Alias Authorizationsettings
	
	DateLastCalculated := new(string)
	if o.DateLastCalculated != nil {
		
		*DateLastCalculated = timeutil.Strftime(o.DateLastCalculated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastCalculated = nil
	}
	
	DateLastActive := new(string)
	if o.DateLastActive != nil {
		*DateLastActive = timeutil.Strftime(o.DateLastActive, "%Y-%m-%d")
	} else {
		DateLastActive = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AnalysisEnabled *bool `json:"analysisEnabled,omitempty"`
		
		AnalysisDays *int `json:"analysisDays,omitempty"`
		
		DateLastCalculated *string `json:"dateLastCalculated,omitempty"`
		
		DateLastActive *string `json:"dateLastActive,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		AnalysisEnabled: o.AnalysisEnabled,
		
		AnalysisDays: o.AnalysisDays,
		
		DateLastCalculated: DateLastCalculated,
		
		DateLastActive: DateLastActive,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Authorizationsettings) UnmarshalJSON(b []byte) error {
	var AuthorizationsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AuthorizationsettingsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthorizationsettingsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if AnalysisEnabled, ok := AuthorizationsettingsMap["analysisEnabled"].(bool); ok {
		o.AnalysisEnabled = &AnalysisEnabled
	}
    
	if AnalysisDays, ok := AuthorizationsettingsMap["analysisDays"].(float64); ok {
		AnalysisDaysInt := int(AnalysisDays)
		o.AnalysisDays = &AnalysisDaysInt
	}
	
	if dateLastCalculatedString, ok := AuthorizationsettingsMap["dateLastCalculated"].(string); ok {
		DateLastCalculated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastCalculatedString)
		o.DateLastCalculated = &DateLastCalculated
	}
	
	if dateLastActiveString, ok := AuthorizationsettingsMap["dateLastActive"].(string); ok {
		DateLastActive, _ := time.Parse("2006-01-02", dateLastActiveString)
		o.DateLastActive = &DateLastActive
	}
	
	if SelfUri, ok := AuthorizationsettingsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Authorizationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
