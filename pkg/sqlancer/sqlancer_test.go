package sqlancer

import (
	"fmt"
	"testing"

	"github.com/chaos-mesh/go-sqlancer/pkg/types"
	"github.com/chaos-mesh/go-sqlancer/pkg/util"
	"github.com/stretchr/testify/assert"
)

func helperMakeTable() *types.Table {
	t := new(types.Table)
	t.Name = types.CIStr(fmt.Sprintf("t%d", util.Rd(100)))
	t.Columns = make([]types.Column, 0)
	for i := util.Rd(5) + 2; i > 0; i-- {
		tp := "int"
		switch util.Rd(4) {
		case 0:
			tp = "varchar"
		case 1:
			tp = "text"
		}
		t.Columns = append(t.Columns, types.Column{
			Name:   types.CIStr(fmt.Sprintf("c%d", util.Rd(100))),
			Type:   tp,
			Length: int(util.RdRange(9, 20)),
		})
	}
	return t
}

func TestDeleteStmt(t *testing.T) {
	p, _ := NewSQLancer(NewConfig())
	p.Tables = []types.Table{*helperMakeTable(), *helperMakeTable()}
	s, err := p.DeleteStmt(p.Tables, p.Tables[0])
	fmt.Println(s)
	assert.NoError(t, err)
}

func TestUpdateStmt(t *testing.T) {
	p, _ := NewSQLancer(NewConfig())
	p.Tables = []types.Table{*helperMakeTable(), *helperMakeTable()}
	s, err := p.UpdateStmt(p.Tables, p.Tables[0])
	fmt.Println(s)
	assert.NoError(t, err)
}

func TestNoRecNormal(t *testing.T) {
	p, _ := NewSQLancer(NewConfig())
	p.Tables = []types.Table{*helperMakeTable(), *helperMakeTable()}
	s, sql, _, err := p.GenSelectStmt()
	fmt.Println(s, sql)
	assert.NoError(t, err)
}

func TestNoRecNoOpt(t *testing.T) {
	p, _ := NewSQLancer(NewConfig())
	p.Tables = []types.Table{*helperMakeTable(), *helperMakeTable()}
	s, sql1, _, err := p.GenSelectStmt()
	fmt.Println(s, sql1)
	assert.NoError(t, err)
	s2, sql2, _, err := p.GenSelectStmt()
	fmt.Println(s2, sql2)
	assert.NoError(t, err)
}
