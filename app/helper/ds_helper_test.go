package helper

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSaveMapTo_TextFile(t *testing.T) {
	err:=SaveMapTo_TextFile("https://google.com/sampleurl/sample1/sample5","c218e3")
	require.NoError(t,err)
}

func TestCheckURL_TextFile(t *testing.T) {
	enc_url, _ := CheckURL_TextFile("https://google.com/sampleurl/sample1/sample5")
	require.Equal(t,enc_url,"c218e3","enc_url successfully retrieved")	
}
