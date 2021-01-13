package telegram

// Sticker contains information about a sticker.
type Sticker struct {
	// FileID is an identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is an unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width sticker width
	Width int `json:"width"`
	// Height sticker height
	Height int `json:"height"`
	// IsAnimated true, if the sticker is animated
	// Optional
	IsAnimated bool `json:"is_animated"`
	// Thumb sticker thumbnail in the .WEBP or .JPG format
	// Optional
	Thumb *PhotoSize `json:"thumb"`
	// Emoji associated with the sticker
	// Optional
	Emoji string `json:"emoji"`
	// SetName of the sticker set to which the sticker belongs
	// Optional
	SetName string `json:"set_name"`
	// MaskPosition for mask stickers, the position where the mask should be placed
	// Optional
	MaskPosition *MaskPosition `json:"mask_position"`
	// FileSize is file size
	// Optional
	FileSize int `json:"file_size"`
}

// StickerSet contains information about an sticker set.
type StickerSet struct {
	// Name sticker set name
	Name string `json:"name"`
	// Title sticker set title
	Title string `json:"title"`
	// IsAnimated true, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated"`
	// ContainsMasks true, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks"`
	// Stickers list of all set stickers
	Stickers *[]Sticker `json:"stickers"`
	// Thumb sticker thumbnail in the .WEBP or .JPG format
	// Optional
	Thumb *PhotoSize `json:"thumb"`
}

// MaskPosition describes the position on faces
// where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size,
	// from left to right. For example, choosing -1.0 will place mask just
	// to the left of the default mask position.
	XShift float64 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size,
	// from top to bottom. For example, 1.0 will place the mask
	// just below the default mask position.
	YShift float64 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}