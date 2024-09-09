package handlers

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpstvs/csitemsapp/internals/database"
)

func CreateUser(c *fiber.Ctx, db *database.Queries) error {
	fullURL := "http://" + r.Host + r.RequestURI

	urlParsed, _ := url.Parse(fullURL)

	params := urlParsed.Query()

	validationParams := url.Values{}
	validationParams.Add("openid.assoc_handle", params.Get("openid.assoc_handle"))
	validationParams.Add("openid.signed", params.Get("openid.signed"))
	validationParams.Add("openid.sig", params.Get("openid.sig"))
	validationParams.Add("openid.ns", params.Get("openid.ns"))
	validationParams.Add("openid.mode", "check_authentication")
	for _, param := range strings.Split(params.Get("openid.signed"), ",") {
		validationParams.Add("openid."+param, params.Get("openid."+param))
	}

	resp, err := http.PostForm("https://steamcommunity.com/openid/login", validationParams)
	if err != nil {
		http.Error(w, "Error validating OpenID response", http.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	if !strings.Contains(string(body), "is_valid:true") {
		http.Error(w, "OpenID validation failed", http.StatusUnauthorized)
		return
	}

	queries, _ := url.ParseQuery(urlParsed.RawQuery)
	openid := queries["openid.claimed_id"][0]
	id := utils.ExtractSteamid(openid)
	profile := steamapi.FetchPlayerData(id)

	name := profile.Response.Players[0].Personaname

	_, err = cfg.DB.GetUserbyId(r.Context(), id)

	if err != nil {
		cfg.DB.CreateUser(r.Context(), database.CreateUser{
			ID:      uuid.New(),
			Name:    name,
			Steamid: id,
		})

		CreateCookie(w, id)
		http.Redirect(w, r, "/v1/api/profile", http.StatusTemporaryRedirect)
	}

	CreateCookie(w, id)
	http.Redirect(w, r, "/v1/api/profile", http.StatusTemporaryRedirect)

}
