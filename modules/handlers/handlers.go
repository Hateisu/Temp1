package handlers

import (
	"encoding/json"
	"log"
	"marketapi/modules/consts"
	"marketapi/modules/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleWelcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func HandleGetAllUsers(c *gin.Context) {
	var users []storage.User
	storage.DB.Find(&users)
	//res,err:=results.Rows()
	/*data, err := json.Marshal(users)
	if err != nil {

	}*/
	c.JSON(http.StatusOK, consts.RESPONSE_OK_200(users))
}

type JSONGetText struct {
	Text string `json:"text"`
}

func CheckForRq(c *gin.Context) {
	var text JSONGetText
	if err := c.BindJSON(&text); err != nil {
		ans, err := json.Marshal(map[string]string{"text": "Check fields"})
		if err != nil {
			return
		}
		c.JSON(500, json.RawMessage(ans))
		return
	}
	log.Println(text.Text)
	c.JSON(http.StatusOK, "")
}

func HandleAddCountry(c *gin.Context) {
	var newCountruy storage.Country

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newCountruy); err != nil {
		return
	}
	storage.DB.Create(&newCountruy)
}

func HandleGetCountry(c *gin.Context) {
	countryname := c.DefaultQuery("name", "")
	if countryname == "" {
		return
	}
	var country storage.Country //= storage.Country{Name: countryname}
	storage.DB.Where("name = ?", countryname).First(&country)
	//res,err:=results.Rows()
	/*data, err := json.Marshal(users)
	if err != nil {

	}*/
	c.JSON(http.StatusOK, consts.RESPONSE_OK_200(country))
}

// Other Functions

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
