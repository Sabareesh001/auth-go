package jwt

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

func AssignJWT(username string, expiry time.Duration,SECRET_KEY []byte) string{
        claims := jwt.MapClaims{
			"username":username,
			"exp":time.Now().Add(expiry).Unix(),
		}

        token  := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

       signedToken,err := token.SignedString(SECRET_KEY);

	   if(err!=nil){
           return ""
	   }
       return signedToken 
}


func ValidateJwt(token string,SECRET_KEY string) *jwt.Token{
     parsedToken,err := jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY),nil
	 })
	 if(err!=nil){
		fmt.Println(err)
		return nil
	 }
    return parsedToken
}