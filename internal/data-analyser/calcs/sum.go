package calcs

import (
	"github.com/spyzhov/ajson"
)

func SumJsonNumbers(data []byte) (float64, error) {
	var result float64
	node, err := ajson.Unmarshal(data)
	if err != nil {
		return result, err
	}
	calcNode(&result, node)
	return result, nil
}

func calcNode(result *float64, node *ajson.Node) {

	if node.IsNumeric() {
		val, _ := node.GetNumeric()
		*result = *result + val
		// worker job is done
	} else if node.IsObject() {
		for _, k := range node.Keys() {
			n, err := node.GetKey(k)
			if err != nil {
				continue
			}
			calcNode(result, n)
		}
	} else if node.IsArray() {
		for _, a := range node.MustArray() {
			// send a new job
			calcNode(result, a)
		}
	}
}
