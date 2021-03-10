package models

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	wallet     *types.Wallet
}

func (s *Suite) SetupSuite() {
	var db *sql.DB
	var err error

	db, s.mock, err = sqlmock.New() // mock sql.DB
	require.NoError(s.T(), err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestRepository_GetWalletByID() {
	id := 1
	balance := "100.00"
	rows := sqlmock.NewRows([]string{"id", "balance"}).AddRow(id, balance)

	s.mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = ?").
		WithArgs(id).
		WillReturnRows(rows)

	res, err := s.repository.GetWalletByID(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&types.Wallet{
		ID:      uint(id),
		Balance: decimal.RequireFromString(balance),
	}, res))
}

func (s *Suite) TestRepository_CreditWallet() {
	id := 2
	balance := "50.00"
	credit := decimal.RequireFromString("0.15")

	updBalance := "50.15"
	rows := sqlmock.NewRows([]string{"id", "balance"}).AddRow(id, balance)
	exp := sqlmock.NewResult(0, 1) // (lastInsertID, rowsAffected)

	s.mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = ?").
		WithArgs(id).
		WillReturnRows(rows)
	s.mock.ExpectExec("UPDATE `wallets` SET `balance`=?").
		WithArgs(updBalance, id, id).
		WillReturnResult(exp)

	res, err := s.repository.CreditWallet(id, credit)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&types.Wallet{
		ID:      uint(id),
		Balance: decimal.RequireFromString(updBalance),
	}, res))
}

func (s *Suite) TestRepository_DebitWallet() {
	id := 3
	balance := "500.00"
	debit := decimal.RequireFromString("1.16")

	updBalance := "498.84"
	rows := sqlmock.NewRows([]string{"id", "balance"}).AddRow(id, balance)
	exp := sqlmock.NewResult(0, 1) // (lastInsertID, rowsAffected)

	s.mock.ExpectQuery("SELECT \\* FROM `wallets` WHERE `wallets`.`id` = ?").
		WithArgs(id).
		WillReturnRows(rows)
	s.mock.ExpectExec("UPDATE `wallets` SET `balance`=?").
		WithArgs(updBalance, id).
		WillReturnResult(exp)

	res, err := s.repository.DebitWallet(id, debit)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&types.Wallet{
		ID:      uint(id),
		Balance: decimal.RequireFromString(updBalance),
	}, res))
}
