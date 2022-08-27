# checkers
**checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Following [cosmos sdk tutorial](https://tutorials.cosmos.network/academy/3-my-own-chain/), this app is made.

This blockchain feature
- We can play [checkers](https://www.ducksters.com/games/checkers_rules.php).
- Players can stake money.

Tutorial official code is 
[b9-checkers-academy-draft](https://github.com/cosmos/b9-checkers-academy-draft/tree/main).

## Set up

```bash
ignite chain init
```

`init` command initialize your chain.

You have to memo account_info.
```bash
🛠️  Building proto...
📦 Installing dependencies...
🛠️  Building the blockchain...
🙂 Created account "alice" with address "cosmos16dzyygs6x02t5ed6p9a9n0p0nds79ndmxz85cw" with mnemonic: "win distance sign gentle census cash animal tip actress polar amount weasel unknown twelve pudding broccoli broccoli island north weapon ball enough inhale summer"
🙂 Created account "bob" with address "cosmos1l8sxl80w3wk8nwje202ljqqtqhhawlz22pye85" with mnemonic: "whip leader habit rice gather copy point choice toward science retreat achieve pride banana exhaust wage drip thumb ghost nice length mosquito knee bottom"
```

Set environment_variable to bash
```bash
# set account_info
export alice=cosmos16dzyygs6x02t5ed6p9a9n0p0nds79ndmxz85cw
export bob=cosmos1l8sxl80w3wk8nwje202ljqqtqhhawlz22pye85
```

```bash
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Get started

Check bob balance.
```bash
checkersd query bank balances $bob
```

This returns:
```bash
balances:
- amount: "100000000"
  denom: stake
- amount: "10000"
  denom: token
pagination:
  next_key: null
  total: "0"
```

You can make use of this other token to create a new game that costs 1 token:
```bash
checkersd tx checkers create-game $alice $bob 1 token --from $alice
```

Which mentions:
```bash
...
- key: Wager
  value: "1"
- key: Token
  value: token
...
```

Have Bob play once:
```bash
checkersd tx checkers play-move 1 1 2 2 3 --from $bob
```

This returns:
```
...
raw_log: '[{"events":[{"type":"message","attributes":[{"key":"action","value":"PlayMove"}]}]}]'
```
Confirm the move went through with your one-line formatter

Has Bob been charged the wager?
```bash
checkersd query bank balances $bob
```

This returns:
```
balances:
- amount: "100000000"
  denom: stake
- amount: "9999"
  denom: token
pagination:
  next_key: null
  total: "0"
```
Correct. You made it possible to wager any token. That includes IBC tokens.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

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
curl https://get.ignite.com/alice/checkers@latest! | sudo bash
```
`alice/checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)