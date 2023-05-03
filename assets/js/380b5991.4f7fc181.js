"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[9050],{3905:(e,n,t)=>{t.d(n,{Zo:()=>d,kt:()=>u});var a=t(7294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function r(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,a,i=function(e,n){if(null==e)return{};var t,a,i={},o=Object.keys(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)t=o[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var l=a.createContext({}),p=function(e){var n=a.useContext(l),t=n;return e&&(t="function"==typeof e?e(n):r(r({},n),e)),t},d=function(e){var n=p(e.components);return a.createElement(l.Provider,{value:n},e.children)},m={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},c=a.forwardRef((function(e,n){var t=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,d=s(e,["components","mdxType","originalType","parentName"]),c=p(t),u=i,h=c["".concat(l,".").concat(u)]||c[u]||m[u]||o;return t?a.createElement(h,r(r({ref:n},d),{},{components:t})):a.createElement(h,r({ref:n},d))}));function u(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=t.length,r=new Array(o);r[0]=c;var s={};for(var l in n)hasOwnProperty.call(n,l)&&(s[l]=n[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,r[1]=s;for(var p=2;p<o;p++)r[p]=t[p];return a.createElement.apply(null,r)}return a.createElement.apply(null,t)}c.displayName="MDXCreateElement"},8531:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>r,default:()=>m,frontMatter:()=>o,metadata:()=>s,toc:()=>p});var a=t(7462),i=(t(7294),t(3905));const o={sidebar_position:1},r="Running a Node",s={unversionedId:"run-node/run-node",id:"run-node/run-node",title:"Running a Node",description:"Now that the application is ready and the keyring populated, it's time to see how to run the blockchain node. In this section, the application we are running is called simapp, and its corresponding CLI binary simd.",source:"@site/docs/run-node/01-run-node.md",sourceDirName:"run-node",slug:"/run-node/run-node",permalink:"/main/run-node/run-node",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Setting up the keyring",permalink:"/main/run-node/keyring"},next:{title:"Interacting with the Node",permalink:"/main/run-node/interact-node"}},l={},p=[{value:"Initialize the Chain",id:"initialize-the-chain",level:2},{value:"Updating Some Default Settings",id:"updating-some-default-settings",level:2},{value:"Client Interaction",id:"client-interaction",level:3},{value:"Adding Genesis Accounts",id:"adding-genesis-accounts",level:2},{value:"Configuring the Node Using <code>app.toml</code> and <code>config.toml</code>",id:"configuring-the-node-using-apptoml-and-configtoml",level:2},{value:"Run a Localnet",id:"run-a-localnet",level:2},{value:"Logging",id:"logging",level:2}],d={toc:p};function m(e){let{components:n,...t}=e;return(0,i.kt)("wrapper",(0,a.Z)({},d,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"running-a-node"},"Running a Node"),(0,i.kt)("admonition",{title:"Synopsis",type:"note"},(0,i.kt)("p",{parentName:"admonition"},"Now that the application is ready and the keyring populated, it's time to see how to run the blockchain node. In this section, the application we are running is called ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/tree/main/simapp"},(0,i.kt)("inlineCode",{parentName:"a"},"simapp")),", and its corresponding CLI binary ",(0,i.kt)("inlineCode",{parentName:"p"},"simd"),".")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,i.kt)("ul",{parentName:"admonition"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/main/basics/app-anatomy"},"Anatomy of a Cosmos SDK Application")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/main/run-node/keyring"},"Setting up the keyring")))),(0,i.kt)("h2",{id:"initialize-the-chain"},"Initialize the Chain"),(0,i.kt)("admonition",{type:"warning"},(0,i.kt)("p",{parentName:"admonition"},"Make sure you can build your own binary, and replace ",(0,i.kt)("inlineCode",{parentName:"p"},"simd")," with the name of your binary in the snippets.")),(0,i.kt)("p",null,"Before actually running the node, we need to initialize the chain, and most importantly its genesis file. This is done with the ",(0,i.kt)("inlineCode",{parentName:"p"},"init")," subcommand:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"# The argument <moniker> is the custom username of your node, it should be human-readable.\nsimd init <moniker> --chain-id my-test-chain\n")),(0,i.kt)("p",null,"The command above creates all the configuration files needed for your node to run, as well as a default genesis file, which defines the initial state of the network. All these configuration files are in ",(0,i.kt)("inlineCode",{parentName:"p"},"~/.simapp")," by default, but you can overwrite the location of this folder by passing the ",(0,i.kt)("inlineCode",{parentName:"p"},"--home")," flag."),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"~/.simapp")," folder has the following structure:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},".                                   # ~/.simapp\n  |- data                           # Contains the databases used by the node.\n  |- config/\n      |- app.toml                   # Application-related configuration file.\n      |- config.toml                # CometBFT-related configuration file.\n      |- genesis.json               # The genesis file.\n      |- node_key.json              # Private key to use for node authentication in the p2p protocol.\n      |- priv_validator_key.json    # Private key to use as a validator in the consensus protocol.\n")),(0,i.kt)("h2",{id:"updating-some-default-settings"},"Updating Some Default Settings"),(0,i.kt)("p",null,"If you want to change any field values in configuration files (for ex: genesis.json) you can use ",(0,i.kt)("inlineCode",{parentName:"p"},"jq")," (",(0,i.kt)("a",{parentName:"p",href:"https://stedolan.github.io/jq/download/"},"installation")," & ",(0,i.kt)("a",{parentName:"p",href:"https://stedolan.github.io/jq/manual/#Assignment"},"docs"),") & ",(0,i.kt)("inlineCode",{parentName:"p"},"sed")," commands to do that. Few examples are listed here."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"# to change the chain-id\njq '.chain_id = \"testing\"' genesis.json > temp.json && mv temp.json genesis.json\n\n# to enable the api server\nsed -i '/\\[api\\]/,+3 s/enable = false/enable = true/' app.toml\n\n# to change the voting_period\njq '.app_state.gov.voting_params.voting_period = \"600s\"' genesis.json > temp.json && mv temp.json genesis.json\n\n# to change the inflation\njq '.app_state.mint.minter.inflation = \"0.300000000000000000\"' genesis.json > temp.json && mv temp.json genesis.json\n")),(0,i.kt)("h3",{id:"client-interaction"},"Client Interaction"),(0,i.kt)("p",null,"When instantiating a node, GRPC and REST are defaulted to localhost to avoid unknown exposure of your node to the public. It is recommended to not expose these endpoints without a proxy that can handle load balancing or authentication is setup between your node and the public. "),(0,i.kt)("admonition",{type:"tip"},(0,i.kt)("p",{parentName:"admonition"},"A commonly used tool for this is ",(0,i.kt)("a",{parentName:"p",href:"https://nginx.org"},"nginx"),".")),(0,i.kt)("h2",{id:"adding-genesis-accounts"},"Adding Genesis Accounts"),(0,i.kt)("p",null,"Before starting the chain, you need to populate the state with at least one account. To do so, first ",(0,i.kt)("a",{parentName:"p",href:"/main/run-node/keyring#adding-keys-to-the-keyring"},"create a new account in the keyring")," named ",(0,i.kt)("inlineCode",{parentName:"p"},"my_validator")," under the ",(0,i.kt)("inlineCode",{parentName:"p"},"test")," keyring backend (feel free to choose another name and another backend)."),(0,i.kt)("p",null,"Now that you have created a local account, go ahead and grant it some ",(0,i.kt)("inlineCode",{parentName:"p"},"stake")," tokens in your chain's genesis file. Doing so will also make sure your chain is aware of this account's existence:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"simd genesis add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake\n")),(0,i.kt)("p",null,"Recall that ",(0,i.kt)("inlineCode",{parentName:"p"},"$MY_VALIDATOR_ADDRESS")," is a variable that holds the address of the ",(0,i.kt)("inlineCode",{parentName:"p"},"my_validator")," key in the ",(0,i.kt)("a",{parentName:"p",href:"/main/run-node/keyring#adding-keys-to-the-keyring"},"keyring"),". Also note that the tokens in the Cosmos SDK have the ",(0,i.kt)("inlineCode",{parentName:"p"},"{amount}{denom}")," format: ",(0,i.kt)("inlineCode",{parentName:"p"},"amount")," is is a 18-digit-precision decimal number, and ",(0,i.kt)("inlineCode",{parentName:"p"},"denom")," is the unique token identifier with its denomination key (e.g. ",(0,i.kt)("inlineCode",{parentName:"p"},"atom")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"uatom"),"). Here, we are granting ",(0,i.kt)("inlineCode",{parentName:"p"},"stake")," tokens, as ",(0,i.kt)("inlineCode",{parentName:"p"},"stake")," is the token identifier used for staking in ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/tree/main/simapp"},(0,i.kt)("inlineCode",{parentName:"a"},"simapp")),". For your own chain with its own staking denom, that token identifier should be used instead."),(0,i.kt)("p",null,"Now that your account has some tokens, you need to add a validator to your chain. Validators are special full-nodes that participate in the consensus process (implemented in the ",(0,i.kt)("a",{parentName:"p",href:"/main/intro/sdk-app-architecture#cometbft"},"underlying consensus engine"),") in order to add new blocks to the chain. Any account can declare its intention to become a validator operator, but only those with sufficient delegation get to enter the active set (for example, only the top 125 validator candidates with the most delegation get to be validators in the Cosmos Hub). For this guide, you will add your local node (created via the ",(0,i.kt)("inlineCode",{parentName:"p"},"init")," command above) as a validator of your chain. Validators can be declared before a chain is first started via a special transaction included in the genesis file called a ",(0,i.kt)("inlineCode",{parentName:"p"},"gentx"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"# Create a gentx.\nsimd genesis gentx my_validator 100000000stake --chain-id my-test-chain --keyring-backend test\n\n# Add the gentx to the genesis file.\nsimd genesis collect-gentxs\n")),(0,i.kt)("p",null,"A ",(0,i.kt)("inlineCode",{parentName:"p"},"gentx")," does three things:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Registers the ",(0,i.kt)("inlineCode",{parentName:"li"},"validator")," account you created as a validator operator account (i.e. the account that controls the validator)."),(0,i.kt)("li",{parentName:"ol"},"Self-delegates the provided ",(0,i.kt)("inlineCode",{parentName:"li"},"amount")," of staking tokens."),(0,i.kt)("li",{parentName:"ol"},"Link the operator account with a CometBFT node pubkey that will be used for signing blocks. If no ",(0,i.kt)("inlineCode",{parentName:"li"},"--pubkey")," flag is provided, it defaults to the local node pubkey created via the ",(0,i.kt)("inlineCode",{parentName:"li"},"simd init")," command above.")),(0,i.kt)("p",null,"For more information on ",(0,i.kt)("inlineCode",{parentName:"p"},"gentx"),", use the following command:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"simd genesis gentx --help\n")),(0,i.kt)("h2",{id:"configuring-the-node-using-apptoml-and-configtoml"},"Configuring the Node Using ",(0,i.kt)("inlineCode",{parentName:"h2"},"app.toml")," and ",(0,i.kt)("inlineCode",{parentName:"h2"},"config.toml")),(0,i.kt)("p",null,"The Cosmos SDK automatically generates two configuration files inside ",(0,i.kt)("inlineCode",{parentName:"p"},"~/.simapp/config"),":"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"config.toml"),": used to configure the CometBFT, learn more on ",(0,i.kt)("a",{parentName:"li",href:"https://docs.cometbft.com/v0.37/core/configuration"},"CometBFT's documentation"),","),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"app.toml"),": generated by the Cosmos SDK, and used to configure your app, such as state pruning strategies, telemetry, gRPC and REST servers configuration, state sync...")),(0,i.kt)("p",null,"Both files are heavily commented, please refer to them directly to tweak your node."),(0,i.kt)("p",null,"One example config to tweak is the ",(0,i.kt)("inlineCode",{parentName:"p"},"minimum-gas-prices")," field inside ",(0,i.kt)("inlineCode",{parentName:"p"},"app.toml"),", which defines the minimum gas prices the validator node is willing to accept for processing a transaction. Depending on the chain, it might be an empty string or not. If it's empty, make sure to edit the field with some value, for example ",(0,i.kt)("inlineCode",{parentName:"p"},"10token"),", or else the node will halt on startup. For the purpose of this tutorial, let's set the minimum gas price to 0:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-toml"},' # The minimum gas prices a validator is willing to accept for processing a\n # transaction. A transaction\'s fees must meet the minimum of any denomination\n # specified in this config (e.g. 0.25token1;0.0001token2).\n minimum-gas-prices = "0stake"\n')),(0,i.kt)("h2",{id:"run-a-localnet"},"Run a Localnet"),(0,i.kt)("p",null,"Now that everything is set up, you can finally start your node:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"simd start\n")),(0,i.kt)("p",null,"You should see blocks come in."),(0,i.kt)("p",null,"The previous command allow you to run a single node. This is enough for the next section on interacting with this node, but you may wish to run multiple nodes at the same time, and see how consensus happens between them."),(0,i.kt)("p",null,"The naive way would be to run the same commands again in separate terminal windows. This is possible, however in the Cosmos SDK, we leverage the power of ",(0,i.kt)("a",{parentName:"p",href:"https://docs.docker.com/compose/"},"Docker Compose")," to run a localnet. If you need inspiration on how to set up your own localnet with Docker Compose, you can have a look at the Cosmos SDK's ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/docker-compose.yml"},(0,i.kt)("inlineCode",{parentName:"a"},"docker-compose.yml")),"."),(0,i.kt)("h2",{id:"logging"},"Logging"),(0,i.kt)("p",null,"Logging provides a way to see what is going on with a node. By default the info level is set. This is a global level and all info logs will be outputted to the terminal. If you would like to filter specific logs to the terminal instead of all, then setting ",(0,i.kt)("inlineCode",{parentName:"p"},"module:log_level")," is how this can work. "),(0,i.kt)("p",null,"Example: "),(0,i.kt)("p",null,"In config.toml:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-toml"},'log_level: "state:info,p2p:info,consensus:info,x/staking:info,x/ibc:info,*error"\n')))}m.isMDXComponent=!0}}]);