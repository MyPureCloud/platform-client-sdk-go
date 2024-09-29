package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyeventssettings - Settings concerning journey events
type Journeyeventssettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether or not journey event collection is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// ExcludedQueryParameters - List of parameters to be excluded from the query string.
	ExcludedQueryParameters *[]string `json:"excludedQueryParameters,omitempty"`

	// ShouldKeepUrlFragment - Whether or not to keep the URL fragment.
	ShouldKeepUrlFragment *bool `json:"shouldKeepUrlFragment,omitempty"`

	// SearchQueryParameters - List of query parameters used for search (e.g. 'q').
	SearchQueryParameters *[]string `json:"searchQueryParameters,omitempty"`

	// PageviewConfig - Controls how the pageview events are tracked.
	PageviewConfig *string `json:"pageviewConfig,omitempty"`

	// ClickEvents - Tracks when and where a visitor clicks on a webpage.
	ClickEvents *[]Selectoreventtrigger `json:"clickEvents,omitempty"`

	// FormsTrackEvents - Controls how the form submitted and form abandoned events are tracked after a visitor interacts with a form element.
	FormsTrackEvents *[]Formstracktrigger `json:"formsTrackEvents,omitempty"`

	// IdleEvents - Tracks when and where a visitor becomes inactive on a webpage.
	IdleEvents *[]Idleeventtrigger `json:"idleEvents,omitempty"`

	// InViewportEvents - Tracks when elements become visible or hidden on screen.
	InViewportEvents *[]Selectoreventtrigger `json:"inViewportEvents,omitempty"`

	// ScrollDepthEvents - Tracks when a visitor scrolls to a specific percentage of a webpage.
	ScrollDepthEvents *[]Scrollpercentageeventtrigger `json:"scrollDepthEvents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyeventssettings) SetField(field string, fieldValue interface{}) {
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

func (o Journeyeventssettings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Journeyeventssettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		ExcludedQueryParameters *[]string `json:"excludedQueryParameters,omitempty"`
		
		ShouldKeepUrlFragment *bool `json:"shouldKeepUrlFragment,omitempty"`
		
		SearchQueryParameters *[]string `json:"searchQueryParameters,omitempty"`
		
		PageviewConfig *string `json:"pageviewConfig,omitempty"`
		
		ClickEvents *[]Selectoreventtrigger `json:"clickEvents,omitempty"`
		
		FormsTrackEvents *[]Formstracktrigger `json:"formsTrackEvents,omitempty"`
		
		IdleEvents *[]Idleeventtrigger `json:"idleEvents,omitempty"`
		
		InViewportEvents *[]Selectoreventtrigger `json:"inViewportEvents,omitempty"`
		
		ScrollDepthEvents *[]Scrollpercentageeventtrigger `json:"scrollDepthEvents,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		ExcludedQueryParameters: o.ExcludedQueryParameters,
		
		ShouldKeepUrlFragment: o.ShouldKeepUrlFragment,
		
		SearchQueryParameters: o.SearchQueryParameters,
		
		PageviewConfig: o.PageviewConfig,
		
		ClickEvents: o.ClickEvents,
		
		FormsTrackEvents: o.FormsTrackEvents,
		
		IdleEvents: o.IdleEvents,
		
		InViewportEvents: o.InViewportEvents,
		
		ScrollDepthEvents: o.ScrollDepthEvents,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyeventssettings) UnmarshalJSON(b []byte) error {
	var JourneyeventssettingsMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyeventssettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := JourneyeventssettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if ExcludedQueryParameters, ok := JourneyeventssettingsMap["excludedQueryParameters"].([]interface{}); ok {
		ExcludedQueryParametersString, _ := json.Marshal(ExcludedQueryParameters)
		json.Unmarshal(ExcludedQueryParametersString, &o.ExcludedQueryParameters)
	}
	
	if ShouldKeepUrlFragment, ok := JourneyeventssettingsMap["shouldKeepUrlFragment"].(bool); ok {
		o.ShouldKeepUrlFragment = &ShouldKeepUrlFragment
	}
    
	if SearchQueryParameters, ok := JourneyeventssettingsMap["searchQueryParameters"].([]interface{}); ok {
		SearchQueryParametersString, _ := json.Marshal(SearchQueryParameters)
		json.Unmarshal(SearchQueryParametersString, &o.SearchQueryParameters)
	}
	
	if PageviewConfig, ok := JourneyeventssettingsMap["pageviewConfig"].(string); ok {
		o.PageviewConfig = &PageviewConfig
	}
    
	if ClickEvents, ok := JourneyeventssettingsMap["clickEvents"].([]interface{}); ok {
		ClickEventsString, _ := json.Marshal(ClickEvents)
		json.Unmarshal(ClickEventsString, &o.ClickEvents)
	}
	
	if FormsTrackEvents, ok := JourneyeventssettingsMap["formsTrackEvents"].([]interface{}); ok {
		FormsTrackEventsString, _ := json.Marshal(FormsTrackEvents)
		json.Unmarshal(FormsTrackEventsString, &o.FormsTrackEvents)
	}
	
	if IdleEvents, ok := JourneyeventssettingsMap["idleEvents"].([]interface{}); ok {
		IdleEventsString, _ := json.Marshal(IdleEvents)
		json.Unmarshal(IdleEventsString, &o.IdleEvents)
	}
	
	if InViewportEvents, ok := JourneyeventssettingsMap["inViewportEvents"].([]interface{}); ok {
		InViewportEventsString, _ := json.Marshal(InViewportEvents)
		json.Unmarshal(InViewportEventsString, &o.InViewportEvents)
	}
	
	if ScrollDepthEvents, ok := JourneyeventssettingsMap["scrollDepthEvents"].([]interface{}); ok {
		ScrollDepthEventsString, _ := json.Marshal(ScrollDepthEvents)
		json.Unmarshal(ScrollDepthEventsString, &o.ScrollDepthEvents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyeventssettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}