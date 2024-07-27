package handlers

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rpstvs/csitemsapp/internals/auth"
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

	steamLink := c.Query("openid.claimed_id")

	tmp := strings.Split(steamLink, "/")
	id := tmp[len(tmp)-1]

	token := auth.CreateToken(id)
	cookie := auth.CreateCookie(token)
	c.Cookie(&cookie)
	c.Redirect("/api/skins/Best")
	return nil
}
