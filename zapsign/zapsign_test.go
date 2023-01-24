package zapsign_test

import (
	"testing"

	"github.com/qbem-repos/zapsign-sdk-go/zapsign"
	"github.com/qbem-repos/zapsign-sdk-go/zapsign/request"
)

// Test_NewClient_EmptyToken_ReturnsError Test new client creation empty token returns error
func Test_NewClient_EmptyToken_ReturnsError(t *testing.T) {
	if _, err := zapsign.NewClient(""); err == nil {
		t.Fail()
	}
}

func TestX(t *testing.T) {
	_, err := zapsign.NewClient("")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	_, err = request.NewDocs("", "", "")

	if err == nil {
		t.Log(err)
		t.Fail()
	}

}

func TestXxx(t *testing.T) {
	zcl, err := zapsign.NewClient("")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	docs, err := request.NewDocs("", "", "")
	if err == nil {
		t.Log(err)
		t.Fail()
	}

	_, err = zcl.CreateDocumentViaUpload(docs)
	if err == nil {
		t.Log(err)
		t.Fail()
	}
}
