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

func (manage *ManageController) Home() {}

func (manage *ManageController) Delete() {}

func (manage *ManageController) Update() {}

func (manage *ManageController) View() {}

func (manage *ManageController) Add() {}
