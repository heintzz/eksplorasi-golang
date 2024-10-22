package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"integration/pkg/images"
	"integration/utils/config"
	"io"
	"log"
	"net/http"
	"os"
)

type CloudService interface {
	// @Param file refer to file buffer
	// @Param pathDestination refer to target directory/bucket in cloud provider
	Upload(ctx context.Context, file interface{}, pathDestination string, quality string) (uri string, compressedUri string, err error)
}

type Services struct {
	cloud CloudService
}

func main() {
	err := config.LoadConfig(".env")
	if err != nil {
		log.Println("error when try to LoadConfig with detail :", err.Error())
	}
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloudProvider := images.NewCloudinary(cloudName, apiKey, apiSecret)

	svc := Services{
		cloud: cloudProvider,
	}

	http.HandleFunc("/files/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)

			response := map[string]string{
				"message": fmt.Sprintf("%s method isn't allowed", r.Method),
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error processing JSON", http.StatusInternalServerError)
				return
			}

			w.Write(jsonResponse)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte("error"))
			return
		}
		defer file.Close()

		const MAX_FILE_SIZE = 5 * 1024 * 1024
		if handler.Size > MAX_FILE_SIZE {
			w.Write([]byte("file terlalu besar"))
			return
		}

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, file)
		if err != nil {
			http.Error(w, "Error reading file into buffer", http.StatusInternalServerError)
			return
		}

		url, compressedUri, err := svc.cloud.Upload(context.Background(), &buffer, r.FormValue("filetype"), r.FormValue("quality"))
		if err != nil {
			response := map[string]interface{}{
				"success": false,
				"message": "upload failed",
				"error":   err,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error processing JSON", http.StatusInternalServerError)
				return
			}

			w.Write(jsonResponse)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			payload := map[string]string{
				"url": url,
			}

			if compressedUri != "" {
				payload["compress_url"] = compressedUri
			}

			response := map[string]interface{}{
				"success": true,
				"message": "upload succeed",
				"payload": payload,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Error processing JSON", http.StatusInternalServerError)
				return
			}

			w.Write(jsonResponse)
			return
		}
	})

	http.ListenAndServe(":4444", nil)
}
