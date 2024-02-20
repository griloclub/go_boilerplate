package service

import (
	"context"
	"errors"

	"github.com/diegoclair/go_boilerplate/application/contract"
	"github.com/diegoclair/go_boilerplate/domain/transfer"
	"github.com/diegoclair/go_boilerplate/infra"
	"github.com/diegoclair/go_utils/resterrors"
	"github.com/twinj/uuid"
)

type transferService struct {
	svc *service
}

func newTransferService(svc *service) contract.TransferService {
	return &transferService{
		svc: svc,
	}
}

func (s *transferService) CreateTransfer(ctx context.Context, t transfer.Transfer) (err error) {
	s.svc.log.Info(ctx, "Process Started")
	defer s.svc.log.Info(ctx, "Process Finished")

	loggedAccountUUID, ok := ctx.Value(infra.AccountUUIDKey).(string)
	if !ok {
		errMsg := "accountUUID should not be empty"
		s.svc.log.Error(ctx, errMsg)
		return errors.New(errMsg)
	}

	fromAccount, err := s.svc.dm.Account().GetAccountByUUID(ctx, loggedAccountUUID)
	if err != nil {
		s.svc.log.Errorf(ctx, "error to get logged account by uuid: %s", err.Error())
		return err
	}

	if !fromAccount.HasSufficientFunds(t.Amount) {
		return resterrors.NewConflictError("Your account don't have sufficient funds to do this operation")
	}

	destAccount, err := s.svc.dm.Account().GetAccountByUUID(ctx, t.AccountDestinationUUID)
	if err != nil {
		s.svc.log.Errorf(ctx, "error to get destination account by uuid: %s", err.Error())
		return err
	}

	if fromAccount.ID == destAccount.ID {
		return resterrors.NewConflictError("You can't transfer to yourself")
	}

	t.TransferUUID = uuid.NewV4().String()

	return s.svc.dm.WithTransaction(ctx, func(tx contract.DataManager) error {

		err = tx.Account().AddTransfer(ctx, t.TransferUUID, fromAccount.ID, destAccount.ID, t.Amount)
		if err != nil {
			s.svc.log.Errorf(ctx, "error to add transfer: %s", err.Error())
			return err
		}

		fromAccount.SubtractBalance(t.Amount)

		err = tx.Account().UpdateAccountBalance(ctx, fromAccount.ID, fromAccount.Balance)
		if err != nil {
			s.svc.log.Errorf(ctx, "error to update origin account balance: %s", err.Error())
			return err
		}

		destAccount.AddBalance(t.Amount)

		err = tx.Account().UpdateAccountBalance(ctx, destAccount.ID, destAccount.Balance)
		if err != nil {
			s.svc.log.Errorf(ctx, "error to update destination account balance: %s", err.Error())
			return err
		}
		return nil
	})
}

func (s *transferService) GetTransfers(ctx context.Context, take, skip int64) (transfers []transfer.Transfer, totalRecords int64, err error) {
	s.svc.log.Info(ctx, "Process Started")
	defer s.svc.log.Info(ctx, "Process Finished")

	loggedAccountUUID, ok := ctx.Value(infra.AccountUUIDKey).(string)
	if !ok {
		errMsg := "accountUUID should not be empty"
		s.svc.log.Error(ctx, errMsg)
		return transfers, totalRecords, errors.New(errMsg)
	}

	account, err := s.svc.dm.Account().GetAccountByUUID(ctx, loggedAccountUUID)
	if err != nil {
		s.svc.log.Errorf(ctx, "error to get logged account by uuid: %s", err.Error())
		return transfers, totalRecords, err
	}

	madeTransfers, madeTotalRecords, err := s.svc.dm.Account().GetTransfersByAccountID(ctx, account.ID, take, skip, true)
	if err != nil {
		s.svc.log.Errorf(ctx, "error to get made transfers: %s", err.Error())
		return transfers, totalRecords, err
	}

	transfers = append(transfers, madeTransfers...)

	receivedTransfers, receivedTotalRecords, err := s.svc.dm.Account().GetTransfersByAccountID(ctx, account.ID, take, skip, false)
	if err != nil {
		s.svc.log.Errorf(ctx, "error to get received transfers: %s", err.Error())
		return transfers, totalRecords, err
	}

	transfers = append(transfers, receivedTransfers...)
	totalRecords = madeTotalRecords + receivedTotalRecords

	return transfers, totalRecords, nil
}
