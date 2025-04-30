package dto

type CalculatePacksRequest struct {
	Items int `json:"items" binding:"required,gt=0"`
}

type PackResponse struct {
	ID       int `json:"id"`
	ItemSize int `json:"item_size"`
}

type AddPackSizeRequest struct {
	ItemSize int `json:"item_size" binding:"required,gt=0"`
}
