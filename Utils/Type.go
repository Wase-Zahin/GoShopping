package Utils

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"title"`
	Category string `json:"category"`
	Rated    bool   `json:"rated"`
	Image    string `json:"img"`
}
