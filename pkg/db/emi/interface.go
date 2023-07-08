package emi

type loanObj struct{

}

type LoanGroup interface{

} 
func LoanDb() LoanGroup{
	return &loanObj{

	}
}