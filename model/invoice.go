package model

import (
	"belajargolang/billingrest/lib"
	"database/sql"
	"time"
)

type Invoice struct {
	IDInvoice       string    `json:"id_invoice"`
	IDRoom          string    `json:"id_room"`
	IDPeriode       string    `json:"id_periode"`
	IDMember        string    `json:"id_member"`
	Quantity        string    `json:"quantity"`
	Charge          string    `json:"charge"`
	Description     string    `json:"description"`
	Total           string    `json:"total"`
	TransactionDate time.Time `json:"transaction_date"`
	StatusBayar     string    `json:"status_bayar"`
	JumlahBayar     string    `json:"jumlah_bayar"`
	DariTgl         time.Time `json:"dari_tgl"`
	SampaiTgl       time.Time `json:"sampai_tgl"`
	Pembayar        string    `json:"pembayar"`
	Petugas         string    `json:"petugas"`
}

var TbInvoice = `
	CREATE TABLE tbinvoice
	(
		id_invoice serial primary key,
		id_room int,
		id_periode int,
		id_member int,
		quantity int,
		charge int,
		description varchar(225),
		total int,
		transaction_date DATE,
		status_bayar smallint,
		jumlah_bayar bigint,
		dari_tgl DATE,
		sampai_tgl DATE,
		pembayar varchar(30),
		petugas varchar(30)
	);
`

//fungi deklarasi nama tabel
func (m *Invoice) Name() string {
	return "tbinvoice"
}

//show field
func (m *Invoice) Field() (fields []string, dst []interface{}) {
	fields = []string{"id_invoice", "id_room", "id_periode", "id_member", "quantity", "charge", "description", "total", "transaction_date", "status_bayar", "jumlah_bayar", "dari_tgl", "sampai_tgl", "pembayar", "petugas"}
	dst = []interface{}{&m.IDInvoice, &m.IDRoom, &m.IDPeriode, &m.IDMember, &m.Quantity, &m.Charge, &m.Description, &m.Total, &m.TransactionDate, &m.StatusBayar, &m.JumlahBayar, &m.DariTgl, &m.SampaiTgl, &m.Pembayar, &m.Petugas}
	return fields, dst
}

//inisialisai primary key kalo ada
func (m *Invoice) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"id_invoice"}
	dst = []interface{}{&m.IDInvoice}
	return fields, dst
}

//conect table
func (m *Invoice) Structur() lib.Table {
	return &Invoice{}
}

// auto number
func (m *Invoice) AutoNumber() bool {
	return true
}

//insert Invoice
func (m *Invoice) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}

func (m *Invoice) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func (m *Invoice) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *Invoice) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

//ambil semua data Invoice
func GetAllInvoice(db *sql.DB, params ...string) ([]*Invoice, error) {
	m := &Invoice{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	invoice := make([]*Invoice, len(data))
	for index, item := range data {
		invoice[index] = item.(*Invoice)
	}
	return invoice, nil
}
