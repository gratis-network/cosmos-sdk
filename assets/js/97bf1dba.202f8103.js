"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[3026],{3905:(e,t,r)=>{r.d(t,{Zo:()=>l,kt:()=>u});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function a(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var c=n.createContext({}),p=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):s(s({},t),e)),r},l=function(e){var t=p(e.components);return n.createElement(c.Provider,{value:t},e.children)},f={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,l=a(e,["components","mdxType","originalType","parentName"]),d=p(r),u=o,m=d["".concat(c,".").concat(u)]||d[u]||f[u]||i;return r?n.createElement(m,s(s({ref:t},l),{},{components:r})):n.createElement(m,s({ref:t},l))}));function u(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,s=new Array(i);s[0]=d;var a={};for(var c in t)hasOwnProperty.call(t,c)&&(a[c]=t[c]);a.originalType=e,a.mdxType="string"==typeof e?e:o,s[1]=a;for(var p=2;p<i;p++)s[p]=r[p];return n.createElement.apply(null,s)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},4783:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>c,contentTitle:()=>s,default:()=>f,frontMatter:()=>i,metadata:()=>a,toc:()=>p});var n=r(7462),o=(r(7294),r(3905));const i={sidebar_position:1},s="Specifications",a={unversionedId:"spec/README",id:"version-v0.47/spec/README",title:"Specifications",description:"This directory contains specifications for the modules of the Cosmos SDK as well as Interchain Standards (ICS) and other specifications.",source:"@site/versioned_docs/version-v0.47/spec/README.md",sourceDirName:"spec",slug:"/spec/",permalink:"/v0.47/spec/",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"ADR {ADR-NUMBER}:",permalink:"/v0.47/architecture/adr-template"},next:{title:"Specification of Specifications",permalink:"/v0.47/spec/SPEC-SPEC"}},c={},p=[{value:"Cosmos SDK specifications",id:"cosmos-sdk-specifications",level:2},{value:"Modules specifications",id:"modules-specifications",level:2},{value:"CometBFT",id:"cometbft",level:2}],l={toc:p};function f(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"specifications"},"Specifications"),(0,o.kt)("p",null,"This directory contains specifications for the modules of the Cosmos SDK as well as Interchain Standards (ICS) and other specifications."),(0,o.kt)("p",null,"Cosmos SDK applications hold this state in a Merkle store. Updates to\nthe store may be made during transactions and at the beginning and end of every\nblock."),(0,o.kt)("h2",{id:"cosmos-sdk-specifications"},"Cosmos SDK specifications"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"./store"},"Store")," - The core Merkle store that holds the state."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/v0.47/spec/addresses/bech32"},"Bech32")," - Address format for Cosmos SDK applications.")),(0,o.kt)("h2",{id:"modules-specifications"},"Modules specifications"),(0,o.kt)("p",null,"Go the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/main/x/README.md"},"module directory")),(0,o.kt)("h2",{id:"cometbft"},"CometBFT"),(0,o.kt)("p",null,"For details on the underlying blockchain and p2p protocols, see\nthe ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cometbft/cometbft/tree/main/spec"},"CometBFT specification"),"."))}f.isMDXComponent=!0}}]);