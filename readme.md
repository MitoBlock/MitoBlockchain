# Mito Blockchain dev

**This is dev / experimental repo of original mitoblock.**

It is difficult for start-ups and small businesses to collaborate with each other to boost their sales / revenue / number of users and provide strong elements for user stickiness. Our application is a blockchain protocol that acts as an ecosystem for startups and small businesses to freely exchange value with each other with the help of Blockchain & Web3. 

**Mito Block / Mito Blockchain dev** is a reliable, secure, and scalable PoA / PoS blockchain with WASM for smart contracts. It is built  using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

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

## Get started

In order to run this blockchain node, you will need the following:
1. [Go](https://go.dev/)
2. [Ignite CLI](https://docs.ignite.com/guide/install)

To run the chain,
```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development. The `-r` flag can be used to reset the state.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of this blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/mitoblock/mitoblockchaindev@latest! | sudo bash
```
`mitoblock/mitoblockchaindev` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more about ignite here

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
