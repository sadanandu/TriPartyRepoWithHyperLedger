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
	fmt.Printf(args[0])
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

	return t.GetEntity(stub, function, args[0])	
	return nil, errors.New("Invalid query function name. Expecting 'query' but found '" + function + "'")
}

func (t TriPartyRepoChaincode) GetEntity(stub shim.ChaincodeStubInterface, entityName string, entityKey string) ([]byte, error){
	var columns []shim.Column
	var col1 shim.Column
	if entityName == "User"{
		fmt.Printf("Getting user")
		col1 = shim.Column{Value: &shim.Column_String_{String_: entityKey}}
	}else{
		fmt.Printf("Getting other")
		entityKeyAsInt , _ := strconv.ParseUint(entityKey, 10, 64)
		fmt.Printf("Entity as int [%d]", entityKeyAsInt)
		col1 = shim.Column{Value: &shim.Column_Uint64{Uint64: entityKeyAsInt}}
	}
	
	columns = append(columns, col1)
	fmt.Printf("Getting %s", entityName)
	//for tables having multiple keys use GetRows else use GetRow
	if t.TableHasMultipleKeys(stub, entityName){
		rows, err := stub.GetRows(entityName, columns)
		if err != nil {
			return nil, fmt.Errorf("Failed retrieving [%s] [%s]: [%s]",entityName, entityKey, err)
		}
		fmt.Printf("Rows is", rows)
		for row := range rows{
			fmt.Printf("[%s] row",entityName, row)
			entity := row.Columns[1].GetUint64()
			fmt.Printf("[%s] is",entityName, entity)
		}
	}else{
		row, err := stub.GetRow(entityName, columns)
		if err != nil {
			return nil, fmt.Errorf("Failed retrieving [%s] [%s]: [%s]",entityName, entityKey, err)
		}
		fmt.Printf("[%s] row",entityName, row)
		entity := row.Columns[1].GetUint64()
		fmt.Printf("[%s] is",entityName, entity)
	}

	return nil, nil
}


//Checking if table has multiple keys
func (t TriPartyRepoChaincode) TableHasMultipleKeys(stub shim.ChaincodeStubInterface, tableName string) (bool){
	table, _ := stub.GetTable(tableName)
	counter := 0
	for _, definition := range table.ColumnDefinitions{
		if definition.Key{
			counter = counter + 1 
		}
	}
	if counter > 1{
		return true
	}
	return false
}

func main() {
	myLogger.Debug("in main")
	err := shim.Start(new(TriPartyRepoChaincode))
	myLogger.Debug("start done")
	if err != nil {
		fmt.Printf("Error starting TriPartyRepoChaincode: %s", err)
	}
}
	
	