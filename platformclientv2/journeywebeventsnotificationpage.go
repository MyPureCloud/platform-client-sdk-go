package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationpage
type Journeywebeventsnotificationpage struct { 
	// Url
	Url *string `json:"url,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// Fragment
	Fragment *string `json:"fragment,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`


	// Keywords
	Keywords *string `json:"keywords,omitempty"`


	// Lang
	Lang *string `json:"lang,omitempty"`


	// Pathname
	Pathname *string `json:"pathname,omitempty"`


	// QueryString
	QueryString *string `json:"queryString,omitempty"`


	// Breadcrumb
	Breadcrumb *[]string `json:"breadcrumb,omitempty"`

}

func (o *Journeywebeventsnotificationpage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationpage
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Fragment *string `json:"fragment,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		
		Keywords *string `json:"keywords,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		
		Pathname *string `json:"pathname,omitempty"`
		
		QueryString *string `json:"queryString,omitempty"`
		
		Breadcrumb *[]string `json:"breadcrumb,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		Title: o.Title,
		
		Domain: o.Domain,
		
		Fragment: o.Fragment,
		
		Hostname: o.Hostname,
		
		Keywords: o.Keywords,
		
		Lang: o.Lang,
		
		Pathname: o.Pathname,
		
		QueryString: o.QueryString,
		
		Breadcrumb: o.Breadcrumb,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationpage) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationpageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationpageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := JourneywebeventsnotificationpageMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Title, ok := JourneywebeventsnotificationpageMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Domain, ok := JourneywebeventsnotificationpageMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Fragment, ok := JourneywebeventsnotificationpageMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Hostname, ok := JourneywebeventsnotificationpageMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := JourneywebeventsnotificationpageMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Lang, ok := JourneywebeventsnotificationpageMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Pathname, ok := JourneywebeventsnotificationpageMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := JourneywebeventsnotificationpageMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Breadcrumb, ok := JourneywebeventsnotificationpageMap["breadcrumb"].([]interface{}); ok {
		BreadcrumbString, _ := json.Marshal(Breadcrumb)
		json.Unmarshal(BreadcrumbString, &o.Breadcrumb)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationpage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
