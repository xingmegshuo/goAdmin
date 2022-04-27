package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetBoxNftTable(ctx *context.Context) table.Table {

	boxNft := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "main"))

	info := boxNft.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)

	info.AddField("ID", "id", db.Bigint).FieldSortable()
	info.AddField("盲盒名称", "name", db.Varchar).FieldEditAble(editType.Text).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("盲盒描述", "description", db.Text).FieldEditAble(editType.Text)
	info.AddField("图片", "img_url", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().Image().
			SetSrc(template.HTML(value.Value)).
			SetHeight("120").SetWidth("120").WithModal().GetContent()
	})
	info.AddField("Storage_id", "storage_id", db.Bigint)
	info.AddField("价格", "price", db.Varchar).FieldEditAble(editType.Text)
	info.AddField("Usdt_price", "usdt_price", db.Decimal)
	info.AddField("支付代币合约地址", "paytoken_address", db.Varchar).FieldEditAble(editType.Text)
	info.AddField("支付代币", "paytoken_name", db.Varchar).FieldEditAble(editType.Text)
	info.AddField("售卖数量💩", "sell_quantity", db.Bigint).FieldEditAble(editType.Text)
	info.AddField("总数量", "quantity", db.Bigint).FieldEditAble(editType.Text)
	info.AddField("签名", "signature", db.Varchar)
	info.AddField("所属人", "owner", db.Varchar)
	info.AddField("是否在售", "onsell", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "未售卖🙅"
		}
		if model.Value == "1" {
			return "售卖中🙆‍♂️"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "未售卖🙅"},
		{Value: "1", Text: "售卖中🙆‍♂️"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "未售卖🙅"},
		{Value: "1", Text: "售卖中🙆‍♂️"},
	})
	info.AddField("是否删除", "deleted", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "未删除✔️"
		}
		if model.Value == "1" {
			return "删除✖️"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "未删除✔️"},
		{Value: "1", Text: "删除✖️"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "未删除✔️"},
		{Value: "1", Text: "删除✖️"},
	})
	info.AddField("创建时间", "create_time", db.Bigint)
	info.AddField("开始时间", "start_time", db.Bigint).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("结束时间", "end_time", db.Bigint).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("Update_time", "update_time", db.Bigint)
	info.AddField("Category_id", "category_id", db.Bigint)
	info.AddField("合约地址", "address", db.Varchar).FieldEditAble(editType.Text)
	info.AddField("同步", "is_sync", db.Tinyint).FieldEditAble(editType.Text)

	info.SetTable("box_nft").SetTitle("盲盒列表").SetDescription("盲盒管理")

	formList := boxNft.GetForm()
	formList.AddField("ID", "id", db.Bigint, form.Default)
	formList.AddField("盲盒名称", "name", db.Varchar, form.Text)
	formList.AddField("盲盒简介", "description", db.Text, form.RichText)
	formList.AddField("图片", "img_url", db.Varchar, form.File)
	formList.AddField("图片id", "storage_id", db.Bigint, form.Number)
	formList.AddField("价格", "price", db.Varchar, form.Text)
	formList.AddField("usdt 价格", "usdt_price", db.Decimal, form.Text)
	formList.AddField("支付代币合约地址", "paytoken_address", db.Varchar, form.Text)
	formList.AddField("支付代币符号", "paytoken_name", db.Varchar, form.Text)
	formList.AddField("售卖数量", "sell_quantity", db.Bigint, form.Number)
	formList.AddField("售卖总量", "quantity", db.Bigint, form.Number)
	formList.AddField("签名", "signature", db.Varchar, form.Text)
	formList.AddField("发起人", "owner", db.Varchar, form.Text)
	formList.AddField("是否在售", "onsell", db.Tinyint, form.Radio).FieldOptions(types.FieldOptions{
		{Text: "未售卖🙅", Value: "0"},
		{Text: "售卖中🙆‍♂️", Value: "1"},
	}).FieldDefault("1")
	formList.AddField("是否删除", "deleted", db.Tinyint, form.Radio).FieldOptions(types.FieldOptions{
		{Text: "未删除✔️", Value: "0"},
		{Text: "删除✖️", Value: "1"},
	}).FieldDefault("0")
	formList.AddField("创建时间", "create_time", db.Bigint, form.Number)
	formList.AddField("开始时间", "start_time", db.Bigint, form.Number)
	formList.AddField("结束时间", "end_time", db.Bigint, form.Number)
	formList.AddField("更新时间", "update_time", db.Bigint, form.Number)
	formList.AddField("分类", "category_id", db.Bigint, form.Number)
	formList.AddField("合约地址", "address", db.Varchar, form.Text)
	formList.AddField("同步", "is_sync", db.Tinyint, form.Number)

	formList.SetTable("box_nft").SetTitle("盲盒详情").SetDescription("盲盒详细信息")

	return boxNft
}
