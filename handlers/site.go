package handler

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/views/site"
	"github.com/labstack/echo/v4"
)

type SiteHandler struct{}

func (h SiteHandler) HandleSiteDetail(c echo.Context) error {
	return render(c, site.Detail())
}

func (h SiteHandler) HandleSiteCollection(c echo.Context) error {
	return render(c, site.Collection())
}

func (h SiteHandler) HandleNewSite(c echo.Context) error {
	return render(c, site.New())
}

func (h SiteHandler) HandleCreateSite(c echo.Context) error {
	newSiteUrl := c.FormValue("url")
	fmt.Println(newSiteUrl)
	return render(c, site.Created(newSiteUrl))
}
