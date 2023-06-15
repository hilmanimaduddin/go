package main

import (
	"context"
	"fmt"
	"go/connection"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
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
}

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

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/project", project)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-delete/:id", deleteBlog)
	e.POST("/update-blog/:id", editBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration FROM tb_projek")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Subject, &each.StartDate, &each.EndDate, &each.Description, &each.Icon1, &each.Icon2, &each.Icon3, &each.Icon4, &each.Myicon1, &each.Myicon2, &each.Myicon3, &each.Myicon4, &each.Duration)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Image = ""

		each.FormatStartDate = each.StartDate.Format("2 January 2006")
		each.FormatEndDate = each.EndDate.Format("2 January 2006")
		

		result = append(result, each)
	}
	
	blogs := map[string]interface{}{
		"Blogs": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
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
	// month := ("startDate - endDate")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projek (subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)", subject, startDate, endDate, description, icon1, icon2, icon3, icon4, label1, label2, label3, label4, duration)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}


func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, subject, start_date, end_date, description, icon1, icon2, icon3, icon4, myicon1, myicon2, myicon3, myicon4, duration FROM tb_projek WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Subject, &BlogDetail.StartDate, &BlogDetail.EndDate, &BlogDetail.Description, &BlogDetail.Icon1, &BlogDetail.Icon2, &BlogDetail.Icon3, &BlogDetail.Icon4, &BlogDetail.Myicon1, &BlogDetail.Myicon2, &BlogDetail.Myicon3, &BlogDetail.Myicon4, &BlogDetail.Duration)


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	BlogDetail.FormatStartDate = BlogDetail.StartDate.Format("2 January 2006")
	BlogDetail.FormatEndDate = BlogDetail.EndDate.Format("2 January 2006")

	BlogDetail.Image = ""

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

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