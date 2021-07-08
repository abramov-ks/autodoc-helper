package autodoc

type BearerToken string

type AutodocConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	BaseUrl  string `yaml:"url"`
	AuthUrl  string `yaml:"auth_url"`
	ApiUrl   string `yaml:"api_url"`
}

// Структура сессии к автодоку
type AutodocSession struct {
	AuthData AuthResult
	BaseUrl  string
	AuthUrl  string
	ApiUrl   string
	Username string
	Password string
}

// Результат авторизации
type AuthResult struct {
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type ManufacterShortInfo struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	PopularityIndex int    `json:"popularityIndex"`
}

// ответ на запрос цен
type PartnumberPriceResponse struct {
	PartNumber          string                                 `json:"partnumber"`
	DisplayPartNumber   string                                 `json:"displayPartNumber"`
	Name                string                                 `json:"name"`
	MinimalPrice        float32                                `json:"minimalPrice"`
	MinimalDeliveryDays int                                    `json:"minimalDeliveryDays"`
	InventoryItems      []PartnumberPriceResponseInventoryItem `json:"inventoryItems"`
	Manufacturer        ManufacterShortInfo                    `json:"manufacturer"`
}

// ценовые предложения в ответе
type PartnumberPriceResponseInventoryItem struct {
	Id                  int                                          `json:"id"`
	Price               float32                                      `json:"price"`
	Quantity            int                                          `json:"quantity"`
	MinimalDeliveryDays int                                          `json:"minimalDeliveryDays"`
	DeliveryDays        int                                          `json:"deliveryDays"`
	UpdateDate          string                                       `json:"updateDate"`
	Supplier            PartnumberPriceResponseInventoryItemSupplier `json:"supplier"`
}

// поставщики ценовых предложений
type PartnumberPriceResponseInventoryItemSupplier struct {
	Name                string `json:"name"`
	NextOrderDate       string `json:"nextOrderDate"`
	NextOrderDateString string `json:"nextOrderDateString"`
	Schedule            string `json:"schedule"`
}

// инфа о производителе
type ManufacterInfo struct {
	Id               int    `json:"id"`
	ManufacturerName string `json:"manufacturerName"`
	PartName         string `json:"partName"`
	ArtNumber        string `json:"artNumber"`
}
