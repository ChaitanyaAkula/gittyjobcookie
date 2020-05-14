package cookieswoexpire
import (
	"net/http"
//	"time"
)

func SetCookie(w http.ResponseWriter,name,value string){
        http.SetCookie(w,&http.Cookie{
		Name:name,
		Value:value,
	//	Expires:expire, // persistent cookie , max age = 1 year
		HttpOnly: true,
	})

}