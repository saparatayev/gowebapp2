package models

import "fmt"

type Product struct {
	Id       uint64
	Name     string
	Price    float64
	Amount   float64
	Quantity int
	Category Category
}

func NewProduct(product Product) (bool, error) {
	con := Connect()
	defer con.Close()

	sql := "INSERT INTO products(name, price, quantity, amount, category) VALUES ($1, $2, $3, $4, $5)"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Quantity, product.Amount, product.Category.Id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetProducts() ([]Product, error) {
	con := Connect()
	defer con.Close()

	sql := `select c.id, c.description,
		p.id, p.name, p.price, p.quantity, p.amount
		from products as p inner join category as c on c.id=p.category order by p.id asc`

	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var products []Product

	for rs.Next() {
		var product Product

		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func SearchProducts(search string) ([]Product, error) {
	search = fmt.Sprintf("%%%s%%", search)

	con := Connect()
	defer con.Close()

	sql := `select c.id, c.description,
		p.id, p.name, p.price, p.quantity, p.amount
		from products as p inner join category as c on c.id=p.category
		where c.description like $1 or p.name like $2`

	stmt, err := con.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rs, err := stmt.Query(search, search)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var products []Product

	for rs.Next() {
		var product Product

		err := rs.Scan(&product.Category.Id,
			&product.Category.Description,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductById(id uint64) (Product, error) {
	con := Connect()
	defer con.Close()

	sql := "select * from products where id = $1"

	rs, err := con.Query(sql, id)
	if err != nil {
		return Product{}, err
	}
	defer rs.Close()

	var product Product

	if rs.Next() {
		err := rs.Scan(&product.Id,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Amount,
			&product.Category.Id,
		)

		if err != nil {
			return Product{}, err
		}
	}

	return product, nil
}

func UpdateProduct(product Product) (rows, error) {
	con := Connect()
	defer con.Close()

	sql := "update products set name = $1, price = $2, quantity = $3, amount = $4, category = $5 where id = $6"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rs, err := stmt.Query(product.Name, product.Price, product.Quantity, product.Amount, product.Category.Id, product.Id)
	if err != nil {
		return 0, err
	}

	rows, err := rs.rowsAffe
	defer rs.Close()
}

func (p *Product) PriceToString() string {
	return fmt.Sprintf("%.2f", p.Price)
}

func (p *Product) AmountToString() string {
	return fmt.Sprintf("%.2f", p.Amount)
}
