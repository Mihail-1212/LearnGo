import "net/http"

// Есть ряд функций
func Get(url string) (resp *Response, err error)	// Get(): отправляет запрос GET
func Head(url string) (resp *Response, err error)	// Head(): отправляет запрос HEAD
func Post(url string, contentType string, body io.Reader) (resp *Response, err error)	// Post(): отправляет запрос POST
func PostForm(url string, data url.Values) (resp *Response, err error)		// PostForm(): отправляет форму в запросе POST