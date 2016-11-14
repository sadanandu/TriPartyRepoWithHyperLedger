package triparty_repo

import (
	"fmt"
	"errors"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("triparty_repo")
func SetUpTables(stub shim.ChaincodeStubInterface) ([]byte, error){
	var err error

	// Create user table
    myLogger.Debug("in create_tables")
	tab, _ := stub.GetTable("User")
	fmt.Printf("%+v", tab)
	if tab == nil{
		err = stub.CreateTable("User", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_STRING, Key: true},
			&shim.ColumnDefinition{Name: "Userid", Type: shim.ColumnDefinition_UINT64, Key: false},
		})
		if err != nil {
			fmt.Printf("%+v", err)
			return nil, errors.New("Failure in creating User table.")
		}
    }


	tab, _ = stub.GetTable("UserAccount")

	if tab == nil{
	
		// Create userAccount table
		err = stub.CreateTable("UserAccount", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "AccNumber", Type: shim.ColumnDefinition_UINT64, Key: true},
			&shim.ColumnDefinition{Name: "Userid", Type: shim.ColumnDefinition_UINT64, Key: true},
		})
		if err != nil {
			return nil, errors.New("Failure in creating UserAccount table.")
		}
	}

	tab, _ = stub.GetTable("Security")
	
	if tab == nil{

		// Create Security table
		err = stub.CreateTable("Security", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: false},
			&shim.ColumnDefinition{Name: "SecurityId", Type: shim.ColumnDefinition_UINT64, Key: true},
			&shim.ColumnDefinition{Name: "SecurityAttributes", Type: shim.ColumnDefinition_STRING, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failure in creating Security table.")
		}
	}

	
	tab, _ = stub.GetTable("AccountSecurityLink")
	
	if tab == nil{

		// Create AccountSecurityLink table
		err = stub.CreateTable("AccountSecurityLink", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "AccNumber", Type: shim.ColumnDefinition_UINT64, Key: true},
			&shim.ColumnDefinition{Name: "SecurityId", Type: shim.ColumnDefinition_UINT64, Key: true},
			&shim.ColumnDefinition{Name: "Quantity", Type: shim.ColumnDefinition_UINT64, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failure in creating AccountSecurityLink table.")
		}
	}

	tab, _ = stub.GetTable("Repo")
	
	if tab == nil{

		// Create Repo table
		err = stub.CreateTable("Repo", []*shim.ColumnDefinition{
			&shim.ColumnDefinition{Name: "RepoId", Type: shim.ColumnDefinition_UINT64, Key: true},
			&shim.ColumnDefinition{Name: "Seller", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "Buyer", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "SecurityId", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "Quantity", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "Haircut", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "Amount", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "EnterDate", Type: shim.ColumnDefinition_UINT64, Key: false},
			&shim.ColumnDefinition{Name: "BuyBackDate", Type: shim.ColumnDefinition_UINT64, Key: false},
		})
		if err != nil {
			return nil, errors.New("Failure in creating Repo table.")
		}

	}

	_, err = stub.InsertRow("User", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "DummySeller"}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 100001}}},
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to User table.")
	}

	_, err = stub.InsertRow("UserAccount", shim.Row{
		Columns:[]*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "Security"}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 100001}}},
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to UserAccount table.")
	}
	
	
	_, err = stub.InsertRow("User", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "DummyBuyer"}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 100002}}},
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to User table.")
	}

	_, err = stub.InsertRow("UserAccount", shim.Row{
		Columns:[]*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "Security"}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10002}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64:100002}}},
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to UserAccount table.")
	}	
	
	_, err = stub.InsertRow("Security", shim.Row{
		Columns:[]*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "Bond"}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
			&shim.Column{Value: &shim.Column_String_{String_: "SomeAttributes"}}},
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to Security table.")
	}	

	_, err = stub.InsertRow("AccountSecurityLink", shim.Row{
		Columns:[]*shim.Column{
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10001}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: 10}}},
	
	})
	if err != nil {
		return nil, errors.New("Failure in adding row to AccountSecurityLink table.")
	}
	return nil, nil
}

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