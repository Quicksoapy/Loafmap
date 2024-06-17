package cookies

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

func SetAccountCookie(w http.ResponseWriter, inputValue string) {
	var s = securecookie.New(securecookie.GenerateRandomKey(50), securecookie.GenerateRandomKey(32))
	//make securecookie instance, first value how to encrypt name, second to encrypt value

	value := map[string]string{
		"AccountId": inputValue,
	}
	if encoded, err := s.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:     "cookie",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
}

func ReadAccountCookie(r *http.Request) string {
	var s = securecookie.New(securecookie.GenerateRandomKey(50), securecookie.GenerateRandomKey(32))

	if cookie, err := r.Cookie("cookie"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("cookie", cookie.Value, &value); err == nil {
			return value["AccountId"]
		}
	}
	return ""
}
