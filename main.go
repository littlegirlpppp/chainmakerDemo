package main

import (
	"chainmaker.org/chainmaker-go/common/log"
	sdk "chainmaker.org/chainmaker-sdk-go"
	"chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
	"chainmaker.org/chainmaker-sdk-go/pb/protogo/config"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

const (
	createContractTimeout = 5
	chainId        = "chain1"
	orgId1         = "wx-org1.chainmaker.org"
	orgId2         = "wx-org2.chainmaker.org"
	orgId3         = "wx-org3.chainmaker.org"
	orgId4         = "wx-org4.chainmaker.org"
	orgId5         = "wx-org5.chainmaker.org"
	orgId6         = "wx-org6.chainmaker.org"

	certPathPrefix = "./testdata"
	tlsHostName    = "chainmaker.org"


	nodeAddr1 = "127.0.0.1:12301"
	connCnt1  = 5

	nodeAddr2 = "127.0.0.1:12301"
	connCnt2  = 5

	multiSignedPayloadFile        = "./testdata/counter-go-demo/collect-signed-all.pb"
	upgradeMultiSignedPayloadFile = "./testdata/counter-go-demo/upgrade-collect-signed-all.pb"

	byteCodePath        = "./testdata/counter-go-demo/counter-go.wasm"
	upgradeByteCodePath = "./testdata/counter-go-demo/counter-go-upgrade.wasm"

	certPathFormat = "/crypto-config/%s/ca"
)
var (
	caPaths = []string{
		certPathPrefix + fmt.Sprintf(certPathFormat, orgId1),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgId2),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgId3),
		certPathPrefix + fmt.Sprintf(certPathFormat, orgId4),
	}

	caCerts = []string{"-----BEGIN CERTIFICATE-----\nMIICsDCCAlWgAwIBAgIDAuGKMAoGCCqBHM9VAYN1MIGKMQswCQYDVQQGEwJDTjEQ\nMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEfMB0GA1UEChMWd3gt\nb3JnMS5jaGFpbm1ha2VyLm9yZzESMBAGA1UECxMJcm9vdC1jZXJ0MSIwIAYDVQQD\nExljYS53eC1vcmcxLmNoYWlubWFrZXIub3JnMB4XDTIxMDMyNTA2NDI1MVoXDTMx\nMDMyMzA2NDI1MVowgYoxCzAJBgNVBAYTAkNOMRAwDgYDVQQIEwdCZWlqaW5nMRAw\nDgYDVQQHEwdCZWlqaW5nMR8wHQYDVQQKExZ3eC1vcmcxLmNoYWlubWFrZXIub3Jn\nMRIwEAYDVQQLEwlyb290LWNlcnQxIjAgBgNVBAMTGWNhLnd4LW9yZzEuY2hhaW5t\nYWtlci5vcmcwWTATBgcqhkjOPQIBBggqgRzPVQGCLQNCAARIG6tdLNtG+eqwTK36\nS/AjzXh9Q0Zwrf7eqyCEQ4Ul7xfgKjCBNVboivH10ieYuh0MAoZj1Ke7z+P6ZUTy\naiuDo4GnMIGkMA4GA1UdDwEB/wQEAwIBpjAPBgNVHSUECDAGBgRVHSUAMA8GA1Ud\nEwEB/wQFMAMBAf8wKQYDVR0OBCIEIJDsy2L0fAK2V4YxOjVEjYj3YKSbX4F24eh0\nZQHoqCr1MEUGA1UdEQQ+MDyCDmNoYWlubWFrZXIub3Jngglsb2NhbGhvc3SCGWNh\nLnd4LW9yZzEuY2hhaW5tYWtlci5vcmeHBH8AAAEwCgYIKoEcz1UBg3UDSQAwRgIh\nAM1oJOU6l4tJVqrCJv5UnMaKLxu4V1dDwu0YsS5Tb1s9AiEA1D8NA3GGy9BEFryq\n5TS0uiqE3QEuDRvs1TrP9H53Sjk=\n-----END CERTIFICATE-----",}

	userKeyPath = certPathPrefix + "/crypto-config/%s/user/client1/client1.tls.key"
	userCrtPath = certPathPrefix + "/crypto-config/%s/user/client1/client1.tls.crt"

	userSignKeyPath = certPathPrefix + "/crypto-config/%s/user/client1/client1.sign.key"
	userSignCrtPath = certPathPrefix + "/crypto-config/%s/user/client1/client1.sign.crt"

	adminKeyPath = certPathPrefix + "/crypto-config/%s/user/admin1/admin1.tls.key"
	adminCrtPath = certPathPrefix + "/crypto-config/%s/user/admin1/admin1.tls.crt"
)

func main()  {

	client, err := createClientWithCertBytes()
    if err!=nil{

	}
	admin1, err := createAdmin(orgId1)
	admin2, err := createAdmin(orgId2)
	admin3, err := createAdmin(orgId3)
	admin4, err := createAdmin(orgId4)

	contractName   := "counter-go-11"
	version        := "1.0.0"
	byteCodePath := "./main1.wasm"

	//创建合约
	UserContractCounterGoCreate(client, admin1, admin2, admin3, admin4, contractName,version,byteCodePath,true)

	//调用合约
	params:=make(map[string]string)
	params["key"]="22222"
	UserContractCounterGoInvoke(client,contractName,"add",params,true)
    //查询合约
	UserContractCounterGoQuery(client,contractName,"getdata",params)
	return
	//新增共识节点  需要先把节点启动起来
	// 3) [TrustRootAdd] 添加信任配置
	trustCount := len(GetChainConfig(client).TrustRoots)
	raw, err := ioutil.ReadFile("testdata/crypto-config/wx-org5.chainmaker.org/ca/ca.crt")
	trustRootOrgId := orgId5
	trustRootCrt := string(raw)
	ChainConfigTrustRootAdd(client, admin1, admin2, admin3, admin4, trustRootOrgId, trustRootCrt)
	time.Sleep(2 * time.Second)
	chainConfig := GetChainConfig(client)
	if trustCount==len(chainConfig.TrustRoots)-1{

	}
	//require.Equal(t, trustCount+1, len(chainConfig.TrustRoots))
	//require.Equal(t, trustRootOrgId, chainConfig.TrustRoots[trustCount].OrgId)
	//require.Equal(t, trustRootCrt, chainConfig.TrustRoots[trustCount].Root)

	nodeOrgId := orgId5
	nodeAddresses := []string{"/ip4/10.190.28.222/tcp/11305/p2p/Qmdixv2PbNxoxSakgPKdbikbHtiFYTwELP2p3EESMCdrCR"}
	ChainConfigConsensusNodeOrgAdd( client, admin1, admin2, admin3, admin4, nodeOrgId, nodeAddresses)
	time.Sleep(2 * time.Second)
	chainConfig = GetChainConfig( client)
}

func ChainConfigConsensusNodeAddrAdd( client,
	admin1, admin2, admin3, admin4 *sdk.ChainClient,
	nodeAddrOrgId string, nodeAddresses []string) {

	// 配置块更新payload生成
	payloadBytes, err := client.CreateChainConfigConsensusNodeAddrAddPayload(nodeAddrOrgId, nodeAddresses)
	if err!=nil{

	}
	signAndSendRequest(client, admin1, admin2, admin3, admin4, payloadBytes)
}

func ChainConfigConsensusNodeOrgAdd( client,
	admin1, admin2, admin3, admin4 *sdk.ChainClient,
	nodeAddrOrgId string, nodeAddresses []string) {

	// 配置块更新payload生成
	payloadBytes, err := client.CreateChainConfigConsensusNodeOrgAddPayload(nodeAddrOrgId, nodeAddresses)
	if err!=nil{

	}

	signAndSendRequest(client, admin1, admin2, admin3, admin4, payloadBytes)
}

func GetChainConfig( client *sdk.ChainClient) *config.ChainConfig {
	resp, err := client.GetChainConfig()
	if err!=nil{

	}
	return resp
}
func ChainConfigTrustRootAdd( client,
	admin1, admin2, admin3, admin4 *sdk.ChainClient,
	trustRootOrgId, trustRootCrt string) {

	// 配置块更新payload生成
	payloadBytes, err := client.CreateChainConfigTrustRootAddPayload(trustRootOrgId, trustRootCrt)
	if err!=nil{

	}
	signAndSendRequest( client, admin1, admin2, admin3, admin4, payloadBytes)
}
//收集签名
func signAndSendRequest( client,
	admin1, admin2, admin3, admin4 *sdk.ChainClient,
	payloadBytes []byte) {
	// 各组织Admin权限用户签名
	signedPayloadBytes1, err := admin1.SignChainConfigPayload(payloadBytes)
    if err!=nil{

	}

	signedPayloadBytes2, err := admin2.SignChainConfigPayload(payloadBytes)


	signedPayloadBytes3, err := admin3.SignChainConfigPayload(payloadBytes)


	signedPayloadBytes4, err := admin4.SignChainConfigPayload(payloadBytes)


	// 收集并合并签名
	mergeSignedPayloadBytes, err := client.MergeChainConfigSignedPayload([][]byte{signedPayloadBytes1,
		signedPayloadBytes2, signedPayloadBytes3, signedPayloadBytes4})


	// 发送配置更新请求
	resp, err := client.SendChainConfigUpdateRequest(mergeSignedPayloadBytes)


	err = checkProposalRequestResp(resp, true)


	fmt.Printf("chain config [CoreUpdate] resp: %+v", resp)
}


// 创建ChainClient（指定证书内容）
func createClientWithCertBytes() (*sdk.ChainClient, error) {

	userCrtBytes, err := ioutil.ReadFile(fmt.Sprintf(userCrtPath, orgId1))
	if err != nil {
		return nil, err
	}

	userKeyBytes, err := ioutil.ReadFile(fmt.Sprintf(userKeyPath, orgId1))
	if err != nil {
		return nil, err
	}

	userSignCrtBytes, err := ioutil.ReadFile(fmt.Sprintf(userSignCrtPath, orgId1))
	if err != nil {
		return nil, err
	}

	userSignKeyBytes, err := ioutil.ReadFile(fmt.Sprintf(userSignKeyPath, orgId1))
	if err != nil {
		return nil, err
	}

	chainClient, err := sdk.NewChainClient(
		sdk.WithConfPath("./testdata/sdk_config.yml"),
		sdk.WithUserCrtBytes(userCrtBytes),
		sdk.WithUserKeyBytes(userKeyBytes),
		sdk.WithUserSignKeyBytes(userSignKeyBytes),
		sdk.WithUserSignCrtBytes(userSignCrtBytes),
	)

	if err != nil {
		return nil, err
	}

	//启用证书压缩（开启证书压缩可以减小交易包大小，提升处理性能）
	err = chainClient.EnableCertHash()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func createAdmin(orgId string) (*sdk.ChainClient, error) {
	if node1 == nil {
		node1 = createNode(nodeAddr1, connCnt1)
	}

	if node2 == nil {
		node2 = createNode(nodeAddr2, connCnt2)
	}

	adminClient, err := sdk.NewChainClient(
		sdk.WithChainClientOrgId(orgId),
		sdk.WithChainClientChainId(chainId),
		sdk.WithChainClientLogger(getDefaultLogger()),
		sdk.WithUserKeyFilePath(fmt.Sprintf(adminKeyPath, orgId)),
		sdk.WithUserCrtFilePath(fmt.Sprintf(adminCrtPath, orgId)),
		sdk.AddChainClientNodeConfig(node1),
		sdk.AddChainClientNodeConfig(node2),
	)
	if err != nil {
		return nil, err
	}

	//启用证书压缩（开启证书压缩可以减小交易包大小，提升处理性能）
	err = adminClient.EnableCertHash()
	if err != nil {
		return nil, err
	}

	return adminClient, nil
}

func getDefaultLogger() *zap.SugaredLogger {
	config := log.LogConfig{
		Module:       "[SDK]",
		LogPath:      "./sdk.log",
		LogLevel:     log.LEVEL_DEBUG,
		MaxAge:       30,
		JsonFormat:   false,
		ShowLine:     true,
		LogInConsole: true,
	}

	logger, _ := log.InitSugarLogger(&config)
	return logger
}
var (
	node1 *sdk.NodeConfig
	node2 *sdk.NodeConfig
)

// 创建节点
func createNode(nodeAddr string, connCnt int) *sdk.NodeConfig {
	node := sdk.NewNodeConfig(
		// 节点地址，格式：127.0.0.1:12301
		sdk.WithNodeAddr(nodeAddr),
		// 节点连接数
		sdk.WithNodeConnCnt(connCnt),
		// 节点是否启用TLS认证
		sdk.WithNodeUseTLS(true),
		// 根证书路径，支持多个
		sdk.WithNodeCAPaths(caPaths),
		// TLS Hostname
		sdk.WithNodeTLSHostName(tlsHostName),
	)

	return node
}

func UserContractCounterGoCreate(client *sdk.ChainClient,
	admin1, admin2, admin3, admin4 *sdk.ChainClient,contractName, version, byteCodePath string, withSyncResult bool) (*common.TxResponse, error) {
	resp, err := createUserContract(client, admin1, admin2, admin3, admin4,
		contractName, version, byteCodePath, common.RuntimeType_GASM, []*common.KeyValuePair{}, withSyncResult)
	return resp,err
}

func createUserContract(client *sdk.ChainClient, admin1, admin2, admin3, admin4 *sdk.ChainClient,
	contractName, version, byteCodePath string, runtime common.RuntimeType, kvs []*common.KeyValuePair, withSyncResult bool) (*common.TxResponse, error) {

	payloadBytes, err := client.CreateContractCreatePayload(contractName, version, byteCodePath, runtime, kvs)
	if err != nil {
		return nil, err
	}

	// 各组织Admin权限用户签名
	signedPayloadBytes1, err := admin1.SignContractManagePayload(payloadBytes)
	if err != nil {
		return nil, err
	}

	signedPayloadBytes2, err := admin2.SignContractManagePayload(payloadBytes)
	if err != nil {
		return nil, err
	}

	signedPayloadBytes3, err := admin3.SignContractManagePayload(payloadBytes)
	if err != nil {
		return nil, err
	}

	signedPayloadBytes4, err := admin4.SignContractManagePayload(payloadBytes)
	if err != nil {
		return nil, err
	}

	// 收集并合并签名
	mergeSignedPayloadBytes, err := client.MergeContractManageSignedPayload([][]byte{signedPayloadBytes1,
		signedPayloadBytes2, signedPayloadBytes3, signedPayloadBytes4})
	if err != nil {
		return nil, err
	}

	// 发送创建合约请求
	resp, err := client.SendContractManageRequest(mergeSignedPayloadBytes, createContractTimeout, withSyncResult)
	if err != nil {
		return nil, err
	}

	err = checkProposalRequestResp(resp, true)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func checkProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != common.ContractResultCode_OK {
		return errors.New(resp.ContractResult.Message)
	}
	return nil
}

func UserContractCounterGoQuery(client *sdk.ChainClient,
	contractName,method string, params map[string]string)(*common.TxResponse, error) {
	resp, err := client.QueryContract(contractName, method, params, -1)
	if err=checkProposalRequestResp(resp,true);err!=nil{
		return nil,err
	}
	return resp,nil
}


func UserContractCounterGoInvoke(client *sdk.ChainClient,
	contractName,method string, params map[string]string, withSyncResult bool)(*common.TxResponse, error) {
	resp,err := invokeUserContract(client, contractName, method, "", params, withSyncResult)
	return resp,err
}

func invokeUserContract(client *sdk.ChainClient, contractName, method, txId string, params map[string]string, withSyncResult bool) (*common.TxResponse, error) {

	resp, err := client.InvokeContract(contractName, method, txId, params, -1, withSyncResult)
	if err != nil {
		return nil,err
	}

     if err=checkProposalRequestResp(resp,true);err!=nil{
     	return nil,err
	 }

	//if !withSyncResult {
	//	fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
	//} else {
	//	fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	//}
	return resp,nil
}

