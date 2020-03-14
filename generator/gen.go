package generator

type Generator interface {
   Generate(topics []string, args[]interface{}) error
}

