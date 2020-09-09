package model

import (
	"belajargolang/billingrest/lib"
	"database/sql"
	"fmt"
)

type Member struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     string `json:"status"`
	NamaMember string `json:"nama_member"`
	NoHp       string `json:"no_hp"`
	Email      string `json:"email"`
	Alamat     string `json:"alamat"`
	Foto       string `json:"foto"`
	IDMember   string `json:"id_member"`
	Flag       string `json:"flag"`
	Nik        string `json:"nik"`
}

var TbMember = `
	CREATE TABLE tbmember
	(
		id_member serial primary key,
		username varchar(50),
		password varchar(30),
		status smallint,
		nama_member varchar(30),
		no_hp bigint,
		email varchar(30),
		alamat varchar(225),
		foto varchar(225),
		flag int,
		nik bigint
	);
`

//fungi deklarasi nama tabel
func (m *Member) Name() string {
	return "tbmember"
}

//show field
func (m *Member) Field() (fields []string, dst []interface{}) {
	fields = []string{"id_member", "username", "password", "status", "nama_member", "no_hp", "email", "alamat", "foto", "flag", "nik"}
	dst = []interface{}{&m.IDMember, &m.Username, &m.Password, &m.Status, &m.NamaMember, &m.NoHp, &m.Email, &m.Alamat, &m.Foto, &m.Flag, &m.Nik}
	return fields, dst
}

//inisialisai primary key kalo ada
func (m *Member) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id_member"}
	dst = []interface{}{&m.IDMember}
	return fields, dst
}

//conect table
func (m *Member) Structur() lib.Table {
	return &Member{}
}

// auto number
func (m *Member) AutoNumber() bool {
	return true
}

//insert Member
func (m *Member) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}

func (m *Member) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func (m *Member) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *Member) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

//ambil semua data Member
func GetAllMember(db *sql.DB, params ...string) ([]*Member, error) {
	m := &Member{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	mmbr := make([]*Member, len(data))
	for index, item := range data {
		mmbr[index] = item.(*Member)
	}
	return mmbr, nil
}

func GetMemberFromAnother(db *sql.DB, id string) error {
	params := fmt.Sprintf("id_member,=,%v", id)
	// fmt.Println(params)
	res, err := GetAllInvoice(db, params)
	if err != nil {
		return err
	}

	if len(res) != 0 {
		// fmt.Println("adaw")
		for _, item := range res {
			fmt.Println(item)
		}
	}
	// fmt.Println("aduh")
	return nil
}
