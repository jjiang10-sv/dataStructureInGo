package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_topsort1(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the path to your circuit file:\n")
	scanner.Scan()
	path := scanner.Text()
	//path := "eightBitCounter.chdl"
	components, terminalComponents := parseFile(path)

	fmt.Printf("Parsed file sucessfully.\n")
	fmt.Printf("Found:\nComponents: %d\nTerminals: %d\n",
		len(components), len(terminalComponents))

	var clk *Component
	numPulses := 0
	pulseCount := 0
	// Kick off components
	for i := range components {
		componentPtr := &components[i]

		var arg int
		switch componentPtr.typeName {

		case "source":
			arg = getInitialValue(scanner, componentPtr.name, "Source value:", true)

		case "dff":
			arg = getInitialValue(scanner, componentPtr.name, "Initial value:", true)

		case "clk":
			clk = componentPtr
			arg = getInitialValue(scanner, componentPtr.name, "Clock frequency:", false)
			numPulses = getInitialValue(scanner, "", "# of clock pulses to run:", false)

		default:
			arg = 0

		}

		go componentPtr.handler(componentPtr, arg)
	}

	// Receive from all terminal channels
	lastValue := -1
	terminator := "\n"
	for {
		var outValues []bool
		for i := range terminalComponents {
			chanPtrs := (*terminalComponents[i]).terminals
			for j := range chanPtrs {
				chanPtr := chanPtrs[j]
				val := <-(*chanPtr)
				outValues = append(outValues, val)
			}
		}
		outNum := parseOutputs(outValues)
		if outNum != lastValue {
			fmt.Printf("[Output] %08b%s", outNum, terminator)
			lastValue = outNum

			// If we don't have a clock, the output isn't going to change again
			// so we can exit
			if clk == nil {
				break
			}

		}

		// If we have a clock, notify it that the circuit propagation is
		// complete
		if clk != nil {
			clk.clkSync <- true
			pulseCount++
			if pulseCount/2 > numPulses {
				break
			}
		}
	}
}
