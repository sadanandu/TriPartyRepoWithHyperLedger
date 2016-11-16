package main

import (
	"errors"
	"fmt"
	"strconv"
    "github.com/hyperledger/fabric/examples/chaincode/go/triparty_repo_chaincode/triparty_repo"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("triparty_repo")
// TriPartyRepoChaincode is simple chaincode implementing a basic TriParty Repo Management system
// 
// Look here for more information on how to implement access control at chaincode level:
// https://github.com/hyperledger/fabric/blob/master/docs/tech/application-ACL.md
// A security's attributes are represented as string for start. Repo will be a table.
type TriPartyRepoChaincode struct {
}

// Init method will be called during deployment.
// The deploy transaction metadata is supposed to contain the administrator cert
func (t *TriPartyRepoChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	myLogger.Debug("Init Chaincode")
	triparty_repo.SetUpTables(stub)
	myLogger.Debug("Init Chaincode...done")

	return nil, nil
}

// Invoke will be called for every transaction.
func (t *TriPartyRepoChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke, Entering a repo")
	msg, _ := triparty_repo.EnterRepo(stub, args)
	if msg != "" {
		return []byte(msg), nil
	}
	return nil, errors.New("No function handling")
}

// Query callback representing the query of a chaincode
func (t *TriPartyRepoChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debugf("Query [%s]", function)
	myLogger.Debugf("Args [%s]", args[0])


	if function == "GetUser"{
		return t.GetUser(stub, args[0])
	}

    if function == "GetUserAccount"{
	    accnumber , _ := strconv.ParseUint(args[0], 10, 64)
		return t.GetUserAccount(stub, accnumber)
	}

    if function == "GetRepo"{
	    repoid , _ := strconv.ParseUint(args[0], 10, 64)
		return t.GetRepo(stub, repoid)
	}
	
	return nil, errors.New("Invalid query function name. Expecting 'query' but found '" + function + "'")
}

func (t TriPartyRepoChaincode) GetUser(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {

	// Verify the identity of the caller
	// Only the owner can transfer one of his assets
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: name}}
	columns = append(columns, col1)
	fmt.Printf("Getting user")
/*    fmt.Printf("%+v\n", columns)
	tab, _ := stub.GetTable("User")
	fmt.Printf("%+v\n", tab)
	rows, _ := stub.GetRows("User", columns)
	fmt.Printf("%+v\n", rows)
*/	
	row, err := stub.GetRow("User", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving user [%s]: [%s]", name, err)
	}
	fmt.Printf("User row", row)
    user := row.Columns[1].GetUint64()
	fmt.Printf("User is", user)
	return nil, nil
}  

func (t TriPartyRepoChaincode) GetUserAccount(stub shim.ChaincodeStubInterface, accnumber uint64) ([]byte, error){
    var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_Uint64{Uint64: accnumber}}
	columns = append(columns, col1)
	fmt.Printf("Getting user account")
	rows, err := stub.GetRows("UserAccount", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving user account [%d]: [%s]", accnumber, err)
	}
	for i:= range(rows){
		fmt.Println("UserAccount is ")
		fmt.Println(i)
		fmt.Println(i.Columns[1].GetUint64())
	}
    //fmt.Printf("UserAccount is", row)
    //user := row[0].Columns[1].GetUint64()
	//fmt.Printf("UserAccount is", user)
	return nil, nil

}

func (t TriPartyRepoChaincode) GetRepo(stub shim.ChaincodeStubInterface, repoid uint64) ([]byte, error){
	var columns []shim.Column
	col1 := shim.Column{Value : &shim.Column_Uint64{Uint64: repoid}}
	columns = append(columns, col1)
	row, err := stub.GetRow("Repo", columns)
	if err != nil{
		return nil, fmt.Errorf("Failed retrieving repo [%d]: [%s]", repoid, err)
	}
	fmt.Println("Repo", row)
    repo := row.Columns[1].GetUint64()
	fmt.Printf("Repo is", repo)
	return nil, nil

}

func main() {
	myLogger.Debug("in main")
	err := shim.Start(new(TriPartyRepoChaincode))
	myLogger.Debug("start done")
	if err != nil {
		fmt.Printf("Error starting TriPartyRepoChaincode: %s", err)
	}
}
	
	