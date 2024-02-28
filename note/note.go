package note

import (
	"errors"
)


type Note struct {

	title string
	content string


	
}

func New(title, content string) (Note, error)  {
	
	
	if(title == "" || content == ""){
		return Note{}, errors.New("empty title or content")
	}

	return Note{
        title: title,
        content: content,
    }, nil
}