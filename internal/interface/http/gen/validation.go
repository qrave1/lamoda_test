package gen

func (r PostReleaseRequestObject) Valid() bool {
	if r.Body == nil {
		return false
	}

	if r.Body.ProductCodes == nil || len(r.Body.ProductCodes) == 0 || r.Body.WarehouseId == 0 {
		return false
	}

	if r.Body.WarehouseId < 0 {
		return false
	}

	for _, val := range r.Body.ProductCodes {
		if val < 0 {
			return false
		}
	}

	return true
}

func (r PostReserveRequestObject) Valid() bool {
	if r.Body == nil {
		return false
	}

	if r.Body.ProductCodes == nil || len(r.Body.ProductCodes) == 0 || r.Body.WarehouseId == 0 || r.Body.Quantity <= 0 {
		return false
	}

	if r.Body.WarehouseId < 0 {
		return false
	}

	for _, val := range r.Body.ProductCodes {
		if val < 0 {
			return false
		}
	}

	return true
}
