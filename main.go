package main

import (
	"context"
	"fmt"
	"go/connection"
	"go/middleware"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Blog struct {
	ID          int
	Subject     string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	Image       string
	Duration    string
	Icon1       string
	Icon2       string
	Icon3       string
	Icon4       string
	Myicon1     string
	Myicon2     string
	Myicon3     string
	Myicon4     string
	FormatStartDate  string
	FormatEndDate  string
	Author      string
}


type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// var dataBlog = []Blog{
// 	{
// 		Subject:     "Kucing Lucu",
// 		StartDate:   "2023-03-17",
// 		EndDate:     "2023-04-18",
// 		Duration:    "2 weeks",
// 		Description: "Alangkah Indahnya Hari ini",
// 		Icon1:       `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon2:       `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon3:       `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon4:       `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i>`,
// 		Myicon1:     `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 		Myicon2:     `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 		Myicon3:     `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 		Myicon4:     `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 	},
// 	{
// 		Subject:     "Kucing Comel",
// 		StartDate:   "2023-06-04",
// 		EndDate:     "2023-08-01",
// 		Duration:    "2 months",
// 		Description: "Makan Dulu aja... Lagi laper,,",
// 		Icon1:       `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon2:       `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon3:       `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i>`,
// 		Icon4:       `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i>`,
// 		Myicon1:     `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i><span>Node Js</span>`,
// 		Myicon2:     `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 		Myicon3:     `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 		Myicon4:     `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`,
// 	},
// }

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/project", project)
	e.GET("/register", formRegister)
	e.GET("/login", formLogin)

	e.POST("/reg", register)
	e.POST("/log", login)

	e.POST("/logout", logout)

	e.POST("/add-blog", middleware.UploadFile(addBlog))
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/update-blog/:id", editBlog)

	e.Logger.Fatal(e.Start("localhost:5400"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_projek.id, subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration, image, tb_user.name AS author FROM tb_projek JOIN tb_user ON tb_projek.author_id = tb_user.id ORDER BY tb_projek.id DESC")


	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Subject, &each.StartDate, &each.EndDate, &each.Description, &each.Icon1, &each.Icon2, &each.Icon3, &each.Icon4, &each.Myicon1, &each.Myicon2, &each.Myicon3, &each.Myicon4, &each.Duration, &each.Image, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormatStartDate = each.StartDate.Format("2 January 2006")
		each.FormatEndDate = each.EndDate.Format("2 January 2006")
		

		result = append(result, each)
	}


	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
		"Blogs": result,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())


	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}


	return tmpl.Execute(c.Response(), datas)
}

func project(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())


	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func contact(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func MyIcon(Valu string) string {
	if Valu == "Node Js " {
		return `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i>`
	} else if Valu == "React Js " {
		return `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i>`
	} else if Valu == "Next Js " {
		return `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i>`
	} else if Valu == "TypeScript " {
		return `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i>`
	} else {
		return ""
	}
}

func MyLabel(Valu string) string {
	if Valu == "Node Js " {
		return `<i class="fa-brands fa-react"" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`
	} else if Valu == "React Js " {
		return `<i class="fa-brands fa-square-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`
	} else if Valu == "Next Js " {
		return `<i class="fa-brands fa-node-js" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`
	} else if Valu == "TypeScript " {
		return `<i class="fa-solid fa-bolt" style="color: #000000; margin-right: 10px"></i><span>React Js</span>`
	} else {
		return ""
	}
}

func addBlog(c echo.Context) error {
	subject := c.FormValue("inputProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := getDuration(endDate, startDate)
	description := c.FormValue("description")
	iconA := c.FormValue("icon1")
	iconB := c.FormValue("icon2")
	iconC := c.FormValue("icon3")
	iconD := c.FormValue("icon4")
	icon1 := MyIcon(iconA)
	icon2 := MyIcon(iconB)
	icon3 := MyIcon(iconC)
	icon4 := MyIcon(iconD)
	label1 := MyLabel(iconA)
	label2 := MyLabel(iconB)
	label3 := MyLabel(iconC)
	label4 := MyLabel(iconD)

	sess, _ := session.Get("session", c)
	author := sess.Values["id"].(int)

	image := c.Get("dataFile").(string)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projek (subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", subject, startDate, endDate, description, icon1, icon2, icon3, icon4, label1, label2, label3, label4, duration, image, author)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}


func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration, image FROM tb_projek WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Subject, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Description, &BlogDetail.Icon1, &BlogDetail.Icon2, &BlogDetail.Icon3, &BlogDetail.Icon4, &BlogDetail.Myicon1, &BlogDetail.Myicon2, &BlogDetail.Myicon3, &BlogDetail.Myicon4, &BlogDetail.Duration, &BlogDetail.Image)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	BlogDetail.FormatStartDate = BlogDetail.StartDate.Format("2 January 2006")
	BlogDetail.FormatEndDate = BlogDetail.EndDate.Format("2 January 2006")

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
		"Blog": BlogDetail,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())


	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)

}

func getDuration(endDate, startDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func editBlog(edit echo.Context) error {
	id, _ := strconv.Atoi(edit.Param("id"))

	fmt.Println("ID: ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projek WHERE id=$1", id)

	if err != nil {
		return edit.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return edit.Redirect(http.StatusMovedPermanently, "/project")
}


func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID: ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projek WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register(c echo.Context) error {
	// to make sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/register")
	}

	return redirectWithMessage(c, "Register success!", true, "/login")
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}