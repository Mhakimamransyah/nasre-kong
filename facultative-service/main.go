package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FacultativeData struct {
	DocName      string
	NetPremi     int32
	isTransfered bool
}

func NewFacultavieData(docName string, premi int32, isTransfer bool) *FacultativeData {
	return &FacultativeData{
		DocName:      docName,
		NetPremi:     premi,
		isTransfered: isTransfer,
	}
}

var list []*FacultativeData

func main() {

	reproduce()

	app := fiber.New()
	router := app.Group("api").Group("/v1")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"from":     "facultative service 1",
			"messages": "Success get all data",
			"data":     list,
		})
	})

	router.Get("/:id", func(c *fiber.Ctx) error {

		res, err := func() (*FacultativeData, error) {
			for _, data := range list {

				if c.Params("id") == data.DocName {
					return data, nil
				}
			}
			return nil, fmt.Errorf("not found")
		}()

		if err == nil {
			return c.JSON(map[string]interface{}{
				"from":     "facultative service 1",
				"messages": "Success get data",
				"data":     res,
			})
		}

		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"from":     "facultative service 1",
			"messages": "Failed get data",
			"error":    err.Error(),
		})

	})

	app.Listen(":3000")
}

func reproduce() {
	list = append(list, NewFacultavieData("1", 100000, true), NewFacultavieData("2", 200000, false))
}
