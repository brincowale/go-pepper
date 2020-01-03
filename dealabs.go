package dealabs

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Dealabs struct {
	Client_key    string
	Client_secret string
	headers       map[string]string
	httpClient    *http.Client
}

func New() *Dealabs {
	client_key := "539f008401dbb"
	client_secret := "539f008401e9c"
	headers := make(map[string]string)
	headers["User-Agent"] = "com.dealabs.apps.android ANDROID [v5.18.03] [22 | SM-G930K] [@2.0x]"
	headers["Pepper-Include-Counters"] = "unread_alerts"
	headers["Pepper-Include-Prev-And-Next-Ids"] = "true"
	headers["Pepper-JSON-Format"] = "thread=list,group=ids,type=light,event=light,user=full,badge=user,formatted_text=html,message=with_code"
	headers["Pepper-Hardware-Id"] = "5bce296a65215d0bb3b9751bb77b0a1d"
	headers["Host"] = "www.dealabs.com"
	config := oauth1.NewConfig(client_key, client_secret)
	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       "",
		TokenSecret: "",
	})
	return &Dealabs{client_key, client_secret, headers, httpClient}
}

func (d *Dealabs) GetHotDeals(paramsOverride map[string]string) *Deals {
	// Params
	path := "https://www.dealabs.com/rest_api/v2/thread"
	params := make(map[string]string)
	params["order_by"] = "hot"
	params["limit"] = "50"
	q := url.Values{}
	for k, v := range paramsOverride {
		params[k] = v
	}
	for k, v := range params {
		q.Add(k, v)
	}
	req, err := http.NewRequest("GET", path, nil)
	// Http request
	req.URL.RawQuery = q.Encode()
	resp, _ := d.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// Parsing
	deals := Deals{}
	err = json.Unmarshal(body, &deals)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &deals
}

func (d *Dealabs) GetNewDeals(paramsOverride map[string]string) *Deals {
	// Params
	path := "https://www.dealabs.com/rest_api/v2/thread"
	params := make(map[string]string)
	params["order_by"] = "new"
	params["limit"] = "50"
	q := url.Values{}
	for k, v := range paramsOverride {
		params[k] = v
	}
	for k, v := range params {
		q.Add(k, v)
	}
	req, err := http.NewRequest("GET", path, nil)
	req.URL.RawQuery = q.Encode()
	// Http request

	resp, _ := d.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// Parsing
	deals := Deals{}
	err = json.Unmarshal(body, &deals)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &deals
}

func (deal *Data) PublishedAt() time.Time {
	if deal.IsHot {
		return time.Unix(deal.HotDate, 0)
	}
	return time.Unix(deal.Submitted, 0)
}