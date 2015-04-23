package digitalocean

import (
	"fmt"
	"strconv"
)

type Image struct {
	Id	int64				     `json:"id"`
	Name	string				     `json:"name"`
	Distro  string				     `json:"distro"`
}

type imageResponse struct {
	Image Image `json:"image"`
}

// Returns the slug for the image
func (i *Image) StringId() string {
	return strconv.FormatInt(i.Id, 10)
}

// RetrieveImage gets returns an Image by their Id and an error in case of their absence.
func (c *Client) RetrieveImage(id string) (Image, error) {
	req, err := c.NewRequest(nil, "GET", fmt.Sprintf("/account/images/%s", id))

	if err != nil {
		return Image{}, err
	}

	resp, err := checkResp(c.Http.Do(req))
	if err != nil {
		return Image{}, fmt.Errorf("Error searching for Image id: %s", err)
	}

	image := new(imageResponse)

	err = decodeBody(resp, &image)

	if err != nil {
		return Image{}, fmt.Errorf("Error: %s", err)
	}

	// The request was successful
	return image.Image, nil
}
