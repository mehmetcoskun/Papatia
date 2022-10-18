package model

type TrendyolResponse struct {
	ProductCode          string
	ProductName          string
	ProductUrl           string
	ProductFavoriteCount string
	BestPrices           []TrendyolBestPrice
}

type TrendyolBestPrice struct {
	Price         string
	FavoriteCount string
}
