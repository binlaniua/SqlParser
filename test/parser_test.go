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
			t1.a,
			t1.b,
			t3.e ccc,
			t3.f ddd
		from
			xx.table1 t1,
			(select t2.b as e, t2.d as f from (select a as b, c as d from yy.table2) t2) t3
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