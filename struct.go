package main

type Data struct {
	Data Pages `json:"data"`
}

type Pages struct {
	PagesConnection Edges `json:"pagesConnection"`
}

 type Edges struct {
	Edges []Node `json:"edges"`
}

type Node struct {
	Node PortfolioPage `json:"node"`
}
