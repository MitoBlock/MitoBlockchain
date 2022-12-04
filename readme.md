# Mito Blockchain

It is difficult for start-ups and small businesses to collaborate with each other to boost their sales / revenue / number of users and provide strong elements for user stickiness. Our application is a blockchain protocol that acts as an ecosystem for startups and small businesses to freely exchange value with each other with the help of Blockchain & Web3. 

**Mito Block** is a reliable, secure, and scalable PoA / PoS blockchain with WASM for smart contracts. It is built  using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Tech Stack of our application 

| Component / Category                                             | Technology being used                  |
| -------------                                                    |:-------------:                         | 
| State machine / Application Layer of the blockchain              | Cosmos SDK                             | 
| Consensus Engine / Consensus and Network Layer of the blockchain | Tendermint Core                        |
| Smart contract Functionality (WASM)                              | CosmWasm                               |
| CLI Tool for creating and managing the blockchain                | Starport/Ignite                        |
| Interoperability                                                 | IBC & Peg-zones                        |
| RPC Framework                                                    | gRPC                                   |
| Languages                                                        | JavaScript, TypeScript, Go, C and Rust |

## Contributors
Daljeet Singh (Lead), Kyro Clestino and Andrew Pereira (WASM)

## Get started

In order to run this blockchain node, you will need the following:
1. [Go](https://go.dev/)
2. [Ignite CLI](https://docs.ignite.com/guide/install)

To run the chain,
```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development. The `-r` flag can be used to reset the state.

## Common API endpoints 

#### GET HTTP method

* Get balance of an account (all coins) - http://0.0.0.0:1317/cosmos/bank/v1beta1/balances/{address}
* Get balance of an account by denom (specific coin) - http://0.0.0.0:1317/cosmos/bank/v1beta1/balances/{address}/by_denom?denom={coin like mitocell}
* Get supply of the coin - http://0.0.0.0:1317/cosmos/bank/v1beta1/supply/by_denom?denom={coin like mitocell}
* Get discount token and its status using ID - http://0.0.0.0:1317/mitoblockchaindev/mitoblockchaindev/discount_token_status_q/{id}
* Get all discount tokens - http://0.0.0.0:1317/mitoblockchaindev/mitoblockchaindev/discount_tokens
Membership token endpoints are similar to discount token

#### POST HTTP method 

You can use faucet to get some coins. Faucet is a service that gives free coins to users. faucet has limited coins or funds.

**Send coins to your account using faucet - http://0.0.0.0:4500**

You will need to pass the address and coins in the body of the request.

Example of a body:
{
  "address": "mito1uzv4v9g9xln2qx2vtqhz99yxum33calja5vruz",
  "coins": [
    "10token"
  ]
}

Parameter content type is application/json

Other TXs like transfering coins from one account to another can only be done through CLI or client code. That's why we thought of creating a SDK using ts-client but current unstability failed the plan. In the end, we created the [go client for startup web apps using GIN](https://github.com/MitoBlock/go-client). Same strategy was used to create the [onboarding program web app](https://github.com/MitoBlock/onboarding-program-webapp/tree/main/backend) that broadcast tx to blockchain through client code (go func in GIN).

## CLI 

`mitoblockchaind` is the daemon name. Currently, there is no documentation for commands of blockchain. The best way to learn commands is `mitoblockchaind --help`. 

## Install
To install the latest version of this blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/mitoblock/mitoblockchain@latest! | sudo bash
```
`mitoblock/mitoblockchain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).
