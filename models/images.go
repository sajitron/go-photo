package models

import (
	"fmt"
	"io"
	"os"
)

// ImageService interface
type ImageService interface {
	Create(galleryID uint, r io.Reader, filename string) error
	// ByGalleryID(galleryID uint) []string
}

// NewImageService function ?
func NewImageService() ImageService {
	return &imageService{}
}

type imageService struct{}

// Create - initiate image upload
func (is *imageService) Create(galleryID uint, r io.Reader, filename string) error {
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}
	// Create a destination file
	dst, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy reader data to the destination file
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := fmt.Sprintf("images/galleries/%v/", galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}
