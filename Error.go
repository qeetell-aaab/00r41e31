package r0r41e31

// Object type Error represents a Golang error. However, this error is also capable of containing
// a secondary error.
//
// What do you mean by secondary error?
//
// Well, if you plan creating an error [primary error] because you encountered some other error,
// the error that prompted the error you're about creating, is what I call a secondary error.
type Error struct {
	errorText string
	secondary error
}

// Example:
//
// 	e := Error_New ("This is an error I created.", theErrorThatPromptedTheCreationOfMyError)
//
func Error_New (errorText string, secondary ... error) (*Error) {
	var sec error

	if len (secondary) > 0 {
		sec = secondary [0]
	} else {
		sec = nil
	}
	return &Error {errorText, sec}
}

func (e *Error) Error () (string) {
	return e.errorText
}

func (e *Error) Unwrap () (error) {
	return e.secondary
}
