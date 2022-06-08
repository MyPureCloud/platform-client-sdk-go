package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationpage
type Journeysessioneventsnotificationpage struct { 
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

func (o *Journeysessioneventsnotificationpage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationpage
	
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

func (o *Journeysessioneventsnotificationpage) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationpageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationpageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := JourneysessioneventsnotificationpageMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Title, ok := JourneysessioneventsnotificationpageMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Domain, ok := JourneysessioneventsnotificationpageMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Fragment, ok := JourneysessioneventsnotificationpageMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Hostname, ok := JourneysessioneventsnotificationpageMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := JourneysessioneventsnotificationpageMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Lang, ok := JourneysessioneventsnotificationpageMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Pathname, ok := JourneysessioneventsnotificationpageMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := JourneysessioneventsnotificationpageMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Breadcrumb, ok := JourneysessioneventsnotificationpageMap["breadcrumb"].([]interface{}); ok {
		BreadcrumbString, _ := json.Marshal(Breadcrumb)
		json.Unmarshal(BreadcrumbString, &o.Breadcrumb)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationpage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
