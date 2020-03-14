package main

import (
   "fmt"
   "./generator"
   "plugin"
)

func main() {
   p, err := plugin.Open("pg1/pg1.so")
   if err != nil {
      panic(err)
   }

   GetGenerator, err := p.Lookup("GetGenerator")
   if err != nil {
      panic(err)
   }
   gen, err := GetGenerator.(func() (generator.Generator, error))()
   fmt.Printf("GetFilter result: %T %v %v\n", gen, gen, err)
   topics := []string{"topic1","topic2"}
   arg1 := "a string arg"
   arg2 := []string{"table1","table2"}
   args := []interface{}{arg1,arg2}
   fmt.Printf("\tGenerate:", gen.Generate(topics, args))
}

