package db

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"go-template/internal/test/mock"
	"gorm.io/gorm"
	"os"
	"reflect"
	"testing"
)

// 依赖数据
var dataMock *mock.DataMock
var dataMockClose func()

// baseDB 每次运行测试会初始化
var baseDB *DB

type baseRepoTestModel struct {
	ID        uint64 `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`
	Name      string `gorm:"column:name;type:varchar(64);not null;comment:暱称" json:"name"`                                    // 暱称
	Email     string `gorm:"column:email;type:varchar(320);uniqueIndex:sys_user_unique_1,priority:1;comment:邮箱" json:"email"` // 邮箱
	CreatedAt int64  // 创建时间（由GORM自动管理）
	UpdatedAt int64  // 最后一次更新时间（由GORM自动管理）
}

// 每个测试前执行
func setup() {
	var err error
	dataMock, dataMockClose, err = mock.NewDataMock()
	if err != nil {
		panic(err)
	}
	baseDB = &DB{db: dataMock.DB}
}

// 每个测试后执行
func teardown() {
	dataMockClose()
}

// 用 TestMain 整合 setup 和 teardown
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestDB_Begin(t *testing.T) {
	type args struct {
		opts []*sql.TxOptions
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Begin(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Commit(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Commit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Count(t *testing.T) {
	tests := []struct {
		name      string
		wantCount int64
		wantErr   bool
		want      *DB
	}{
		{
			name:      "all",
			wantCount: 1,
			wantErr:   false,
		}, {
			name:      "where",
			wantCount: 1,
			wantErr:   false,
		},
	}
	dataMock.SqlMock.ExpectQuery("SELECT count\\(\\*\\) FROM `base_repo_test_models`").WillReturnRows(
		sqlmock.NewRows([]string{"count(*)"}).AddRow("1"),
	)
	dataMock.SqlMock.ExpectQuery("SELECT count\\(\\*\\) FROM `base_repo_test_models` WHERE name=\\?").WillReturnRows(
		sqlmock.NewRows([]string{"count(*)"}).AddRow("1"),
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var count int64
			d := baseDB
			got := d.Count(&count)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if count != tt.wantCount {
				t.Errorf("Count() gotCount = %v, want %v", count, tt.wantCount)
			}
		})
	}
}

func TestDB_Create(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "all",
			args: args{
				value: &baseRepoTestModel{
					Email: "a@a.c",
					Name:  "te_name",
				},
			},
			wantErr: false,
		},
	}
	dataMock.SqlMock.ExpectBegin()
	dataMock.SqlMock.ExpectExec(
		"INSERT INTO `base_repo_test_models` \\(`name`,`email`,`created_at`\\) VALUES \\(\\?,\\?,\\?\\)").
		WithArgs("te_name", "a@a.c", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			got := d.Create(tt.args.value)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Delete(t *testing.T) {
	type args struct {
		value any
		conds []any
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "del",
			args: args{
				value: "id=?",
				conds: []any{1},
			},
			wantErr: false,
		},
	}
	dataMock.SqlMock.ExpectBegin()
	dataMock.SqlMock.ExpectExec(
		"DELETE FROM `base_repo_test_models` WHERE id=\\?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			got := d.Delete(tt.args.value, tt.args.conds...)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Error(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if err := d.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Find(t *testing.T) {
	type args struct {
		dest  interface{}
		conds []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "find",
			args: args{
				"name=?",
				[]any{"a"},
			},
			wantErr: false,
		},
	}
	dataMock.SqlMock.ExpectQuery(
		"SELECT \\* FROM `base_repo_test_models` WHERE name=\\?").
		WithArgs("a").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
			AddRow(1, "a", "a@a.c", 1720351176, 1720351176))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			got := d.Find(tt.args.dest, tt.args.conds...)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_First(t *testing.T) {
	type args struct {
		dest  interface{}
		conds []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "first",
			args: args{
				"name=?",
				[]any{"a"},
			},
			wantErr: false,
		},
	}
	dataMock.SqlMock.ExpectQuery(
		"SELECT \\* FROM `base_repo_test_models` WHERE name=\\? ORDER BY `base_repo_test_models`.`id` LIMIT \\?").
		WithArgs("a", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
			AddRow(1, "a", "a@a.c", 1720351176, 1720351176))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			got := d.First(tt.args.dest, tt.args.conds...)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Joins(t *testing.T) {
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Joins(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Joins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Last(t *testing.T) {
	type args struct {
		dest  interface{}
		conds []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Last(tt.args.dest, tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Limit(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Model(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Model(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Not(t *testing.T) {
	type args struct {
		query interface{}
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Not(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Not() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Offset(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Offset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Omit(t *testing.T) {
	type args struct {
		columns []string
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Omit(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Omit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Or(t *testing.T) {
	type args struct {
		query interface{}
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Or(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Order(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Order(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Rollback(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Rollback(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rollback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_RollbackTo(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.RollbackTo(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RollbackTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_RowsAffected(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.RowsAffected(); got != tt.want {
				t.Errorf("RowsAffected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Save(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
	}{
		{
			name: "del",
			args: args{
				value: &baseRepoTestModel{
					Name:  "a",
					Email: "a@a.c",
				},
			},
			wantErr: false,
		},
	}
	dataMock.SqlMock.ExpectBegin()
	dataMock.SqlMock.ExpectExec(
		"INSERT INTO `base_repo_test_models` \\(`name`,`email`,`updated_at`\\) VALUES \\(\\?,\\?,\\?\\)").
		WithArgs("a", "a@a.c", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	dataMock.SqlMock.ExpectCommit()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			got := d.Save(tt.args.value)
			err := got.Error()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_SavePoint(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.SavePoint(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SavePoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Scan(t *testing.T) {
	type args struct {
		dest any
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Scan(tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Select(t *testing.T) {
	type args struct {
		query interface{}
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Select(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Take(t *testing.T) {
	type args struct {
		dest  interface{}
		conds []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Take(tt.args.dest, tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Transaction(t *testing.T) {
	type args struct {
		fc   func(tx *DB) error
		opts []*sql.TxOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if err := d.Transaction(tt.args.fc, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Transaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Update(t *testing.T) {
	type args struct {
		column string
		value  interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Update(tt.args.column, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Updates(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Updates(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Updates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Where(t *testing.T) {
	type args struct {
		query interface{}
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := baseDB
			if got := d.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDB(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want IDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
