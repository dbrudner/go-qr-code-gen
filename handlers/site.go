package handler

import (
	"fmt"

	site "github.com/dbrudner/go-qr-code-gen/internal/site"
	siteView "github.com/dbrudner/go-qr-code-gen/views/site"
	"github.com/labstack/echo/v4"
)

type SiteHandler struct{}

func (h SiteHandler) HandleSiteDetail(c echo.Context) error {
	return render(c, siteView.Detail())
}

func (h SiteHandler) HandleSiteCollection(c echo.Context) error {
	return render(c, siteView.Collection())
}

func (h SiteHandler) HandleNewSite(c echo.Context) error {
	return render(c, siteView.New())
}

func (h SiteHandler) HandleCreateSite(c echo.Context) error {
	newSiteURL := c.FormValue("url")
	newSiteDescription := c.FormValue("description")

	newSite, err := site.NewSite(newSiteDescription, newSiteURL)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return render(c, siteView.New())
	}

	fmt.Println(newSiteURL)
	return render(c, siteView.Created(newSite.URL, newSite.Description))
}
