package beer_test

import (
	"database/sql"
	"testing"

	"github.com/JVFayad/go-beer-api/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := exempleBeer(1, "Heineken")

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	defer db.Close()

	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
}

func TestGet(t *testing.T) {
	b := exempleBeer(1, "Heineken")

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	defer db.Close()

	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}

	saved, err := service.Get(1)

	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %d", 1, saved.ID)
	}

}

func TestGetAll(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	defer db.Close()

	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)

	for i := 1; i < 3; i++ {
		name := "Heineken" + string(i)
		b := exempleBeer(int64(i), name)
		err = service.Store(b)
	}

	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}

	all_saved, err := service.GetAll()

	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if len(all_saved) != 2 {
		t.Fatalf("Falta de dados. Esperado %d, recebido %d", 2, len(all_saved))
	}
}

func TestUpdate(t *testing.T) {
	b := exempleBeer(1, "Heineken")

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	defer db.Close()

	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}

	new_name := "Heineken_new"

	new_b := exempleBeer(1, new_name)

	err = service.Update(new_b)

	if err != nil {
		t.Fatalf("Erro atualizando no banco de dados: %s", err.Error())
	}

	saved, err := service.Get(1)

	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}

	if saved.Name != "Heineken_new" {
		t.Fatalf("Nome inválido. Esperado %s, recebido %s", new_name, saved.Name)
	}

}

func TestRemove(t *testing.T) {
	b := exempleBeer(1, "Heineken")

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	defer db.Close()

	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}

	err = service.Remove(1)

	if err != nil {
		t.Fatalf("Erro removendo no banco de dados: %s", err.Error())
	}

	_, err = service.Get(1)

	if err == nil {
		t.Fatalf("O registro com id %d deveria ter sido removido", 1)
	}

}

func exempleBeer(id int64, name string) *beer.Beer {
	b := &beer.Beer{
		ID:    id,
		Name:  name,
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	return b
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from beer")
	tx.Commit()
	return err
}
