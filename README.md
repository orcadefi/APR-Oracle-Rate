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

#### API

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

#### Example output
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

#### Smart Contract
All platforms (Only Lend data and some pairs):
```bash
go run apr_oracle_smart_contract.go
```

All platforms (Lend/Borrow):
```bash
go run apr_oracle_custom_smart_contract.go
```

#### Example output
```
   Aave Platform Lend/Borrow
   ABAT
   733411039147732
   5963026150844428
   ADAI
   30411433388962662
   53389709338899109
   AETH
   13140628688712371
   33636230041130821
   AKNC
   4261491725355918
   11412548441584914
   ALEND
   268913486904
   114243701253093
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
