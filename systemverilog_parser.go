package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/nickjones/systemverilog_parser/services/parser"
)

var (
	debug = flag.Bool("debug", false, "Enable debug logging")
)

func init() {
	flag.Parse()
}

func main() {
	sampleInput := `
package pkg;

// Import UVM package
import uvm_pkg::*;

/* Example multiline
 * comment.
*/
class my_component extends uvm_component;
  // Some member variable examples
  static int class_var = 2;
  int int_var;
  int int_var_init = 3;
  longint unsigned lint_unsigned = 4;

  // Constructor
  function new(string name, uvm_component parent);
    super.new(name, parent);
  endfunction

  task run_phase(uvm_phase phase);
    phase.raise_objection(this);
    phase.drop_objection(this);
  endtask

endclass

endpackage

// Top-level module
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
	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	_ = parser.Parse("sample.sv", sampleInput)
	// prettyJSON, err := json.MarshalIndent(parsedINIFile, "", "   ")

	// if err != nil {
	// 	log.Println("Error marshalling JSON:", err.Error())
	// 	return
	// }

	// log.Println(string(prettyJSON))
}
