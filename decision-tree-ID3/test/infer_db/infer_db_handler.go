package infer_db

import (
	"bufio"
	"bytes"
	"database/sql"
	"io"
	"os"
	"regexp"
	"strconv"
	"fmt"
	"errors"
)

type CDbHandler struct  {
	m_db *sql.DB
}

func (this *CDbHandler) Connect(host string, port uint, username string, userpwd string, dbname string, dbtype string) (err error) {
	b := bytes.Buffer{}
	b.WriteString(username)
	b.WriteString(":")
	b.WriteString(userpwd)
	b.WriteString("@tcp(")
	b.WriteString(host)
	b.WriteString(":")
	b.WriteString(strconv.FormatUint(uint64(port), 10))
	b.WriteString(")/")
	b.WriteString(dbname)
	var name string
	if dbtype == "mysql" {
		name = b.String()
	} else if dbtype == "sqlite3" {
		name = dbname
	} else {
		return errors.New("dbtype not support")
	}
	this.m_db, err = sql.Open(dbtype, name)
	if err != nil {
		return err
	}
	this.m_db.SetMaxOpenConns(2000)
	this.m_db.SetMaxIdleConns(1000)
	this.m_db.Ping()
	return nil
}

func (this *CDbHandler) ConnectByRule(rule string, dbtype string) (err error) {
	this.m_db, err = sql.Open(dbtype, rule)
	if err != nil {
		return err
	}
	this.m_db.SetMaxOpenConns(2000)
	this.m_db.SetMaxIdleConns(1000)
	this.m_db.Ping()
	return nil
}

func (this *CDbHandler) ConnectByCfg(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var host string = "localhost"
	var port uint = 3306
	var username string = "root"
	var userpwd string = "123456"
	var dbname string = "test"
	var dbtype string = "mysql"
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		content := string(a)
		r, _ := regexp.Compile("(.*)?=(.*)?")
		ret := r.FindStringSubmatch(content)
		if len(ret) != 3 {
			continue
		}
		k := ret[1]
		v := ret[2]
		switch k {
		case "host":
			host = v
		case "port":
			port_tmp, _ := strconv.ParseUint(v, 10, 32)
			port = uint(port_tmp)
		case "username":
			username = v
		case "userpwd":
			userpwd = v
		case "dbname":
			dbname = v
		case "dbtype":
			dbtype = v
		}
	}
	return this.Connect(host, port, username, userpwd, dbname, dbtype)
}

func (this *CDbHandler) Disconnect() {
	this.m_db.Close()
}

func (this *CDbHandler) Create() (error) {
	var err error = nil
	var _ error = err
	_, err = this.m_db.Exec(`create table if not exists t_data (
    data_uuid varchar(64),
    lable_kv json,
    classify varchar(128),
    create_time datetime,
    update_time datetime
);`)
	if err != nil {
		return err
	}
	_, err = this.m_db.Exec(`
create table if not exists t_tree (
    tree json
);`)
	if err != nil {
		return err
	}
	return nil
}

func (this *CDbHandler) AddData(input0 *[]CAddDataInput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	for _, v := range *input0 {
		result, err = this.m_db.Exec(fmt.Sprintf(`insert into t_data values(?, ?, ?, ?, ?);`), v.DataUuid, v.LableKV, v.Classify, v.CreateTime, v.UpdateTime)
		if err != nil {
			tx.Rollback()
			return err, rowCount
		}
		var _ = result
	}
	tx.Commit()
	return nil, rowCount
}

func (this *CDbHandler) UpdateData(input0 *[]CUpdateDataInput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	for _, v := range *input0 {
		result, err = this.m_db.Exec(fmt.Sprintf(`update t_data set lable_kv = ?, classify = ?, update_time = ?
where data_uuid = ?;`), v.LableKV, v.Classify, v.UpdateTime, v.DataUuid)
		if err != nil {
			tx.Rollback()
			return err, rowCount
		}
		var _ = result
	}
	tx.Commit()
	return nil, rowCount
}

func (this *CDbHandler) UpdateThenAddData(input0 *[]CUpdateThenAddDataInput, input1 *[]CAddDataInput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	for _, v := range *input0 {
		result, err = this.m_db.Exec(fmt.Sprintf(`update t_data set lable_kv = ?, classify = ?, update_time = ?
where data_uuid = ?;`), v.LableKV, v.Classify, v.UpdateTime, v.DataUuid)
		if err != nil {
			tx.Rollback()
			return err, rowCount
		}
		var _ = result
	}
	for _, v := range *input1 {
		result, err = this.m_db.Exec(fmt.Sprintf(`insert into t_data values(?, ?, ?, ?, ?);`), v.DataUuid, v.LableKV, v.Classify, v.CreateTime, v.UpdateTime)
		if err != nil {
			tx.Rollback()
			return err, rowCount
		}
		var _ = result
	}
	tx.Commit()
	return nil, rowCount
}

func (this *CDbHandler) GetAllData(output0 *[]CGetAllDataOutput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	rows0, err := this.m_db.Query(fmt.Sprintf(`select data_uuid, lable_kv, classify from t_data;`))
	if err != nil {
		tx.Rollback()
		return err, rowCount
	}
	tx.Commit()
	defer rows0.Close()
	for rows0.Next() {
		rowCount += 1
		tmp := CGetAllDataOutput{}
		var dataUuid sql.NullString
		var lableKV sql.NullString
		var classify sql.NullString
		scanErr := rows0.Scan(&dataUuid, &lableKV, &classify)
		if scanErr != nil {
			continue
		}
		tmp.DataUuid = dataUuid.String
		tmp.DataUuidIsValid = dataUuid.Valid
		tmp.LableKV = lableKV.String
		tmp.LableKVIsValid = lableKV.Valid
		tmp.Classify = classify.String
		tmp.ClassifyIsValid = classify.Valid
		*output0 = append(*output0, tmp)
	}
	return nil, rowCount
}

func (this *CDbHandler) GetOneData(output0 *CGetOneDataOutput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	rows0, err := this.m_db.Query(fmt.Sprintf(`select lable_kv, classify from t_data
limit 1;`))
	if err != nil {
		tx.Rollback()
		return err, rowCount
	}
	tx.Commit()
	defer rows0.Close()
	for rows0.Next() {
		rowCount += 1
		var lableKV sql.NullString
		var classify sql.NullString
		scanErr := rows0.Scan(&lableKV, &classify)
		if scanErr != nil {
			continue
		}
		output0.LableKV = lableKV.String
		output0.LableKVIsValid = lableKV.Valid
		output0.Classify = classify.String
		output0.ClassifyIsValid = classify.Valid
	}
	return nil, rowCount
}

func (this *CDbHandler) UpdateTree(input0 *CUpdateTreeInput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	result, err = this.m_db.Exec(fmt.Sprintf(`delete from t_tree;
insert into t_tree values(?);`), input0.Tree)
	if err != nil {
		tx.Rollback()
		return err, rowCount
	}
	tx.Commit()
	var _ = result
	return nil, rowCount
}

func (this *CDbHandler) GetTree(output0 *CGetTreeOutput) (error, uint64) {
	var rowCount uint64 = 0
	tx, _ := this.m_db.Begin()
	var result sql.Result
	var _ = result
	var err error
	var _ error = err
	rows0, err := this.m_db.Query(fmt.Sprintf(`select * from t_tree;`))
	if err != nil {
		tx.Rollback()
		return err, rowCount
	}
	tx.Commit()
	defer rows0.Close()
	for rows0.Next() {
		rowCount += 1
		var tree sql.NullString
		scanErr := rows0.Scan(&tree)
		if scanErr != nil {
			continue
		}
		output0.Tree = tree.String
		output0.TreeIsValid = tree.Valid
	}
	return nil, rowCount
}

