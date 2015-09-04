package main

import (
	"encoding/json"
	"log"

	"github.com/nickjones/systemverilog_parser/services/parser"
)

func main() {
	sampleInput := `
		package pkg;

		import uvm_pkg::*;

		class my_component extends uvm_component;

		  function new(string name, uvm_component parent);
		    super.new(name, parent);
		  endfunction

		  task run_phase(uvm_phase phase);
		    phase.raise_objection(this);
		    phase.drop_objection(this);
		  endtask

		endclass

		endpackage

		module test;
		   import uvm_pkg::*;
		   import pkg::*;
		   my_component t;

		   initial begin
		      t = new("Top", null);
		      run_test();
		   end
		endmodule
	`

	parsedINIFile := parser.Parse("sample.sv", sampleInput)
	prettyJSON, err := json.MarshalIndent(parsedINIFile, "", "   ")

	if err != nil {
		log.Println("Error marshalling JSON:", err.Error())
		return
	}

	log.Println(string(prettyJSON))
}
