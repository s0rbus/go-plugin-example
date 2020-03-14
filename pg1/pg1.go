package main

import (
   "fmt"
   "../generator"
)

type plgGenerator struct{}

func (plgGenerator) Generate(topics []string, args []interface{}) error {
   fmt.Printf("pg1 plugin Generate has topics: %v\n",topics)
   fmt.Printf("pg1 plugin Generate has args: %v\n",args)
   return nil
}

func GetGenerator() (g  generator.Generator, err error) {
   g = plgGenerator{}
   fmt.Printf("[plugin GetGenerator] Returning generator: %T %v\n", g, g)
   return
}
 
