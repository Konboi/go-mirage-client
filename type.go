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
	SubDomain string `json:"sub_domain"`
	Branch    string `json:"branch"`
	Image     string `json:"image"`
	IPAddress string `json:"ipaddress"`
}

type RequestParam struct {
	Subdomain string `json:"subdomain"`
	Image     string `json:"image"`
	Branch    string `json:"branch`
}
