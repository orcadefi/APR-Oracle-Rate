# APR-Oracle-Rate

APR Oracle Rate is a Go project for get % Lend/Borrow APR Rate on-chain realtime data with Smart Contracts and the Official APIs.

Supported platforms:
- Aave
    - Lend / Borrow
      - TUSD
      - YFI
      - BAT
      - MANA
      - REP
      - WBTC
      - REN
      - LINK
      - DAI
      - LEND 
      - MKR 
      - SNX 
      - KNC 
      - ZRX
      - ETH
      - ENJ
- Compound
    - Lend / Borrow
        - REP (Only Borrow)
        - ZRX
        - SAI (Only Borrow)
        - BAT
        - WBTC
        - USDC
        - DAI
        - ETH
        - USDT
- Dydx
    - Lend / Borrow
        - ETH
        - USDC
        - DAI
- Fulcrum
  - Lend / Borrow
    - ETH
    - USDC
    - WBTC
    - KNC
    - LINK
    - DAI
    - USDT
    - LEND (Only Borrow)
    - MKR (Only Borrow)
    - BZRX (Only Borrow)
    - YFI (Only Borrow)
    - ETHv1 (Only Borrow)

## Requisites

[Install Go v1.15](https://golang.org/doc/install)

Go packages:
```
go get github.com/ethereum/go-ethereum/common
go get github.com/ethereum/go-ethereum/ethclient
```

## Installation

Use the git to clone this repository.

```bash
git clone https://github.com/orcadefi/APR-Oracle-Rate.git
cd APR-Oracle-Rate
```

## Usage

Get data of the DeFI pair rates of each platform:

####API

Aave:
```bash
go run aave_api.go
```

Compound:
```bash
go run compound_api.go
```

Dydx:
```bash
go run dydx_api.go
```

Fulcrum:
```bash
go run fulcrum_api.go
```

####Example output
```
   Dydx % APR API
   ---------------------------
   ETH
   Lend Rate 0.034511014316548225 %
   Borrow Rate:  0.6027220206864 %
   ---------------------------
   SAI
   Lend Rate 0 %
   Borrow Rate:  0 %
   ---------------------------
   USDC
   Lend Rate 2.9472773888341472 %
   Borrow Rate:  5.5699167780864 %
   ---------------------------
   DAI
   Lend Rate 3.8155297992568338 %
   Borrow Rate:  6.3374679973584 %
   
   Process finished with exit code 0

```

####Smart Contract
All platforms:
```bash
go run apr_oracle_smart_contract.go
```

####Example output
```
   % APR Lending Rate
   Aave:
   ABAT: 22132134023949084
   ADAI: 299276540560146987
   AETH: 13672812867432560
   AKNC: 70066385857411234
   ALINK: 2008293004005908
   AREP: 213856676943092
   ASNX: 5210520108752593
   ASUSD: 3344101109887360
   ATUSD: 2977567471409835
   AUSDC: 34627109141419945
   AUSDT: 17114076855615664
   AZRX: 111938687161188
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)