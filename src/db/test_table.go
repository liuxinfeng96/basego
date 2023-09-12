package db

const TableNameTestTable = "test_table"

type TestTable struct {
	GeneralField
	TestColumn1 string
	TestColumn2 string
}

func (b TestTable) TableName() string {
	return TableNameTestTable
}

func init() {
	t := new(TestTable)
	TableSlice = append(TableSlice, t)
}
