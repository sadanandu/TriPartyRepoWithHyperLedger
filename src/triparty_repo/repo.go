package triparty_repo

import(
	"encoding/json"
	"fmt"
	"errors"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type Repo struct {
	Id  uint64 "repoid"
	Seller  uint64 `json:"seller"`
	Buyer uint64   `json:"buyer"`
	SecurityId   uint64   `json:"securityid"`
	Quantity    uint64      `json:"quantity"`
	Haircut    uint64  `json:"haircut"`
	Amount     uint64  `json:"amount"`
	EnterDate  uint64  `json:"enterdate"`
	BuyBackDate uint64 `json:"buybackdate"`
}

//EnterRepo accepts a json object and converts it into Repo struct. 
//The input in the form of json object will be helpful when calling 
//this function from gui/web gui.
//Can I use protobuff here instead of json object?
func EnterRepo(stub shim.ChaincodeStubInterface, args []string) (string, error){
	fmt.Printf("EnterRepo")
	var err error
	repo := Repo{}
	
	json.Unmarshal([]byte(args[0]), &repo)
	fmt.Printf("Repo to be entered", repo)
	b, _ := ValidateRepoDetails("DummySeller", 10001, 10) 
	if b == true{
		fmt.Printf("ValidRepo")
		fmt.Printf(string(repo.Id))
		_, err = stub.InsertRow("Repo", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Id}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Seller}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Buyer}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.SecurityId}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Quantity}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Haircut}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.Amount}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.EnterDate}},
				&shim.Column{Value: &shim.Column_Uint64{Uint64: repo.BuyBackDate}}},
		})
		
		if err != nil {
			return "Error inRepo", errors.New("Failure in adding row to AccountSecurityLink table.")
		}
		fmt.Printf("EnteredRepo")
		return "EnteredRepo" , nil
	}
	return "" ,nil
}

func ValidateRepoDetails(seller string, securityid uint64, quantity uint64) (bool, error){
	fmt.Printf("ValidatingRepo")
	return true, nil
}


