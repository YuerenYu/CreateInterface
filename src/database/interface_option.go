package database

import (
	"CreateInterface/src/utils"
	"fmt"
)

func AddInterface() {
	var interfaceInfo InterfaceInfo

	utils.PrintAddPattern()
	_, err := fmt.Scanf("%s", &interfaceInfo.Pattern)
	utils.CheckError(err)

	utils.PrintAddContent()
	_, err = fmt.Scanf("%s", &interfaceInfo.Content)
	utils.CheckError(err)

	db, err := OpenDB()
	defer CloseDB(db)
	id := InsertInfo(db, interfaceInfo)
	if id > 0 {
		utils.PrintAddDone()
	}
}

func ListInterface() {
	var interfaceInfo InterfaceInfo
	db, err := OpenDB()
	utils.CheckError(err)
	defer CloseDB(db)

	rows, err := db.Query("select * from interface_info")
	utils.CheckError(err)

	for rows.Next() {
		err = rows.Scan(&interfaceInfo.Id, &interfaceInfo.Pattern, &interfaceInfo.Content, &interfaceInfo.InsertTime)
		fmt.Println(interfaceInfo)
	}

	err = rows.Close()
	utils.CheckError(err)
}

func DeleteInterface() {
	ListInterface()
	fmt.Println("Input the id you want to delete.")
	var id int
	_, err := fmt.Scanf("%d", &id)
	utils.CheckError(err)

	db, err := OpenDB()
	utils.CheckError(err)
	defer CloseDB(db)

	stmt, err := db.Prepare("delete from interface_info where id = ?")
	utils.CheckError(err)

	res, err := stmt.Exec(id)
	utils.CheckError(err)

	affect, err := res.RowsAffected()

	if affect > 0 {
		fmt.Printf("%dth's interface has been deleted\n", id)
		utils.PrintDeleteDone()
	}
}
