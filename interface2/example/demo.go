package example

import "fmt"

func Execute() {
	repo := NewStaticRepository()
	h := NewHandler(repo)
	s := h.Do(0)
	fmt.Println(s)

}

type Language struct {
	ID   uint
	Name string
}

type Repository interface {
	QueryLang(uint) Language
}

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Do(id uint) string {
	lang := h.repo.QueryLang(id)
	return fmt.Sprintf("%s language", lang.Name)
}

var data = map[uint]Language{
	1: {ID: 1, Name: "Go"},
	2: {ID: 2, Name: "Java"},
	3: {ID: 3, Name: "Python"},
	4: {ID: 4, Name: "Rust"},
}

type StaticRepo struct {
	data map[uint]Language
}

func NewStaticRepository() Repository {
	var staticRepo = StaticRepo{}
	staticRepo.data = data
	return &staticRepo
}

func (s *StaticRepo) QueryLang(id uint) Language {
	return s.data[id]
}
