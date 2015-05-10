package db

import (
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/go-horizon/test"
	"testing"
)

func TestCoreAccountByAddressQuery(t *testing.T) {

	Convey("CoreAccountByAddress", t, func() {
		test.LoadScenario("base")
		db := OpenStellarCoreTestDatabase()

		Convey("Existing record behavior", func() {
			address := "gspbxqXqEUZkiCCEFFCN9Vu4FLucdjLLdLcsV6E82Qc1T7ehsTC"
			q := CoreAccountByAddressQuery{
				GormQuery{&db},
				address,
			}
			result, _ := First(q)
			account := result.(CoreAccountRecord)

			So(account.Accountid, ShouldEqual, address)
			So(account.Balance, ShouldEqual, 99999996999999970)
		})

		Convey("Missing record behavior", func() {
			address := "not real"
			q := CoreAccountByAddressQuery{
				GormQuery{&db},
				address,
			}
			result, err := First(q)
			So(result, ShouldBeNil)
			So(err, ShouldBeNil)
		})

	})
}
