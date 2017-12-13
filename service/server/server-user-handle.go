package server

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"

	"github.com/tpisntgod/Agenda/service/entity/user"
)

// error.toString
func toString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

type B struct {
	A string
	B string
}

type Todo struct {
	Success bool   `json:"Success"`
	Result  string `json:Result`
}

// 标准Json，只包含Success和Result
func stdJson(succ bool, res string) Todo {
	return Todo{succ, res}
}

func initMydb(args []string) {
	// if len(args) != 5 && len(args) != 1 {
	// 	fmt.Fprintln(os.Stderr, "Please input the database information!")
	// 	fmt.Fprintln(os.Stderr, "\t./app username password port databasename")
	// 	fmt.Fprintln(os.Stderr, "Or use: \n\t./app\nwe will use (root) (root) (2048) (test)")
	// 	os.Exit(1)
	// }
	//
	// // 声明四个变量
	// name := "root"
	// password := "root"
	// port := "2048"
	// dname := "test"
	//
	// if len(args) != 1 {
	// 	name = args[1]
	// 	password = args[2]
	// 	port = args[3]
	// 	dname = args[4]
	// }
	//
	// // 创建数据库
	// entities.InitMydb(name, password, port, dname)
}

// 显示CurrentUser的id和name
func showCurrentUserHandle(formatter *render.Render) http.HandlerFunc {
	fmt.Println("Enter?")
	return func(w http.ResponseWriter, r *http.Request) {
		if user.IsLogin() {
			// 在登录状态
			formatter.JSON(w, http.StatusOK, struct {
				Success bool
				Name    string
				ID      int
			}{
				true,
				user.CurrentUser.Name,
				user.CurrentUser.ID})
		} else {
			// 不在登录状态
			formatter.JSON(w, http.StatusOK, struct {
				Success bool
				Result  string
			}{
				false,
				"ERROR: no logined user!"})
		}
	}
}

// 创建一个新的用户
func createUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := user.RegisterUser(r.FormValue("Name"), r.FormValue("Passname"), r.FormValue("Email"), r.FormValue("Phone"))
		succ := (bool)(err == nil)
		res := toString(err)

		formatter.JSON(w, http.StatusOK, struct {
			Success bool
			Result  string
		}{
			succ,
			res})
	}
}

// 登录用户
func loginUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := user.LoginUser(r.FormValue("Name"), r.FormValue("Password"))
		succ := (bool)(err == nil)
		res := toString(err)

		if succ {
			formatter.JSON(w, http.StatusOK, struct {
				Success bool
				Result  string
				ID      int
				Name    string
				Email   string
				Phone   string
			}{
				succ,
				res,
				user.CurrentUser.ID,
				user.CurrentUser.Name,
				user.CurrentUser.Email,
				user.CurrentUser.PhoneNumber})
		} else {
			formatter.JSON(w, http.StatusOK, struct {
				Success bool
				Result  string
			}{
				succ,
				res})
		}
	}
}

// 登出用户
func logoutUserHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := user.LogoutUser()
		succ := (bool)(err == nil)
		res := toString(err)
	}
}

func undefinedHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
	}
}