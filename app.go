package main

import "github.com/labstack/echo"
import "net/http"
import "fmt"
import "os/exec"
import "bytes"

func main() {
	e := echo.New()
	e.GET("/get", getPage)
	e.Logger.Fatal(e.Start(":8080"))
}
func getPage(c echo.Context) error {
	url := c.FormValue("url")
	fmt.Println(url)
	cmd := exec.Command("adb", "shell", fmt.Sprintf("curl --cacert /sdcard/cacert-2020-01-01.pem %s", url))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, out.String())
	return nil
}
