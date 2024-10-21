package images

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (c Cloudinary) Upload(ctx context.Context, file interface{}, path string) (string, error) {
	filename := uuid.NewString()
	res, err := c.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "heintzz/" + path + "/" + filename,
		Eager:    "q_10",
	})

	if err != nil {
		return "", err
	}

	if len(res.Eager) > 0 {
		return res.Eager[0].SecureURL, nil
	}

	url := res.SecureURL

	return url, nil
}

func (c Cloudinary) Remove(ctx context.Context, path string) error {
	res, err := c.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: path,
	})

	if err != nil {
		return err
	}

	if strings.Contains(res.Result, "not found") {
		return errors.New("image not found")
	}

	fmt.Printf("%+v\n", res)

	return err
}
