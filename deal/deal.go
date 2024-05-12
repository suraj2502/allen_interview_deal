package deal

import (
	"fmt"
	"math/rand"
	"time"
)

// Limited Time Deals

// A limited-time deal implies that a seller will put up an item on sale for a limited time period, say, 2 hours, and will keep a
//  maximum limit on the number of items that would be sold as part of that deal.

// Users cannot buy the deal if the deal time is over
// Users cannot buy if the maximum allowed deal has already been bought by other users.
// One user cannot buy more than one item as part of the deal.

// The task is to create APIs to enable the following operations

// Create a deal with the price and number of items to be sold as part of the deal
// End a deal
// Update a deal to increase the number of items or end time
// Claim a deal

var UsersList []*User
var ProductsList []*Product
var DealsList []*Deal

type Product struct {
	Id                 int    `json:"id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	InventoryCount     int    `json:"inventory_count"`
	Deal               *Deal  `json:"deal"`
}

func (P *Product) NewProduct(productName string, productDescription string) *Product {
	// create a random id for the product
	Id := rand.Intn(100)

	product := &Product{
		Id:                 Id,
		ProductName:        productName,
		ProductDescription: productDescription,
		// Deal:               deal,
	}
	ProductsList = append(ProductsList, product)
	return product
}

func (P *Product) AssignDealToProduct(deal *Deal) {
	P.Deal = deal

}

func (P *Product) GetProduct(id int) *Product {

	for _, p := range ProductsList {
		if p.Id == id {
			return p
		}
	}
	return nil
}

type Deal struct {
	Id              int       `json:"id"`
	DealStartTime   string    `json:"deal_start_time"`
	DealEndTime     string    `json:"deal_end_time"`
	DealName        string    `json:"deal_name"`
	DealDescription string    `json:"deal_description"`
	IsaDealActive   bool      `json:"is_deal_active"`
	DealPrice       float64   `json:"deal_price"`
	MaxDealQuantity int       `json:"max_deal_quantity"`
	DealQuantity    int       `json:"deal_quantity"`
	Products        []Product `json:"products"`
}

// NewDeal creates a new deal
func (d *Deal) NewDeal(dealStartTime string, dealEndTime string, dealName string, dealDescription string, dealPrice float64, maxDealQuantity int, dealQuantity int, products []Product) *Deal {
	deal := &Deal{
		Id:              rand.Intn(100),
		DealStartTime:   dealStartTime,
		DealEndTime:     dealEndTime,
		DealName:        dealName,
		DealDescription: dealDescription,
		DealPrice:       dealPrice,
		MaxDealQuantity: maxDealQuantity,
		DealQuantity:    dealQuantity,
		Products:        products,
		IsaDealActive:   true,
	}
	for _, product := range products {
		product.AssignDealToProduct(deal)
	}
	DealsList = append(DealsList, deal)
	return deal
}

func (d *Deal) GetDeal(id int) *Deal {

	for _, d := range DealsList {
		if d.Id == id {
			return d
		}
	}
	return nil
}

func (d *Deal) StartDeal() {
	d.IsaDealActive = true
}

// UpdateDeal updates the deal
func (d *Deal) UpdateDeal(dealEndTime string, maxDealQuantity int) {
	d.DealEndTime = dealEndTime
	d.MaxDealQuantity = maxDealQuantity
}

func (d *Deal) EndDeal(dealEnddateTime string) {
	d.DealEndTime = dealEnddateTime
	d.IsaDealActive = false
}

type User struct {
	Id                    int       `json:"id"`
	UserName              string    `json:"user_name"`
	Email                 string    `json:"email"`
	IsProductBoughtByUser bool      `json:"is_product_bought_by_user"`
	ProductsBought        []Product `json:"products_bought"`
}

func (U *User) NewUser(userName string, email string) *User {
	// create a random id for the product
	Id := rand.Intn(100)

	u := &User{
		Id:       Id,
		UserName: userName,
		Email:    email,
	}
	UsersList = append(UsersList, u)
	return u

}

func (u *User) GetUser(id int) *User {

	for _, user := range UsersList {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (u *User) ClaimDeal(P *Product) (bool, error) {
	if P.Deal == nil {
		fmt.Println("Product does not have a deal")
		return false, fmt.Errorf("Product does not have a deal")
	}
	if P.Deal.DealEndTime < time.Now().String() {
		fmt.Println("Deal has already ended")
		return false, fmt.Errorf("Deal has already ended")
	}

	if P.Deal.DealStartTime > time.Now().String() {
		fmt.Println("Deal has not started yet")
		return false, fmt.Errorf("Deal has not started yet")
	}
	if P.Deal.IsaDealActive {
		for _, product := range u.ProductsBought {

			if product.Deal != nil && P.Deal != nil && product.Deal.Id == P.Deal.Id {
				fmt.Println("User has already claimed the deal : ", P.Deal.DealName)
				return false, fmt.Errorf("User has already claimed the deal : %s", P.Deal.DealName)

			}
		}

		if P.Deal != nil && P.Deal.DealQuantity < P.Deal.MaxDealQuantity {

			P.Deal.DealQuantity++
			u.ProductsBought = append(u.ProductsBought, *P)
		} else {
			fmt.Println("Deal is already claimed by all users")
			return false, fmt.Errorf("Deal is already claimed by all users")

		}

	} else {
		fmt.Println("Deal is not active or user has already claimed the deal")
		return false, fmt.Errorf("Deal is not active or user has already claimed the deal")
	}
	return true, nil
}

func (u *User) BuyProduct(P *Product) {
	if P.InventoryCount > 0 {
		P.InventoryCount--
	} else {
		fmt.Println("Product is out of stock")
	}
}
