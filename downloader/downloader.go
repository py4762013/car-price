package downloader

import (
	"car-prices/fake"
	"github.com/axgle/mahonia"
	"io"
	"log"
	"net/http"
)

func Get(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest err: %v", err)
	}

	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Referer", "https://cart.autohome.com.cn")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err: %v", err)
	}

	math := mahonia.NewDecoder("gbk")
	return math.NewReader(resp.Body)
}