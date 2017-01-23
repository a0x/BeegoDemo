package controllers

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/astaxie/beego/validation"
  models "sitepointgoapp/models"
  "strconv"
)

type ManageController struct {
  beego.Controller
}

func (manage *ManageController) Home() {
  manage.Layout = "basic-layout.tpl"
  manage.LayoutSections = make(map[string]string)
  manage.LayoutSections["Header"] = "header.tpl"
  manage.LayoutSections["Footer"] = "footer.tpl"
  manage.TplName = "manage/home.tpl"
}

/**
 * Delete the user, by their id, if that user's available.
 */
func (manage *ManageController) Delete() {
  manage.Layout = "basic-layout.tpl"
  manage.LayoutSections = make(map[string]string)
  manage.LayoutSections["Header"] = "header.tpl"
  manage.LayoutSections["Footer"] = "footer.tpl"
  manage.TplName = "manage/home.tpl"

  // convert the string value to an int
  articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))

  o := orm.NewOrm()
  o.Using("default")

  article := models.Article{}

  // Check if the article exists first
  if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
    if num, err := o.Delete(&models.Article{Id: articleId}); err == nil {
      beego.Info("Record Delete. ", num)
    } else {
      beego.Error("Record couldn't be deleted. Reason: ", err)
    }
  } else {
    beego.Info("Record Not Found.")
  }
}

func (manage *ManageController) Update() {
  o := orm.NewOrm()
  o.Using("default")
  flash := beego.NewFlash()

  // convert the string value to an int
  if articleId, err := strconv.Atoi(manage.Ctx.Input.Param(":id")); err == nil {
    article := models.Article{Id: articleId}
    // attempt to load the record from the database
    if o.Read(&article) == nil {
      // set the Client and Url properties to arbitrary values
      article.Client = "Sitepoint"
      article.Url = "http://bing.com"
      if num, err := o.Update(&article); err == nil {
        flash.Notice("Record was updated.")
        flash.Store(&manage.Controller)
        beego.Info("Record was updated. ", num)
      }
    } else {
      flash.Notice("Record was NOT updated.")
      flash.Store(&manage.Controller)
      beego.Error("Cound not find article matching id: ", articleId)
    }
  } else {
    flash.Notice("Record was NOT updated.")
    flash.Store(&manage.Controller)
    beego.Error("Could not convert id from a string to a number. ", err)
  }

  // redirect afterwards
  manage.Redirect("/manage/view", 302)
}

func (manage *ManageController) View() {
  manage.Layout = "basic-layout.tpl"
  manage.LayoutSections = make(map[string]string)
  manage.LayoutSections["Header"] = "header.tpl"
  manage.LayoutSections["Footer"] = "footer.tpl"
  manage.TplName = "manage/view.tpl"

  flash := beego.ReadFromRequest(&manage.Controller)

  if ok := flash.Data["error"]; ok != "" {
    // Display error messages
    manage.Data["errors"] = ok
  }

  if ok := flash.Data["notice"]; ok != "" {
    // Display notice messages
    manage.Data["notices"] = ok
  }

  o := orm.NewOrm()
  o.Using("default")

  var articles []*models.Article
  num, err := o.QueryTable("articles").All(&articles)

  if err != orm.ErrNoRows && num > 0 {
    manage.Data["records"] = articles
  }
}

func (manage *ManageController) Add() {
  manage.Data["Form"] = &models.Article{}
  manage.Layout = "basic-layout.tpl"
  manage.LayoutSections = make(map[string]string)
  manage.LayoutSections["Header"] = "header.tpl"
  manage.LayoutSections["Footer"] = "footer.tpl"
  manage.TplName = "manage/add.tpl"

  flash := beego.ReadFromRequest(&manage.Controller)

  if ok := flash.Data["error"]; ok != "" {
    // Display error messages
    manage.Data["flash"] = ok
  }

  o := orm.NewOrm()
  o.Using("default")

  article := models.Article{}

  if err := manage.ParseForm(&article); err != nil {
    beego.Error("Could not parse the form. Reason: ", err)
  } else {
    manage.Data["Articles"] = article
    valid := validation.Validation{}
    isValid, _ := valid.Valid(article)

    if manage.Ctx.Input.Method() == "POST" {
      if !isValid {
        manage.Data["Errors"] = valid.ErrorsMap
        beego.Error("Form did not validate.")
      } else {
        searchArticle := models.Article{Name: article.Name}
        beego.Debug("Article name supplied:", article.Name)
        err := o.Read(&searchArticle)
        beego.Debug("Err:", err)
        flash := beego.NewFlash()

        if err == orm.ErrNoRows || err == orm.ErrMissPK {
          beego.Debug("No article found matching details supplied. Attempt to insert article: ", article)
          id, err := o.Insert(&article)
          if err == nil {
            msg := fmt.Sprintf("Article inserted with id: ", id)
            beego.Debug(msg)
            flash.Notice(msg)
            flash.Store(&manage.Controller)
          } else {
            msg := fmt.Sprintf("Could not insert new article. Reason: ", err)
            beego.Debug(msg)
            flash.Error(msg)
            flash.Store(&manage.Controller)
          }
        } else {
          beego.Debug("Article found matching details supplied. Cannot insert.")
        }
      }
    }
  }
}
