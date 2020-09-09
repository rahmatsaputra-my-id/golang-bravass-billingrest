package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
	"time"
)

func TestInsertInvoice(t *testing.T) {
	var dataInsertInvoice = []model.Invoice{

		model.Invoice{IDRoom: "2", IDPeriode: "3", IDMember: "2", Quantity: "30", Charge: "10000", Description: "urutan invoice pertama", Total: "20000000", TransactionDate: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), StatusBayar: "1", JumlahBayar: "7000000", DariTgl: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), SampaiTgl: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), Pembayar: "rahmat", Petugas: "teo"},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Invoice", func(t *testing.T) {
		for _, dataInsert := range dataInsertInvoice {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Invoice{IDInvoice: dataInsert.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"id_room":          "2",
			"id_periode":       "5",
			"id_member":        "3",
			"quantity":         "40",
			"charge":           "700000",
			"description":      "keterangan update test",
			"total":            "80000000",
			"transaction_date": time.Now(),
			"status_bayar":     "1",
			"jumlah_bayar":     "70000000",
			"dari_tgl":         time.Now(),
			"sampai_tgl":       time.Now(),
			"pembayar":         "rahmat",
			"petugas":          "teo",
		}

		dataUpdate := model.Invoice{IDInvoice: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllInvoice(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Invoice{IDInvoice: item.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "id_room,=,2"
		result, err := model.GetAllInvoice(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Invoice{IDInvoice: item.IDInvoice}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareInvoice(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Invoice{IDInvoice: "2"}
		// m := model.User{ID: dataInsertUser[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func compareInvoice(t *testing.T, got, want model.Invoice) {
	if got.IDInvoice != want.IDInvoice {
		t.Fatalf("got : %s want :%s id_invoice tidak sama", got.IDInvoice, want.IDInvoice)
	}
	if got.IDRoom != want.IDRoom {
		t.Fatalf("got :%s want :%s id_room tidak Sama", got.IDRoom, want.IDRoom)
	}
	if got.IDMember != want.IDMember {
		t.Fatalf("got :%s want :%s id member tidak Sama", got.IDMember, want.IDMember)
	}
	if got.Quantity != want.Quantity {
		t.Fatalf("got :%s want :%s quantity tidak Sama", got.Quantity, want.Quantity)
	}
	if got.Charge != want.Charge {
		t.Fatalf("got :%s want :%s charge tidak Sama", got.Charge, want.Charge)
	}
	if got.Description != want.Description {
		t.Fatalf("got :%s want :%s description tidak Sama", got.Description, want.Description)
	}
	if got.Total != want.Total {
		t.Fatalf("got :%s want :%s total tidak Sama", got.Total, want.Total)
	}
	if got.StatusBayar != want.StatusBayar {
		t.Fatalf("got :%s want :%s status bayar tidak Sama", got.StatusBayar, want.StatusBayar)
	}
	if got.JumlahBayar != want.JumlahBayar {
		t.Fatalf("got :%s want :%s jumlah bayar tidak Sama", got.JumlahBayar, want.JumlahBayar)
	}
	if got.Pembayar != want.Pembayar {
		t.Fatalf("got :%s want :%s pembayar tidak Sama", got.Pembayar, want.Pembayar)
	}
	if got.Petugas != want.Petugas {
		t.Fatalf("got :%s want :%s petugas tidak Sama", got.Petugas, want.Petugas)
	}
	// if got.TransactionDate != want.TransactionDate {
	// 	t.Fatalf("got :%v want :%v transaction_date tidak Sama", got.TransactionDate, want.TransactionDate)
	// }
	if got.IDPeriode != want.IDPeriode {
		t.Fatalf("got :%s want :%s id_periode tidak Sama", got.IDPeriode, want.IDPeriode)
	}
}
