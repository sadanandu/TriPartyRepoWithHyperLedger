package triparty_repo

import(
	"fmt"
	"errors"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func EnterRepo(stub shim.ChaincodeStubInterface) (string, error){
	fmt.Printf("EnterRepo")
	var err error
	
	b, _ := ValidateRepoDetails("DummySeller", 10001, 10) 
	if b == true{
		fmt.Printf("ValidRepo")
		_, err = stub.InsertRow("Repo", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 10002}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 10}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 2}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 1000000}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 12112016}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: 14112016}}},
		})
		
		if err != nil {
			return "", errors.New("Failure in adding row to AccountSecurityLink table.")
		}
		return "EnteredRepo" , nil
	}
	return "" ,nil
}

func ValidateRepoDetails(seller string, securityid uint64, quantity uint64) (bool, error){
	fmt.Printf("ValidatingRepo")
	return true, nil
}


