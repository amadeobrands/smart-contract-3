package response

import (
	"context"

	"github.com/tokenized/smart-contract/internal/app/inspector"
	"github.com/tokenized/smart-contract/internal/app/state/contract"
)

type reconciliationHandler struct{}

func newReconciliationHandler() reconciliationHandler {
	return reconciliationHandler{}
}

func (h reconciliationHandler) process(ctx context.Context,
	itx *inspector.Transaction, c *contract.Contract) error {

	return nil
}
