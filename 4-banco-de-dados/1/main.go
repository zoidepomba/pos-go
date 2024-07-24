package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// aqui estou criando a struct para armazenamento dos dados
type Product struct {
	ID    string
	Name  string
	Price float64
}

// aqui estou criando a função que vai acessar o ponteiro da struct Product para criar um uuid inserir os dados nela
func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}

}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert") //aqui estou abrindo a conexão com o banco
	if err != nil {
		panic(err)
	} //aqui estou verificando se a conexão foi com sucesso, se não eu imprimo o erro
	defer db.Close() //aqui eu garanto que a conexão ira ser fechada no final da execução do programa

	product := NewProduct("Notebook", 1899.90) //aqui estou chamando a função que cria um uuid para poder escrever na struct e ja passando dois dados nela
	err = insertProducts(db, *product)         //aqui estou chamando a função insertProducts passando a string de conexão com o banco e passando um ponteiro da variavel product onde ela já esta com os valores e jogando tudo para o erro
	if err != nil {
		panic(err)
	} //aqui estou verificando se a minha transação foi tudo ok
	product.Price = 100.0 //nesse  aqui agora vou alterar o valor
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
//	p, err := selectProduct(db, product.ID)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Product: %v, possui o preço ded %.2f", p.Name, p.Price)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products{
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}
	err = deleteProduct(db, product.ID)
	if err != nil{
		panic(err)
	}

}

// Função para Inserir Produtos no Banco de Dados
func insertProducts(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)") //estou chamando o stmt para esconder os dados que vou inserir no banco
	if err != nil {
		return err
	} // se der um erro no insert aqui vou ser informado
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price) //aqui estou chamando a variavel stmt e passando aonde que vou inserir os dados, no caso seria na struct
	if err != nil {
		return err
	} //aqui se meu insert der erro ja vou ser informado
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price  from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil{
			return nil, err
		}
		product = append(product, p)
	}
	return product, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err :=  db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}