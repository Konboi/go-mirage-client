package mirage

type Status struct {
	Result string `json:"result"`
}

type List struct {
	Result []Info `json:"result"`
}

type Info struct {
	ID        string `json:"id"`
	ShortID   string `json:"short_id"`
	SubDomain string `json:"subdomain"`
	Branch    string `json:"branch"`
	Image     string `json:"image"`
	IPAddress string `json:"ipaddress"`
}
