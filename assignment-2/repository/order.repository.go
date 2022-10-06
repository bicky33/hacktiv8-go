package repository

import (
	"assignment-2/domain"
	"database/sql"
	"fmt"
	"strings"
)

type OrderRepositoryImp struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) OrderRepository {
	return &OrderRepositoryImp{DB}

}

func (repository *OrderRepositoryImp) Create(data domain.Order) error {
	var stringArgs []string
	var args []interface{}
	counter := 0
	var orderId int
	orderQuery := "INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id"
	itemQuery := "INSERT INTO items (item_code, description, quantity, order_id) VALUES"
	err := repository.DB.QueryRow(orderQuery, data.CustomerName, data.OrderedAt).Scan(&orderId)
	if err != nil {
		return err
	}

	if len(data.Items) == 0 {
		return nil
	}

	for _, v := range data.Items {
		counter++
		stringArgs = append(stringArgs, fmt.Sprintf("($%d, $%d, $%d, $%d)", counter, counter+1, counter+2, counter+3))
		counter += 3
		args = append(args, v.ItemCode, v.Description, v.Quantity, orderId)
	}

	itemParams := fmt.Sprintf("%s %s", itemQuery, strings.Join(stringArgs, ","))
	_, err = repository.DB.Exec(itemParams, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepositoryImp) Delete(orderId int) error {
	orderQuery := "DELETE FROM orders WHERE order_id = $1"
	itemQuery := "DELETE FROM items WHERE order_id = $1"
	_, err := repository.DB.Exec(itemQuery, orderId)
	if err != nil {
		return err
	}
	_, err = repository.DB.Exec(orderQuery, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (repository *OrderRepositoryImp) GetAll() ([]domain.Order, error) {
	var orders []domain.Order
	var items []domain.Item
	ordersQuery := "SELECT order_id, customer_name, ordered_at FROM orders"
	itemsQuery := "SELECT item_id, item_code, description, quantity, order_id FROM items"
	orderRows, err := repository.DB.Query(ordersQuery)
	if err != nil {
		return nil, err
	}
	itemRows, err := repository.DB.Query(itemsQuery)
	if err != nil {
		return nil, err
	}
	defer orderRows.Close()
	defer itemRows.Close()

	for orderRows.Next() {
		var order domain.Order
		err = orderRows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	for itemRows.Next() {
		var item domain.Item
		var orderId sql.NullInt64
		err := itemRows.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &orderId)
		if err != nil {
			return nil, err
		}
		if orderId.Valid {
			item.OrderId = int(orderId.Int64)
		}
		items = append(items, item)
	}
	for i, v := range orders {
		for j, z := range items {
			if v.OrderId == z.OrderId {
				orders[i].Items = append(orders[i].Items, items[j])
			}
		}
	}
	return orders, nil

}

func (repository *OrderRepositoryImp) Update(orderId int, data domain.Order) error {
	orderQuery := "UPDATE orders set customer_name = $1, ordered_at = $2 WHERE order_id = $3"
	_, err := repository.DB.Exec(orderQuery, data.CustomerName, data.OrderedAt, orderId)
	if err != nil {
		return err
	}
	if len(data.Items) == 0 {
		return nil
	}
	var stringArgs []string
	var args []interface{}
	counter := 0

	for _, v := range data.Items {
		stringArgs = append(stringArgs, fmt.Sprintf("($%d, $%d, $%d, $%d)", counter+1, counter+2, counter+3, counter+4))
		args = append(args, v.ItemCode)
		args = append(args, v.Description)
		args = append(args, v.Quantity)
		args = append(args, v.ItemId)
		counter += 4
	}
	fmt.Println(stringArgs)
	itemParams := strings.Join(stringArgs, ",")
	itemQuery := fmt.Sprintf("update items as a set item_code = b.item_code, description = b.description, quantity = b.quantity::integer from ( values %s) as b(item_code, description, quantity, item_id) where a.item_id = b.item_id::integer", itemParams)
	_, err = repository.DB.Exec(itemQuery, args...)
	if err != nil {
		return err
	}

	return nil

}
