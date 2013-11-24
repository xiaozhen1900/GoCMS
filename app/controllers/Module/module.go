package controllers

//模块首页

import "strconv"
import "github.com/robfig/revel"
import "admin/app/models"
import "admin/app/routes"

type Module struct {
	*revel.Controller
}

func (c Module) Index() revel.Result {
	title := "模块--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Module/Index.html")
}

//检测是否登陆
//init.go调用
func (c *Module) inject() revel.Result {

	if UserID, ok := c.Session["UserID"]; ok {
		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)
		if admin_info.Id <= 0 {
			c.Flash.Error("请先登录")
			return c.Redirect(routes.User.Login())
		}

		//控制器
		c.RenderArgs["controller"] = c.Name
		//动作
		c.RenderArgs["action"] = c.Action
		//模型
		c.RenderArgs["model"] = c.MethodName

		//导航菜单
		menu := new(models.Menu)
		c.RenderArgs["menus"] = menu.GetMenuAll(0, true)

		//登陆用户信息
		c.RenderArgs["admin_info"] = admin_info

	} else {
		c.Flash.Error("请先登录")
		return c.Redirect(routes.User.Login())
	}

	return nil
}
