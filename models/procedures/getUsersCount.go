package procedures

type GetUsersCount struct {
	Count int  `json:"rate"`
}

func (GetUsersCount) create() error {
	//sql := `CREATE PROCEDURE GetUsersCount()
	//BEGIN
	//  SELECT SUM(followers_count) as followers FROM modules where modules.status = 1;
	//END`
	//return app.Application.DB.Exec(sql).Error
	return nil
}

func (GetUsersCount) drop() error {
	//sql := `DROP PROCEDURE GetUsersCount;`
	//return app.Application.DB.Exec(sql).Error
	return nil
}

func (GetUsersCount) dropIfExist() error {
	//sql := `DROP PROCEDURE IF EXISTS GetUsersCount;`
	//return app.Application.DB.Exec(sql).Error
	return nil
}

func GetSumModulesFollowersCount() GetUsersCount {
	var returnVal GetUsersCount
	//app.Application.DB.Raw("CALL GetSumModulesFollowersCount()").Scan(&returnVal)
	return returnVal
}

