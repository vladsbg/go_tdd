package _07_maps

const (
	ErrorNotFound      = ErrorDictionary("Não foi possível encontrar a palavra no dicionário")
	ErrorAlreadyExists = ErrorDictionary("Não foi possível adicionar a palavra, pois ela já existe no dicionário")
	ErrorNotExists     = ErrorDictionary("Não foi possível atualizar a palavra, pois ela não existe no dicionário")
)

type ErrorDictionary string

func (e ErrorDictionary) Error() string {
	return string(e)
}
