package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Subject     string
	StartDate   string
	EndDate     string
	Month       string
	Description string
	checkbox1   bool
	checkbox2   bool
	checkbox3   bool
	checkbox4   bool
}

var dataBlog = []Blog{
	{
		Subject:     "Kucing Lucu",
		StartDate:   "5 feb 2023",
		EndDate:     "5 Mar 2023",
		Month:       "1 Month",
		Description: "Alangkah Indahnya Hari ini",
		checkbox1:   true,
		checkbox2:   true,
		checkbox3:   true,
		checkbox4:   true,
	},
	{
		Subject:     "Kucing Comel",
		StartDate:   "17 Jun 2023",
		EndDate:     "18 Jul 2023",
		Month:       "1 Month",
		Description: "Makan Dulu aja... Lagi laper,,",
		checkbox1:   true,
		checkbox2:   true,
		checkbox3:   true,
		checkbox4:   true,
	},
}

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
	e.POST("/blog-delete/:id", deleteBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": dataBlog,
	}

	return tmpl.Execute(c.Response(), blogs)
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


func addBlog(c echo.Context) error {
	project := c.FormValue("inputProject")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	checkbox1 := c.FormValue("checkbox1")
	checkbox2 := c.FormValue("checkbox2")
	checkbox3 := c.FormValue("checkbox3")
	checkbox4 := c.FormValue("checkbox4")
	description := c.FormValue("description")
	// month := ("startDate - endDate")
	
	
	

	println("Project : " + project)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Technologies : " + checkbox1 + checkbox2 + checkbox3 + checkbox4)
	println("Description : " + description)

	var newBlog = Blog{
		Subject:     project,
		StartDate:   startDate,
		EndDate:     endDate,
		Month:       "1 month",
		Description: description,
	}

	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"Id":	    id,
	// 	"Title":	"Kucing Lagi Maen",
	// 	"Duration": "19 May 2023 - 19 Jun 2023",
	// 	"Month":    "1 Month",
	// 	"Content":	"Lorem ipsum dolor sit amet consectetur adipisicing elit. Quod, veniam porro tenetur maiores in nesciunt, aperiam labore iste corporis officiis hic illum fugiat commodi sit. Ullam eos explicabo delectus aperiam. Lorem ipsum dolor sit amet consectetur, adipisicing elit. Quaerat dolore animi amet consequuntur reprehenderit temporibus cupiditate voluptatem, voluptate ducimus? Velit officia, nemo ducimus eius reiciendis illum quia voluptate assumenda iure.",
	// }

	var blogDetail = Blog{}

	for i, data := range dataBlog {
		if id == i {
			blogDetail = Blog{
				Subject:     data.Subject,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Month:       "1 Month",
				Description: data.Description,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": blogDetail,
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : ", id)

	dataBlog =append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}