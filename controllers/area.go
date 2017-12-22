package controllers

import (
	"zd112_backstage/models"
	"time"
)

type ContinentController struct {
	BaseWebController
}

func (this *ContinentController) List() {
	this.pageTitle("洲列表")
	this.display(this.getBgAreaAction("continent/list"))
}

func (this *ContinentController) Add() {
	this.pageTitle("新增洲名称")
	this.display(this.getBgAreaAction("continent/add"))
}

func (this *ContinentController) Edit() {
	this.pageTitle("编辑洲名称")
	continent := new(models.Continent)
	continent.Id = this.getId64(0)
	if err := continent.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.row(nil, continent,true)
	this.display(this.getBgAreaAction("continent/edit"))
}

func (this *ContinentController) AjaxSave() {
	continent := new(models.Continent)
	continent.Id = this.getId64(0)
	continent.Name = this.getString("name", "名称不能为空!", 1)
	continent.Icon = this.getString("file", "Icon不能为空!", defaultMinSize)
	var err error
	if continent.Id == 0 {
		continent.CreateId = this.userId
		continent.CreateTime = time.Now().Unix()
		_, err = continent.Add()
	} else {
		continent.UpdateId = this.userId
		continent.UpdateTime = time.Now().Unix()
		_, err = continent.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ContinentController) Table() {
	continent := new(models.Continent)
	result, count := continent.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		this.parse(list, nil, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ContinentController) AjaxDel() {
	continent := new(models.Continent)
	continent.Id = this.getId64(0)
	if _, err := continent.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type StateController struct {
	BaseWebController
}

func (this *StateController) List() {
	this.pageTitle("国列表")
	this.display(this.getBgAreaAction("state/list"))
}

func (this *StateController) Add() {
	this.pageTitle("增加国名称")
	this.parent(0)
	this.display(this.getBgAreaAction("state/add"))
}

func (this *StateController) Edit() {
	this.pageTitle("编辑国名")
	state := new(models.State)
	state.Id = this.getId64(0)
	if err := state.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(state.ParentId)
	this.row(nil, state,true)
	this.display(this.getBgAreaAction("state/edit"))
}

func (this *StateController) AjaxSave() {
	state := new(models.State)
	state.Name = this.getString("name", "名称不能为空!", 1)
	state.Icon = this.getString("file", "Icon不能为空!", defaultMinSize)
	state.Id = this.getId64(0)
	state.ParentId = this.getInt64("group_id", 0)
	var err error
	if state.Id == 0 {
		state.CreateId = this.userId
		state.CreateTime = time.Now().Unix()
		_, err = state.Add()
	} else {
		state.UpdateId = this.userId
		state.UpdateTime = time.Now().Unix()
		_, err = state.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *StateController) parent(id int64) {
	continent := new(models.Continent)
	result, count := continent.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *StateController) Table() {
	state := new(models.State)
	result, count := state.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		continent := new(models.Continent)
		continent.Id = v.ParentId
		if err := continent.Query(); err == nil {
			row["Parent"] = continent.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *StateController) AjaxDel() {
	state := new(models.State)
	state.Id = this.getId64(0)
	if _, err := state.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type ProvinceController struct {
	BaseWebController
}

func (this *ProvinceController) List() {
	this.pageTitle("省列表")
	this.display(this.getBgAreaAction("province/list"))
}

func (this *ProvinceController) Add() {
	this.pageTitle("增加省名称")
	this.parent(0)
	this.display(this.getBgAreaAction("province/add"))
}

func (this *ProvinceController) Edit() {
	this.pageTitle("编辑省名称")
	province := new(models.Province)
	province.Id = this.getId64(0)
	if err := province.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(province.ParentId)
	this.row(nil, province,true)
	this.display(this.getBgAreaAction("province/edit"))
}

func (this *ProvinceController) AjaxSave() {
	province := new(models.Province)
	province.Name = this.getString("name", "名称不能为空!", 1)
	province.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	province.Id = this.getId64(0)
	province.ParentId = this.getInt64("group_id", 0)
	var err error
	if province.Id == 0 {
		province.CreateId = this.userId
		province.CreateTime = time.Now().Unix()
		_, err = province.Add()
	} else {
		province.UpdateId = this.userId
		province.UpdateTime = time.Now().Unix()
		_, err = province.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProvinceController) parent(id int64) {
	state := new(models.State)
	result, count := state.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *ProvinceController) Table() {
	province := new(models.Province)
	result, count := province.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		state := new(models.State)
		state.Id = v.ParentId
		if err := state.Query(); err == nil {
			row["Parent"] = state.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProvinceController) AjaxDel() {
	province := new(models.Province)
	province.Id = this.getId64(0)
	if _, err := province.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CityController struct {
	BaseWebController
}

func (this *CityController) List() {
	this.pageTitle("城市列表")
	this.display(this.getBgAreaAction("city/list"))
}

func (this *CityController) Add() {
	this.pageTitle("增加城市名称")
	this.parent(0)
	this.display(this.getBgAreaAction("city/add"))
}

func (this *CityController) Edit() {
	this.pageTitle("编辑城市名称")
	city := new(models.City)
	city.Id = this.getId64(0)
	if err := city.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(city.ParentId)
	this.row(nil, city,true)
	this.display(this.getBgAreaAction("city/edit"))
}

func (this *CityController) AjaxSave() {
	city := new(models.City)
	city.Name = this.getString("name", "名称不能为空!", 1)
	city.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	city.Id = this.getId64(0)
	city.ParentId = this.getInt64("group_id", 0)
	var err error
	if city.Id == 0 {
		city.CreateId = this.userId
		city.CreateTime = time.Now().Unix()
		_, err = city.Add()
	} else {
		city.UpdateId = this.userId
		city.UpdateTime = time.Now().Unix()
		_, err = city.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CityController) parent(id int64) {
	province := new(models.Province)
	result, count := province.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *CityController) Table() {
	city := new(models.City)
	result, count := city.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		province := new(models.Province)
		province.Id = v.ParentId
		if err := province.Query(); err == nil {
			row["Parent"] = province.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CityController) AjaxDel() {
	city := new(models.City)
	city.Id = this.getId64(0)
	if _, err := city.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type RegionController struct {
	BaseWebController
}

func (this *RegionController) List() {
	this.pageTitle("地区列表")
	this.display(this.getBgAreaAction("region/list"))
}

func (this *RegionController) Add() {
	this.pageTitle("增加地区名称")
	this.parent(0)
	this.display(this.getBgAreaAction("region/add"))
}

func (this *RegionController) Edit() {
	this.pageTitle("编辑地区名称")
	region := new(models.Region)
	region.Id = this.getId64(0)
	if err := region.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(region.ParentId)
	this.row(nil, region,true)
	this.display(this.getBgAreaAction("region/edit"))
}

func (this *RegionController) AjaxSave() {
	region := new(models.Region)
	region.Name = this.getString("name", "名称不能为空!", 1)
	region.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	region.Id = this.getId64(0)
	region.ParentId = this.getInt64("group_id", 0)
	var err error
	if region.Id == 0 {
		region.CreateId = this.userId
		region.CreateTime = time.Now().Unix()
		_, err = region.Add()
	} else {
		region.UpdateId = this.userId
		region.UpdateTime = time.Now().Unix()
		_, err = region.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *RegionController) parent(id int64) {
	city := new(models.City)
	result, count := city.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *RegionController) Table() {
	region := new(models.Region)
	result, count := region.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		city := new(models.City)
		city.Id = v.ParentId
		if err := city.Query(); err == nil {
			row["Parent"] = city.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *RegionController) AjaxDel() {
	region := new(models.Region)
	region.Id = this.getId64(0)
	if _, err := region.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CountyController struct {
	BaseWebController
}

func (this *CountyController) List() {
	this.pageTitle("县列表")
	this.display(this.getBgAreaAction("county/list"))
}

func (this *CountyController) Add() {
	this.pageTitle("增加县名称")
	this.parent(0)
	this.display(this.getBgAreaAction("county/add"))
}

func (this *CountyController) Edit() {
	county := new(models.County)
	county.Id = this.getId64(0)
	if err := county.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(county.ParentId)
	this.row(nil, county,true)
	this.display(this.getBgAreaAction("county/edit"))
}

func (this *CountyController) AjaxSave() {
	county := new(models.County)
	county.Name = this.getString("name", "名称不能为空!", 1)
	county.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	county.Id = this.getId64(0)
	county.ParentId = this.getInt64("group_id", 0)
	var err error
	if county.Id == 0 {
		county.CreateId = this.userId
		county.CreateTime = time.Now().Unix()
		_, err = county.Add()
	} else {
		county.UpdateId = this.userId
		county.UpdateTime = time.Now().Unix()
		_, err = county.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountyController) parent(id int64) {
	city := new(models.City)
	result, count := city.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *CountyController) Table() {
	county := new(models.County)
	result, count := county.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		region := new(models.Region)
		region.Id = v.ParentId
		if err := region.Query(); err == nil {
			row["Parent"] = region.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CountyController) AjaxDel() {
	county := new(models.County)
	county.Id = this.getId64(0)
	if _, err := county.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TownController struct {
	BaseWebController
}

func (this *TownController) List() {
	this.pageTitle("镇列表")
	this.display(this.getBgAreaAction("town/list"))
}

func (this *TownController) Add() {
	this.pageTitle("增加镇名称")
	this.parent(0)
	this.display(this.getBgAreaAction("town/add"))
}

func (this *TownController) Edit() {
	this.pageTitle("编辑镇名称")
	town := new(models.Town)
	town.Id = this.getId64(0)
	if err := town.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(town.ParentId)
	this.row(nil, town,true)
	this.display(this.getBgAreaAction("town/edit"))
}

func (this *TownController) AjaxSave() {
	town := new(models.Town)
	town.Name = this.getString("name", "名称不能为空!", 1)
	town.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	town.Id = this.getId64(0)
	town.ParentId = this.getInt64("group_id", 0)
	var err error
	if town.Id == 0 {
		town.CreateId = this.userId
		town.CreateTime = time.Now().Unix()
		_, err = town.Add()
	} else {
		town.UpdateId = this.userId
		town.UpdateTime = time.Now().Unix()
		_, err = town.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TownController) parent(id int64) {
	county := new(models.County)
	result, count := county.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *TownController) Table() {
	town := new(models.Town)
	result, count := town.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		county := new(models.County)
		county.Id = v.ParentId
		if err := county.Query(); err == nil {
			row["Parent"] = county.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TownController) AjaxDel() {
	town := new(models.Town)
	town.Id = this.getId64(0)
	if _, err := town.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type CountryController struct {
	BaseWebController
}

func (this *CountryController) List() {
	this.pageTitle("乡列表")
	this.display(this.getBgAreaAction("country/list"))
}

func (this *CountryController) Add() {
	this.pageTitle("增加乡名称")
	this.parent(0)
	this.display(this.getBgAreaAction("country/add"))
}

func (this *CountryController) Edit() {
	this.pageTitle("编辑乡名称")
	country := new(models.Country)
	country.Id = this.getInt64("id", 0)
	if err := country.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(country.ParentId)
	this.row(nil, country,true)
	this.display(this.getBgAreaAction("country/edit"))
}

func (this *CountryController) AjaxSave() {
	country := new(models.Country)
	country.Name = this.getString("name", "名称不能为空!", 1)
	country.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	country.Id = this.getId64(0)
	country.ParentId = this.getInt64("group_id", 0)
	var err error
	if country.Id == 0 {
		country.CreateId = this.userId
		country.CreateTime = time.Now().Unix()
		_, err = country.Add()
	} else {
		country.UpdateId = this.userId
		country.UpdateTime = time.Now().Unix()
		_, err = country.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *CountryController) parent(id int64) {
	town := new(models.Town)
	result, count := town.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *CountryController) Table() {
	country := new(models.Country)
	result, count := country.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		town := new(models.Town)
		town.Id = v.ParentId
		if err := town.Query(); err == nil {
			row["Parent"] = town.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *CountryController) AjaxDel() {
	country := new(models.Country)
	country.Id = this.getId64(0)
	if _, err := country.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type VillageController struct {
	BaseWebController
}

func (this *VillageController) List() {
	this.pageTitle("村列表")
	this.display(this.getBgAreaAction("village/list"))
}

func (this *VillageController) Add() {
	this.pageTitle("增加村名称")
	this.parent(0)
	this.display(this.getBgAreaAction("village/add"))
}

func (this *VillageController) Edit() {
	this.pageTitle("编辑村名称")
	village := new(models.Village)
	village.Id = this.getId64(0)
	if err := village.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(village.ParentId)
	this.row(nil, village,true)
	this.display(this.getBgAreaAction("village/edit"))
}

func (this *VillageController) AjaxSave() {
	village := new(models.Village)
	village.Name = this.getString("name", "名称不能为空!", 1)
	village.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	village.Id = this.getId64(0)
	village.ParentId = this.getInt64("group_id", 0)
	var err error
	if village.Id == 0 {
		village.CreateId = this.userId
		village.CreateTime = time.Now().Unix()
		_, err = village.Add()
	} else {
		village.UpdateId = this.userId
		village.UpdateTime = time.Now().Unix()
		_, err = village.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *VillageController) parent(id int64) {
	country := new(models.Country)
	result, count := country.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *VillageController) Table() {
	village := new(models.Village)
	result, count := village.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		country := new(models.Country)
		country.Id = v.ParentId
		if err := country.Query(); err == nil {
			row["Parent"] = country.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *VillageController) AjaxDel() {
	village := new(models.Village)
	village.Id = this.getId64(0)
	if _, err := village.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type GroupController struct {
	BaseWebController
}

func (this *GroupController) List() {
	this.pageTitle("组列表")
	this.display(this.getBgAreaAction("group/list"))
}

func (this *GroupController) Add() {
	this.pageTitle("增加组名称")
	this.parent(0)
	this.display(this.getBgAreaAction("group/add"))
}

func (this *GroupController) Edit() {
	this.pageTitle("编辑组名称")
	group := new(models.Group)
	group.Id = this.getId64(0)
	if err := group.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(group.ParentId)
	this.row(nil, group,true)
	this.display(this.getBgAreaAction("group/edit"))
}

func (this *GroupController) AjaxSave() {
	group := new(models.Group)
	group.Name = this.getString("name", "名称不能为空!", 1)
	group.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	group.Id = this.getId64(0)
	group.ParentId = this.getInt64("group_id", 0)
	var err error
	if group.Id == 0 {
		group.CreateId = this.userId
		group.CreateTime = time.Now().Unix()
		_, err = group.Add()
	} else {
		group.UpdateId = this.userId
		group.UpdateTime = time.Now().Unix()
		_, err = group.Update()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *GroupController) parent(id int64) {
	village := new(models.Village)
	result, count := village.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *GroupController) Table() {
	group := new(models.Group)
	result, count := group.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		village := new(models.Village)
		village.Id = v.ParentId
		if err := village.Query(); err == nil {
			row["Parent"] = village.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *GroupController) AjaxDel() {
	group := new(models.Group)
	group.Id = this.getId64(0)
	if _, err := group.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

type TeamController struct {
	BaseWebController
}

func (this *TeamController) List() {
	this.pageTitle("队列表")
	this.display(this.getBgAreaAction("team/list"))
}

func (this *TeamController) Add() {
	this.pageTitle("增加队名称")
	this.parent(0)
	this.display(this.getBgAreaAction("team/add"))
}

func (this *TeamController) Edit() {
	this.pageTitle("编辑队名称")
	team := new(models.Team)
	team.Id = this.getId64(0)
	if err := team.Query(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.parent(team.ParentId)
	this.row(nil, team,true)
	this.display(this.getBgAreaAction("team/edit"))
}

func (this *TeamController) AjaxSave() {
	team := new(models.Team)
	team.Id = this.getId64(0)
	team.Name = this.getString("name", "名称不能为空!", 1)
	team.Icon = this.getString("file", "File不能为空!", defaultMinSize)
	team.ParentId = this.getInt64("group_id", 0)
	var err error
	if team.Id == 0 {
		team.CreateId = this.userId
		team.CreateTime = time.Now().Unix()
		_, err = team.Add()
	} else {
		team.UpdateId = this.userId
		team.UpdateTime = time.Now().Unix()
	}
	if err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *TeamController) parent(id int64) {
	group := new(models.Group)
	result, count := group.List(-1, -1)
	list := make([]map[string]interface{}, count)
	for k, v := range result {
		this.group(list, nil, k, v, id,false)
	}
	this.Data["Group"] = list
}

func (this *TeamController) Table() {
	team := new(models.Team)
	result, count := team.List(this.pageSize, this.offSet)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{}, 0)
		group := new(models.Group)
		group.Id = v.ParentId
		if err := group.Query(); err == nil {
			row["Parent"] = group.Name
		}
		this.parse(list, row, k, v,false)
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *TeamController) AjaxDel() {
	team := new(models.Team)
	team.Id = this.getInt64("id", 0)
	if _, err := team.Del(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
