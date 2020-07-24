package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (result string, err error) {
	result, exist := d[word]
	if !exist {
		err = ErrorNotFound
	}
	return
}

func (d Dictionary) Add(word string, definition string) (err error) {
	_, err = d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorAlreadyExists
	default:
		return
	}
	return
}

func (d Dictionary) Update(word string, newDefinition string) (err error) {
	_, err = d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
		return
	case ErrorNotFound:
		return ErrorNotExists
	default:
		return
	}
}

func (d Dictionary) Delete(word string) (err error) {
	_, err = d.Search(word)
	switch err {
	case nil:
		delete(d, word)
		return
	case ErrorNotFound:
		return ErrorNotFound
	default:
		return
	}
}
