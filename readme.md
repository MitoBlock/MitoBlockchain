# Mito Block

It is difficult for start-ups and small businesses to collaborate with each other to boost their sales / revenue / number of users and provide strong elements for user stickiness. Our application is a blockchain protocol that acts as an ecosystem for startups and small businesses to freely exchange value with each other with the help of Blockchain & Web3. 

**Mito Block / Mito Blockchain** is a reliable, secure, and scalable PoA / PoS blockchain with WASM for smart contracts. It is built using Cosmos SDK and created with [Starport](https://starport.com).

## Tech Stack of our application 

| Component / Category                                             | Technology being used                  |
| -------------                                                    |:-------------:                         | 
| State machine / Application Layer of the blockchain              | Cosmos SDK                             | 
| Consensus Engine / Consensus and Network Layer of the blockchain | Tendermint Core                        |
| Smart contract Functionality (WASM)                              | CosmWasm                               |
| CLI Tool for creating and managing the blockchain                | Starport                               |
| Interoperability                                                 | IBC & Peg-zones                        |
| RPC Framework                                                    | gRPC                                   |
| Languages                                                        | JavaScript, TypeScript, Go, C and Rust |

## Expected performance of our application

| Category           | Expected Value    |
| -------------      |:-------------:    | 
| Block time         | <8 sec            | 
| Throughput         | >1000 tps         |

## Get started with the application

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/mitoblock/mitoblockchain@latest! | sudo bash
```
`mitoblock/mitoblockchain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)
