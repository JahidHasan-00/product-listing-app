package database

var productList []Product

type Product struct {
	ID          int     ` json:"id" `
	Title       string  ` json:"title" `
	Description string  ` json:"description" `
	Price       float64 ` json:"price" `
	ImgUrl      string  ` json:"imageUrl" `
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is full of Vitamin C, Orange color is like mix of yellow and red",
		Price:       100.0,
		ImgUrl:      "https://orgfarm.store/cdn/shop/files/malda_orange.jpg?v=1721886788&width=1214",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green color",
		Price:       60,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTiI-HlOSLzPb53nc6-NueK8wv5dyPezl4diw&s",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is a good for health. Banana is yello color",
		Price:       5,
		ImgUrl:      "https://www.lovefoodhatewaste.com/sites/default/files/styles/16_9_two_column/public/2022-07/Bananas.jpg.webp?itok=bqo03oZ1",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Grape",
		Description: "I love grape",
		Price:       300,
		ImgUrl:      "https://erewhonappsftp.s3.us-west-2.amazonaws.com/images/233571000-1.png?date=2025-04-1600:56:58.372",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "I love mango, This is our seasonal fruit",
		Price:       500,
		ImgUrl:      "https://images.jdmagicbox.com/rep/b2b/alphonso-mangoes/alphonso-mangoes-15.jpg",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Lichi",
		Description: "Lichi is a seasonal fruits in BD, I love to eat Lichi",
		Price:       200,
		ImgUrl:      "https://m.media-amazon.com/images/I/81F6bQaa2NL._UF1000,1000_QL80_.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)

}
