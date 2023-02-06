package utils

// func PareError(ctx context.Context, d defaultError, err error) middleware.Responder {
// 	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
// 		d.SetStatusCode(http.StatusRequestTimeout)
// 		d.SetPayload(&models.Error{
// 			Details: err.Error(),
// 		})
// 		return d
// 	}
// 	if e, ok := mcomErrors.As(err); ok {
// 		d.SetStatusCode(http.StatusBadRequest)
// 		d.SetPayload(&models.Error{
// 			Code:    int64(e.Code),
// 			Details: e.Details,
// 		})
// 	} else {
// 		d.SetStatusCode(http.StatusInternalServerError)
// 		d.SetPayload(&models.Error{
// 			Details: err.Error(),
// 		})
// 	}

// 	return d
// }

// type defaultError interface {
// 	middleware.Responder

// 	setStatusCode(int)
// 	SetPayload(*models.Error)
// }
