package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func childrenPrinter(children []*node) string {
	if l := len(children); l == 0 {
		return ""
	} else {
		str := nodePrinter(children[0])
		for _, node := range children[1:] {
			str += "," + nodePrinter(node)
		}

		return str
	}
}

func nodePrinter(node *node) string {
	return fmt.Sprintf(`
{
  "id": %d,
  "name": %q,
  "action": %q,
  "status": %q,
  "message": %q,
  "children": [%s]
}`, node.id, node.name, node.action, node.status, node.message, childrenPrinter(node.children))
}

func TestBuildTree(t *testing.T) {
	tasks, err := newTaskFromJSONFile("./tasks/user-login.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	tree := tasks.createTree(&ddtracerwrapper{})
	jsonstr := nodePrinter(tree.root)
	if !json.Valid([]byte(jsonstr)) {
		log.Fatalln("invalid JSON string")
	}

	log.Println(nodePrinter(tree.root))
}
