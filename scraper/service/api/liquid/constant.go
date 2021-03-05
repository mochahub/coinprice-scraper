package liquid

type Interval string

const (
	LIQUID   = "LIQUID"
	maxLimit = 1000
	// calls per second
	callsPerMinute = 300

	// Endpoints
	baseUrl     = "https://api.liquid.com"
	getProducts = "/products"
)
