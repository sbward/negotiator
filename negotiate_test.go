package negotiator

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddCustomResponseProcessors(t *testing.T) {

	var fakeResponseProcessor = &fakeProcessor{}

	New(fakeResponseProcessor)

	assert.Equal(t, 3, len(processors))
}

func TestShouldAddCustomResponseProcessorsToBeginning(t *testing.T) {

	var fakeResponseProcessor = &fakeProcessor{}

	New(fakeResponseProcessor)

	firstProcessor := processors[0]
	processorName := reflect.TypeOf(firstProcessor).String()

	assert.Equal(t, "*negotiator.fakeProcessor", processorName)
}

type fakeProcessor struct {
}

func (*fakeProcessor) CanProcess(mediaRange string) bool {
	return true
}

func (*fakeProcessor) Process(w http.ResponseWriter, model interface{}) {

}