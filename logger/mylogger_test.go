package logger

import (
	"github.com/guebu/common-utils/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfo(t *testing.T) {
	msg0 := "Test-Info"

	msg1 := "Test-Info-1 with one tag"
	msg2 := "Test-Info-2 with two tags"
	msg3 := "Test-Info-3 with three tags"

	tag1 := "tag1:T1"
	tag2 := "tag2:T2"
	tag3 := "tag3:T3"

	Info(msg0)
	assert.EqualValues(t, "Test-Info", msg0 )

	Info(msg1, tag1)
	assert.EqualValues(t, "Test-Info-1 with one tag", msg1 )

	Info(msg2, tag1+tag2)
	assert.EqualValues(t, "Test-Info-2 with two tags", msg2 )

	Info(msg3, tag1+tag2+tag3)
	assert.EqualValues(t, "Test-Info-3 with three tags", msg3 )
}

func TestError(t *testing.T) {

	err := errors.NewNotYetImplementedError("Not yet implemented...")

	msg0 := "Error-Info"

	msg1 := "Error-Info-1 with one tag"
	msg2 := "Error-Info-2 with two tags"
	msg3 := "Error-Info-3 with three tags"

	tag1 := "tag1:T1"
	tag2 := "tag2:T2"
	tag3 := "tag3:T3"

	Error(msg0, err)
	assert.EqualValues(t, "Error-Info", msg0 )

	Error(msg1, err, tag1)
	assert.EqualValues(t, "Error-Info-1 with one tag", msg1 )

	Error(msg2, err, tag1+tag2)
	assert.EqualValues(t, "Error-Info-2 with two tags", msg2 )

	Error(msg3, err, tag1+tag2+tag3)
	assert.EqualValues(t, "Error-Info-3 with three tags", msg3 )
}


func TestDebug(t *testing.T) {

	msg0 := "Debug-Info"

	msg1 := "Debug-Info-1 with one tag"
	msg2 := "Debug-Info-2 with two tags"
	msg3 := "Debug-Info-3 with three tags"

	tag1 := "tag1:T1"
	tag2 := "tag2:T2"
	tag3 := "tag3:T3"

	Debug(msg0)
	assert.EqualValues(t, "Debug-Info", msg0 )

	Debug(msg1, tag1)
	assert.EqualValues(t, "Debug-Info-1 with one tag", msg1 )

	Debug(msg2, tag1+tag2)
	assert.EqualValues(t, "Debug-Info-2 with two tags", msg2 )

	Debug(msg3, tag1+tag2+tag3)
	assert.EqualValues(t, "Debug-Info-3 with three tags", msg3 )

}

func TestPrintf(t *testing.T){
	var logger myLogger
	logger.Printf("One String")
	logger.Printf ("One String with %s", "value")
}

