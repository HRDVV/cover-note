package model

type List struct {
	Id      int    `json:"id"`
	Status  int    `json:"status"`
	Content string `json:"content"`
}

func (l *List) addList() {

}
