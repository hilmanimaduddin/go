package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/project", project)
	e.POST("/add-blog", addBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func project(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":	    id,
		"Title":	"Kucing Lagi Maen",
		"Content":	"Lorem ipsum dolor sit amet consectetur adipisicing elit. Quod, veniam porro tenetur maiores in nesciunt, aperiam labore iste corporis officiis hic illum fugiat commodi sit. Ullam eos explicabo delectus aperiam. Lorem ipsum dolor sit amet consectetur, adipisicing elit. Quaerat dolore animi amet consequuntur reprehenderit temporibus cupiditate voluptatem, voluptate ducimus? Velit officia, nemo ducimus eius reiciendis illum quia voluptate assumenda iure.",
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	project := c.FormValue("inputProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	checkbox1 := c.FormValue("checkbox")
	checkbox2 := c.FormValue("checkBox2")
	checkbox3 := c.FormValue("checkBox3")
	checkbox4 := c.FormValue("checkBox4")
	description := c.FormValue("description")

	println("Project : " + project)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Technologies : " + checkbox1 + checkbox2 + checkbox3 + checkbox4)
	println("Description : " + description)

	return c.Redirect(http.StatusMovedPermanently, "/")
}