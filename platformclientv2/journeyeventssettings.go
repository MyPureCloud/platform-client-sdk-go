package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyeventssettings - Settings concerning journey events
type Journeyeventssettings struct { 
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

// String returns a JSON representation of the model
func (o *Journeyeventssettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
