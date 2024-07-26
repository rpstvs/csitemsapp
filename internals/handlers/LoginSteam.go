package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CallbackLogin(c *fiber.Ctx) error {

	validationParams := url.Values{}
	validationParams.Add("openid.assoc_handle", c.Query("openid.assoc_handle"))
	validationParams.Add("openid.signed", c.Query("openid.signed"))
	validationParams.Add("openid.sig", c.Query("openid.sig"))
	validationParams.Add("openid.ns", c.Query("openid.ns"))
	validationParams.Add("openid.mode", "check_authentication")
	for _, param := range strings.Split(c.Query("openid.signed"), ",") {
		validationParams.Add("openid."+param, c.Query("openid."+param))
	}

	resp, err := http.PostForm("https://steamcommunity.com/openid/login", validationParams)
	if err != nil {

		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		return nil
	}

	if !strings.Contains(string(body), "is_valid:true") {

		return nil
	}

	fmt.Println("We are loggedin")
	c.Redirect("/api/skins/Best")
	return nil
}
