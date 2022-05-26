package registration

import (
	"flag"
	"github.com/datumtechs/datum-network-carrier/common/flags"
	"github.com/datumtechs/datum-network-carrier/params"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func TestP2PPreregistration_DefaultDataDir(t *testing.T) {
	app := cli.App{}
	set := flag.NewFlagSet("test", 0)
	set.String(flags.DataDirFlag.Name, "", "")
	ctx := cli.NewContext(&app, set, nil)

	_, dataDir, err := P2PPreregistration(ctx)
	require.NoError(t, err)
	assert.Equal(t, flags.DefaultDataDir(), dataDir)
}

func TestP2PPreregistration(t *testing.T) {
	sampleNode := "- enr:-TESTNODE"
	testDataDir := "testDataDir"

	file, err := ioutil.TempFile(t.TempDir(), "bootstrapFile*.yaml")
	defer file.Close()
	require.NoError(t, err)
	err = ioutil.WriteFile(file.Name(), []byte(sampleNode), 0644)
	require.NoError(t, err, "Error in WriteFile call")
	//params.SetupTestConfigCleanup(t)
	config := params.CarrierNetworkConfig()
	config.BootstrapNodes = []string{file.Name()}
	params.OverrideCarrierNetworkConfig(config)

	app := cli.App{}
	set := flag.NewFlagSet("test", 0)
	set.String(flags.DataDirFlag.Name, testDataDir, "")
	ctx := cli.NewContext(&app, set, nil)

	bootstrapNodeAddrs, dataDir, err := P2PPreregistration(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(bootstrapNodeAddrs))
	assert.Equal(t, sampleNode[2:], bootstrapNodeAddrs[0])
	assert.Equal(t, testDataDir, dataDir)
}

func TestBootStrapNodeFile(t *testing.T) {
	file, err := ioutil.TempFile(t.TempDir(), "bootstrapFile")
	defer file.Close()
	require.NoError(t, err)

	sampleNode0 := "- enr:-Ku4QMKVC_MowDsmEa20d5uGjrChI0h8_KsKXDmgVQbIbngZV0i" +
		"dV6_RL7fEtZGo-kTNZ5o7_EJI_vCPJ6scrhwX0Z4Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD" +
		"1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQJxCnE6v_x2ekgY_uo" +
		"E1rtwzvGy40mq9eD66XfHPBWgIIN1ZHCCD6A"
	sampleNode1 := "- enr:-TESTNODE2"
	sampleNode2 := "- enr:-TESTNODE3"
	err = ioutil.WriteFile(file.Name(), []byte(sampleNode0+"\n"+sampleNode1+"\n"+sampleNode2), 0644)
	require.NoError(t, err, "Error in WriteFile call")
	nodeList, _, err := readbootNodes(file.Name())
	require.NoError(t, err, "Error in readbootNodes call")
	assert.Equal(t, sampleNode0[2:], nodeList[0], "Unexpected nodes")
	assert.Equal(t, sampleNode1[2:], nodeList[1], "Unexpected nodes")
	assert.Equal(t, sampleNode2[2:], nodeList[2], "Unexpected nodes")
}

func TestStaticNodeFile(t *testing.T) {
	file, err := ioutil.TempFile(t.TempDir(), "staticFile")
	defer file.Close()
	require.NoError(t, err)

	sampleNode := `[
	"enr:-TESTNODE1",
	"enr:-TESTNODE2"
]`
	err = ioutil.WriteFile(file.Name(), []byte(sampleNode), 0644)
	require.NoError(t, err, "Error in WriteFile call")
	nodeList, err := readStaticNodesFromJSON(file.Name())
	require.NoError(t, err, "Error in staticNodes call")
	t.Logf("%v", nodeList)
}