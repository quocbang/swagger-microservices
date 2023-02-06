package pgimpl

import (
	"context"
)

type DataManager interface {
	// Notice that nested transactions are NOT available.
	// You will get Code_ALREADY_IN_TRANSACTION error if you are trying to create a transaction on a DataHandler which is already in another transaction.
	// You can continue doing anything with the error `Code_ALREADY_IN_TRANSACTION` but not others since the error is not fatal.
	//
	// usage example:
	//
	//  func Example(ctx context.Context, dm DataManagerV1) error {
	//  	tx,err := dm.Begin(ctx);
	//  	if err != nil {
	//  		if errors.Is(mcomErr.Error{Code: Code_ALREADY_IN_TRANSACTION},err){
	//  			print("duplicated transaction")
	//  		}else{
	//  			return err
	//  		}
	//  	}
	//  	defer tx.RollBack()
	//
	//  	return tx.Commit()
	//  }
	// Begin(context.Context, ...*sql.TxOptions) (DataManager, error)
	// Commit() error
	// RollBack() error
	// Notice that nested transactions are NOT available.
	// You will get Code_ALREADY_IN_TRANSACTION error if you are trying to create a transaction on a DataHandler which is already in another transaction.
	// Transaction(context.Context, func(DataManager) error, ...*sql.TxOptions) error
	// Close() error

	// dataManagerServices

	GetLimitaryHour(context.Context, GetLimitaryHourRequest) (GetLimitaryHourReply, error)
}
