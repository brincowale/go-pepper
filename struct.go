package dealabs

type Data struct {
	DealType          string  `json:"_type"`
	Representation    string  `json:"_representation"`
	ThreadID          int     `json:"thread_id"`
	Description       string  `json:"description"`
	Title             string  `json:"title"`
	Origin            string  `json:"origin"`
	Submitted         int     `json:"submitted"`
	Updated           int     `json:"updated"`
	Expired           bool    `json:"expired"`
	Status            string  `json:"status"`
	Local             bool    `json:"local"`
	LastCommented     int     `json:"last_commented"`
	CommentCount      int     `json:"comment_count"`
	GroupCount        int     `json:"group_count"`
	Price             float64 `json:"price,omitempty"`
	DealURI           string  `json:"deal_uri"`
	VisitURI          string  `json:"visit_uri,omitempty"`
	ShortURI          string  `json:"short_uri"`
	TemperatureRating int     `json:"temperature_rating"`
	TemperatureLevel  string  `json:"temperature_level"`
	IsHot             bool    `json:"is_hot"`
	Type              struct {
		Type              string `json:"_type"`
		Representation    string `json:"_representation"`
		TypeID            int    `json:"type_id"`
		Name              string `json:"name"`
		ShortName         string `json:"short_name"`
		DisplayOrder      int    `json:"display_order"`
		ThreadTypeURLName string `json:"thread_type_url_name"`
		Dealable          bool   `json:"dealable"`
		TypeUniversalID   string `json:"type_universal_id"`
		TopicFilterable   bool   `json:"topic_filterable"`
		Filterable        bool   `json:"filterable"`
	} `json:"type"`
	NextBestPrice        float64 `json:"next_best_price,omitempty"`
	PriceDiscount        int     `json:"price_discount,omitempty"`
	Flaggable            bool    `json:"flaggable"`
	Votable              bool    `json:"votable"`
	Editable             bool    `json:"editable"`
	GroupDisplaySummary  string  `json:"group_display_summary"`
	ShowBumpedStatus     bool    `json:"show_bumped_status"`
	DisplayPriceAsFree   bool    `json:"display_price_as_free,omitempty"`
	HasSuggestionCards   bool    `json:"has_suggestion_cards"`
	IconListURL          string  `json:"icon_list_url"`
	IconDetailURL        string  `json:"icon_detail_url"`
	HasSuggestedKeywords bool    `json:"has_suggested_keywords"`
	Clearance            bool    `json:"clearance"`
	Featured             int     `json:"featured"`
	Created              int     `json:"created"`
	IsNew                bool    `json:"is_new"`
	IsNsfw               bool    `json:"is_nsfw"`
	Image                Image   `json:"image"`
	MainGroup            struct {
		Type       string `json:"_type"`
		GroupID    int    `json:"group_id"`
		Name       string `json:"name"`
		URLName    string `json:"url_name"`
		Default    bool   `json:"default"`
		Followable bool   `json:"followable"`
		Followed   bool   `json:"followed"`
		Statistics struct {
			ActiveDealCount    int `json:"active_deal_count"`
			ActiveVoucherCount int `json:"active_voucher_count"`
			CommentCount       int `json:"comment_count"`
			ThreadCount        int `json:"thread_count"`
		} `json:"statistics"`
		SubmittableThreadTypes  []int    `json:"submittable_thread_types"`
		Sections                []string `json:"sections"`
		PepperURL               string   `json:"pepper_url"`
		HeaderDescription       string   `json:"header_description"`
		HeroBannerSmartphoneURL string   `json:"hero_banner_smartphone_url"`
		HeroBannerTabletURL     string   `json:"hero_banner_tablet_url"`
		IconListURL             string   `json:"icon_list_url"`
		IconDetailURL           string   `json:"icon_detail_url"`
		Image                   Image    `json:"image"`
	} `json:"main_group"`
	Merchant             Merchant `json:"merchant,omitempty"`
	AreShippingCostsFree bool     `json:"are_shipping_costs_free,omitempty"`
	ShippingCosts        float64  `json:"shipping_costs,omitempty"`
	Poster               Poster   `json:"poster,omitempty"`
	ExpiryDate           int      `json:"expiry_date,omitempty"`
}

type Image struct {
	Type        string `json:"_type"`
	ImageID     int    `json:"image_id"`
	URI         string `json:"uri"`
	LargeURI    string `json:"large_uri"`
	LargeWidth  int    `json:"large_width"`
	LargeHeight int    `json:"large_height"`
}

type Poster struct {
	Type           string `json:"_type"`
	Representation string `json:"_representation"`
	UserID         int    `json:"user_id"`
	Username       string `json:"username"`
	Followable     bool   `json:"followable"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	IconListURL    string `json:"icon_list_url"`
	IconDetailURL  string `json:"icon_detail_url"`
	Avatar         Image  `json:"avatar"`
}
type Merchant struct {
	Type                string `json:"_type"`
	MerchantID          int    `json:"merchant_id"`
	Name                string `json:"name"`
	URLName             string `json:"url_name"`
	PepperURL           string `json:"pepper_url"`
	MerchantExternalURL string `json:"merchant_external_url"`
	HeaderDescription   string `json:"header_description"`
	IconListURL         string `json:"icon_list_url"`
	IconDetailURL       string `json:"icon_detail_url"`
	Image               struct {
		Type        string `json:"_type"`
		ImageID     int    `json:"image_id"`
		URI         string `json:"uri"`
		LargeURI    string `json:"large_uri"`
		LargeWidth  int    `json:"large_width"`
		LargeHeight int    `json:"large_height"`
	} `json:"image"`
}

type Deals struct {
	Data       []Data        `json:"data"`
	Messages   []interface{} `json:"messages"`
	Validation []interface{} `json:"validation"`
	Cursors    struct {
		Previous struct {
			Before string `json:"before"`
		} `json:"previous"`
		Next struct {
			After string `json:"after"`
		} `json:"next"`
	} `json:"cursors"`
	AdditionalData struct {
		PinnedThreads []struct {
			DealType          string  `json:"_type"`
			Representation    string  `json:"_representation"`
			ThreadID          int     `json:"thread_id"`
			Description       string  `json:"description"`
			Title             string  `json:"title"`
			Origin            string  `json:"origin"`
			Submitted         int     `json:"submitted"`
			Updated           int     `json:"updated"`
			Expired           bool    `json:"expired"`
			Status            string  `json:"status"`
			Local             bool    `json:"local"`
			LastCommented     int     `json:"last_commented"`
			CommentCount      int     `json:"comment_count"`
			GroupCount        int     `json:"group_count"`
			HotDate           int     `json:"hot_date"`
			Price             float64 `json:"price"`
			DealURI           string  `json:"deal_uri"`
			VisitURI          string  `json:"visit_uri"`
			ShortURI          string  `json:"short_uri"`
			TemperatureRating int     `json:"temperature_rating"`
			TemperatureLevel  string  `json:"temperature_level"`
			IsHot             bool    `json:"is_hot"`
			Type              struct {
				Type              string `json:"_type"`
				Representation    string `json:"_representation"`
				TypeID            int    `json:"type_id"`
				Name              string `json:"name"`
				ShortName         string `json:"short_name"`
				DisplayOrder      int    `json:"display_order"`
				ThreadTypeURLName string `json:"thread_type_url_name"`
				Dealable          bool   `json:"dealable"`
				TypeUniversalID   string `json:"type_universal_id"`
				TopicFilterable   bool   `json:"topic_filterable"`
				Filterable        bool   `json:"filterable"`
			} `json:"type"`
			ExpiryDate           int    `json:"expiry_date"`
			Flaggable            bool   `json:"flaggable"`
			Votable              bool   `json:"votable"`
			Editable             bool   `json:"editable"`
			AreShippingCostsFree bool   `json:"are_shipping_costs_free"`
			GroupDisplaySummary  string `json:"group_display_summary"`
			ShowBumpedStatus     bool   `json:"show_bumped_status"`
			DisplayPriceAsFree   bool   `json:"display_price_as_free"`
			HasSuggestionCards   bool   `json:"has_suggestion_cards"`
			IconListURL          string `json:"icon_list_url"`
			IconDetailURL        string `json:"icon_detail_url"`
			HasSuggestedKeywords bool   `json:"has_suggested_keywords"`
			Clearance            bool   `json:"clearance"`
			Featured             int    `json:"featured"`
			Created              int    `json:"created"`
			IsNew                bool   `json:"is_new"`
			IsNsfw               bool   `json:"is_nsfw"`
			Poster               Poster `json:"poster"`
			Image                Image  `json:"image"`
			MainGroup            struct {
				Type       string `json:"_type"`
				GroupID    int    `json:"group_id"`
				Name       string `json:"name"`
				URLName    string `json:"url_name"`
				Default    bool   `json:"default"`
				Followable bool   `json:"followable"`
				Followed   bool   `json:"followed"`
				Statistics struct {
					ActiveDealCount    int `json:"active_deal_count"`
					ActiveVoucherCount int `json:"active_voucher_count"`
					CommentCount       int `json:"comment_count"`
					ThreadCount        int `json:"thread_count"`
				} `json:"statistics"`
				SubmittableThreadTypes  []int    `json:"submittable_thread_types"`
				Sections                []string `json:"sections"`
				PepperURL               string   `json:"pepper_url"`
				HeaderDescription       string   `json:"header_description"`
				HeroBannerSmartphoneURL string   `json:"hero_banner_smartphone_url"`
				HeroBannerTabletURL     string   `json:"hero_banner_tablet_url"`
				IconListURL             string   `json:"icon_list_url"`
				IconDetailURL           string   `json:"icon_detail_url"`
				Image                   Image    `json:"image"`
			} `json:"main_group"`
			Merchant   Merchant `json:"merchant"`
			PinnedInfo struct {
				Type            string `json:"_type"`
				Position        int    `json:"position"`
				Title           string `json:"title"`
				ForegroundColor string `json:"foreground_color"`
				BackgroundColor string `json:"background_color"`
			} `json:"pinned_info"`
		} `json:"pinned_threads"`
		ListBanners []struct {
			Type                string `json:"_type"`
			StartTime           int    `json:"start_time"`
			AnalyticsIdentifier string `json:"analytics_identifier"`
			Position            int    `json:"position"`
			CanBeHidden         bool   `json:"can_be_hidden"`
			Banner              struct {
				Type        string `json:"_type"`
				ID          string `json:"id"`
				Destination struct {
					Hash            string   `json:"hash"`
					DestinationType string   `json:"destination_type"`
					PreCachedUrls   []string `json:"pre_cached_urls"`
					CanonicalURL    string   `json:"canonical_url"`
				} `json:"destination"`
				DisplayType        string        `json:"display_type"`
				DisplayInformation []interface{} `json:"display_information"`
			} `json:"banner"`
		} `json:"list_banners"`
	} `json:"additional_data"`
}
