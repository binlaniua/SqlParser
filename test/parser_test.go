package test

import (
	"testing"
	"github.com/binlaniua/SqlParser"
)


//-------------------------------------
//
//
//
//-------------------------------------
func TestSelect(t *testing.T)  {
	p := sqlparser.NewSQLParser(`
		select
			t2.e ccc,
			t2.f ddd,
		from
			(select a as e, b as f from table2) t2
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func TestInsert(t *testing.T)  {
	p := sqlparser.NewSQLParser(`
		insert into table1 (a,b,c,d) values (1,2,3,4)
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func TestUpdate(t *testing.T)  {
	p := sqlparser.NewSQLParser(`
		update table1 set a = 1, b = 2, c = 3
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func TestDelete(t *testing.T)  {
	p := sqlparser.NewSQLParser(`
		delete from table1 where a = 1
	`)
	r, err := p.DoParser()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r.String())
	}
}