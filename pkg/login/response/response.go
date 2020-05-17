package response

type Parser interface {
	Parse(html []byte)(string, error)
}

// <input type="hidden" name="authenticity_token" value="Y3lNZWD/VIzVlwokHL6HW6mJMwsnL4dwQVoCPbU36xZpC/C1uxWnt+xB9yP7bRAEhpOB3tnerSNTIYeY+OQ8ZQ=="

type parserImpl struct {}

func (impl *parserImpl)Parse(html []byte)(string, error){
	
}