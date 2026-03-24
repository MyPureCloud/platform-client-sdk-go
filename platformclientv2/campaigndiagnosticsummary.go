package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigndiagnosticsummary
type Campaigndiagnosticsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CampaignId - Campaign ID
	CampaignId *string `json:"campaignId,omitempty"`

	// DateStart - Start of the interval. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - End of the interval. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`

	// CampaignStates - Array of campaign states
	CampaignStates *[]Campaigndiagnosticcampaignstate `json:"campaignStates,omitempty"`

	// CampaignInfo - Array of diagnostic windows
	CampaignInfo *[]Campaigndiagnosticwindow `json:"campaignInfo,omitempty"`

	// CampaignHealthStates - Array of campaign health states
	CampaignHealthStates *[]Campaigndiagnosticcampaignhealthstate `json:"campaignHealthStates,omitempty"`

	// ConfigChanges - Configuration changes occurring within the time window
	ConfigChanges *[]Campaigndiagnosticconfigchange `json:"configChanges,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaigndiagnosticsummary) SetField(field string, fieldValue interface{}) {
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

func (o Campaigndiagnosticsummary) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd", }
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
	type Alias Campaigndiagnosticsummary
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		CampaignId *string `json:"campaignId,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		CampaignStates *[]Campaigndiagnosticcampaignstate `json:"campaignStates,omitempty"`
		
		CampaignInfo *[]Campaigndiagnosticwindow `json:"campaignInfo,omitempty"`
		
		CampaignHealthStates *[]Campaigndiagnosticcampaignhealthstate `json:"campaignHealthStates,omitempty"`
		
		ConfigChanges *[]Campaigndiagnosticconfigchange `json:"configChanges,omitempty"`
		Alias
	}{ 
		CampaignId: o.CampaignId,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		CampaignStates: o.CampaignStates,
		
		CampaignInfo: o.CampaignInfo,
		
		CampaignHealthStates: o.CampaignHealthStates,
		
		ConfigChanges: o.ConfigChanges,
		Alias:    (Alias)(o),
	})
}

func (o *Campaigndiagnosticsummary) UnmarshalJSON(b []byte) error {
	var CampaigndiagnosticsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigndiagnosticsummaryMap)
	if err != nil {
		return err
	}
	
	if CampaignId, ok := CampaigndiagnosticsummaryMap["campaignId"].(string); ok {
		o.CampaignId = &CampaignId
	}
    
	if dateStartString, ok := CampaigndiagnosticsummaryMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := CampaigndiagnosticsummaryMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if CampaignStates, ok := CampaigndiagnosticsummaryMap["campaignStates"].([]interface{}); ok {
		CampaignStatesString, _ := json.Marshal(CampaignStates)
		json.Unmarshal(CampaignStatesString, &o.CampaignStates)
	}
	
	if CampaignInfo, ok := CampaigndiagnosticsummaryMap["campaignInfo"].([]interface{}); ok {
		CampaignInfoString, _ := json.Marshal(CampaignInfo)
		json.Unmarshal(CampaignInfoString, &o.CampaignInfo)
	}
	
	if CampaignHealthStates, ok := CampaigndiagnosticsummaryMap["campaignHealthStates"].([]interface{}); ok {
		CampaignHealthStatesString, _ := json.Marshal(CampaignHealthStates)
		json.Unmarshal(CampaignHealthStatesString, &o.CampaignHealthStates)
	}
	
	if ConfigChanges, ok := CampaigndiagnosticsummaryMap["configChanges"].([]interface{}); ok {
		ConfigChangesString, _ := json.Marshal(ConfigChanges)
		json.Unmarshal(ConfigChangesString, &o.ConfigChanges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigndiagnosticsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
