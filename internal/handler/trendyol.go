package handler

import (
	"Papatia/internal/db"
	"Papatia/internal/helper"
	"Papatia/internal/model"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

func Trendyol(c *fiber.Ctx) error {
	products := db.GetTrendyolProducts()
	var scrappedProducts []model.TrendyolResponse

	for key, product := range products {
		debug := fmt.Sprintf("%d. Request", key)
		fmt.Println(debug)

		helper.Sleeper(15, key)

		wg.Add(1)
		ProductData := model.TrendyolResponse{}

		go func(product model.Product, key int) {
			defer wg.Done()

			res := helper.Request(product.ProductUrl, "GET")

			defer res.Body.Close()

			if res.StatusCode == 200 {
				doc, err := goquery.NewDocumentFromReader(res.Body)
				if err != nil {
					log.Fatal(err)
				}

				ProductData.ProductCode = product.ProductCode
				ProductData.ProductUrl = product.ProductUrl

				doc.Find("h1.pr-new-br").Each(func(i int, s *goquery.Selection) {
					ProductData.ProductName = s.Text()
				})

				doc.Find(".delivery-and-favorite-wrapper .fv-dt").Each(func(i int, s *goquery.Selection) {
					ProductData.ProductFavoriteCount = strings.Split(s.Text(), " ")[0]
				})

				doc.Find(".pr-mc-w").Each(func(i int, x *goquery.Selection) {
					x.Find(".mc-ct-rght").Each(func(i int, y *goquery.Selection) {
						y.Find(".pr-bx-w .prc-dsc").Each(func(i int, z *goquery.Selection) {
							ProductData.BestPrices = append(ProductData.BestPrices, model.TrendyolBestPrice{
								Price:         strings.Split(z.Text(), " ")[0],
								FavoriteCount: x.Find(".mc-ct-lft .pr-mb .pr-mb-mn .sl-pn").Text(),
							})
						})
					})
				})

				scrappedProducts = append(scrappedProducts, ProductData)
			} else if res.StatusCode == 429 {
				panic(fmt.Sprintf("We got banned at (Product code: %s)", product.ProductCode))
			} else {
				scrappedProducts = append(scrappedProducts, model.TrendyolResponse{
					ProductCode:          product.ProductCode,
					ProductUrl:           product.ProductUrl,
					ProductName:          "!! Skipped !!",
					ProductFavoriteCount: fmt.Sprintf("%d", res.StatusCode),
					BestPrices:           []model.TrendyolBestPrice{},
				})
			}
		}(product, key)
	}

	wg.Wait()

	status, err := helper.CreateFile("trendyol.json", scrappedProducts, products)

	if err != nil {
		log.Fatal(err)
	}

	if status {
		fmt.Println("Trendyol process is done")
	}

	return c.JSON(scrappedProducts)
}
