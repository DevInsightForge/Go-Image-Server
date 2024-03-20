package dto

type ResizeRequestQueryDto struct {
	ImageUrl string `json:"imageUrl"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Quality  int    `json:"quality,omitempty"`
}
