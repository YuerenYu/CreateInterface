package database

import (
	"CreateInterface/src/utils"
	"bufio"
	"fmt"
	"os"
)

func AddInterface() {
	var interfaceInfo InterfaceInfo

	utils.PrintAddPattern()
	reader := bufio.NewReader(os.Stdin)
	var err error
	interfaceInfo.Pattern, err = reader.ReadString('$')
	interfaceInfo.Pattern = interfaceInfo.Pattern[:len(interfaceInfo.Pattern)-1]
	utils.CheckError(err)

	utils.PrintAddContent()
	reader = bufio.NewReader(os.Stdin)
	interfaceInfo.Content, err = reader.ReadString('$')
	interfaceInfo.Content = interfaceInfo.Content[:len(interfaceInfo.Content)-1]
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
