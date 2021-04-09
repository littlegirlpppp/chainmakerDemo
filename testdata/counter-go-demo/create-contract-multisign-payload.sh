#cmc payload create contract --chain-id chain1 --contract-name counter-go-1 --version 1.0.0 --runtime WASMER_RUST --method init  --kv-pairs "" --byte-code-path ./counter-go.wasm --output collect.pb
cmc payload create contract --chain-id chain1 --contract-name counter-go-1 --version 1.0.0 --runtime GASM_CPP --method init  --kv-pairs "" --byte-code-path ./counter-go.wasm --output collect.pb

cmc payload sign contract --input collect.pb --org-id wx-org1.chainmaker.org --admin-key-path ./crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.crt --output collect-signed-org1.pb
cmc payload sign contract --input collect.pb --org-id wx-org2.chainmaker.org --admin-key-path ./crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.crt --output collect-signed-org2.pb
cmc payload sign contract --input collect.pb --org-id wx-org3.chainmaker.org --admin-key-path ./crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.crt --output collect-signed-org3.pb
cmc payload sign contract --input collect.pb --org-id wx-org4.chainmaker.org --admin-key-path ./crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.crt --output collect-signed-org4.pb

cmc payload merge contract --input collect-signed-org1.pb --input collect-signed-org2.pb --input collect-signed-org3.pb --input collect-signed-org4.pb --output collect-signed-all.pb



cmc payload create contract --chain-id chain1 --contract-name counter-go-1 --version 2.0.0 --runtime GASM_CPP --method upgrade  --kv-pairs "" --byte-code-path ./counter-go-upgrade.wasm --output upgrade-collect.pb

cmc payload sign contract --input upgrade-collect.pb --org-id wx-org1.chainmaker.org --admin-key-path ./crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org1.chainmaker.org/user/admin1/admin1.sign.crt --output upgrade-collect-signed-org1.pb
cmc payload sign contract --input upgrade-collect.pb --org-id wx-org2.chainmaker.org --admin-key-path ./crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org2.chainmaker.org/user/admin1/admin1.sign.crt --output upgrade-collect-signed-org2.pb
cmc payload sign contract --input upgrade-collect.pb --org-id wx-org3.chainmaker.org --admin-key-path ./crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org3.chainmaker.org/user/admin1/admin1.sign.crt --output upgrade-collect-signed-org3.pb
cmc payload sign contract --input upgrade-collect.pb --org-id wx-org4.chainmaker.org --admin-key-path ./crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.key --admin-crt-path ./crypto-config/wx-org4.chainmaker.org/user/admin1/admin1.sign.crt --output upgrade-collect-signed-org4.pb

cmc payload merge contract --input upgrade-collect-signed-org1.pb --input upgrade-collect-signed-org2.pb --input upgrade-collect-signed-org3.pb --input upgrade-collect-signed-org4.pb --output upgrade-collect-signed-all.pb
