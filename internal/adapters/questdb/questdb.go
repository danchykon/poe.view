package questdb

import qdb "github.com/questdb/go-questdb-client/v3"

type OffersRepository struct {
	client qdb.LineSender
}

func (r *OffersRepository) Add(offer Offer) error {

}
