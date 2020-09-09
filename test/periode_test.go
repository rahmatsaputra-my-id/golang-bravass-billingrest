package test

import (
	"belajargolang/billingrest/model"
	"fmt"
	"testing"
	"time"
)

func TestInsertPeriode(t *testing.T) {
	var dataInsertPeriode = []model.Periode{
		// db := PrepareTest(t)
		// defer db.Close()
		// data := []*model.User{
		model.Periode{Periode: "periode 1", IDRoom: "11"},
		model.Periode{Periode: "periode 2", IDRoom: "12"},
		model.Periode{Periode: "periode 3", IDRoom: "13"},
	}
	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	t.Run("Testing Insert Get Periode", func(t *testing.T) {
		for _, dataInsert := range dataInsertPeriode {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Periode{IDPeriode: dataInsert.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, dataInsert)
		}
	})

	t.Run("Testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"periode": "20",
			"id_room": "5",
		}

		dataUpdate := model.Periode{IDPeriode: "2"}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := model.GetAllPeriode(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Periode{IDPeriode: item.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, *item)
		}
	})

	t.Run("Testing Gets with Paramaters", func(t *testing.T) {
		params := "periode,=,DIA"
		result, err := model.GetAllPeriode(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Periode{IDPeriode: item.IDPeriode}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			comparePeriode(t, got, *item)
		}

	})
	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Periode{IDPeriode: "2"}
		// m := model.Periode{ID: dataInsertPeriode[0].ID}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})

}

func TestGetPeriode(t *testing.T) {
	var dataInsertInvoice = []model.Invoice{
		model.Invoice{IDRoom: "2", IDPeriode: "3", IDMember: "2", Quantity: "30", Charge: "10000", Description: "urutan invoice pertama", Total: "20000000", TransactionDate: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), StatusBayar: "1", JumlahBayar: "7000000", DariTgl: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), SampaiTgl: time.Date(2019, 11, 22, 0, 0, 0, 0, time.UTC), Pembayar: "rahmat", Petugas: "teo"},
	}

	var dataInsertPeriode = []model.Periode{
		model.Periode{Periode: "periode 1", IDRoom: "11"},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	for _, dataInsert := range dataInsertInvoice {
		err := dataInsert.Insert(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, dataInsert := range dataInsertPeriode {
		err := dataInsert.Insert(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	if err := model.GetPeriodeFromAnother(db, dataInsertInvoice[0].IDPeriode); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Sukses brooo Periode")
}

func comparePeriode(t *testing.T, got, want model.Periode) {
	if got.Periode != want.Periode {
		t.Fatalf("got : %s want :%s periode tidak sama", got.Periode, want.Periode)
	}
	if got.IDRoom != want.IDRoom {
		t.Fatalf("got : %s want :%s id_room tidak sama", got.IDRoom, want.IDRoom)
	}

}
