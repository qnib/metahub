package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	googleAuth "golang.org/x/oauth2/google"
)

func googleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/*{"code":"4/KwGtc043rxAVLiK4b24Qau8ET7x25V4HVlvo9Jb20704AC8lvdWdmSQwOvlXhcpgl6wvxqnfL1wc3qB9LltE__g",
	"clientId":"936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com",
	"redirectUri":"http://localhost:8081"}*/

	decoder := json.NewDecoder(r.Body)
	var body struct {
		Code        string `json:"code"`
		ClientID    string `json:"clientId"`
		RedirectURI string `json:"redirectUri"`
	}
	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("error decoding code: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//log.Printf("code: %s", body.Code)

	googleConfig := &oauth2.Config{
		ClientID:     "936040293434-i3m9p4km8it5np2bs253a7rvedchofs6.apps.googleusercontent.com",
		ClientSecret: "E0wI5Bb0KbZ1__DgztogDu1O",
		Scopes:       []string{"profile", "email", "openid"},
		Endpoint:     googleAuth.Endpoint,
		RedirectURL:  body.RedirectURI,
	}

	if body.ClientID != googleConfig.ClientID {
		log.Printf("invalid client ID %q", body.ClientID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := googleConfig.Exchange(ctx, body.Code)
	if err != nil {
		log.Printf("Exchange failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !token.Valid() {
		log.Printf("token invalid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("failed getting user info: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed read response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*{
	  "id": "112043567300478037431",
	  "email": "robert.fujara@gmail.com",
	  "verified_email": true,
	  "name": "Robert Fujara",
	  "given_name": "Robert",
	  "family_name": "Fujara",
	  "link": "https://plus.google.com/112043567300478037431",
	  "picture": "https://lh4.googleusercontent.com/-2GeSbg8UaEc/AAAAAAAAAAI/AAAAAAABNnU/G8nvHyqrdx8/photo.jpg",
	  "locale": "de"
	}*/
	userInfoReader := strings.NewReader(string(contents))
	userInfoDecoder := json.NewDecoder(userInfoReader)
	var userInfo struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := userInfoDecoder.Decode(&userInfo); err != nil {
		log.Printf("error decoding user info: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwt, err := tokenToJSON(token)
	if err != nil {
		log.Printf("error creating JWT: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print(jwt)

	w.Write([]byte(jwt))
}
