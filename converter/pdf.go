package converter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func ConvertToPdf() {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("files", "index.html")

	if err != nil {
		panic(err)
	}

	fd, err := os.Open("index.html")
	defer fd.Close()

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, fd)

	if err != nil {
		panic(err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:3000/forms/chromium/convert/html", buf)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(fmt.Errorf("status: %s", res.Status))
	}

	defer res.Body.Close()
	out, err := os.Create("resume.pdf")

	if err != nil {
		panic(err)
	}

	defer out.Close()
	_, err = io.Copy(out, res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conversion to HTML Successful")
}
