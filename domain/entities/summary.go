// Package defines the entities of the project.
package entities

import "fmt"

// Summary defines a summary of account transactions.
type Summary struct {
	Account                string
	TotalBalance           float64
	NumMonthlyTransactions [12]uint
	AvgCredit              float64
	AvgDebit               float64
}

const htmlContent = `
<style>
  .parent {
    display: grid;
    place-items: center;
  }

  .card {
    width: clamp(23ch, 50vw, 46ch);
    display: flex;
    flex-direction: column;
    padding: 1rem;
  }

  .visual {
      width: 350;
      width: 100vw;
  }

  td {
    text-align: center;
    border: 3px solid #829ee4;
  }

  table {
    width:  350px;
    border: 3px solid #829ee4;
  }

</style>
 <div class="parent white">
  <div class="card purple">
    <h2>Your report for account #%s</h2>
    <div class="visual yellow">
      <img src="https://play-lh.googleusercontent.com/aZ_yKQAb02P6DJJm8NHVS7EgRRBUtAI4b1xRhclMLsEEqN3m8ycPg96CB7yDH2-bzOU" width="350">
    </div>
    <p>Total Balance: %f</p>
    <p>Montly Average Credit: %f</p>
    <p>Montly Average Debit: %f</p>
    <h4>Table of monthly transactions.</h4>
    <table>
      <tr>
        <th>Month</th>
        <th>Transactions</th>
      </tr>
	  %s
    </table>
  </div>
</div>
`

// Print the summary in HTML
func (s Summary) PrintHTML() string {
	mapMonth := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	transTable := ""
	for i := 0; i < 12; i++ {
		if s.NumMonthlyTransactions[i] > 0 {
			transTable += "<tr>\n<td>" + mapMonth[i] + "</td>\n<td>" + fmt.Sprint(s.NumMonthlyTransactions[i]) + "</td>\n</tr>\n"
		}
	}

	return fmt.Sprintf(htmlContent, s.Account, s.TotalBalance, s.AvgCredit, s.AvgDebit, transTable)
}
