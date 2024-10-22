package images

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Cloudinary struct {
	Cloud   *cloudinary.Cloudinary
	IsError error
}

func NewCloudinary(cloud, apiKey, apiSecret string) Cloudinary {
	c, err := cloudinary.NewFromParams(cloud, apiKey, apiSecret)

	return Cloudinary{
		Cloud:   c,
		IsError: err,
	}
}

func (c Cloudinary) Upload(ctx context.Context, file interface{}, path string, quality string) (string, string, error) {
	var eager string
	if quality == "" {
		eager = "q_auto:eco"
	} else {
		eager = fmt.Sprintf("q_%s", quality)
	}

	filename := uuid.NewString()
	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		AssetFolder: "heintzz/" + path,
		PublicID:    "heintzz/" + path + "/" + filename,
		Eager:       eager,
	})

	if err != nil {
		return "", "", err
	}

	compressedURL := ""
	if len(res.Eager) > 0 {
		compressedURL = res.Eager[0].SecureURL
	}

	url := res.SecureURL

	return url, compressedURL, nil
}
