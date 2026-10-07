package print_tools

import "fmt"

func PrintLogicalOperation(operation string, param1 int, param2 int) {
	var result int

	switch operation {
		case "&":
			result = param1 & param2
		case "|":
			result = param1 | param2
		case "&^":
			result = param1 &^ param2
		case "^":
			result = param1 ^ param2
		default: 
			fmt.Printf("%s - не верный оператор\n", operation)
			return
	}

	fmt.Printf("%d (%b) %s %d (%b) = %d (%b)\n", param1, param1, operation, param2, param2, result, result)
}