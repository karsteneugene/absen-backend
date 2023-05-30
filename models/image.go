package models

type Image struct {
	ID           uint   `json:"id" xml:"id" form:"id" gorm:"primaryKey"`
	FileName     string `json:"file_name" xml:"file_name" form:"file_name"`
	FilePath     string `json:"file_path" xml:"file_path" form:"file_path"`
	OriginalURL  string `json:"original_url" xml:"original_url" form:"original_url"`
	ThumbnailURL string `json:"thumbnail_url" xml:"thumbnail_url" form:"thumbnail_url"`
	Width        string `json:"width" xml:"width" form:"width"`
	Height       string `json:"height" xml:"height" form:"height"`
	URL300       string `json:"url_300" xml:"url_300" form:"url_300"`
}
